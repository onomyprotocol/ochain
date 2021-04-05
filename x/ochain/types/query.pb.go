// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ochain/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("ochain/query.proto", fileDescriptor_377f6fc793efd374) }

var fileDescriptor_377f6fc793efd374 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x4f, 0xce, 0x48,
	0xcc, 0xcc, 0xd3, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92,
	0xce, 0xcf, 0xcb, 0xcf, 0xad, 0x04, 0xb3, 0x93, 0xf3, 0x73, 0xf4, 0x20, 0x2a, 0xa0, 0x94, 0x94,
	0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e,
	0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0xab, 0x94, 0x56, 0x72, 0x7e, 0x71, 0x6e, 0x7e,
	0xb1, 0x7e, 0x52, 0x62, 0x71, 0x2a, 0xc4, 0x4c, 0xfd, 0x32, 0xc3, 0xa4, 0xd4, 0x92, 0x44, 0x43,
	0xfd, 0x82, 0xc4, 0xf4, 0xcc, 0x3c, 0xb0, 0x62, 0x88, 0x5a, 0x23, 0x76, 0x2e, 0xd6, 0x40, 0x90,
	0x0a, 0x27, 0x8f, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4b, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0x71, 0x94, 0x3e, 0xd4, 0xd9, 0x15,
	0x30, 0x46, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x58, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xda, 0xc0, 0xe3, 0xff, 0xd6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "onomyprotocol.ochain.ochain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "ochain/query.proto",
}
