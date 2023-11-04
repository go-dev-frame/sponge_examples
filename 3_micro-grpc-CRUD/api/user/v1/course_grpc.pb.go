// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/user/v1/course.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CourseClient is the client API for Course service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseClient interface {
	// create course
	Create(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseReply, error)
	// delete course by id
	DeleteByID(ctx context.Context, in *DeleteCourseByIDRequest, opts ...grpc.CallOption) (*DeleteCourseByIDReply, error)
	// delete course by batch id
	DeleteByIDs(ctx context.Context, in *DeleteCourseByIDsRequest, opts ...grpc.CallOption) (*DeleteCourseByIDsReply, error)
	// update course by id
	UpdateByID(ctx context.Context, in *UpdateCourseByIDRequest, opts ...grpc.CallOption) (*UpdateCourseByIDReply, error)
	// get course by id
	GetByID(ctx context.Context, in *GetCourseByIDRequest, opts ...grpc.CallOption) (*GetCourseByIDReply, error)
	// get course by condition
	GetByCondition(ctx context.Context, in *GetCourseByConditionRequest, opts ...grpc.CallOption) (*GetCourseByConditionReply, error)
	// list of course by batch id
	ListByIDs(ctx context.Context, in *ListCourseByIDsRequest, opts ...grpc.CallOption) (*ListCourseByIDsReply, error)
	// list of course by query parameters
	List(ctx context.Context, in *ListCourseRequest, opts ...grpc.CallOption) (*ListCourseReply, error)
}

type courseClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseClient(cc grpc.ClientConnInterface) CourseClient {
	return &courseClient{cc}
}

func (c *courseClient) Create(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CreateCourseReply, error) {
	out := new(CreateCourseReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.course/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) DeleteByID(ctx context.Context, in *DeleteCourseByIDRequest, opts ...grpc.CallOption) (*DeleteCourseByIDReply, error) {
	out := new(DeleteCourseByIDReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.course/DeleteByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) DeleteByIDs(ctx context.Context, in *DeleteCourseByIDsRequest, opts ...grpc.CallOption) (*DeleteCourseByIDsReply, error) {
	out := new(DeleteCourseByIDsReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.course/DeleteByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) UpdateByID(ctx context.Context, in *UpdateCourseByIDRequest, opts ...grpc.CallOption) (*UpdateCourseByIDReply, error) {
	out := new(UpdateCourseByIDReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.course/UpdateByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) GetByID(ctx context.Context, in *GetCourseByIDRequest, opts ...grpc.CallOption) (*GetCourseByIDReply, error) {
	out := new(GetCourseByIDReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.course/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) GetByCondition(ctx context.Context, in *GetCourseByConditionRequest, opts ...grpc.CallOption) (*GetCourseByConditionReply, error) {
	out := new(GetCourseByConditionReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.course/GetByCondition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) ListByIDs(ctx context.Context, in *ListCourseByIDsRequest, opts ...grpc.CallOption) (*ListCourseByIDsReply, error) {
	out := new(ListCourseByIDsReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.course/ListByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseClient) List(ctx context.Context, in *ListCourseRequest, opts ...grpc.CallOption) (*ListCourseReply, error) {
	out := new(ListCourseReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.course/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServer is the server API for Course service.
// All implementations must embed UnimplementedCourseServer
// for forward compatibility
type CourseServer interface {
	// create course
	Create(context.Context, *CreateCourseRequest) (*CreateCourseReply, error)
	// delete course by id
	DeleteByID(context.Context, *DeleteCourseByIDRequest) (*DeleteCourseByIDReply, error)
	// delete course by batch id
	DeleteByIDs(context.Context, *DeleteCourseByIDsRequest) (*DeleteCourseByIDsReply, error)
	// update course by id
	UpdateByID(context.Context, *UpdateCourseByIDRequest) (*UpdateCourseByIDReply, error)
	// get course by id
	GetByID(context.Context, *GetCourseByIDRequest) (*GetCourseByIDReply, error)
	// get course by condition
	GetByCondition(context.Context, *GetCourseByConditionRequest) (*GetCourseByConditionReply, error)
	// list of course by batch id
	ListByIDs(context.Context, *ListCourseByIDsRequest) (*ListCourseByIDsReply, error)
	// list of course by query parameters
	List(context.Context, *ListCourseRequest) (*ListCourseReply, error)
	mustEmbedUnimplementedCourseServer()
}

// UnimplementedCourseServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServer struct {
}

func (UnimplementedCourseServer) Create(context.Context, *CreateCourseRequest) (*CreateCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCourseServer) DeleteByID(context.Context, *DeleteCourseByIDRequest) (*DeleteCourseByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByID not implemented")
}
func (UnimplementedCourseServer) DeleteByIDs(context.Context, *DeleteCourseByIDsRequest) (*DeleteCourseByIDsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByIDs not implemented")
}
func (UnimplementedCourseServer) UpdateByID(context.Context, *UpdateCourseByIDRequest) (*UpdateCourseByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByID not implemented")
}
func (UnimplementedCourseServer) GetByID(context.Context, *GetCourseByIDRequest) (*GetCourseByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedCourseServer) GetByCondition(context.Context, *GetCourseByConditionRequest) (*GetCourseByConditionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCondition not implemented")
}
func (UnimplementedCourseServer) ListByIDs(context.Context, *ListCourseByIDsRequest) (*ListCourseByIDsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListByIDs not implemented")
}
func (UnimplementedCourseServer) List(context.Context, *ListCourseRequest) (*ListCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCourseServer) mustEmbedUnimplementedCourseServer() {}

// UnsafeCourseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServer will
// result in compilation errors.
type UnsafeCourseServer interface {
	mustEmbedUnimplementedCourseServer()
}

func RegisterCourseServer(s grpc.ServiceRegistrar, srv CourseServer) {
	s.RegisterService(&Course_ServiceDesc, srv)
}

func _Course_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.course/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).Create(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_DeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourseByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).DeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.course/DeleteByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).DeleteByID(ctx, req.(*DeleteCourseByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_DeleteByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourseByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).DeleteByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.course/DeleteByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).DeleteByIDs(ctx, req.(*DeleteCourseByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_UpdateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).UpdateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.course/UpdateByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).UpdateByID(ctx, req.(*UpdateCourseByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.course/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).GetByID(ctx, req.(*GetCourseByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_GetByCondition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseByConditionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).GetByCondition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.course/GetByCondition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).GetByCondition(ctx, req.(*GetCourseByConditionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_ListByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCourseByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).ListByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.course/ListByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).ListByIDs(ctx, req.(*ListCourseByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Course_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.course/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServer).List(ctx, req.(*ListCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Course_ServiceDesc is the grpc.ServiceDesc for Course service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Course_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.v1.course",
	HandlerType: (*CourseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Course_Create_Handler,
		},
		{
			MethodName: "DeleteByID",
			Handler:    _Course_DeleteByID_Handler,
		},
		{
			MethodName: "DeleteByIDs",
			Handler:    _Course_DeleteByIDs_Handler,
		},
		{
			MethodName: "UpdateByID",
			Handler:    _Course_UpdateByID_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Course_GetByID_Handler,
		},
		{
			MethodName: "GetByCondition",
			Handler:    _Course_GetByCondition_Handler,
		},
		{
			MethodName: "ListByIDs",
			Handler:    _Course_ListByIDs_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Course_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/course.proto",
}
