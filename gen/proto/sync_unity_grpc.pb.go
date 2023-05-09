// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sync_unity.proto

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
	SyncUnity_Echo_FullMethodName = "/main.SyncUnity/Echo"
)

// SyncUnityClient is the client API for SyncUnity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncUnityClient interface {
	Echo(ctx context.Context, in *NoSql, opts ...grpc.CallOption) (*NoSql, error)
}

type syncUnityClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncUnityClient(cc grpc.ClientConnInterface) SyncUnityClient {
	return &syncUnityClient{cc}
}

func (c *syncUnityClient) Echo(ctx context.Context, in *NoSql, opts ...grpc.CallOption) (*NoSql, error) {
	out := new(NoSql)
	err := c.cc.Invoke(ctx, SyncUnity_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncUnityServer is the server API for SyncUnity service.
// All implementations must embed UnimplementedSyncUnityServer
// for forward compatibility
type SyncUnityServer interface {
	Echo(context.Context, *NoSql) (*NoSql, error)
	mustEmbedUnimplementedSyncUnityServer()
}

// UnimplementedSyncUnityServer must be embedded to have forward compatible implementations.
type UnimplementedSyncUnityServer struct {
}

func (UnimplementedSyncUnityServer) Echo(context.Context, *NoSql) (*NoSql, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedSyncUnityServer) mustEmbedUnimplementedSyncUnityServer() {}

// UnsafeSyncUnityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncUnityServer will
// result in compilation errors.
type UnsafeSyncUnityServer interface {
	mustEmbedUnimplementedSyncUnityServer()
}

func RegisterSyncUnityServer(s grpc.ServiceRegistrar, srv SyncUnityServer) {
	s.RegisterService(&SyncUnity_ServiceDesc, srv)
}

func _SyncUnity_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoSql)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncUnityServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncUnity_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncUnityServer).Echo(ctx, req.(*NoSql))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncUnity_ServiceDesc is the grpc.ServiceDesc for SyncUnity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncUnity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.SyncUnity",
	HandlerType: (*SyncUnityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _SyncUnity_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sync_unity.proto",
}
