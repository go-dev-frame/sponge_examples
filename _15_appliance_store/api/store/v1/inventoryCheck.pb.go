// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.2
// source: api/store/v1/inventoryCheck.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	types "store/api/types"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateInventoryCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreID    uint32 `protobuf:"varint,1,opt,name=storeID,proto3" json:"storeID"`       // 门店ID
	OperatorID uint32 `protobuf:"varint,2,opt,name=operatorID,proto3" json:"operatorID"` // 操作人ID
	CheckTime  string `protobuf:"bytes,3,opt,name=checkTime,proto3" json:"checkTime"`    // 盘点时间
	TotalDiff  int32  `protobuf:"varint,4,opt,name=totalDiff,proto3" json:"totalDiff"`   // 总差异数量
}

func (x *CreateInventoryCheckRequest) Reset() {
	*x = CreateInventoryCheckRequest{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateInventoryCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInventoryCheckRequest) ProtoMessage() {}

func (x *CreateInventoryCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInventoryCheckRequest.ProtoReflect.Descriptor instead.
func (*CreateInventoryCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInventoryCheckRequest) GetStoreID() uint32 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

func (x *CreateInventoryCheckRequest) GetOperatorID() uint32 {
	if x != nil {
		return x.OperatorID
	}
	return 0
}

func (x *CreateInventoryCheckRequest) GetCheckTime() string {
	if x != nil {
		return x.CheckTime
	}
	return ""
}

func (x *CreateInventoryCheckRequest) GetTotalDiff() int32 {
	if x != nil {
		return x.TotalDiff
	}
	return 0
}

type CreateInventoryCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *CreateInventoryCheckReply) Reset() {
	*x = CreateInventoryCheckReply{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateInventoryCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInventoryCheckReply) ProtoMessage() {}

func (x *CreateInventoryCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInventoryCheckReply.ProtoReflect.Descriptor instead.
func (*CreateInventoryCheckReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInventoryCheckReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteInventoryCheckByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" uri:"id"`
}

func (x *DeleteInventoryCheckByIDRequest) Reset() {
	*x = DeleteInventoryCheckByIDRequest{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteInventoryCheckByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInventoryCheckByIDRequest) ProtoMessage() {}

func (x *DeleteInventoryCheckByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInventoryCheckByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteInventoryCheckByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteInventoryCheckByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteInventoryCheckByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteInventoryCheckByIDReply) Reset() {
	*x = DeleteInventoryCheckByIDReply{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteInventoryCheckByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInventoryCheckByIDReply) ProtoMessage() {}

func (x *DeleteInventoryCheckByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInventoryCheckByIDReply.ProtoReflect.Descriptor instead.
func (*DeleteInventoryCheckByIDReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{3}
}

type UpdateInventoryCheckByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" uri:"id"`         // 盘点单号
	StoreID    uint32 `protobuf:"varint,2,opt,name=storeID,proto3" json:"storeID"`       // 门店ID
	OperatorID uint32 `protobuf:"varint,3,opt,name=operatorID,proto3" json:"operatorID"` // 操作人ID
	CheckTime  string `protobuf:"bytes,4,opt,name=checkTime,proto3" json:"checkTime"`    // 盘点时间
	TotalDiff  int32  `protobuf:"varint,5,opt,name=totalDiff,proto3" json:"totalDiff"`   // 总差异数量
}

func (x *UpdateInventoryCheckByIDRequest) Reset() {
	*x = UpdateInventoryCheckByIDRequest{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateInventoryCheckByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInventoryCheckByIDRequest) ProtoMessage() {}

func (x *UpdateInventoryCheckByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInventoryCheckByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateInventoryCheckByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateInventoryCheckByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateInventoryCheckByIDRequest) GetStoreID() uint32 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

func (x *UpdateInventoryCheckByIDRequest) GetOperatorID() uint32 {
	if x != nil {
		return x.OperatorID
	}
	return 0
}

func (x *UpdateInventoryCheckByIDRequest) GetCheckTime() string {
	if x != nil {
		return x.CheckTime
	}
	return ""
}

func (x *UpdateInventoryCheckByIDRequest) GetTotalDiff() int32 {
	if x != nil {
		return x.TotalDiff
	}
	return 0
}

type UpdateInventoryCheckByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateInventoryCheckByIDReply) Reset() {
	*x = UpdateInventoryCheckByIDReply{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateInventoryCheckByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInventoryCheckByIDReply) ProtoMessage() {}

func (x *UpdateInventoryCheckByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInventoryCheckByIDReply.ProtoReflect.Descriptor instead.
func (*UpdateInventoryCheckByIDReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{5}
}

type InventoryCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`                  // 盘点单号
	StoreID    uint32 `protobuf:"varint,2,opt,name=storeID,proto3" json:"storeID"`       // 门店ID
	OperatorID uint32 `protobuf:"varint,3,opt,name=operatorID,proto3" json:"operatorID"` // 操作人ID
	CheckTime  string `protobuf:"bytes,4,opt,name=checkTime,proto3" json:"checkTime"`    // 盘点时间
	TotalDiff  int32  `protobuf:"varint,5,opt,name=totalDiff,proto3" json:"totalDiff"`   // 总差异数量
}

func (x *InventoryCheck) Reset() {
	*x = InventoryCheck{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InventoryCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryCheck) ProtoMessage() {}

func (x *InventoryCheck) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryCheck.ProtoReflect.Descriptor instead.
func (*InventoryCheck) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{6}
}

func (x *InventoryCheck) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InventoryCheck) GetStoreID() uint32 {
	if x != nil {
		return x.StoreID
	}
	return 0
}

func (x *InventoryCheck) GetOperatorID() uint32 {
	if x != nil {
		return x.OperatorID
	}
	return 0
}

func (x *InventoryCheck) GetCheckTime() string {
	if x != nil {
		return x.CheckTime
	}
	return ""
}

func (x *InventoryCheck) GetTotalDiff() int32 {
	if x != nil {
		return x.TotalDiff
	}
	return 0
}

type GetInventoryCheckByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" uri:"id"`
}

func (x *GetInventoryCheckByIDRequest) Reset() {
	*x = GetInventoryCheckByIDRequest{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInventoryCheckByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInventoryCheckByIDRequest) ProtoMessage() {}

func (x *GetInventoryCheckByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInventoryCheckByIDRequest.ProtoReflect.Descriptor instead.
func (*GetInventoryCheckByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{7}
}

func (x *GetInventoryCheckByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetInventoryCheckByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InventoryCheck *InventoryCheck `protobuf:"bytes,1,opt,name=inventoryCheck,proto3" json:"inventoryCheck"`
}

func (x *GetInventoryCheckByIDReply) Reset() {
	*x = GetInventoryCheckByIDReply{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInventoryCheckByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInventoryCheckByIDReply) ProtoMessage() {}

func (x *GetInventoryCheckByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInventoryCheckByIDReply.ProtoReflect.Descriptor instead.
func (*GetInventoryCheckByIDReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{8}
}

func (x *GetInventoryCheckByIDReply) GetInventoryCheck() *InventoryCheck {
	if x != nil {
		return x.InventoryCheck
	}
	return nil
}

type ListInventoryCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *types.Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (x *ListInventoryCheckRequest) Reset() {
	*x = ListInventoryCheckRequest{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInventoryCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInventoryCheckRequest) ProtoMessage() {}

func (x *ListInventoryCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInventoryCheckRequest.ProtoReflect.Descriptor instead.
func (*ListInventoryCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{9}
}

func (x *ListInventoryCheckRequest) GetParams() *types.Params {
	if x != nil {
		return x.Params
	}
	return nil
}

type ListInventoryCheckReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total           int64             `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	InventoryChecks []*InventoryCheck `protobuf:"bytes,2,rep,name=inventoryChecks,proto3" json:"inventoryChecks"`
}

func (x *ListInventoryCheckReply) Reset() {
	*x = ListInventoryCheckReply{}
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInventoryCheckReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInventoryCheckReply) ProtoMessage() {}

func (x *ListInventoryCheckReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_inventoryCheck_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInventoryCheckReply.ProtoReflect.Descriptor instead.
func (*ListInventoryCheckReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_inventoryCheck_proto_rawDescGZIP(), []int{10}
}

func (x *ListInventoryCheckReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListInventoryCheckReply) GetInventoryChecks() []*InventoryCheck {
	if x != nil {
		return x.InventoryChecks
	}
	return nil
}

var File_api_store_v1_inventoryCheck_proto protoreflect.FileDescriptor

var file_api_store_v1_inventoryCheck_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x22, 0x2b, 0x0a, 0x19, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x1f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x9a,
	0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0xbd, 0x01, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x14, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75,
	0x72, 0x69, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x66, 0x66,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x66,
	0x66, 0x22, 0x1f, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x96, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x66, 0x66, 0x22, 0x44, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x62, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x44, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x46, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x77, 0x0a,
	0x17, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x46,
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x32, 0xbb, 0x05, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x7f, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a,
	0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x8d, 0x01, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x90, 0x01, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a,
	0x1a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x84, 0x01,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x42, 0xb4, 0x01, 0x92, 0x41, 0x99, 0x01, 0x12, 0x15, 0x0a, 0x0e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x70, 0x69, 0x20, 0x64, 0x6f, 0x63, 0x73, 0x32, 0x03, 0x32,
	0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30,
	0x38, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x48, 0x0a, 0x46, 0x0a, 0x0a,
	0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x38, 0x08, 0x02, 0x12, 0x23,
	0x54, 0x79, 0x70, 0x65, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x79, 0x6f, 0x75, 0x72,
	0x2d, 0x6a, 0x77, 0x74, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x02, 0x5a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_store_v1_inventoryCheck_proto_rawDescOnce sync.Once
	file_api_store_v1_inventoryCheck_proto_rawDescData = file_api_store_v1_inventoryCheck_proto_rawDesc
)

func file_api_store_v1_inventoryCheck_proto_rawDescGZIP() []byte {
	file_api_store_v1_inventoryCheck_proto_rawDescOnce.Do(func() {
		file_api_store_v1_inventoryCheck_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_store_v1_inventoryCheck_proto_rawDescData)
	})
	return file_api_store_v1_inventoryCheck_proto_rawDescData
}

var file_api_store_v1_inventoryCheck_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_store_v1_inventoryCheck_proto_goTypes = []any{
	(*CreateInventoryCheckRequest)(nil),     // 0: api.store.v1.CreateInventoryCheckRequest
	(*CreateInventoryCheckReply)(nil),       // 1: api.store.v1.CreateInventoryCheckReply
	(*DeleteInventoryCheckByIDRequest)(nil), // 2: api.store.v1.DeleteInventoryCheckByIDRequest
	(*DeleteInventoryCheckByIDReply)(nil),   // 3: api.store.v1.DeleteInventoryCheckByIDReply
	(*UpdateInventoryCheckByIDRequest)(nil), // 4: api.store.v1.UpdateInventoryCheckByIDRequest
	(*UpdateInventoryCheckByIDReply)(nil),   // 5: api.store.v1.UpdateInventoryCheckByIDReply
	(*InventoryCheck)(nil),                  // 6: api.store.v1.InventoryCheck
	(*GetInventoryCheckByIDRequest)(nil),    // 7: api.store.v1.GetInventoryCheckByIDRequest
	(*GetInventoryCheckByIDReply)(nil),      // 8: api.store.v1.GetInventoryCheckByIDReply
	(*ListInventoryCheckRequest)(nil),       // 9: api.store.v1.ListInventoryCheckRequest
	(*ListInventoryCheckReply)(nil),         // 10: api.store.v1.ListInventoryCheckReply
	(*types.Params)(nil),                    // 11: api.types.Params
}
var file_api_store_v1_inventoryCheck_proto_depIdxs = []int32{
	6,  // 0: api.store.v1.GetInventoryCheckByIDReply.inventoryCheck:type_name -> api.store.v1.InventoryCheck
	11, // 1: api.store.v1.ListInventoryCheckRequest.params:type_name -> api.types.Params
	6,  // 2: api.store.v1.ListInventoryCheckReply.inventoryChecks:type_name -> api.store.v1.InventoryCheck
	0,  // 3: api.store.v1.inventoryCheck.Create:input_type -> api.store.v1.CreateInventoryCheckRequest
	2,  // 4: api.store.v1.inventoryCheck.DeleteByID:input_type -> api.store.v1.DeleteInventoryCheckByIDRequest
	4,  // 5: api.store.v1.inventoryCheck.UpdateByID:input_type -> api.store.v1.UpdateInventoryCheckByIDRequest
	7,  // 6: api.store.v1.inventoryCheck.GetByID:input_type -> api.store.v1.GetInventoryCheckByIDRequest
	9,  // 7: api.store.v1.inventoryCheck.List:input_type -> api.store.v1.ListInventoryCheckRequest
	1,  // 8: api.store.v1.inventoryCheck.Create:output_type -> api.store.v1.CreateInventoryCheckReply
	3,  // 9: api.store.v1.inventoryCheck.DeleteByID:output_type -> api.store.v1.DeleteInventoryCheckByIDReply
	5,  // 10: api.store.v1.inventoryCheck.UpdateByID:output_type -> api.store.v1.UpdateInventoryCheckByIDReply
	8,  // 11: api.store.v1.inventoryCheck.GetByID:output_type -> api.store.v1.GetInventoryCheckByIDReply
	10, // 12: api.store.v1.inventoryCheck.List:output_type -> api.store.v1.ListInventoryCheckReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_store_v1_inventoryCheck_proto_init() }
func file_api_store_v1_inventoryCheck_proto_init() {
	if File_api_store_v1_inventoryCheck_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_store_v1_inventoryCheck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_store_v1_inventoryCheck_proto_goTypes,
		DependencyIndexes: file_api_store_v1_inventoryCheck_proto_depIdxs,
		MessageInfos:      file_api_store_v1_inventoryCheck_proto_msgTypes,
	}.Build()
	File_api_store_v1_inventoryCheck_proto = out.File
	file_api_store_v1_inventoryCheck_proto_rawDesc = nil
	file_api_store_v1_inventoryCheck_proto_goTypes = nil
	file_api_store_v1_inventoryCheck_proto_depIdxs = nil
}
