// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: protocol/proto/hello.proto

package protocol

import (
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

type SexType int32

const (
	SexType_UnKnow SexType = 0 //proto3版本中，首成员必须为0，成员不应有相同的值
	SexType_MALE   SexType = 1 //1男
	SexType_FEMALE SexType = 2 //2女  0未知
)

// Enum value maps for SexType.
var (
	SexType_name = map[int32]string{
		0: "UnKnow",
		1: "MALE",
		2: "FEMALE",
	}
	SexType_value = map[string]int32{
		"UnKnow": 0,
		"MALE":   1,
		"FEMALE": 2,
	}
)

func (x SexType) Enum() *SexType {
	p := new(SexType)
	*p = x
	return p
}

func (x SexType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SexType) Descriptor() protoreflect.EnumDescriptor {
	return file_protocol_proto_hello_proto_enumTypes[0].Descriptor()
}

func (SexType) Type() protoreflect.EnumType {
	return &file_protocol_proto_hello_proto_enumTypes[0]
}

func (x SexType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SexType.Descriptor instead.
func (SexType) EnumDescriptor() ([]byte, []int) {
	return file_protocol_proto_hello_proto_rawDescGZIP(), []int{0}
}

type Say struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                //int类型
	Hello  string   `protobuf:"bytes,2,opt,name=hello,proto3" json:"hello,omitempty"`           //string类型
	Word   []string `protobuf:"bytes,3,rep,name=word,proto3" json:"word,omitempty"`             //string[] 数组
	Arrays []int32  `protobuf:"varint,4,rep,packed,name=arrays,proto3" json:"arrays,omitempty"` //int32[] 数组
}

func (x *Say) Reset() {
	*x = Say{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Say) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Say) ProtoMessage() {}

func (x *Say) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Say.ProtoReflect.Descriptor instead.
func (*Say) Descriptor() ([]byte, []int) {
	return file_protocol_proto_hello_proto_rawDescGZIP(), []int{0}
}

func (x *Say) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Say) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

func (x *Say) GetWord() []string {
	if x != nil {
		return x.Word
	}
	return nil
}

func (x *Say) GetArrays() []int32 {
	if x != nil {
		return x.Arrays
	}
	return nil
}

// 定义一个用户消息
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                                                                            // 姓名字段
	Sex      SexType          `protobuf:"varint,2,opt,name=sex,proto3,enum=protocol.SexType" json:"sex,omitempty"`                                                                       // 性别字段，使用SexType枚举类型
	Id       int64            `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`                                                                                               //id
	Duration int32            `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`                                                                                   //学习的时长 单位秒
	Score    map[string]int32 `protobuf:"bytes,5,rep,name=score,proto3" json:"score,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` //学科 分数的map Map 字段是不能使用repeated关键字修饰。
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_protocol_proto_hello_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetSex() SexType {
	if x != nil {
		return x.Sex
	}
	return SexType_UnKnow
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *UserInfo) GetScore() map[string]int32 {
	if x != nil {
		return x.Score
	}
	return nil
}

// 定义Quote消息
type Quote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Tags  []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"` // 字符串数组类型
}

func (x *Quote) Reset() {
	*x = Quote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quote) ProtoMessage() {}

func (x *Quote) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quote.ProtoReflect.Descriptor instead.
func (*Quote) Descriptor() ([]byte, []int) {
	return file_protocol_proto_hello_proto_rawDescGZIP(), []int{2}
}

func (x *Quote) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Quote) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Quote) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// 定义ListQuote消息
type ListQuote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 引用上面定义的Quote消息类型，作为results字段的类型
	Articles []*Quote         `protobuf:"bytes,1,rep,name=articles,proto3" json:"articles,omitempty"` // repeated关键词标记，说明Quotes字段是一个数组
	Others   *ListQuote_Other `protobuf:"bytes,2,opt,name=others,proto3" json:"others,omitempty"`
}

func (x *ListQuote) Reset() {
	*x = ListQuote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuote) ProtoMessage() {}

func (x *ListQuote) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuote.ProtoReflect.Descriptor instead.
func (*ListQuote) Descriptor() ([]byte, []int) {
	return file_protocol_proto_hello_proto_rawDescGZIP(), []int{3}
}

func (x *ListQuote) GetArticles() []*Quote {
	if x != nil {
		return x.Articles
	}
	return nil
}

func (x *ListQuote) GetOthers() *ListQuote_Other {
	if x != nil {
		return x.Others
	}
	return nil
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_hello_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_hello_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_protocol_proto_hello_proto_rawDescGZIP(), []int{4}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_hello_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_hello_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_protocol_proto_hello_proto_rawDescGZIP(), []int{5}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 嵌套消息定义
type ListQuote_Other struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ListQuote_Other) Reset() {
	*x = ListQuote_Other{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_proto_hello_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQuote_Other) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQuote_Other) ProtoMessage() {}

func (x *ListQuote_Other) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_proto_hello_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQuote_Other.ProtoReflect.Descriptor instead.
func (*ListQuote_Other) Descriptor() ([]byte, []int) {
	return file_protocol_proto_hello_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListQuote_Other) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_protocol_proto_hello_proto protoreflect.FileDescriptor

var file_protocol_proto_hello_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x57, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x61, 0x72, 0x72, 0x61, 0x79, 0x73, 0x22,
	0xde, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x43, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x12, 0x31, 0x0a, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x73, 0x1a, 0x19, 0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x22,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x2b, 0x0a, 0x07, 0x53, 0x65,
	0x78, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x32, 0x45, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x12, 0x3a, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_proto_hello_proto_rawDescOnce sync.Once
	file_protocol_proto_hello_proto_rawDescData = file_protocol_proto_hello_proto_rawDesc
)

func file_protocol_proto_hello_proto_rawDescGZIP() []byte {
	file_protocol_proto_hello_proto_rawDescOnce.Do(func() {
		file_protocol_proto_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_proto_hello_proto_rawDescData)
	})
	return file_protocol_proto_hello_proto_rawDescData
}

var file_protocol_proto_hello_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protocol_proto_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protocol_proto_hello_proto_goTypes = []interface{}{
	(SexType)(0),            // 0: protocol.SexType
	(*Say)(nil),             // 1: protocol.Say
	(*UserInfo)(nil),        // 2: protocol.UserInfo
	(*Quote)(nil),           // 3: protocol.Quote
	(*ListQuote)(nil),       // 4: protocol.ListQuote
	(*HelloRequest)(nil),    // 5: protocol.HelloRequest
	(*HelloReply)(nil),      // 6: protocol.HelloReply
	nil,                     // 7: protocol.UserInfo.ScoreEntry
	(*ListQuote_Other)(nil), // 8: protocol.ListQuote.Other
}
var file_protocol_proto_hello_proto_depIdxs = []int32{
	0, // 0: protocol.UserInfo.sex:type_name -> protocol.SexType
	7, // 1: protocol.UserInfo.score:type_name -> protocol.UserInfo.ScoreEntry
	3, // 2: protocol.ListQuote.articles:type_name -> protocol.Quote
	8, // 3: protocol.ListQuote.others:type_name -> protocol.ListQuote.Other
	5, // 4: protocol.Greeter.SayHello:input_type -> protocol.HelloRequest
	6, // 5: protocol.Greeter.SayHello:output_type -> protocol.HelloReply
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protocol_proto_hello_proto_init() }
func file_protocol_proto_hello_proto_init() {
	if File_protocol_proto_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_proto_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Say); i {
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
		file_protocol_proto_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_protocol_proto_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quote); i {
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
		file_protocol_proto_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuote); i {
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
		file_protocol_proto_hello_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_protocol_proto_hello_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_protocol_proto_hello_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQuote_Other); i {
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
			RawDescriptor: file_protocol_proto_hello_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_proto_hello_proto_goTypes,
		DependencyIndexes: file_protocol_proto_hello_proto_depIdxs,
		EnumInfos:         file_protocol_proto_hello_proto_enumTypes,
		MessageInfos:      file_protocol_proto_hello_proto_msgTypes,
	}.Build()
	File_protocol_proto_hello_proto = out.File
	file_protocol_proto_hello_proto_rawDesc = nil
	file_protocol_proto_hello_proto_goTypes = nil
	file_protocol_proto_hello_proto_depIdxs = nil
}
