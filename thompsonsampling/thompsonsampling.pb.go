// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0-devel
// 	protoc        v3.17.1
// source: thompsonsampling.proto

package thompsonsampling

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UpdateQuery_Type int32

const (
	UpdateQuery_INCR  UpdateQuery_Type = 0
	UpdateQuery_RESET UpdateQuery_Type = 1
)

// Enum value maps for UpdateQuery_Type.
var (
	UpdateQuery_Type_name = map[int32]string{
		0: "INCR",
		1: "RESET",
	}
	UpdateQuery_Type_value = map[string]int32{
		"INCR":  0,
		"RESET": 1,
	}
)

func (x UpdateQuery_Type) Enum() *UpdateQuery_Type {
	p := new(UpdateQuery_Type)
	*p = x
	return p
}

func (x UpdateQuery_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateQuery_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_thompsonsampling_proto_enumTypes[0].Descriptor()
}

func (UpdateQuery_Type) Type() protoreflect.EnumType {
	return &file_thompsonsampling_proto_enumTypes[0]
}

func (x UpdateQuery_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateQuery_Type.Descriptor instead.
func (UpdateQuery_Type) EnumDescriptor() ([]byte, []int) {
	return file_thompsonsampling_proto_rawDescGZIP(), []int{0, 0}
}

type UpdateQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateType        UpdateQuery_Type `protobuf:"varint,1,opt,name=update_type,json=updateType,proto3,enum=thompsonsampling.UpdateQuery_Type" json:"update_type,omitempty"` // 更新类型,是重置还是增量
	BusinessNamespace string           `protobuf:"bytes,2,opt,name=business_namespace,json=businessNamespace,proto3" json:"business_namespace,omitempty"`                    //针对业务控制的命名空间,空字符串含义为全局,redis中以`__global__`替代
	TargetNamespace   string           `protobuf:"bytes,3,opt,name=target_namespace,json=targetNamespace,proto3" json:"target_namespace,omitempty"`                          //针对目标控制的命名空间,指代用户,用户分群等逻辑概念,空字符串含义为全局,redis中以`__global__`替代
	Candidate         string           `protobuf:"bytes,4,opt,name=candidate,proto3" json:"candidate,omitempty"`                                                             // 针对的目标候选人
	Alpha             float64          `protobuf:"fixed64,5,opt,name=alpha,proto3" json:"alpha,omitempty"`                                                                   //候选人的alpha参数
	Beta              float64          `protobuf:"fixed64,6,opt,name=beta,proto3" json:"beta,omitempty"`                                                                     //候选人的beta参数
	Ttl               int64            `protobuf:"varint,7,opt,name=ttl,proto3" json:"ttl,omitempty"`                                                                        //reset键的过期时长,单位s
}

func (x *UpdateQuery) Reset() {
	*x = UpdateQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thompsonsampling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQuery) ProtoMessage() {}

func (x *UpdateQuery) ProtoReflect() protoreflect.Message {
	mi := &file_thompsonsampling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQuery.ProtoReflect.Descriptor instead.
func (*UpdateQuery) Descriptor() ([]byte, []int) {
	return file_thompsonsampling_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateQuery) GetUpdateType() UpdateQuery_Type {
	if x != nil {
		return x.UpdateType
	}
	return UpdateQuery_INCR
}

func (x *UpdateQuery) GetBusinessNamespace() string {
	if x != nil {
		return x.BusinessNamespace
	}
	return ""
}

func (x *UpdateQuery) GetTargetNamespace() string {
	if x != nil {
		return x.TargetNamespace
	}
	return ""
}

func (x *UpdateQuery) GetCandidate() string {
	if x != nil {
		return x.Candidate
	}
	return ""
}

func (x *UpdateQuery) GetAlpha() float64 {
	if x != nil {
		return x.Alpha
	}
	return 0
}

func (x *UpdateQuery) GetBeta() float64 {
	if x != nil {
		return x.Beta
	}
	return 0
}

func (x *UpdateQuery) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alpha float64 `protobuf:"fixed64,1,opt,name=alpha,proto3" json:"alpha,omitempty"` //候选人的alpha参数当前值
	Beta  float64 `protobuf:"fixed64,2,opt,name=beta,proto3" json:"beta,omitempty"`   //候选人的beta参数当前值
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thompsonsampling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thompsonsampling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_thompsonsampling_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateResponse) GetAlpha() float64 {
	if x != nil {
		return x.Alpha
	}
	return 0
}

func (x *UpdateResponse) GetBeta() float64 {
	if x != nil {
		return x.Beta
	}
	return 0
}

type RankQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc              bool     `protobuf:"varint,1,opt,name=desc,proto3" json:"desc,omitempty"`                                                   //是否倒叙从大到小
	BusinessNamespace string   `protobuf:"bytes,2,opt,name=business_namespace,json=businessNamespace,proto3" json:"business_namespace,omitempty"` //针对业务控制的命名空间,空字符串含义为全局,redis中以`__global__`替代
	TargetNamespace   string   `protobuf:"bytes,3,opt,name=target_namespace,json=targetNamespace,proto3" json:"target_namespace,omitempty"`       //针对目标控制的命名空间,指代用户,用户分群等逻辑概念,空字符串含义为全局,redis中以`__global__`替代
	Candidates        []string `protobuf:"bytes,4,rep,name=candidates,proto3" json:"candidates,omitempty"`
}

func (x *RankQuery) Reset() {
	*x = RankQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thompsonsampling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankQuery) ProtoMessage() {}

func (x *RankQuery) ProtoReflect() protoreflect.Message {
	mi := &file_thompsonsampling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankQuery.ProtoReflect.Descriptor instead.
func (*RankQuery) Descriptor() ([]byte, []int) {
	return file_thompsonsampling_proto_rawDescGZIP(), []int{2}
}

func (x *RankQuery) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

func (x *RankQuery) GetBusinessNamespace() string {
	if x != nil {
		return x.BusinessNamespace
	}
	return ""
}

func (x *RankQuery) GetTargetNamespace() string {
	if x != nil {
		return x.TargetNamespace
	}
	return ""
}

func (x *RankQuery) GetCandidates() []string {
	if x != nil {
		return x.Candidates
	}
	return nil
}

type WeightedCandidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidate string  `protobuf:"bytes,1,opt,name=Candidate,proto3" json:"Candidate,omitempty"`
	Weight    float64 `protobuf:"fixed64,2,opt,name=Weight,proto3" json:"Weight,omitempty"`
}

func (x *WeightedCandidate) Reset() {
	*x = WeightedCandidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thompsonsampling_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightedCandidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightedCandidate) ProtoMessage() {}

func (x *WeightedCandidate) ProtoReflect() protoreflect.Message {
	mi := &file_thompsonsampling_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightedCandidate.ProtoReflect.Descriptor instead.
func (*WeightedCandidate) Descriptor() ([]byte, []int) {
	return file_thompsonsampling_proto_rawDescGZIP(), []int{3}
}

func (x *WeightedCandidate) GetCandidate() string {
	if x != nil {
		return x.Candidate
	}
	return ""
}

func (x *WeightedCandidate) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type RankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderedCandidates []*WeightedCandidate `protobuf:"bytes,1,rep,name=ordered_candidates,json=orderedCandidates,proto3" json:"ordered_candidates,omitempty"`
}

func (x *RankResponse) Reset() {
	*x = RankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thompsonsampling_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankResponse) ProtoMessage() {}

func (x *RankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thompsonsampling_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankResponse.ProtoReflect.Descriptor instead.
func (*RankResponse) Descriptor() ([]byte, []int) {
	return file_thompsonsampling_proto_rawDescGZIP(), []int{4}
}

func (x *RankResponse) GetOrderedCandidates() []*WeightedCandidate {
	if x != nil {
		return x.OrderedCandidates
	}
	return nil
}

type TopQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessNamespace string   `protobuf:"bytes,1,opt,name=business_namespace,json=businessNamespace,proto3" json:"business_namespace,omitempty"` //针对业务控制的命名空间,空字符串含义为全局,redis中以`__global__`替代
	TargetNamespace   string   `protobuf:"bytes,2,opt,name=target_namespace,json=targetNamespace,proto3" json:"target_namespace,omitempty"`       //针对目标控制的命名空间,指代用户,用户分群等逻辑概念,空字符串含义为全局,redis中以`__global__`替代
	Candidates        []string `protobuf:"bytes,3,rep,name=candidates,proto3" json:"candidates,omitempty"`
}

func (x *TopQuery) Reset() {
	*x = TopQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thompsonsampling_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopQuery) ProtoMessage() {}

func (x *TopQuery) ProtoReflect() protoreflect.Message {
	mi := &file_thompsonsampling_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopQuery.ProtoReflect.Descriptor instead.
func (*TopQuery) Descriptor() ([]byte, []int) {
	return file_thompsonsampling_proto_rawDescGZIP(), []int{5}
}

func (x *TopQuery) GetBusinessNamespace() string {
	if x != nil {
		return x.BusinessNamespace
	}
	return ""
}

func (x *TopQuery) GetTargetNamespace() string {
	if x != nil {
		return x.TargetNamespace
	}
	return ""
}

func (x *TopQuery) GetCandidates() []string {
	if x != nil {
		return x.Candidates
	}
	return nil
}

type TopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidate string `protobuf:"bytes,1,opt,name=candidate,proto3" json:"candidate,omitempty"`
}

func (x *TopResponse) Reset() {
	*x = TopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thompsonsampling_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopResponse) ProtoMessage() {}

func (x *TopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thompsonsampling_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopResponse.ProtoReflect.Descriptor instead.
func (*TopResponse) Descriptor() ([]byte, []int) {
	return file_thompsonsampling_proto_rawDescGZIP(), []int{6}
}

func (x *TopResponse) GetCandidate() string {
	if x != nil {
		return x.Candidate
	}
	return ""
}

var File_thompsonsampling_proto protoreflect.FileDescriptor

var file_thompsonsampling_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x68, 0x6f, 0x6d, 0x70, 0x73, 0x6f, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74, 0x68, 0x6f, 0x6d, 0x70, 0x73,
	0x6f, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0xa3, 0x02, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x43, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x74, 0x68, 0x6f, 0x6d, 0x70, 0x73, 0x6f, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2d, 0x0a, 0x12, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x62, 0x65, 0x74,
	0x61, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x74, 0x74, 0x6c, 0x22, 0x1b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x4e, 0x43, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x01,
	0x22, 0x3a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x65, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x62, 0x65, 0x74, 0x61, 0x22, 0x99, 0x01, 0x0a,
	0x09, 0x52, 0x61, 0x6e, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x2d,
	0x0a, 0x12, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x11, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x65, 0x64, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0x62, 0x0a, 0x0c, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x74, 0x68, 0x6f, 0x6d, 0x70, 0x73, 0x6f, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x2d, 0x0a, 0x12, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x2b,
	0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x32, 0xea, 0x01, 0x0a, 0x10,
	0x54, 0x48, 0x4f, 0x4d, 0x50, 0x53, 0x4f, 0x4e, 0x53, 0x41, 0x4d, 0x50, 0x4c, 0x49, 0x4e, 0x47,
	0x12, 0x4b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x74, 0x68, 0x6f,
	0x6d, 0x70, 0x73, 0x6f, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x20, 0x2e, 0x74, 0x68, 0x6f, 0x6d,
	0x70, 0x73, 0x6f, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x1b, 0x2e, 0x74, 0x68, 0x6f, 0x6d, 0x70, 0x73, 0x6f, 0x6e,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x1e, 0x2e, 0x74, 0x68, 0x6f, 0x6d, 0x70, 0x73, 0x6f, 0x6e, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x03, 0x54, 0x6f, 0x70, 0x12, 0x1a, 0x2e, 0x74, 0x68,
	0x6f, 0x6d, 0x70, 0x73, 0x6f, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x54,
	0x6f, 0x70, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1d, 0x2e, 0x74, 0x68, 0x6f, 0x6d, 0x70, 0x73,
	0x6f, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x6f, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x74, 0x68,
	0x6f, 0x6d, 0x70, 0x73, 0x6f, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thompsonsampling_proto_rawDescOnce sync.Once
	file_thompsonsampling_proto_rawDescData = file_thompsonsampling_proto_rawDesc
)

func file_thompsonsampling_proto_rawDescGZIP() []byte {
	file_thompsonsampling_proto_rawDescOnce.Do(func() {
		file_thompsonsampling_proto_rawDescData = protoimpl.X.CompressGZIP(file_thompsonsampling_proto_rawDescData)
	})
	return file_thompsonsampling_proto_rawDescData
}

var file_thompsonsampling_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_thompsonsampling_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_thompsonsampling_proto_goTypes = []interface{}{
	(UpdateQuery_Type)(0),     // 0: thompsonsampling.UpdateQuery.Type
	(*UpdateQuery)(nil),       // 1: thompsonsampling.UpdateQuery
	(*UpdateResponse)(nil),    // 2: thompsonsampling.UpdateResponse
	(*RankQuery)(nil),         // 3: thompsonsampling.RankQuery
	(*WeightedCandidate)(nil), // 4: thompsonsampling.WeightedCandidate
	(*RankResponse)(nil),      // 5: thompsonsampling.RankResponse
	(*TopQuery)(nil),          // 6: thompsonsampling.TopQuery
	(*TopResponse)(nil),       // 7: thompsonsampling.TopResponse
}
var file_thompsonsampling_proto_depIdxs = []int32{
	0, // 0: thompsonsampling.UpdateQuery.update_type:type_name -> thompsonsampling.UpdateQuery.Type
	4, // 1: thompsonsampling.RankResponse.ordered_candidates:type_name -> thompsonsampling.WeightedCandidate
	1, // 2: thompsonsampling.THOMPSONSAMPLING.Update:input_type -> thompsonsampling.UpdateQuery
	3, // 3: thompsonsampling.THOMPSONSAMPLING.Rank:input_type -> thompsonsampling.RankQuery
	6, // 4: thompsonsampling.THOMPSONSAMPLING.Top:input_type -> thompsonsampling.TopQuery
	2, // 5: thompsonsampling.THOMPSONSAMPLING.Update:output_type -> thompsonsampling.UpdateResponse
	5, // 6: thompsonsampling.THOMPSONSAMPLING.Rank:output_type -> thompsonsampling.RankResponse
	7, // 7: thompsonsampling.THOMPSONSAMPLING.Top:output_type -> thompsonsampling.TopResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_thompsonsampling_proto_init() }
func file_thompsonsampling_proto_init() {
	if File_thompsonsampling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thompsonsampling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thompsonsampling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thompsonsampling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thompsonsampling_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightedCandidate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thompsonsampling_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thompsonsampling_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_thompsonsampling_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_thompsonsampling_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thompsonsampling_proto_goTypes,
		DependencyIndexes: file_thompsonsampling_proto_depIdxs,
		EnumInfos:         file_thompsonsampling_proto_enumTypes,
		MessageInfos:      file_thompsonsampling_proto_msgTypes,
	}.Build()
	File_thompsonsampling_proto = out.File
	file_thompsonsampling_proto_rawDesc = nil
	file_thompsonsampling_proto_goTypes = nil
	file_thompsonsampling_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// THOMPSONSAMPLINGClient is the client API for THOMPSONSAMPLING service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type THOMPSONSAMPLINGClient interface {
	//Update 更新参数
	Update(ctx context.Context, in *UpdateQuery, opts ...grpc.CallOption) (*UpdateResponse, error)
	//Rank 排序候选集
	Rank(ctx context.Context, in *RankQuery, opts ...grpc.CallOption) (*RankResponse, error)
	//Top 返回可能性最大的后选者
	Top(ctx context.Context, in *TopQuery, opts ...grpc.CallOption) (*TopResponse, error)
}

type tHOMPSONSAMPLINGClient struct {
	cc grpc.ClientConnInterface
}

func NewTHOMPSONSAMPLINGClient(cc grpc.ClientConnInterface) THOMPSONSAMPLINGClient {
	return &tHOMPSONSAMPLINGClient{cc}
}

func (c *tHOMPSONSAMPLINGClient) Update(ctx context.Context, in *UpdateQuery, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/thompsonsampling.THOMPSONSAMPLING/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tHOMPSONSAMPLINGClient) Rank(ctx context.Context, in *RankQuery, opts ...grpc.CallOption) (*RankResponse, error) {
	out := new(RankResponse)
	err := c.cc.Invoke(ctx, "/thompsonsampling.THOMPSONSAMPLING/Rank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tHOMPSONSAMPLINGClient) Top(ctx context.Context, in *TopQuery, opts ...grpc.CallOption) (*TopResponse, error) {
	out := new(TopResponse)
	err := c.cc.Invoke(ctx, "/thompsonsampling.THOMPSONSAMPLING/Top", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// THOMPSONSAMPLINGServer is the server API for THOMPSONSAMPLING service.
type THOMPSONSAMPLINGServer interface {
	//Update 更新参数
	Update(context.Context, *UpdateQuery) (*UpdateResponse, error)
	//Rank 排序候选集
	Rank(context.Context, *RankQuery) (*RankResponse, error)
	//Top 返回可能性最大的后选者
	Top(context.Context, *TopQuery) (*TopResponse, error)
}

// UnimplementedTHOMPSONSAMPLINGServer can be embedded to have forward compatible implementations.
type UnimplementedTHOMPSONSAMPLINGServer struct {
}

func (*UnimplementedTHOMPSONSAMPLINGServer) Update(context.Context, *UpdateQuery) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedTHOMPSONSAMPLINGServer) Rank(context.Context, *RankQuery) (*RankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rank not implemented")
}
func (*UnimplementedTHOMPSONSAMPLINGServer) Top(context.Context, *TopQuery) (*TopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Top not implemented")
}

func RegisterTHOMPSONSAMPLINGServer(s *grpc.Server, srv THOMPSONSAMPLINGServer) {
	s.RegisterService(&_THOMPSONSAMPLING_serviceDesc, srv)
}

func _THOMPSONSAMPLING_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(THOMPSONSAMPLINGServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thompsonsampling.THOMPSONSAMPLING/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(THOMPSONSAMPLINGServer).Update(ctx, req.(*UpdateQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _THOMPSONSAMPLING_Rank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(THOMPSONSAMPLINGServer).Rank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thompsonsampling.THOMPSONSAMPLING/Rank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(THOMPSONSAMPLINGServer).Rank(ctx, req.(*RankQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _THOMPSONSAMPLING_Top_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(THOMPSONSAMPLINGServer).Top(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thompsonsampling.THOMPSONSAMPLING/Top",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(THOMPSONSAMPLINGServer).Top(ctx, req.(*TopQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _THOMPSONSAMPLING_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thompsonsampling.THOMPSONSAMPLING",
	HandlerType: (*THOMPSONSAMPLINGServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _THOMPSONSAMPLING_Update_Handler,
		},
		{
			MethodName: "Rank",
			Handler:    _THOMPSONSAMPLING_Rank_Handler,
		},
		{
			MethodName: "Top",
			Handler:    _THOMPSONSAMPLING_Top_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thompsonsampling.proto",
}
