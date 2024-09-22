// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.2
// source: api/stock/v1/atomic.proto

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

type UpdateAtomicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`
	Stock uint32 `protobuf:"varint,2,opt,name=stock,proto3" json:"stock"` // 库存数量
}

func (x *UpdateAtomicRequest) Reset() {
	*x = UpdateAtomicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stock_v1_atomic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAtomicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAtomicRequest) ProtoMessage() {}

func (x *UpdateAtomicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stock_v1_atomic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAtomicRequest.ProtoReflect.Descriptor instead.
func (*UpdateAtomicRequest) Descriptor() ([]byte, []int) {
	return file_api_stock_v1_atomic_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAtomicRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAtomicRequest) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type UpdateAtomicRequestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAtomicRequestReply) Reset() {
	*x = UpdateAtomicRequestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stock_v1_atomic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAtomicRequestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAtomicRequestReply) ProtoMessage() {}

func (x *UpdateAtomicRequestReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_stock_v1_atomic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAtomicRequestReply.ProtoReflect.Descriptor instead.
func (*UpdateAtomicRequestReply) Descriptor() ([]byte, []int) {
	return file_api_stock_v1_atomic_proto_rawDescGZIP(), []int{1}
}

type QueryAtomicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`
}

func (x *QueryAtomicRequest) Reset() {
	*x = QueryAtomicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stock_v1_atomic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAtomicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAtomicRequest) ProtoMessage() {}

func (x *QueryAtomicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stock_v1_atomic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAtomicRequest.ProtoReflect.Descriptor instead.
func (*QueryAtomicRequest) Descriptor() ([]byte, []int) {
	return file_api_stock_v1_atomic_proto_rawDescGZIP(), []int{2}
}

func (x *QueryAtomicRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QueryAtomicReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock uint32 `protobuf:"varint,1,opt,name=stock,proto3" json:"stock"` // 库存数量
}

func (x *QueryAtomicReply) Reset() {
	*x = QueryAtomicReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stock_v1_atomic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAtomicReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAtomicReply) ProtoMessage() {}

func (x *QueryAtomicReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_stock_v1_atomic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAtomicReply.ProtoReflect.Descriptor instead.
func (*QueryAtomicReply) Descriptor() ([]byte, []int) {
	return file_api_stock_v1_atomic_proto_rawDescGZIP(), []int{3}
}

func (x *QueryAtomicReply) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

var File_api_stock_v1_atomic_proto protoreflect.FileDescriptor

var file_api_stock_v1_atomic_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x14, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x6f, 0x6d, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3a, 0x0a,
	0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x14, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69,
	0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x32, 0xee, 0x02, 0x0a, 0x06, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x12, 0xc2,
	0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x6d, 0x92, 0x41, 0x46, 0x0a, 0x11, 0x63, 0x61, 0x73, 0x65, 0x20,
	0x32, 0x3a, 0x20, 0xe5, 0x8e, 0x9f, 0xe5, 0xad, 0x90, 0xe6, 0x80, 0xa7, 0x12, 0x0c, 0xe6, 0x9b,
	0xb4, 0xe6, 0x96, 0xb0, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0x1a, 0x23, 0xe6, 0x9b, 0xb4, 0xe6,
	0x96, 0xb0, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xef, 0xbc, 0x8c, 0x44, 0x42, 0xe5, 0x92, 0x8c,
	0xe7, 0xbc, 0x93, 0xe5, 0xad, 0x98, 0xe5, 0x8e, 0x9f, 0xe5, 0xad, 0x90, 0xe6, 0x80, 0xa7, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x1a, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x74, 0x6f, 0x6d, 0x69, 0x63,
	0x3a, 0x01, 0x2a, 0x12, 0x9e, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x53, 0x92, 0x41, 0x2f, 0x0a, 0x11, 0x63, 0x61, 0x73, 0x65, 0x20, 0x32, 0x3a, 0x20, 0xe5, 0x8e,
	0x9f, 0xe5, 0xad, 0x90, 0xe6, 0x80, 0xa7, 0x12, 0x0c, 0xe6, 0x9f, 0xa5, 0xe8, 0xaf, 0xa2, 0xe6,
	0x95, 0xb0, 0xe6, 0x8d, 0xae, 0x1a, 0x0c, 0xe6, 0x9f, 0xa5, 0xe8, 0xaf, 0xa2, 0xe6, 0x95, 0xb0,
	0xe6, 0x8d, 0xae, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x74,
	0x6f, 0x6d, 0x69, 0x63, 0x42, 0xb4, 0x01, 0x5a, 0x15, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x92, 0x41,
	0x99, 0x01, 0x12, 0x15, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x20, 0x61, 0x70, 0x69, 0x20,
	0x64, 0x6f, 0x63, 0x73, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x5a, 0x48, 0x0a, 0x46, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x38, 0x08, 0x02, 0x12, 0x23, 0x54, 0x79, 0x70, 0x65, 0x20, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x2d, 0x6a, 0x77, 0x74, 0x2d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x20, 0x74, 0x6f, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_stock_v1_atomic_proto_rawDescOnce sync.Once
	file_api_stock_v1_atomic_proto_rawDescData = file_api_stock_v1_atomic_proto_rawDesc
)

func file_api_stock_v1_atomic_proto_rawDescGZIP() []byte {
	file_api_stock_v1_atomic_proto_rawDescOnce.Do(func() {
		file_api_stock_v1_atomic_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_stock_v1_atomic_proto_rawDescData)
	})
	return file_api_stock_v1_atomic_proto_rawDescData
}

var file_api_stock_v1_atomic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_stock_v1_atomic_proto_goTypes = []interface{}{
	(*UpdateAtomicRequest)(nil),      // 0: api.stock.v1.UpdateAtomicRequest
	(*UpdateAtomicRequestReply)(nil), // 1: api.stock.v1.UpdateAtomicRequestReply
	(*QueryAtomicRequest)(nil),       // 2: api.stock.v1.QueryAtomicRequest
	(*QueryAtomicReply)(nil),         // 3: api.stock.v1.QueryAtomicReply
}
var file_api_stock_v1_atomic_proto_depIdxs = []int32{
	0, // 0: api.stock.v1.atomic.Update:input_type -> api.stock.v1.UpdateAtomicRequest
	2, // 1: api.stock.v1.atomic.Query:input_type -> api.stock.v1.QueryAtomicRequest
	1, // 2: api.stock.v1.atomic.Update:output_type -> api.stock.v1.UpdateAtomicRequestReply
	3, // 3: api.stock.v1.atomic.Query:output_type -> api.stock.v1.QueryAtomicReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_stock_v1_atomic_proto_init() }
func file_api_stock_v1_atomic_proto_init() {
	if File_api_stock_v1_atomic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_stock_v1_atomic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAtomicRequest); i {
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
		file_api_stock_v1_atomic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAtomicRequestReply); i {
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
		file_api_stock_v1_atomic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAtomicRequest); i {
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
		file_api_stock_v1_atomic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAtomicReply); i {
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
			RawDescriptor: file_api_stock_v1_atomic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_stock_v1_atomic_proto_goTypes,
		DependencyIndexes: file_api_stock_v1_atomic_proto_depIdxs,
		MessageInfos:      file_api_stock_v1_atomic_proto_msgTypes,
	}.Build()
	File_api_stock_v1_atomic_proto = out.File
	file_api_stock_v1_atomic_proto_rawDesc = nil
	file_api_stock_v1_atomic_proto_goTypes = nil
	file_api_stock_v1_atomic_proto_depIdxs = nil
}
