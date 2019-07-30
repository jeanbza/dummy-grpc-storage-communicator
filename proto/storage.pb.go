// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package storage_bench

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ObjectRead struct {
	// The bucket string identifier.
	BucketName string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	// The object/blob string identifier.
	ObjectName           string   `protobuf:"bytes,2,opt,name=objectName,proto3" json:"objectName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectRead) Reset()         { *m = ObjectRead{} }
func (m *ObjectRead) String() string { return proto.CompactTextString(m) }
func (*ObjectRead) ProtoMessage()    {}
func (*ObjectRead) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *ObjectRead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectRead.Unmarshal(m, b)
}
func (m *ObjectRead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectRead.Marshal(b, m, deterministic)
}
func (m *ObjectRead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectRead.Merge(m, src)
}
func (m *ObjectRead) XXX_Size() int {
	return xxx_messageInfo_ObjectRead.Size(m)
}
func (m *ObjectRead) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectRead.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectRead proto.InternalMessageInfo

func (m *ObjectRead) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *ObjectRead) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

type ObjectWrite struct {
	// The bucket string identifier.
	BucketName string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	// The object/blob string identifiers.
	ObjectName string `protobuf:"bytes,2,opt,name=objectName,proto3" json:"objectName,omitempty"`
	// The string containing the upload file path.
	Destination          string   `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectWrite) Reset()         { *m = ObjectWrite{} }
func (m *ObjectWrite) String() string { return proto.CompactTextString(m) }
func (*ObjectWrite) ProtoMessage()    {}
func (*ObjectWrite) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *ObjectWrite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectWrite.Unmarshal(m, b)
}
func (m *ObjectWrite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectWrite.Marshal(b, m, deterministic)
}
func (m *ObjectWrite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectWrite.Merge(m, src)
}
func (m *ObjectWrite) XXX_Size() int {
	return xxx_messageInfo_ObjectWrite.Size(m)
}
func (m *ObjectWrite) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectWrite.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectWrite proto.InternalMessageInfo

func (m *ObjectWrite) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *ObjectWrite) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

func (m *ObjectWrite) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ObjectRead)(nil), "storage_bench.ObjectRead")
	proto.RegisterType((*ObjectWrite)(nil), "storage_bench.ObjectWrite")
	proto.RegisterType((*EmptyResponse)(nil), "storage_bench.EmptyResponse")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x82, 0x71, 0xe3, 0x93, 0x52, 0xf3,
	0x92, 0x33, 0x94, 0x7c, 0xb8, 0xb8, 0xfc, 0x93, 0xb2, 0x52, 0x93, 0x4b, 0x82, 0x52, 0x13, 0x53,
	0x84, 0xe4, 0xb8, 0xb8, 0x92, 0x4a, 0x93, 0xb3, 0x53, 0x4b, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x90, 0x44, 0x40, 0xf2, 0xf9, 0x60, 0xd5, 0x60, 0x79, 0x26, 0x88,
	0x3c, 0x42, 0x44, 0x29, 0x9f, 0x8b, 0x1b, 0x62, 0x5a, 0x78, 0x51, 0x66, 0x49, 0x2a, 0xa5, 0xc6,
	0x09, 0x29, 0x70, 0x71, 0xa7, 0xa4, 0x16, 0x97, 0x64, 0xe6, 0x25, 0x96, 0x64, 0xe6, 0xe7, 0x49,
	0x30, 0x83, 0x15, 0x20, 0x0b, 0x29, 0xf1, 0x73, 0xf1, 0xba, 0xe6, 0x16, 0x94, 0x54, 0x06, 0xa5,
	0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0xcd, 0x65, 0xe4, 0x12, 0x0e, 0x86, 0xf8, 0xd0, 0x09,
	0xe4, 0xc1, 0xf0, 0xa2, 0xc4, 0x82, 0x82, 0xd4, 0x22, 0x21, 0x67, 0x2e, 0x56, 0x88, 0x9b, 0xa4,
	0xf4, 0x50, 0x02, 0x40, 0x0f, 0xc9, 0xbd, 0x52, 0x32, 0x68, 0x72, 0x28, 0x46, 0x2b, 0x31, 0x08,
	0x39, 0x72, 0xb1, 0x80, 0x83, 0x49, 0x12, 0xab, 0x19, 0x20, 0x29, 0x42, 0x46, 0x38, 0x89, 0x70,
	0x09, 0x24, 0xe7, 0xe7, 0xea, 0x81, 0x25, 0xcb, 0x21, 0x6e, 0x0b, 0x60, 0x4c, 0x62, 0x03, 0xc7,
	0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xb0, 0xd3, 0x42, 0xac, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageBenchWrapperClient is the client API for StorageBenchWrapper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageBenchWrapperClient interface {
	// Performs an upload from a specific object.
	Write(ctx context.Context, in *ObjectWrite, opts ...grpc.CallOption) (*EmptyResponse, error)
	// Read a specific object.
	Read(ctx context.Context, in *ObjectRead, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type storageBenchWrapperClient struct {
	cc *grpc.ClientConn
}

func NewStorageBenchWrapperClient(cc *grpc.ClientConn) StorageBenchWrapperClient {
	return &storageBenchWrapperClient{cc}
}

func (c *storageBenchWrapperClient) Write(ctx context.Context, in *ObjectWrite, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/storage_bench.StorageBenchWrapper/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageBenchWrapperClient) Read(ctx context.Context, in *ObjectRead, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/storage_bench.StorageBenchWrapper/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageBenchWrapperServer is the server API for StorageBenchWrapper service.
type StorageBenchWrapperServer interface {
	// Performs an upload from a specific object.
	Write(context.Context, *ObjectWrite) (*EmptyResponse, error)
	// Read a specific object.
	Read(context.Context, *ObjectRead) (*EmptyResponse, error)
}

// UnimplementedStorageBenchWrapperServer can be embedded to have forward compatible implementations.
type UnimplementedStorageBenchWrapperServer struct {
}

func (*UnimplementedStorageBenchWrapperServer) Write(ctx context.Context, req *ObjectWrite) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedStorageBenchWrapperServer) Read(ctx context.Context, req *ObjectRead) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func RegisterStorageBenchWrapperServer(s *grpc.Server, srv StorageBenchWrapperServer) {
	s.RegisterService(&_StorageBenchWrapper_serviceDesc, srv)
}

func _StorageBenchWrapper_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectWrite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBenchWrapperServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage_bench.StorageBenchWrapper/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBenchWrapperServer).Write(ctx, req.(*ObjectWrite))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageBenchWrapper_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRead)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageBenchWrapperServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage_bench.StorageBenchWrapper/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageBenchWrapperServer).Read(ctx, req.(*ObjectRead))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageBenchWrapper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage_bench.StorageBenchWrapper",
	HandlerType: (*StorageBenchWrapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _StorageBenchWrapper_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _StorageBenchWrapper_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}