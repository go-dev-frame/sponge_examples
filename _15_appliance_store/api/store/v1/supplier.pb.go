// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.2
// source: api/store/v1/supplier.proto

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

type CreateSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`                   // 供应商名称
	ContactPerson string `protobuf:"bytes,2,opt,name=contactPerson,proto3" json:"contactPerson"` // 联系人
	Phone         string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone"`                 // 联系电话
	PaymentTerms  string `protobuf:"bytes,4,opt,name=paymentTerms,proto3" json:"paymentTerms"`   // 结算条款
}

func (x *CreateSupplierRequest) Reset() {
	*x = CreateSupplierRequest{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupplierRequest) ProtoMessage() {}

func (x *CreateSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupplierRequest.ProtoReflect.Descriptor instead.
func (*CreateSupplierRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSupplierRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSupplierRequest) GetContactPerson() string {
	if x != nil {
		return x.ContactPerson
	}
	return ""
}

func (x *CreateSupplierRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateSupplierRequest) GetPaymentTerms() string {
	if x != nil {
		return x.PaymentTerms
	}
	return ""
}

type CreateSupplierReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *CreateSupplierReply) Reset() {
	*x = CreateSupplierReply{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSupplierReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSupplierReply) ProtoMessage() {}

func (x *CreateSupplierReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSupplierReply.ProtoReflect.Descriptor instead.
func (*CreateSupplierReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSupplierReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteSupplierByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`
}

func (x *DeleteSupplierByIDRequest) Reset() {
	*x = DeleteSupplierByIDRequest{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSupplierByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplierByIDRequest) ProtoMessage() {}

func (x *DeleteSupplierByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplierByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteSupplierByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteSupplierByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteSupplierByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSupplierByIDReply) Reset() {
	*x = DeleteSupplierByIDReply{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSupplierByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSupplierByIDReply) ProtoMessage() {}

func (x *DeleteSupplierByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSupplierByIDReply.ProtoReflect.Descriptor instead.
func (*DeleteSupplierByIDReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{3}
}

type UpdateSupplierByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`             // 供应商ID
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                   // 供应商名称
	ContactPerson string `protobuf:"bytes,3,opt,name=contactPerson,proto3" json:"contactPerson"` // 联系人
	Phone         string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone"`                 // 联系电话
	PaymentTerms  string `protobuf:"bytes,5,opt,name=paymentTerms,proto3" json:"paymentTerms"`   // 结算条款
}

func (x *UpdateSupplierByIDRequest) Reset() {
	*x = UpdateSupplierByIDRequest{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSupplierByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupplierByIDRequest) ProtoMessage() {}

func (x *UpdateSupplierByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupplierByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateSupplierByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSupplierByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateSupplierByIDRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSupplierByIDRequest) GetContactPerson() string {
	if x != nil {
		return x.ContactPerson
	}
	return ""
}

func (x *UpdateSupplierByIDRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateSupplierByIDRequest) GetPaymentTerms() string {
	if x != nil {
		return x.PaymentTerms
	}
	return ""
}

type UpdateSupplierByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSupplierByIDReply) Reset() {
	*x = UpdateSupplierByIDReply{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSupplierByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSupplierByIDReply) ProtoMessage() {}

func (x *UpdateSupplierByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSupplierByIDReply.ProtoReflect.Descriptor instead.
func (*UpdateSupplierByIDReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{5}
}

type Supplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                      // 供应商ID
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                   // 供应商名称
	ContactPerson string `protobuf:"bytes,3,opt,name=contactPerson,proto3" json:"contactPerson"` // 联系人
	Phone         string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone"`                 // 联系电话
	PaymentTerms  string `protobuf:"bytes,5,opt,name=paymentTerms,proto3" json:"paymentTerms"`   // 结算条款
	CreatedAt     string `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt"`         // 创建时间
	UpdatedAt     string `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt"`         // 更新时间
}

func (x *Supplier) Reset() {
	*x = Supplier{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Supplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Supplier) ProtoMessage() {}

func (x *Supplier) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Supplier.ProtoReflect.Descriptor instead.
func (*Supplier) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{6}
}

func (x *Supplier) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Supplier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Supplier) GetContactPerson() string {
	if x != nil {
		return x.ContactPerson
	}
	return ""
}

func (x *Supplier) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Supplier) GetPaymentTerms() string {
	if x != nil {
		return x.PaymentTerms
	}
	return ""
}

func (x *Supplier) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Supplier) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetSupplierByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`
}

func (x *GetSupplierByIDRequest) Reset() {
	*x = GetSupplierByIDRequest{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSupplierByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupplierByIDRequest) ProtoMessage() {}

func (x *GetSupplierByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupplierByIDRequest.ProtoReflect.Descriptor instead.
func (*GetSupplierByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{7}
}

func (x *GetSupplierByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSupplierByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Supplier *Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier"`
}

func (x *GetSupplierByIDReply) Reset() {
	*x = GetSupplierByIDReply{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSupplierByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSupplierByIDReply) ProtoMessage() {}

func (x *GetSupplierByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSupplierByIDReply.ProtoReflect.Descriptor instead.
func (*GetSupplierByIDReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{8}
}

func (x *GetSupplierByIDReply) GetSupplier() *Supplier {
	if x != nil {
		return x.Supplier
	}
	return nil
}

type ListSupplierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *types.Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (x *ListSupplierRequest) Reset() {
	*x = ListSupplierRequest{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSupplierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupplierRequest) ProtoMessage() {}

func (x *ListSupplierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupplierRequest.ProtoReflect.Descriptor instead.
func (*ListSupplierRequest) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{9}
}

func (x *ListSupplierRequest) GetParams() *types.Params {
	if x != nil {
		return x.Params
	}
	return nil
}

type ListSupplierReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     int64       `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Suppliers []*Supplier `protobuf:"bytes,2,rep,name=suppliers,proto3" json:"suppliers"`
}

func (x *ListSupplierReply) Reset() {
	*x = ListSupplierReply{}
	mi := &file_api_store_v1_supplier_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSupplierReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupplierReply) ProtoMessage() {}

func (x *ListSupplierReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_store_v1_supplier_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupplierReply.ProtoReflect.Descriptor instead.
func (*ListSupplierReply) Descriptor() ([]byte, []int) {
	return file_api_store_v1_supplier_proto_rawDescGZIP(), []int{10}
}

func (x *ListSupplierReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListSupplierReply) GetSuppliers() []*Supplier {
	if x != nil {
		return x.Suppliers
	}
	return nil
}

var File_api_store_v1_supplier_proto protoreflect.FileDescriptor

var file_api_store_v1_supplier_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x61, 0x70, 0x69,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b,
	0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x22, 0x25, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x14, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69,
	0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0xb5, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x14, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69, 0x64,
	0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x72, 0x6d,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x65, 0x72, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x3e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x14, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x08, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x69, 0x65, 0x72, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x22, 0x40, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0x5f, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x73,
	0x32, 0xd8, 0x04, 0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x6d, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x7b, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a,
	0x01, 0x2a, 0x1a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x72, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0xb4, 0x01, 0x92, 0x41,
	0x99, 0x01, 0x12, 0x15, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x70, 0x69, 0x20,
	0x64, 0x6f, 0x63, 0x73, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x5a, 0x48, 0x0a, 0x46, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x38, 0x08, 0x02, 0x12, 0x23, 0x54, 0x79, 0x70, 0x65, 0x20, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x2d, 0x6a, 0x77, 0x74, 0x2d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x20, 0x74, 0x6f, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x5a, 0x15, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_store_v1_supplier_proto_rawDescOnce sync.Once
	file_api_store_v1_supplier_proto_rawDescData = file_api_store_v1_supplier_proto_rawDesc
)

func file_api_store_v1_supplier_proto_rawDescGZIP() []byte {
	file_api_store_v1_supplier_proto_rawDescOnce.Do(func() {
		file_api_store_v1_supplier_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_store_v1_supplier_proto_rawDescData)
	})
	return file_api_store_v1_supplier_proto_rawDescData
}

var file_api_store_v1_supplier_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_store_v1_supplier_proto_goTypes = []any{
	(*CreateSupplierRequest)(nil),     // 0: api.store.v1.CreateSupplierRequest
	(*CreateSupplierReply)(nil),       // 1: api.store.v1.CreateSupplierReply
	(*DeleteSupplierByIDRequest)(nil), // 2: api.store.v1.DeleteSupplierByIDRequest
	(*DeleteSupplierByIDReply)(nil),   // 3: api.store.v1.DeleteSupplierByIDReply
	(*UpdateSupplierByIDRequest)(nil), // 4: api.store.v1.UpdateSupplierByIDRequest
	(*UpdateSupplierByIDReply)(nil),   // 5: api.store.v1.UpdateSupplierByIDReply
	(*Supplier)(nil),                  // 6: api.store.v1.Supplier
	(*GetSupplierByIDRequest)(nil),    // 7: api.store.v1.GetSupplierByIDRequest
	(*GetSupplierByIDReply)(nil),      // 8: api.store.v1.GetSupplierByIDReply
	(*ListSupplierRequest)(nil),       // 9: api.store.v1.ListSupplierRequest
	(*ListSupplierReply)(nil),         // 10: api.store.v1.ListSupplierReply
	(*types.Params)(nil),              // 11: api.types.Params
}
var file_api_store_v1_supplier_proto_depIdxs = []int32{
	6,  // 0: api.store.v1.GetSupplierByIDReply.supplier:type_name -> api.store.v1.Supplier
	11, // 1: api.store.v1.ListSupplierRequest.params:type_name -> api.types.Params
	6,  // 2: api.store.v1.ListSupplierReply.suppliers:type_name -> api.store.v1.Supplier
	0,  // 3: api.store.v1.supplier.Create:input_type -> api.store.v1.CreateSupplierRequest
	2,  // 4: api.store.v1.supplier.DeleteByID:input_type -> api.store.v1.DeleteSupplierByIDRequest
	4,  // 5: api.store.v1.supplier.UpdateByID:input_type -> api.store.v1.UpdateSupplierByIDRequest
	7,  // 6: api.store.v1.supplier.GetByID:input_type -> api.store.v1.GetSupplierByIDRequest
	9,  // 7: api.store.v1.supplier.List:input_type -> api.store.v1.ListSupplierRequest
	1,  // 8: api.store.v1.supplier.Create:output_type -> api.store.v1.CreateSupplierReply
	3,  // 9: api.store.v1.supplier.DeleteByID:output_type -> api.store.v1.DeleteSupplierByIDReply
	5,  // 10: api.store.v1.supplier.UpdateByID:output_type -> api.store.v1.UpdateSupplierByIDReply
	8,  // 11: api.store.v1.supplier.GetByID:output_type -> api.store.v1.GetSupplierByIDReply
	10, // 12: api.store.v1.supplier.List:output_type -> api.store.v1.ListSupplierReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_store_v1_supplier_proto_init() }
func file_api_store_v1_supplier_proto_init() {
	if File_api_store_v1_supplier_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_store_v1_supplier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_store_v1_supplier_proto_goTypes,
		DependencyIndexes: file_api_store_v1_supplier_proto_depIdxs,
		MessageInfos:      file_api_store_v1_supplier_proto_msgTypes,
	}.Build()
	File_api_store_v1_supplier_proto = out.File
	file_api_store_v1_supplier_proto_rawDesc = nil
	file_api_store_v1_supplier_proto_goTypes = nil
	file_api_store_v1_supplier_proto_depIdxs = nil
}
