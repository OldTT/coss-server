// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: api/grpc/v1/group_msg_read.proto

package v1

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

// 设置群聊消息已读请求
type SetGroupMessageReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"msg_id"
	MsgId uint32 `protobuf:"varint,1,opt,name=MsgId,proto3" json:"msg_id"` // 消息ID
	// @inject_tag: json:"dialog_id"
	DialogId uint32 `protobuf:"varint,2,opt,name=DialogId,proto3" json:"dialog_id"` // 对话ID
	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,3,opt,name=GroupId,proto3" json:"group_id"` // 群聊ID
	// @inject_tag: json:"read_at"
	ReadAt int64 `protobuf:"varint,4,opt,name=ReadAt,proto3" json:"read_at"` // 已读时间
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,5,opt,name=UserId,proto3" json:"user_id"` // 用户ID
}

func (x *SetGroupMessageReadRequest) Reset() {
	*x = SetGroupMessageReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGroupMessageReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGroupMessageReadRequest) ProtoMessage() {}

func (x *SetGroupMessageReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGroupMessageReadRequest.ProtoReflect.Descriptor instead.
func (*SetGroupMessageReadRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_msg_read_proto_rawDescGZIP(), []int{0}
}

func (x *SetGroupMessageReadRequest) GetMsgId() uint32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *SetGroupMessageReadRequest) GetDialogId() uint32 {
	if x != nil {
		return x.DialogId
	}
	return 0
}

func (x *SetGroupMessageReadRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *SetGroupMessageReadRequest) GetReadAt() int64 {
	if x != nil {
		return x.ReadAt
	}
	return 0
}

func (x *SetGroupMessageReadRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SetGroupMessagesReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"list"
	List []*SetGroupMessageReadRequest `protobuf:"bytes,1,rep,name=List,proto3" json:"list"`
	// @inject_tag: json:"open_burn_after_reading_time_out"
	OpenBurnAfterReadingTimeOut int64 `protobuf:"varint,2,opt,name=OpenBurnAfterReadingTimeOut,proto3" json:"open_burn_after_reading_time_out"`
}

func (x *SetGroupMessagesReadRequest) Reset() {
	*x = SetGroupMessagesReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGroupMessagesReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGroupMessagesReadRequest) ProtoMessage() {}

func (x *SetGroupMessagesReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGroupMessagesReadRequest.ProtoReflect.Descriptor instead.
func (*SetGroupMessagesReadRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_msg_read_proto_rawDescGZIP(), []int{1}
}

func (x *SetGroupMessagesReadRequest) GetList() []*SetGroupMessageReadRequest {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SetGroupMessagesReadRequest) GetOpenBurnAfterReadingTimeOut() int64 {
	if x != nil {
		return x.OpenBurnAfterReadingTimeOut
	}
	return 0
}

// 设置群聊消息已读响应
type SetGroupMessageReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetGroupMessageReadResponse) Reset() {
	*x = SetGroupMessageReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGroupMessageReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGroupMessageReadResponse) ProtoMessage() {}

func (x *SetGroupMessageReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGroupMessageReadResponse.ProtoReflect.Descriptor instead.
func (*SetGroupMessageReadResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_msg_read_proto_rawDescGZIP(), []int{2}
}

// 获取指定消息已读人员请求
type GetGroupMessageReadersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"msg_id"
	MsgId uint32 `protobuf:"varint,1,opt,name=MsgId,proto3" json:"msg_id"` // 消息ID
	// @inject_tag: json:"dialog_id"
	DialogId uint32 `protobuf:"varint,2,opt,name=DialogId,proto3" json:"dialog_id"` // 对话ID
	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,3,opt,name=GroupId,proto3" json:"group_id"` // 群聊ID
}

func (x *GetGroupMessageReadersRequest) Reset() {
	*x = GetGroupMessageReadersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupMessageReadersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupMessageReadersRequest) ProtoMessage() {}

func (x *GetGroupMessageReadersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupMessageReadersRequest.ProtoReflect.Descriptor instead.
func (*GetGroupMessageReadersRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_msg_read_proto_rawDescGZIP(), []int{3}
}

func (x *GetGroupMessageReadersRequest) GetMsgId() uint32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *GetGroupMessageReadersRequest) GetDialogId() uint32 {
	if x != nil {
		return x.DialogId
	}
	return 0
}

func (x *GetGroupMessageReadersRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

// 获取指定消息已读人员响应
type GetGroupMessageReadersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"user_ids"
	UserIds []string `protobuf:"bytes,1,rep,name=UserIds,proto3" json:"user_ids"` // 已读人员的用户ID列表
}

func (x *GetGroupMessageReadersResponse) Reset() {
	*x = GetGroupMessageReadersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupMessageReadersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupMessageReadersResponse) ProtoMessage() {}

func (x *GetGroupMessageReadersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupMessageReadersResponse.ProtoReflect.Descriptor instead.
func (*GetGroupMessageReadersResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_msg_read_proto_rawDescGZIP(), []int{4}
}

func (x *GetGroupMessageReadersResponse) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type GetGroupMessageReadByMsgIdAndUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"msg_id"
	MsgId uint32 `protobuf:"varint,1,opt,name=MsgId,proto3" json:"msg_id"` // 消息ID
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"user_id"`
}

func (x *GetGroupMessageReadByMsgIdAndUserIdRequest) Reset() {
	*x = GetGroupMessageReadByMsgIdAndUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupMessageReadByMsgIdAndUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupMessageReadByMsgIdAndUserIdRequest) ProtoMessage() {}

func (x *GetGroupMessageReadByMsgIdAndUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupMessageReadByMsgIdAndUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetGroupMessageReadByMsgIdAndUserIdRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_msg_read_proto_rawDescGZIP(), []int{5}
}

func (x *GetGroupMessageReadByMsgIdAndUserIdRequest) GetMsgId() uint32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *GetGroupMessageReadByMsgIdAndUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetGroupMessageReadByMsgIdAndUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"read_at"
	ReadAt int64 `protobuf:"varint,1,opt,name=ReadAt,proto3" json:"read_at"`
}

func (x *GetGroupMessageReadByMsgIdAndUserIdResponse) Reset() {
	*x = GetGroupMessageReadByMsgIdAndUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupMessageReadByMsgIdAndUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupMessageReadByMsgIdAndUserIdResponse) ProtoMessage() {}

func (x *GetGroupMessageReadByMsgIdAndUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_v1_group_msg_read_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupMessageReadByMsgIdAndUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetGroupMessageReadByMsgIdAndUserIdResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_v1_group_msg_read_proto_rawDescGZIP(), []int{6}
}

func (x *GetGroupMessageReadByMsgIdAndUserIdResponse) GetReadAt() int64 {
	if x != nil {
		return x.ReadAt
	}
	return 0
}

var File_api_grpc_v1_group_msg_read_proto protoreflect.FileDescriptor

var file_api_grpc_v1_group_msg_read_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x76, 0x31, 0x22, 0x98, 0x01, 0x0a, 0x1a, 0x53,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x73, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a,
	0x1b, 0x4f, 0x70, 0x65, 0x6e, 0x42, 0x75, 0x72, 0x6e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x1b, 0x4f, 0x70, 0x65, 0x6e, 0x42, 0x75, 0x72, 0x6e, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x75, 0x74, 0x22,
	0x1d, 0x0a, 0x1b, 0x53, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x1e, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x5a, 0x0a, 0x2a, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x42, 0x79,
	0x4d, 0x73, 0x67, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x2b, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x42, 0x79, 0x4d, 0x73, 0x67, 0x49,
	0x64, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x32, 0xf0, 0x02, 0x0a, 0x13, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5f, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x12, 0x23, 0x2e, 0x6d, 0x73, 0x67, 0x5f,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x6d, 0x73, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e,
	0x6d, 0x73, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x73, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8e, 0x01, 0x0a,
	0x23, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x42, 0x79, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x32, 0x2e, 0x6d, 0x73, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x61,
	0x64, 0x42, 0x79, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6d, 0x73, 0x67, 0x5f, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x61, 0x64, 0x42, 0x79, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x73,
	0x69, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x73, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_v1_group_msg_read_proto_rawDescOnce sync.Once
	file_api_grpc_v1_group_msg_read_proto_rawDescData = file_api_grpc_v1_group_msg_read_proto_rawDesc
)

func file_api_grpc_v1_group_msg_read_proto_rawDescGZIP() []byte {
	file_api_grpc_v1_group_msg_read_proto_rawDescOnce.Do(func() {
		file_api_grpc_v1_group_msg_read_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_v1_group_msg_read_proto_rawDescData)
	})
	return file_api_grpc_v1_group_msg_read_proto_rawDescData
}

var file_api_grpc_v1_group_msg_read_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_grpc_v1_group_msg_read_proto_goTypes = []interface{}{
	(*SetGroupMessageReadRequest)(nil),                  // 0: msg_v1.SetGroupMessageReadRequest
	(*SetGroupMessagesReadRequest)(nil),                 // 1: msg_v1.SetGroupMessagesReadRequest
	(*SetGroupMessageReadResponse)(nil),                 // 2: msg_v1.SetGroupMessageReadResponse
	(*GetGroupMessageReadersRequest)(nil),               // 3: msg_v1.GetGroupMessageReadersRequest
	(*GetGroupMessageReadersResponse)(nil),              // 4: msg_v1.GetGroupMessageReadersResponse
	(*GetGroupMessageReadByMsgIdAndUserIdRequest)(nil),  // 5: msg_v1.GetGroupMessageReadByMsgIdAndUserIdRequest
	(*GetGroupMessageReadByMsgIdAndUserIdResponse)(nil), // 6: msg_v1.GetGroupMessageReadByMsgIdAndUserIdResponse
}
var file_api_grpc_v1_group_msg_read_proto_depIdxs = []int32{
	0, // 0: msg_v1.SetGroupMessagesReadRequest.List:type_name -> msg_v1.SetGroupMessageReadRequest
	1, // 1: msg_v1.GroupMessageService.SetGroupMessageRead:input_type -> msg_v1.SetGroupMessagesReadRequest
	3, // 2: msg_v1.GroupMessageService.GetGroupMessageReaders:input_type -> msg_v1.GetGroupMessageReadersRequest
	5, // 3: msg_v1.GroupMessageService.GetGroupMessageReadByMsgIdAndUserId:input_type -> msg_v1.GetGroupMessageReadByMsgIdAndUserIdRequest
	2, // 4: msg_v1.GroupMessageService.SetGroupMessageRead:output_type -> msg_v1.SetGroupMessageReadResponse
	4, // 5: msg_v1.GroupMessageService.GetGroupMessageReaders:output_type -> msg_v1.GetGroupMessageReadersResponse
	6, // 6: msg_v1.GroupMessageService.GetGroupMessageReadByMsgIdAndUserId:output_type -> msg_v1.GetGroupMessageReadByMsgIdAndUserIdResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_grpc_v1_group_msg_read_proto_init() }
func file_api_grpc_v1_group_msg_read_proto_init() {
	if File_api_grpc_v1_group_msg_read_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_v1_group_msg_read_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGroupMessageReadRequest); i {
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
		file_api_grpc_v1_group_msg_read_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGroupMessagesReadRequest); i {
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
		file_api_grpc_v1_group_msg_read_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGroupMessageReadResponse); i {
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
		file_api_grpc_v1_group_msg_read_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupMessageReadersRequest); i {
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
		file_api_grpc_v1_group_msg_read_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupMessageReadersResponse); i {
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
		file_api_grpc_v1_group_msg_read_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupMessageReadByMsgIdAndUserIdRequest); i {
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
		file_api_grpc_v1_group_msg_read_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupMessageReadByMsgIdAndUserIdResponse); i {
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
			RawDescriptor: file_api_grpc_v1_group_msg_read_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_v1_group_msg_read_proto_goTypes,
		DependencyIndexes: file_api_grpc_v1_group_msg_read_proto_depIdxs,
		MessageInfos:      file_api_grpc_v1_group_msg_read_proto_msgTypes,
	}.Build()
	File_api_grpc_v1_group_msg_read_proto = out.File
	file_api_grpc_v1_group_msg_read_proto_rawDesc = nil
	file_api_grpc_v1_group_msg_read_proto_goTypes = nil
	file_api_grpc_v1_group_msg_read_proto_depIdxs = nil
}
