// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: proto/favouritegrpc.proto

package proto

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

const (
	FavouriteGRPC_GetFavouritesByShowID_FullMethodName = "/FavouriteGRPC/GetFavouritesByShowID"
)

// FavouriteGRPCClient is the client API for FavouriteGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavouriteGRPCClient interface {
	GetFavouritesByShowID(ctx context.Context, in *GetFavouritesByShowIDRequest, opts ...grpc.CallOption) (*GetFavouritesByShowIDResponse, error)
}

type favouriteGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewFavouriteGRPCClient(cc grpc.ClientConnInterface) FavouriteGRPCClient {
	return &favouriteGRPCClient{cc}
}

func (c *favouriteGRPCClient) GetFavouritesByShowID(ctx context.Context, in *GetFavouritesByShowIDRequest, opts ...grpc.CallOption) (*GetFavouritesByShowIDResponse, error) {
	out := new(GetFavouritesByShowIDResponse)
	err := c.cc.Invoke(ctx, FavouriteGRPC_GetFavouritesByShowID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FavouriteGRPCServer is the server API for FavouriteGRPC service.
// All implementations must embed UnimplementedFavouriteGRPCServer
// for forward compatibility
type FavouriteGRPCServer interface {
	GetFavouritesByShowID(context.Context, *GetFavouritesByShowIDRequest) (*GetFavouritesByShowIDResponse, error)
	mustEmbedUnimplementedFavouriteGRPCServer()
}

// UnimplementedFavouriteGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedFavouriteGRPCServer struct {
}

func (UnimplementedFavouriteGRPCServer) GetFavouritesByShowID(context.Context, *GetFavouritesByShowIDRequest) (*GetFavouritesByShowIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavouritesByShowID not implemented")
}
func (UnimplementedFavouriteGRPCServer) mustEmbedUnimplementedFavouriteGRPCServer() {}

// UnsafeFavouriteGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavouriteGRPCServer will
// result in compilation errors.
type UnsafeFavouriteGRPCServer interface {
	mustEmbedUnimplementedFavouriteGRPCServer()
}

func RegisterFavouriteGRPCServer(s grpc.ServiceRegistrar, srv FavouriteGRPCServer) {
	s.RegisterService(&FavouriteGRPC_ServiceDesc, srv)
}

func _FavouriteGRPC_GetFavouritesByShowID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavouritesByShowIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavouriteGRPCServer).GetFavouritesByShowID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavouriteGRPC_GetFavouritesByShowID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavouriteGRPCServer).GetFavouritesByShowID(ctx, req.(*GetFavouritesByShowIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FavouriteGRPC_ServiceDesc is the grpc.ServiceDesc for FavouriteGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavouriteGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FavouriteGRPC",
	HandlerType: (*FavouriteGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFavouritesByShowID",
			Handler:    _FavouriteGRPC_GetFavouritesByShowID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/favouritegrpc.proto",
}