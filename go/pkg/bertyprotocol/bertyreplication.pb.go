// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bertyreplication.proto

package bertyprotocol

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	bertytypes "berty.tech/berty/v2/go/pkg/bertytypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("bertyreplication.proto", fileDescriptor_07ad8506da4f8389) }

var fileDescriptor_07ad8506da4f8389 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4a, 0x2d, 0x2a,
	0xa9, 0x2c, 0x4a, 0x2d, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x04, 0x8b, 0x43, 0x38, 0xc9, 0xf9, 0x39, 0x7a, 0x65, 0x86, 0x52, 0x02,
	0x60, 0xa1, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x88, 0xb8, 0xd1, 0x4c, 0x46, 0x2e, 0xa1, 0x20, 0x84,
	0xd6, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0xa1, 0x26, 0x46, 0x2e, 0x3e, 0x98, 0x70, 0xaa,
	0x7b, 0x51, 0x7e, 0x69, 0x81, 0x90, 0x85, 0x1e, 0xc4, 0x3c, 0x88, 0xee, 0x32, 0x43, 0x3d, 0x4c,
	0x6d, 0xa8, 0x3a, 0xf4, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0xcc, 0xc8, 0xd0, 0x59,
	0x90, 0x53, 0xe9, 0xa4, 0x1e, 0xa5, 0x0a, 0xd5, 0x98, 0x9a, 0x9c, 0xa1, 0x0f, 0x66, 0xea, 0xa7,
	0xe7, 0xeb, 0x17, 0x64, 0xa7, 0x43, 0x38, 0x30, 0x9f, 0x25, 0xb1, 0x81, 0x59, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6e, 0x4a, 0x04, 0x6b, 0x0a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ReplicationServiceClient is the client API for ReplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReplicationServiceClient interface {
	// ReplicateGroup
	ReplicateGroup(ctx context.Context, in *bertytypes.ReplicationServiceReplicateGroup_Request, opts ...grpc.CallOption) (*bertytypes.ReplicationServiceReplicateGroup_Reply, error)
}

type replicationServiceClient struct {
	cc *grpc.ClientConn
}

func NewReplicationServiceClient(cc *grpc.ClientConn) ReplicationServiceClient {
	return &replicationServiceClient{cc}
}

func (c *replicationServiceClient) ReplicateGroup(ctx context.Context, in *bertytypes.ReplicationServiceReplicateGroup_Request, opts ...grpc.CallOption) (*bertytypes.ReplicationServiceReplicateGroup_Reply, error) {
	out := new(bertytypes.ReplicationServiceReplicateGroup_Reply)
	err := c.cc.Invoke(ctx, "/berty.protocol.v1.ReplicationService/ReplicateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServiceServer is the server API for ReplicationService service.
type ReplicationServiceServer interface {
	// ReplicateGroup
	ReplicateGroup(context.Context, *bertytypes.ReplicationServiceReplicateGroup_Request) (*bertytypes.ReplicationServiceReplicateGroup_Reply, error)
}

// UnimplementedReplicationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReplicationServiceServer struct {
}

func (*UnimplementedReplicationServiceServer) ReplicateGroup(ctx context.Context, req *bertytypes.ReplicationServiceReplicateGroup_Request) (*bertytypes.ReplicationServiceReplicateGroup_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicateGroup not implemented")
}

func RegisterReplicationServiceServer(s *grpc.Server, srv ReplicationServiceServer) {
	s.RegisterService(&_ReplicationService_serviceDesc, srv)
}

func _ReplicationService_ReplicateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bertytypes.ReplicationServiceReplicateGroup_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).ReplicateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.protocol.v1.ReplicationService/ReplicateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).ReplicateGroup(ctx, req.(*bertytypes.ReplicationServiceReplicateGroup_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReplicationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "berty.protocol.v1.ReplicationService",
	HandlerType: (*ReplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplicateGroup",
			Handler:    _ReplicationService_ReplicateGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bertyreplication.proto",
}
