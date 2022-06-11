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
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type objectServiceAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectServiceAPIClient(cc grpc.ClientConnInterface) ObjectServiceAPIClient {
	return &objectServiceAPIClient{cc}
}

func (c *objectServiceAPIClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/object_service_api.ObjectServiceAPI/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectServiceAPIServer is the server API for ObjectServiceAPI service.
// All implementations must embed UnimplementedObjectServiceAPIServer
// for forward compatibility
type ObjectServiceAPIServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	mustEmbedUnimplementedObjectServiceAPIServer()
}

// UnimplementedObjectServiceAPIServer must be embedded to have forward compatible implementations.
type UnimplementedObjectServiceAPIServer struct {
}

func (UnimplementedObjectServiceAPIServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
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

func _ObjectServiceAPI_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceAPIServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object_service_api.ObjectServiceAPI/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceAPIServer).Status(ctx, req.(*StatusRequest))
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
			MethodName: "Status",
			Handler:    _ObjectServiceAPI_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/object_service.proto",
}