// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/community/v1/like.proto

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

type LikeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	UserId    uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId"`
	ObjType   int32  `protobuf:"varint,3,opt,name=objType,proto3" json:"objType"`
	ObjId     uint64 `protobuf:"varint,4,opt,name=objId,proto3" json:"objId"`
	Status    int32  `protobuf:"varint,5,opt,name=status,proto3" json:"status"`
	CreatedAt int64  `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt int64  `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt"`
}

func (x *LikeInfo) Reset() {
	*x = LikeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_community_v1_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeInfo) ProtoMessage() {}

func (x *LikeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_community_v1_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeInfo.ProtoReflect.Descriptor instead.
func (*LikeInfo) Descriptor() ([]byte, []int) {
	return file_api_community_v1_like_proto_rawDescGZIP(), []int{0}
}

func (x *LikeInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LikeInfo) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeInfo) GetObjType() int32 {
	if x != nil {
		return x.ObjType
	}
	return 0
}

func (x *LikeInfo) GetObjId() uint64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *LikeInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *LikeInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *LikeInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId"`
	ObjType int32  `protobuf:"varint,2,opt,name=objType,proto3" json:"objType"`
	ObjId   uint64 `protobuf:"varint,3,opt,name=objId,proto3" json:"objId"`
}

func (x *CreateLikeRequest) Reset() {
	*x = CreateLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_community_v1_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLikeRequest) ProtoMessage() {}

func (x *CreateLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_community_v1_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLikeRequest.ProtoReflect.Descriptor instead.
func (*CreateLikeRequest) Descriptor() ([]byte, []int) {
	return file_api_community_v1_like_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLikeRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateLikeRequest) GetObjType() int32 {
	if x != nil {
		return x.ObjType
	}
	return 0
}

func (x *CreateLikeRequest) GetObjId() uint64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

type CreateLikeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateLikeReply) Reset() {
	*x = CreateLikeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_community_v1_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLikeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLikeReply) ProtoMessage() {}

func (x *CreateLikeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_community_v1_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLikeReply.ProtoReflect.Descriptor instead.
func (*CreateLikeReply) Descriptor() ([]byte, []int) {
	return file_api_community_v1_like_proto_rawDescGZIP(), []int{2}
}

type DeleteLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId"`
	ObjType int32  `protobuf:"varint,2,opt,name=objType,proto3" json:"objType"`
	ObjId   uint64 `protobuf:"varint,3,opt,name=objId,proto3" json:"objId"`
}

func (x *DeleteLikeRequest) Reset() {
	*x = DeleteLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_community_v1_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLikeRequest) ProtoMessage() {}

func (x *DeleteLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_community_v1_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLikeRequest.ProtoReflect.Descriptor instead.
func (*DeleteLikeRequest) Descriptor() ([]byte, []int) {
	return file_api_community_v1_like_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteLikeRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DeleteLikeRequest) GetObjType() int32 {
	if x != nil {
		return x.ObjType
	}
	return 0
}

func (x *DeleteLikeRequest) GetObjId() uint64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

type DeleteLikeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteLikeReply) Reset() {
	*x = DeleteLikeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_community_v1_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLikeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLikeReply) ProtoMessage() {}

func (x *DeleteLikeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_community_v1_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLikeReply.ProtoReflect.Descriptor instead.
func (*DeleteLikeReply) Descriptor() ([]byte, []int) {
	return file_api_community_v1_like_proto_rawDescGZIP(), []int{4}
}

type ListPostLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId uint64 `protobuf:"varint,2,opt,name=postId,proto3" json:"postId"`
	Page   int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page"`
	Limit  int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit"`
}

func (x *ListPostLikeRequest) Reset() {
	*x = ListPostLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_community_v1_like_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostLikeRequest) ProtoMessage() {}

func (x *ListPostLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_community_v1_like_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostLikeRequest.ProtoReflect.Descriptor instead.
func (*ListPostLikeRequest) Descriptor() ([]byte, []int) {
	return file_api_community_v1_like_proto_rawDescGZIP(), []int{5}
}

func (x *ListPostLikeRequest) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *ListPostLikeRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPostLikeRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListPostLikeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Likes []*LikeInfo `protobuf:"bytes,1,rep,name=likes,proto3" json:"likes"`
	Total int64       `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *ListPostLikeReply) Reset() {
	*x = ListPostLikeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_community_v1_like_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostLikeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostLikeReply) ProtoMessage() {}

func (x *ListPostLikeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_community_v1_like_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostLikeReply.ProtoReflect.Descriptor instead.
func (*ListPostLikeReply) Descriptor() ([]byte, []int) {
	return file_api_community_v1_like_proto_rawDescGZIP(), []int{6}
}

func (x *ListPostLikeReply) GetLikes() []*LikeInfo {
	if x != nil {
		return x.Likes
	}
	return nil
}

func (x *ListPostLikeReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListCommentLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId uint64 `protobuf:"varint,2,opt,name=commentId,proto3" json:"commentId"`
	Page      int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page"`
	Limit     int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit"`
}

func (x *ListCommentLikeRequest) Reset() {
	*x = ListCommentLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_community_v1_like_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentLikeRequest) ProtoMessage() {}

func (x *ListCommentLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_community_v1_like_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentLikeRequest.ProtoReflect.Descriptor instead.
func (*ListCommentLikeRequest) Descriptor() ([]byte, []int) {
	return file_api_community_v1_like_proto_rawDescGZIP(), []int{7}
}

func (x *ListCommentLikeRequest) GetCommentId() uint64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *ListCommentLikeRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListCommentLikeRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListCommentLikeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Likes []*LikeInfo `protobuf:"bytes,1,rep,name=likes,proto3" json:"likes"`
	Total int64       `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *ListCommentLikeReply) Reset() {
	*x = ListCommentLikeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_community_v1_like_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentLikeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentLikeReply) ProtoMessage() {}

func (x *ListCommentLikeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_community_v1_like_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentLikeReply.ProtoReflect.Descriptor instead.
func (*ListCommentLikeReply) Descriptor() ([]byte, []int) {
	return file_api_community_v1_like_proto_rawDescGZIP(), []int{8}
}

func (x *ListCommentLikeReply) GetLikes() []*LikeInfo {
	if x != nil {
		return x.Likes
	}
	return nil
}

func (x *ListCommentLikeReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_api_community_v1_like_proto protoreflect.FileDescriptor

var file_api_community_v1_like_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x6b, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6f,
	0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x62,
	0x6a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x78, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x02, 0x28,
	0x01, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x6f, 0x62,
	0x6a, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x28, 0x01, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x78, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x07, 0x6f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x02, 0x28, 0x01, 0x52, 0x07,
	0x6f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52,
	0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x74, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x1a, 0x04, 0x18, 0x64, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x5b, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7d, 0x0a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x28, 0x01, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x28, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04,
	0x18, 0x64, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5e, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xd8, 0x05, 0x0a, 0x0b,
	0x4c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x6c, 0x69, 0x6b, 0x65, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x22, 0x12, 0x06, 0xe7, 0x82, 0xb9, 0xe8,
	0xb5, 0x9e, 0x1a, 0x06, 0xe7, 0x82, 0xb9, 0xe8, 0xb5, 0x9e, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x12, 0xa1, 0x01, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x4f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x92,
	0x41, 0x2e, 0x12, 0x0c, 0xe5, 0x8f, 0x96, 0xe6, 0xb6, 0x88, 0xe7, 0x82, 0xb9, 0xe8, 0xb5, 0x9e,
	0x1a, 0x0c, 0xe5, 0x8f, 0x96, 0xe6, 0xb6, 0x88, 0xe7, 0x82, 0xb9, 0xe8, 0xb5, 0x9e, 0x62, 0x10,
	0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00,
	0x12, 0xc2, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x6a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1b, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x46, 0x12,
	0x18, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe8, 0xb4, 0xb4, 0xe5, 0xad, 0x90, 0xe7, 0x82, 0xb9,
	0xe8, 0xb5, 0x9e, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x1a, 0x18, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f,
	0x96, 0xe8, 0xb4, 0xb4, 0xe5, 0xad, 0x90, 0xe7, 0x82, 0xb9, 0xe8, 0xb5, 0x9e, 0xe5, 0x88, 0x97,
	0xe8, 0xa1, 0xa8, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x00, 0x12, 0xce, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22,
	0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x46,
	0x12, 0x18, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe8, 0xaf, 0x84, 0xe8, 0xae, 0xba, 0xe7, 0x82,
	0xb9, 0xe8, 0xb5, 0x9e, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0x1a, 0x18, 0xe8, 0x8e, 0xb7, 0xe5,
	0x8f, 0x96, 0xe8, 0xaf, 0x84, 0xe8, 0xae, 0xba, 0xe7, 0x82, 0xb9, 0xe8, 0xb5, 0x9e, 0xe5, 0x88,
	0x97, 0xe8, 0xa1, 0xa8, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x42, 0xc3, 0x01, 0x5a, 0x1d, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x92, 0x41, 0xa0, 0x01, 0x12, 0x17, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x70, 0x69, 0x20, 0x64, 0x6f, 0x63, 0x73,
	0x32, 0x03, 0x32, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74,
	0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x4d, 0x0a,
	0x4b, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x3d, 0x08,
	0x02, 0x12, 0x28, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x61, 0x20, 0x22, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x2d, 0x6a, 0x77, 0x74, 0x2d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x20, 0x74, 0x6f, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_community_v1_like_proto_rawDescOnce sync.Once
	file_api_community_v1_like_proto_rawDescData = file_api_community_v1_like_proto_rawDesc
)

func file_api_community_v1_like_proto_rawDescGZIP() []byte {
	file_api_community_v1_like_proto_rawDescOnce.Do(func() {
		file_api_community_v1_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_community_v1_like_proto_rawDescData)
	})
	return file_api_community_v1_like_proto_rawDescData
}

var file_api_community_v1_like_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_community_v1_like_proto_goTypes = []interface{}{
	(*LikeInfo)(nil),               // 0: api.community.v1.LikeInfo
	(*CreateLikeRequest)(nil),      // 1: api.community.v1.CreateLikeRequest
	(*CreateLikeReply)(nil),        // 2: api.community.v1.CreateLikeReply
	(*DeleteLikeRequest)(nil),      // 3: api.community.v1.DeleteLikeRequest
	(*DeleteLikeReply)(nil),        // 4: api.community.v1.DeleteLikeReply
	(*ListPostLikeRequest)(nil),    // 5: api.community.v1.ListPostLikeRequest
	(*ListPostLikeReply)(nil),      // 6: api.community.v1.ListPostLikeReply
	(*ListCommentLikeRequest)(nil), // 7: api.community.v1.ListCommentLikeRequest
	(*ListCommentLikeReply)(nil),   // 8: api.community.v1.ListCommentLikeReply
}
var file_api_community_v1_like_proto_depIdxs = []int32{
	0, // 0: api.community.v1.ListPostLikeReply.likes:type_name -> api.community.v1.LikeInfo
	0, // 1: api.community.v1.ListCommentLikeReply.likes:type_name -> api.community.v1.LikeInfo
	1, // 2: api.community.v1.LikeService.Create:input_type -> api.community.v1.CreateLikeRequest
	3, // 3: api.community.v1.LikeService.Delete:input_type -> api.community.v1.DeleteLikeRequest
	5, // 4: api.community.v1.LikeService.ListPost:input_type -> api.community.v1.ListPostLikeRequest
	7, // 5: api.community.v1.LikeService.ListComment:input_type -> api.community.v1.ListCommentLikeRequest
	2, // 6: api.community.v1.LikeService.Create:output_type -> api.community.v1.CreateLikeReply
	4, // 7: api.community.v1.LikeService.Delete:output_type -> api.community.v1.DeleteLikeReply
	6, // 8: api.community.v1.LikeService.ListPost:output_type -> api.community.v1.ListPostLikeReply
	8, // 9: api.community.v1.LikeService.ListComment:output_type -> api.community.v1.ListCommentLikeReply
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_community_v1_like_proto_init() }
func file_api_community_v1_like_proto_init() {
	if File_api_community_v1_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_community_v1_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeInfo); i {
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
		file_api_community_v1_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLikeRequest); i {
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
		file_api_community_v1_like_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLikeReply); i {
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
		file_api_community_v1_like_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLikeRequest); i {
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
		file_api_community_v1_like_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLikeReply); i {
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
		file_api_community_v1_like_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostLikeRequest); i {
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
		file_api_community_v1_like_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostLikeReply); i {
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
		file_api_community_v1_like_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentLikeRequest); i {
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
		file_api_community_v1_like_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentLikeReply); i {
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
			RawDescriptor: file_api_community_v1_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_community_v1_like_proto_goTypes,
		DependencyIndexes: file_api_community_v1_like_proto_depIdxs,
		MessageInfos:      file_api_community_v1_like_proto_msgTypes,
	}.Build()
	File_api_community_v1_like_proto = out.File
	file_api_community_v1_like_proto_rawDesc = nil
	file_api_community_v1_like_proto_goTypes = nil
	file_api_community_v1_like_proto_depIdxs = nil
}
