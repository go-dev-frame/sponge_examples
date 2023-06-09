// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/user_gw/v1/user_gw.proto

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_api_user_gw_v1_user_gw_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_api_user_gw_v1_user_gw_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_user_gw_v1_user_gw_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_api_user_gw_v1_user_gw_proto_rawDescGZIP(), []int{3}
}

func (x *LoginReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_api_user_gw_v1_user_gw_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutReply) Reset() {
	*x = LogoutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReply) ProtoMessage() {}

func (x *LogoutReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReply.ProtoReflect.Descriptor instead.
func (*LogoutReply) Descriptor() ([]byte, []int) {
	return file_api_user_gw_v1_user_gw_proto_rawDescGZIP(), []int{5}
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_api_user_gw_v1_user_gw_proto_rawDescGZIP(), []int{6}
}

func (x *ChangePasswordRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChangePasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ChangeRegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeRegisterReply) Reset() {
	*x = ChangeRegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeRegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRegisterReply) ProtoMessage() {}

func (x *ChangeRegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_gw_v1_user_gw_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRegisterReply.ProtoReflect.Descriptor instead.
func (*ChangeRegisterReply) Descriptor() ([]byte, []int) {
	return file_api_user_gw_v1_user_gw_proto_rawDescGZIP(), []int{7}
}

var File_api_user_gw_v1_user_gw_proto protoreflect.FileDescriptor

var file_api_user_gw_v1_user_gw_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2f, 0x76, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x06, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1f, 0x0a, 0x0d,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x06, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x32, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x14, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0d,
	0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x55, 0x0a,
	0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x23, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x06, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xc1, 0x04, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x7f, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x10, 0x12, 0x06, 0xe6, 0xb3, 0xa8, 0xe5, 0x86, 0x8c, 0x1a, 0x06, 0xe6,
	0xb3, 0xa8, 0xe5, 0x86, 0x8c, 0x12, 0x73, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x10, 0x12, 0x06, 0xe7, 0x99, 0xbb, 0xe5,
	0xbd, 0x95, 0x1a, 0x06, 0xe7, 0x99, 0xbb, 0xe5, 0xbd, 0x95, 0x12, 0x89, 0x01, 0x0a, 0x06, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x43, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x3a, 0x01,
	0x2a, 0x92, 0x41, 0x22, 0x12, 0x06, 0xe7, 0x99, 0xbb, 0xe5, 0x87, 0xba, 0x1a, 0x06, 0xe7, 0x99,
	0xbb, 0xe5, 0x87, 0xba, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x12, 0xb6, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x58, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x01, 0x2a, 0x92, 0x41,
	0x2e, 0x12, 0x0c, 0xe4, 0xbf, 0xae, 0xe6, 0x94, 0xb9, 0xe5, 0xaf, 0x86, 0xe7, 0xa0, 0x81, 0x1a,
	0x0c, 0xe4, 0xbf, 0xae, 0xe6, 0x94, 0xb9, 0xe5, 0xaf, 0x86, 0xe7, 0xa0, 0x81, 0x62, 0x10, 0x0a,
	0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x42,
	0xbb, 0x01, 0x5a, 0x18, 0x65, 0x64, 0x75, 0x73, 0x79, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x92, 0x41, 0x9d, 0x01,
	0x12, 0x14, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x70, 0x69, 0x20, 0x64, 0x6f, 0x63,
	0x73, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73,
	0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x4d,
	0x0a, 0x4b, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x3d,
	0x08, 0x02, 0x12, 0x28, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x61, 0x20, 0x22, 0x42, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x2d, 0x6a, 0x77, 0x74, 0x2d, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x20, 0x74, 0x6f, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_user_gw_v1_user_gw_proto_rawDescOnce sync.Once
	file_api_user_gw_v1_user_gw_proto_rawDescData = file_api_user_gw_v1_user_gw_proto_rawDesc
)

func file_api_user_gw_v1_user_gw_proto_rawDescGZIP() []byte {
	file_api_user_gw_v1_user_gw_proto_rawDescOnce.Do(func() {
		file_api_user_gw_v1_user_gw_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_user_gw_v1_user_gw_proto_rawDescData)
	})
	return file_api_user_gw_v1_user_gw_proto_rawDescData
}

var file_api_user_gw_v1_user_gw_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_user_gw_v1_user_gw_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),       // 0: api.user_gw.v1.RegisterRequest
	(*RegisterReply)(nil),         // 1: api.user_gw.v1.RegisterReply
	(*LoginRequest)(nil),          // 2: api.user_gw.v1.LoginRequest
	(*LoginReply)(nil),            // 3: api.user_gw.v1.LoginReply
	(*LogoutRequest)(nil),         // 4: api.user_gw.v1.LogoutRequest
	(*LogoutReply)(nil),           // 5: api.user_gw.v1.LogoutReply
	(*ChangePasswordRequest)(nil), // 6: api.user_gw.v1.ChangePasswordRequest
	(*ChangeRegisterReply)(nil),   // 7: api.user_gw.v1.ChangeRegisterReply
}
var file_api_user_gw_v1_user_gw_proto_depIdxs = []int32{
	0, // 0: api.user_gw.v1.user.Register:input_type -> api.user_gw.v1.RegisterRequest
	2, // 1: api.user_gw.v1.user.Login:input_type -> api.user_gw.v1.LoginRequest
	4, // 2: api.user_gw.v1.user.Logout:input_type -> api.user_gw.v1.LogoutRequest
	6, // 3: api.user_gw.v1.user.ChangePassword:input_type -> api.user_gw.v1.ChangePasswordRequest
	1, // 4: api.user_gw.v1.user.Register:output_type -> api.user_gw.v1.RegisterReply
	3, // 5: api.user_gw.v1.user.Login:output_type -> api.user_gw.v1.LoginReply
	5, // 6: api.user_gw.v1.user.Logout:output_type -> api.user_gw.v1.LogoutReply
	7, // 7: api.user_gw.v1.user.ChangePassword:output_type -> api.user_gw.v1.ChangeRegisterReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_gw_v1_user_gw_proto_init() }
func file_api_user_gw_v1_user_gw_proto_init() {
	if File_api_user_gw_v1_user_gw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_user_gw_v1_user_gw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_api_user_gw_v1_user_gw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReply); i {
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
		file_api_user_gw_v1_user_gw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_api_user_gw_v1_user_gw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_api_user_gw_v1_user_gw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_api_user_gw_v1_user_gw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReply); i {
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
		file_api_user_gw_v1_user_gw_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_api_user_gw_v1_user_gw_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeRegisterReply); i {
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
			RawDescriptor: file_api_user_gw_v1_user_gw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_gw_v1_user_gw_proto_goTypes,
		DependencyIndexes: file_api_user_gw_v1_user_gw_proto_depIdxs,
		MessageInfos:      file_api_user_gw_v1_user_gw_proto_msgTypes,
	}.Build()
	File_api_user_gw_v1_user_gw_proto = out.File
	file_api_user_gw_v1_user_gw_proto_rawDesc = nil
	file_api_user_gw_v1_user_gw_proto_goTypes = nil
	file_api_user_gw_v1_user_gw_proto_depIdxs = nil
}
