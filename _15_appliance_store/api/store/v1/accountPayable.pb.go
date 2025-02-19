// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.2
// source: api/store/v1/accountPayable.proto

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

type CreateAccountPayableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SupplierID uint32 `protobuf:"varint,1,opt,name=supplierID,proto3" json:"supplierID"` // 供应商ID
	PurchaseID string `protobuf:"bytes,2,opt,name=purchaseID,proto3" json:"purchaseID"`  // 采购单号
	DueAmount  string `protobuf:"bytes,3,opt,name=dueAmount,proto3" json:"dueAmount"`    // 应付款金额
	PaidAmount string `protobuf:"bytes,4,opt,name=paidAmount,proto3" json:"paidAmount"`  // 已付款金额
	DueDate    string `protobuf:"bytes,5,opt,name=dueDate,proto3" json:"dueDate"`        // 应付款日期
}

func (x *CreateAccountPayableRequest) Reset() {
	*x = CreateAccountPayableRequest{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountPayableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountPayableRequest) ProtoMessage() {}

func (x *CreateAccountPayableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountPayableRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountPayableRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountPayableRequest) GetSupplierID() uint32 {
	if x != nil {
		return x.SupplierID
	}
	return 0
}

func (x *CreateAccountPayableRequest) GetPurchaseID() string {
	if x != nil {
		return x.PurchaseID
	}
	return ""
}

func (x *CreateAccountPayableRequest) GetDueAmount() string {
	if x != nil {
		return x.DueAmount
	}
	return ""
}

func (x *CreateAccountPayableRequest) GetPaidAmount() string {
	if x != nil {
		return x.PaidAmount
	}
	return ""
}

func (x *CreateAccountPayableRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

type CreateAccountPayableReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *CreateAccountPayableReply) Reset() {
	*x = CreateAccountPayableReply{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountPayableReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountPayableReply) ProtoMessage() {}

func (x *CreateAccountPayableReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountPayableReply.ProtoReflect.Descriptor instead.
func (*CreateAccountPayableReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountPayableReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAccountPayableByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`
}

func (x *DeleteAccountPayableByIDRequest) Reset() {
	*x = DeleteAccountPayableByIDRequest{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAccountPayableByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountPayableByIDRequest) ProtoMessage() {}

func (x *DeleteAccountPayableByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountPayableByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccountPayableByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteAccountPayableByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAccountPayableByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAccountPayableByIDReply) Reset() {
	*x = DeleteAccountPayableByIDReply{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAccountPayableByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountPayableByIDReply) ProtoMessage() {}

func (x *DeleteAccountPayableByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountPayableByIDReply.ProtoReflect.Descriptor instead.
func (*DeleteAccountPayableByIDReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{3}
}

type UpdateAccountPayableByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`        // 记录ID
	SupplierID uint32 `protobuf:"varint,2,opt,name=supplierID,proto3" json:"supplierID"` // 供应商ID
	PurchaseID string `protobuf:"bytes,3,opt,name=purchaseID,proto3" json:"purchaseID"`  // 采购单号
	DueAmount  string `protobuf:"bytes,4,opt,name=dueAmount,proto3" json:"dueAmount"`    // 应付款金额
	PaidAmount string `protobuf:"bytes,5,opt,name=paidAmount,proto3" json:"paidAmount"`  // 已付款金额
	DueDate    string `protobuf:"bytes,6,opt,name=dueDate,proto3" json:"dueDate"`        // 应付款日期
}

func (x *UpdateAccountPayableByIDRequest) Reset() {
	*x = UpdateAccountPayableByIDRequest{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAccountPayableByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountPayableByIDRequest) ProtoMessage() {}

func (x *UpdateAccountPayableByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountPayableByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountPayableByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAccountPayableByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAccountPayableByIDRequest) GetSupplierID() uint32 {
	if x != nil {
		return x.SupplierID
	}
	return 0
}

func (x *UpdateAccountPayableByIDRequest) GetPurchaseID() string {
	if x != nil {
		return x.PurchaseID
	}
	return ""
}

func (x *UpdateAccountPayableByIDRequest) GetDueAmount() string {
	if x != nil {
		return x.DueAmount
	}
	return ""
}

func (x *UpdateAccountPayableByIDRequest) GetPaidAmount() string {
	if x != nil {
		return x.PaidAmount
	}
	return ""
}

func (x *UpdateAccountPayableByIDRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

type UpdateAccountPayableByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAccountPayableByIDReply) Reset() {
	*x = UpdateAccountPayableByIDReply{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAccountPayableByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountPayableByIDReply) ProtoMessage() {}

func (x *UpdateAccountPayableByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountPayableByIDReply.ProtoReflect.Descriptor instead.
func (*UpdateAccountPayableByIDReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{5}
}

type AccountPayable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                 // 记录ID
	SupplierID uint32 `protobuf:"varint,2,opt,name=supplierID,proto3" json:"supplierID"` // 供应商ID
	PurchaseID string `protobuf:"bytes,3,opt,name=purchaseID,proto3" json:"purchaseID"`  // 采购单号
	DueAmount  string `protobuf:"bytes,4,opt,name=dueAmount,proto3" json:"dueAmount"`    // 应付款金额
	PaidAmount string `protobuf:"bytes,5,opt,name=paidAmount,proto3" json:"paidAmount"`  // 已付款金额
	DueDate    string `protobuf:"bytes,6,opt,name=dueDate,proto3" json:"dueDate"`        // 应付款日期
}

func (x *AccountPayable) Reset() {
	*x = AccountPayable{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountPayable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountPayable) ProtoMessage() {}

func (x *AccountPayable) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountPayable.ProtoReflect.Descriptor instead.
func (*AccountPayable) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{6}
}

func (x *AccountPayable) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountPayable) GetSupplierID() uint32 {
	if x != nil {
		return x.SupplierID
	}
	return 0
}

func (x *AccountPayable) GetPurchaseID() string {
	if x != nil {
		return x.PurchaseID
	}
	return ""
}

func (x *AccountPayable) GetDueAmount() string {
	if x != nil {
		return x.DueAmount
	}
	return ""
}

func (x *AccountPayable) GetPaidAmount() string {
	if x != nil {
		return x.PaidAmount
	}
	return ""
}

func (x *AccountPayable) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

type GetAccountPayableByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`
}

func (x *GetAccountPayableByIDRequest) Reset() {
	*x = GetAccountPayableByIDRequest{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountPayableByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountPayableByIDRequest) ProtoMessage() {}

func (x *GetAccountPayableByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountPayableByIDRequest.ProtoReflect.Descriptor instead.
func (*GetAccountPayableByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{7}
}

func (x *GetAccountPayableByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAccountPayableByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountPayable *AccountPayable `protobuf:"bytes,1,opt,name=accountPayable,proto3" json:"accountPayable"`
}

func (x *GetAccountPayableByIDReply) Reset() {
	*x = GetAccountPayableByIDReply{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountPayableByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountPayableByIDReply) ProtoMessage() {}

func (x *GetAccountPayableByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountPayableByIDReply.ProtoReflect.Descriptor instead.
func (*GetAccountPayableByIDReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{8}
}

func (x *GetAccountPayableByIDReply) GetAccountPayable() *AccountPayable {
	if x != nil {
		return x.AccountPayable
	}
	return nil
}

type ListAccountPayableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *types.Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (x *ListAccountPayableRequest) Reset() {
	*x = ListAccountPayableRequest{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccountPayableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountPayableRequest) ProtoMessage() {}

func (x *ListAccountPayableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountPayableRequest.ProtoReflect.Descriptor instead.
func (*ListAccountPayableRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{9}
}

func (x *ListAccountPayableRequest) GetParams() *types.Params {
	if x != nil {
		return x.Params
	}
	return nil
}

type ListAccountPayableReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total           int64             `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	AccountPayables []*AccountPayable `protobuf:"bytes,2,rep,name=accountPayables,proto3" json:"accountPayables"`
}

func (x *ListAccountPayableReply) Reset() {
	*x = ListAccountPayableReply{}
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccountPayableReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountPayableReply) ProtoMessage() {}

func (x *ListAccountPayableReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_accountPayable_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountPayableReply.ProtoReflect.Descriptor instead.
func (*ListAccountPayableReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_accountPayable_proto_rawDescGZIP(), []int{10}
}

func (x *ListAccountPayableReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListAccountPayableReply) GetAccountPayables() []*AccountPayable {
	if x != nil {
		return x.AccountPayables
	}
	return nil
}

var File_api_store_v1_accountPayable_proto protoreflect.FileDescriptor

var file_api_store_v1_accountPayable_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72,
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
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x75, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x75, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x19,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x1f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x14, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0xdf, 0x01, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x14, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x9a, 0x84, 0x9e, 0x03,
	0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x64, 0x75, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x75, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x61, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x61, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xb8, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x75, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x75,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x69, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x69,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x44, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x14, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22,
	0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x19, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x22, 0x77, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x46, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x32, 0xbb, 0x05, 0x0a,
	0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x7f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x8d, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61,
	0x62, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62,
	0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x90, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61,
	0x62, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62,
	0x6c, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x1a, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x84, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61,
	0x79, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61,
	0x79, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0xb4, 0x01, 0x92, 0x41, 0x99,
	0x01, 0x12, 0x15, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x70, 0x69, 0x20, 0x64,
	0x6f, 0x63, 0x73, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68,
	0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x5a, 0x48, 0x0a, 0x46, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x38, 0x08, 0x02, 0x12, 0x23, 0x54, 0x79, 0x70, 0x65, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65,
	0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x2d, 0x6a, 0x77, 0x74, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x20, 0x74, 0x6f, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x5a, 0x15, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_store_v1_accountPayable_proto_rawDescOnce sync.Once
	file_api_store_v1_accountPayable_proto_rawDescData = file_api_store_v1_accountPayable_proto_rawDesc
)

func file_api_store_v1_accountPayable_proto_rawDescGZIP() []byte {
	file_api_store_v1_accountPayable_proto_rawDescOnce.Do(func() {
		file_api_store_v1_accountPayable_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_store_v1_accountPayable_proto_rawDescData)
	})
	return file_api_store_v1_accountPayable_proto_rawDescData
}

var file_api_store_v1_accountPayable_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_store_v1_accountPayable_proto_goTypes = []any{
	(*CreateAccountPayableRequest)(nil),     // 0: api.store.v1.CreateAccountPayableRequest
	(*CreateAccountPayableReply)(nil),       // 1: api.store.v1.CreateAccountPayableReply
	(*DeleteAccountPayableByIDRequest)(nil), // 2: api.store.v1.DeleteAccountPayableByIDRequest
	(*DeleteAccountPayableByIDReply)(nil),   // 3: api.store.v1.DeleteAccountPayableByIDReply
	(*UpdateAccountPayableByIDRequest)(nil), // 4: api.store.v1.UpdateAccountPayableByIDRequest
	(*UpdateAccountPayableByIDReply)(nil),   // 5: api.store.v1.UpdateAccountPayableByIDReply
	(*AccountPayable)(nil),                  // 6: api.store.v1.AccountPayable
	(*GetAccountPayableByIDRequest)(nil),    // 7: api.store.v1.GetAccountPayableByIDRequest
	(*GetAccountPayableByIDReply)(nil),      // 8: api.store.v1.GetAccountPayableByIDReply
	(*ListAccountPayableRequest)(nil),       // 9: api.store.v1.ListAccountPayableRequest
	(*ListAccountPayableReply)(nil),         // 10: api.store.v1.ListAccountPayableReply
	(*types.Params)(nil),                    // 11: api.types.Params
}
var file_api_store_v1_accountPayable_proto_depIdxs = []int32{
	6,  // 0: api.store.v1.GetAccountPayableByIDReply.accountPayable:type_name -> api.store.v1.AccountPayable
	11, // 1: api.store.v1.ListAccountPayableRequest.params:type_name -> api.types.Params
	6,  // 2: api.store.v1.ListAccountPayableReply.accountPayables:type_name -> api.store.v1.AccountPayable
	0,  // 3: api.store.v1.accountPayable.Create:input_type -> api.store.v1.CreateAccountPayableRequest
	2,  // 4: api.store.v1.accountPayable.DeleteByID:input_type -> api.store.v1.DeleteAccountPayableByIDRequest
	4,  // 5: api.store.v1.accountPayable.UpdateByID:input_type -> api.store.v1.UpdateAccountPayableByIDRequest
	7,  // 6: api.store.v1.accountPayable.GetByID:input_type -> api.store.v1.GetAccountPayableByIDRequest
	9,  // 7: api.store.v1.accountPayable.List:input_type -> api.store.v1.ListAccountPayableRequest
	1,  // 8: api.store.v1.accountPayable.Create:output_type -> api.store.v1.CreateAccountPayableReply
	3,  // 9: api.store.v1.accountPayable.DeleteByID:output_type -> api.store.v1.DeleteAccountPayableByIDReply
	5,  // 10: api.store.v1.accountPayable.UpdateByID:output_type -> api.store.v1.UpdateAccountPayableByIDReply
	8,  // 11: api.store.v1.accountPayable.GetByID:output_type -> api.store.v1.GetAccountPayableByIDReply
	10, // 12: api.store.v1.accountPayable.List:output_type -> api.store.v1.ListAccountPayableReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_store_v1_accountPayable_proto_init() }
func file_api_store_v1_accountPayable_proto_init() {
	if File_api_store_v1_accountPayable_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_store_v1_accountPayable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_store_v1_accountPayable_proto_goTypes,
		DependencyIndexes: file_api_store_v1_accountPayable_proto_depIdxs,
		MessageInfos:      file_api_store_v1_accountPayable_proto_msgTypes,
	}.Build()
	File_api_store_v1_accountPayable_proto = out.File
	file_api_store_v1_accountPayable_proto_rawDesc = nil
	file_api_store_v1_accountPayable_proto_goTypes = nil
	file_api_store_v1_accountPayable_proto_depIdxs = nil
}
