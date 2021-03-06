package thompsonsampling

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	sync "sync"
	"syscall"
	"time"

	log "github.com/Golang-Tools/loggerhelper"
	se "github.com/Golang-Tools/schema-entry-go"
	"gonum.org/v1/gonum/stat/distuv"

	"github.com/liyue201/grpc-lb/common"
	"github.com/liyue201/grpc-lb/registry"
	zk "github.com/liyue201/grpc-lb/registry/zookeeper"

	rp "github.com/Golang-Tools/redishelper/proxy"
	grpc "google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

//Server grpc的服务器结构体
//服务集成了如下特性:
//设置收发最大消息长度
//健康检测
//gzip做消息压缩
//接口反射
//channelz支持
//TLS支持
//keep alive 支持
type Server struct {
	AppName       string   `json:"app_name,omitempty" jsonschema:"description=服务名"`
	AppVersion    string   `json:"app_version,omitempty" jsonschema:"description=服务版本"`
	Address       string   `json:"address,omitempty" jsonschema:"description=服务的主机和端口"`
	LogLevel      string   `json:"log_level,omitempty" jsonschema:"description=项目的log等级,enum=TRACE,enum=DEBUG,enum=INFO,enum=WARN,enum=ERROR"`
	ZookeeperURL  []string `json:"zookeeper_url,omitempty" jsonschema:"description=负载均衡使用的zookeeper地址序列以逗号分隔"`
	BalanceWeight string   `json:"balance_weight,omitempty" jsonschema:"description=负载均衡的权重"`

	MaxRecvMsgSize int `json:"max_rec_msg_size,omitempty" jsonschema:"description=允许接收的最大消息长度"`
	MaxSendMsgSize int `json:"max_send_msg_size,omitempty" jsonschema:"description=允许发送的最大消息长度"`

	InitialWindowSize     int `json:"initial_window_size,omitempty" jsonschema:"description=基于Stream的滑动窗口大小"`
	InitialConnWindowSize int `json:"initial_conn_window_size,omitempty" jsonschema:"description=基于Connection的滑动窗口大小"`
	MaxConcurrentStreams  int `json:"max_concurrent_streams,omitempty" jsonschema:"description=一个连接中最大并发Stream数"`

	UseChannelz                             bool   `json:"use_channelz,omitempty" jsonschema:"description=是否使用channelz协助优化"`
	ServerCertPath                          string `json:"server_cert_path ,omitempty" jsonschema:"description=使用TLS时服务端的证书位置"`
	ServerKeyPath                           string `json:"server_key_path,omitempty" jsonschema:"description=使用TLS时服务端证书的私钥位置"`
	MaxConnectionIdle                       int    `json:"max_connection_idle,omitempty" jsonschema:"description=客户端连接的最大空闲时长"`
	MaxConnectionAge                        int    `json:"max_connection_age,omitempty" jsonschema:"description=如果连接存活超过n则发送goaway"`
	MaxConnectionAgeGrace                   int    `json:"max_connection_age_grace,omitempty" jsonschema:"description=强制关闭连接之前允许等待的rpc在n秒内完成"`
	KeepaliveTime                           int    `json:"keepalive_time,omitempty" jsonschema:"description=空闲连接每隔n秒ping一次客户端已确保连接存活"`
	KeepaliveTimeout                        int    `json:"keepalive_timeout,omitempty" jsonschema:"description=ping时长超过n则认为连接已死"`
	KeepaliveEnforcementMinTime             int    `json:"keepalive_enforement_min_time,omitempty" jsonschema:"description=如果客户端超过每n秒ping一次则终止连接"`
	KeepaliveEnforcementPermitWithoutStream bool   `json:"keepalive_enforement_permit_without_stream,omitempty" jsonschema:"description=即使没有活动流也允许ping"`

	RedisURL          string `json:"redis_url,omitempty" jsonschema:"required,description=保存点击和未点击数据的redis位置"`
	QueryRedisTimeout int    `json:"query_redis_timeout,omitempty" jsonschema:"description=请求redis的超时时长"`
	DefaultKeyTTL     int    `json:"default_key_ttl,omitempty" jsonschema:"description=保存键的默认过期时间"`

	service       *registry.ServiceInfo
	healthservice *health.Server
	registrar     *zk.Registrar

	betapool *sync.Pool
}

func (s *Server) QueryRedisCtx() (context.Context, context.CancelFunc) {
	var ctx context.Context
	var cancel context.CancelFunc
	if s.QueryRedisTimeout > 0 {
		timeout := time.Duration(s.QueryRedisTimeout) * time.Millisecond
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	return ctx, cancel
}

//Main 服务的入口函数
func (s *Server) Main() {
	// 初始化log
	log.Init(s.LogLevel, log.Dict{
		"app_name":    s.AppName,
		"app_version": s.AppVersion,
	})
	log.Info("获得参数", nil, log.Dict{"ServiceConfig": s}, nil)

	rp.Proxy.InitFromURL(s.RedisURL)
	defer rp.Proxy.Close()

	s.betapool = &sync.Pool{
		New: func() interface{} {
			return new(distuv.Beta)
		},
	}
	s.Run()
}

//RunServer 启动服务
func (s *Server) RunServer() {
	lis, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Error("Failed to Listen", log.Dict{"error": err.Error(), "address": s.Address})
		os.Exit(1)
	}
	opts := []grpc.ServerOption{}
	if s.MaxRecvMsgSize != 0 {
		opts = append(opts, grpc.MaxRecvMsgSize(s.MaxRecvMsgSize))
	}
	if s.MaxSendMsgSize != 0 {
		opts = append(opts, grpc.MaxSendMsgSize(s.MaxSendMsgSize))
	}
	if s.InitialWindowSize != 0 {
		opts = append(opts, grpc.InitialWindowSize(int32(s.InitialWindowSize)))
	}
	if s.InitialConnWindowSize != 0 {
		opts = append(opts, grpc.InitialConnWindowSize(int32(s.InitialConnWindowSize)))
	}
	if s.MaxConcurrentStreams != 0 {
		opts = append(opts, grpc.MaxConcurrentStreams(uint32(s.MaxConcurrentStreams)))
	}
	if s.ServerCertPath != "" && s.ServerKeyPath != "" {
		creds, err := credentials.NewServerTLSFromFile(s.ServerCertPath, s.ServerKeyPath)
		if err != nil {
			log.Warn("Failed to Listen as a TLS Server", log.Dict{"error": err.Error()})
		}
		opts = append(opts, grpc.Creds(creds))
	}
	if s.MaxConnectionIdle != 0 || s.MaxConnectionAge != 0 || s.MaxConnectionAgeGrace != 0 || s.KeepaliveTime != 0 || s.KeepaliveTimeout != 0 {
		kasp := keepalive.ServerParameters{
			MaxConnectionIdle:     time.Duration(s.MaxConnectionIdle) * time.Second,
			MaxConnectionAge:      time.Duration(s.MaxConnectionAge) * time.Second,
			MaxConnectionAgeGrace: time.Duration(s.MaxConnectionAgeGrace) * time.Second,
			Time:                  time.Duration(s.KeepaliveTime) * time.Second,
			Timeout:               time.Duration(s.KeepaliveTimeout) * time.Second,
		}
		opts = append(opts, grpc.KeepaliveParams(kasp))
	}

	if s.KeepaliveEnforcementMinTime != 0 || s.KeepaliveEnforcementPermitWithoutStream == true {
		kaep := keepalive.EnforcementPolicy{
			MinTime:             time.Duration(s.KeepaliveEnforcementMinTime) * time.Second,
			PermitWithoutStream: s.KeepaliveEnforcementPermitWithoutStream,
		}
		opts = append(opts, grpc.KeepaliveEnforcementPolicy(kaep))
	}
	gs := grpc.NewServer(opts...)
	defer gs.Stop()
	s.healthservice = health.NewServer()
	healthpb.RegisterHealthServer(gs, s.healthservice)

	RegisterTHOMPSONSAMPLINGServer(gs, s)
	reflection.Register(gs)
	if s.UseChannelz {
		channelz.RegisterChannelzServiceToServer(gs)
	}
	log.Info("Server Start", log.Dict{"address": s.Address})
	err = gs.Serve(lis)
	if err != nil {
		log.Error("Failed to Serve", log.Dict{"error": err})
		os.Exit(1)
	}
}

//RegistService 注册服务到zookeeper
func (s *Server) RegistService() {
	if s.registrar != nil && s.service != nil {
		log.Warn("服务注册已经初始化")
		return
	}
	port := strings.Split(s.Address, ":")[1]
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Error("获取本地ip失败", log.Dict{"place": "RegistService", "err": err})
		os.Exit(1)
	}
	ip := ""
	for _, _ip := range addrs {
		IP := _ip.String()
		if strings.HasPrefix(IP, "172.16.1.") {
			if strings.Contains(IP, "/") {
				ip = strings.Split(IP, "/")[0]
			} else {
				ip = IP
			}
			break
		}
	}
	if ip == "" {
		log.Error("未找到ip", log.Dict{"place": "RegistService"})
		os.Exit(1)
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.Error("获取本地容器hostname失败", log.Dict{"place": "RegistService", "err": err})
		os.Exit(1)
	}
	service := &registry.ServiceInfo{
		InstanceId: hostname,
		Name:       s.AppName,
		Version:    s.AppVersion,
		Address:    fmt.Sprintf("%s:%s", ip, port),
		Metadata:   metadata.Pairs(common.WeightKey, s.BalanceWeight),
	}
	log.Info("注册的服务", log.Dict{"service": *service})
	registrar, err := zk.NewRegistrar(
		&zk.Config{
			ZkServers:      s.ZookeeperURL,
			RegistryDir:    "/backend/services",
			SessionTimeout: time.Second,
		})
	if err != nil {
		log.Error("regist error", log.Dict{"err": err})
		os.Exit(1)
	}
	s.registrar = registrar
	s.service = service

}

//Run 执行grpc服务
func (s *Server) Run() {
	if len(s.ZookeeperURL) == 0 || s.ZookeeperURL == nil {
		s.RunServer()
	} else {
		s.RegistService()
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			s.RunServer()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			s.registrar.Register(s.service)
			wg.Done()
		}()
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
		s.registrar.Unregister(s.service)
		// serv.Stop()
		wg.Wait()
	}
}

var ServNode, _ = se.New(
	&se.EntryPointMeta{Name: "thompson_sampling", Usage: "thompson_sampling [options]"},
	&Server{
		AppName:           "thompsonsampling",
		AppVersion:        "0.0.0",
		LogLevel:          "DEBUG",
		Address:           "0.0.0.0:5000",
		RedisURL:          "redis://localhost",
		QueryRedisTimeout: 50,
		DefaultKeyTTL:     3600,
	},
)
