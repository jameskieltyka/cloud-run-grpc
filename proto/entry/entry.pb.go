// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entry.proto

package au_com_kieltyka_james_entry

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EntryDataParameter struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryDataParameter) Reset()         { *m = EntryDataParameter{} }
func (m *EntryDataParameter) String() string { return proto.CompactTextString(m) }
func (*EntryDataParameter) ProtoMessage()    {}
func (*EntryDataParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_entry_80a2749a2daea5e2, []int{0}
}
func (m *EntryDataParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryDataParameter.Unmarshal(m, b)
}
func (m *EntryDataParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryDataParameter.Marshal(b, m, deterministic)
}
func (dst *EntryDataParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryDataParameter.Merge(dst, src)
}
func (m *EntryDataParameter) XXX_Size() int {
	return xxx_messageInfo_EntryDataParameter.Size(m)
}
func (m *EntryDataParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryDataParameter.DiscardUnknown(m)
}

var xxx_messageInfo_EntryDataParameter proto.InternalMessageInfo

func (m *EntryDataParameter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EntryDataParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EntryResult struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryResult) Reset()         { *m = EntryResult{} }
func (m *EntryResult) String() string { return proto.CompactTextString(m) }
func (*EntryResult) ProtoMessage()    {}
func (*EntryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_entry_80a2749a2daea5e2, []int{1}
}
func (m *EntryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryResult.Unmarshal(m, b)
}
func (m *EntryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryResult.Marshal(b, m, deterministic)
}
func (dst *EntryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryResult.Merge(dst, src)
}
func (m *EntryResult) XXX_Size() int {
	return xxx_messageInfo_EntryResult.Size(m)
}
func (m *EntryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryResult.DiscardUnknown(m)
}

var xxx_messageInfo_EntryResult proto.InternalMessageInfo

func (m *EntryResult) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EntryDataParameter)(nil), "au.com.kieltyka.james.entry.EntryDataParameter")
	proto.RegisterType((*EntryResult)(nil), "au.com.kieltyka.james.entry.EntryResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EntryServiceClient is the client API for EntryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EntryServiceClient interface {
	CallEntry(ctx context.Context, in *EntryDataParameter, opts ...grpc.CallOption) (*EntryResult, error)
}

type entryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEntryServiceClient(cc *grpc.ClientConn) EntryServiceClient {
	return &entryServiceClient{cc}
}

func (c *entryServiceClient) CallEntry(ctx context.Context, in *EntryDataParameter, opts ...grpc.CallOption) (*EntryResult, error) {
	out := new(EntryResult)
	err := c.cc.Invoke(ctx, "/au.com.kieltyka.james.entry.EntryService/CallEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntryServiceServer is the server API for EntryService service.
type EntryServiceServer interface {
	CallEntry(context.Context, *EntryDataParameter) (*EntryResult, error)
}

func RegisterEntryServiceServer(s *grpc.Server, srv EntryServiceServer) {
	s.RegisterService(&_EntryService_serviceDesc, srv)
}

func _EntryService_CallEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntryDataParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntryServiceServer).CallEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/au.com.kieltyka.james.entry.EntryService/CallEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntryServiceServer).CallEntry(ctx, req.(*EntryDataParameter))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "au.com.kieltyka.james.entry.EntryService",
	HandlerType: (*EntryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallEntry",
			Handler:    _EntryService_CallEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entry.proto",
}

func init() { proto.RegisterFile("entry.proto", fileDescriptor_entry_80a2749a2daea5e2) }

var fileDescriptor_entry_80a2749a2daea5e2 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x85, 0x6d, 0x11, 0xa5, 0x57, 0x71, 0xb8, 0xa9, 0xe8, 0x22, 0x5d, 0xec, 0x14, 0x41, 0x17,
	0x77, 0x75, 0x97, 0xfa, 0x0b, 0xce, 0xf6, 0xd0, 0xd8, 0xa4, 0x95, 0x24, 0x15, 0xfb, 0xef, 0xc5,
	0xa8, 0x93, 0xa0, 0xdb, 0xbb, 0xc7, 0xfb, 0xe0, 0xe3, 0x20, 0xe6, 0xda, 0x99, 0x4e, 0x5c, 0x4d,
	0xe3, 0x1a, 0x9c, 0x52, 0x2b, 0x8a, 0x46, 0x8b, 0x4a, 0xb2, 0x72, 0x5d, 0x45, 0xe2, 0x42, 0x9a,
	0xad, 0xf0, 0x93, 0x74, 0x0d, 0xb8, 0x7b, 0x86, 0x2d, 0x39, 0xda, 0x93, 0x21, 0xcd, 0x8e, 0x0d,
	0x8e, 0x21, 0x94, 0x65, 0x12, 0xcc, 0x82, 0x2c, 0xca, 0x43, 0x59, 0x22, 0x42, 0xbf, 0x26, 0xcd,
	0x49, 0xe8, 0x1b, 0x9f, 0xd3, 0x39, 0xc4, 0x9e, 0xcc, 0xd9, 0xb6, 0xca, 0x61, 0x02, 0x43, 0xcd,
	0xd6, 0xd2, 0x89, 0xdf, 0xdc, 0xe7, 0x5c, 0xde, 0x61, 0xe4, 0x87, 0x07, 0x36, 0x37, 0x59, 0x30,
	0x9e, 0x21, 0xda, 0x90, 0x52, 0xbe, 0xc3, 0x85, 0xf8, 0x61, 0x27, 0xbe, 0xd5, 0x26, 0xd9, 0x7f,
	0xe0, 0x65, 0x94, 0xf6, 0x8e, 0x03, 0xff, 0x80, 0xd5, 0x23, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xfb,
	0x48, 0x5b, 0x0f, 0x01, 0x00, 0x00,
}
