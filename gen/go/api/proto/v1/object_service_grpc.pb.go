// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: api/proto/v1/object_service.proto

package gen

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

// ObjectServiceAPIClient is the client API for ObjectServiceAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectServiceAPIClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	GetNearest(ctx context.Context, in *GetNearestRequest, opts ...grpc.CallOption) (*GetNearestResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type objectServiceAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectServiceAPIClient(cc grpc.ClientConnInterface) ObjectServiceAPIClient {
	return &objectServiceAPIClient{cc}
}

func (c *objectServiceAPIClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/object_service_api.ObjectServiceAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceAPIClient) GetById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetByIdResponse, error) {
	out := new(GetByIdResponse)
	err := c.cc.Invoke(ctx, "/object_service_api.ObjectServiceAPI/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceAPIClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/object_service_api.ObjectServiceAPI/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceAPIClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/object_service_api.ObjectServiceAPI/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceAPIClient) GetNearest(ctx context.Context, in *GetNearestRequest, opts ...grpc.CallOption) (*GetNearestResponse, error) {
	out := new(GetNearestResponse)
	err := c.cc.Invoke(ctx, "/object_service_api.ObjectServiceAPI/GetNearest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/object_service_api.ObjectServiceAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/object_service_api.ObjectServiceAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectServiceAPIServer is the server API for ObjectServiceAPI service.
// All implementations must embed UnimplementedObjectServiceAPIServer
// for forward compatibility
type ObjectServiceAPIServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error)
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	GetNearest(context.Context, *GetNearestRequest) (*GetNearestResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedObjectServiceAPIServer()
}

// UnimplementedObjectServiceAPIServer must be embedded to have forward compatible implementations.
type UnimplementedObjectServiceAPIServer struct {
}

func (UnimplementedObjectServiceAPIServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedObjectServiceAPIServer) GetById(context.Context, *GetByIdRequest) (*GetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedObjectServiceAPIServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedObjectServiceAPIServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedObjectServiceAPIServer) GetNearest(context.Context, *GetNearestRequest) (*GetNearestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearest not implemented")
}
func (UnimplementedObjectServiceAPIServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedObjectServiceAPIServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedObjectServiceAPIServer) mustEmbedUnimplementedObjectServiceAPIServer() {}

// UnsafeObjectServiceAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectServiceAPIServer will
// result in compilation errors.
type UnsafeObjectServiceAPIServer interface {
	mustEmbedUnimplementedObjectServiceAPIServer()
}

func RegisterObjectServiceAPIServer(s grpc.ServiceRegistrar, srv ObjectServiceAPIServer) {
	s.RegisterService(&ObjectServiceAPI_ServiceDesc, srv)
}

func _ObjectServiceAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_service_api.ObjectServiceAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceAPIServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectServiceAPI_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceAPIServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_service_api.ObjectServiceAPI/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceAPIServer).GetById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectServiceAPI_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceAPIServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_service_api.ObjectServiceAPI/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceAPIServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectServiceAPI_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceAPIServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_service_api.ObjectServiceAPI/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceAPIServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectServiceAPI_GetNearest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNearestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceAPIServer).GetNearest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_service_api.ObjectServiceAPI/GetNearest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceAPIServer).GetNearest(ctx, req.(*GetNearestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectServiceAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_service_api.ObjectServiceAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectServiceAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_service_api.ObjectServiceAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectServiceAPI_ServiceDesc is the grpc.ServiceDesc for ObjectServiceAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectServiceAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object_service_api.ObjectServiceAPI",
	HandlerType: (*ObjectServiceAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ObjectServiceAPI_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ObjectServiceAPI_GetById_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ObjectServiceAPI_GetList_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ObjectServiceAPI_Search_Handler,
		},
		{
			MethodName: "GetNearest",
			Handler:    _ObjectServiceAPI_GetNearest_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ObjectServiceAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ObjectServiceAPI_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/object_service.proto",
}
