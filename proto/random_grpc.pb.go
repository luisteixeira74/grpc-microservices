// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: random.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RandomService_GetRandomStream_FullMethodName = "/random.RandomService/GetRandomStream"
)

// RandomServiceClient is the client API for RandomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RandomServiceClient interface {
	GetRandomStream(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RandomResponse], error)
}

type randomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRandomServiceClient(cc grpc.ClientConnInterface) RandomServiceClient {
	return &randomServiceClient{cc}
}

func (c *randomServiceClient) GetRandomStream(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[RandomResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RandomService_ServiceDesc.Streams[0], RandomService_GetRandomStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[RandomRequest, RandomResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomService_GetRandomStreamClient = grpc.ServerStreamingClient[RandomResponse]

// RandomServiceServer is the server API for RandomService service.
// All implementations must embed UnimplementedRandomServiceServer
// for forward compatibility.
type RandomServiceServer interface {
	GetRandomStream(*RandomRequest, grpc.ServerStreamingServer[RandomResponse]) error
	mustEmbedUnimplementedRandomServiceServer()
}

// UnimplementedRandomServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRandomServiceServer struct{}

func (UnimplementedRandomServiceServer) GetRandomStream(*RandomRequest, grpc.ServerStreamingServer[RandomResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetRandomStream not implemented")
}
func (UnimplementedRandomServiceServer) mustEmbedUnimplementedRandomServiceServer() {}
func (UnimplementedRandomServiceServer) testEmbeddedByValue()                       {}

// UnsafeRandomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RandomServiceServer will
// result in compilation errors.
type UnsafeRandomServiceServer interface {
	mustEmbedUnimplementedRandomServiceServer()
}

func RegisterRandomServiceServer(s grpc.ServiceRegistrar, srv RandomServiceServer) {
	// If the following call pancis, it indicates UnimplementedRandomServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RandomService_ServiceDesc, srv)
}

func _RandomService_GetRandomStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RandomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RandomServiceServer).GetRandomStream(m, &grpc.GenericServerStream[RandomRequest, RandomResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RandomService_GetRandomStreamServer = grpc.ServerStreamingServer[RandomResponse]

// RandomService_ServiceDesc is the grpc.ServiceDesc for RandomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RandomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "random.RandomService",
	HandlerType: (*RandomServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRandomStream",
			Handler:       _RandomService_GetRandomStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "random.proto",
}
