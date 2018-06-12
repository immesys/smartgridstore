// Code generated by protoc-gen-go. DO NOT EDIT.
// source: adminapi.proto

/*
Package adminapi is a generated protocol buffer package.

It is generated from these files:
	adminapi.proto

It has these top-level messages:
	ManifestAddParams
	ManifestAddResponse
	MetaKeyValue
	ManifestDelParams
	ManifestDelResponse
	ManifestDelPrefixParams
	ManifestDelPrefixResponse
	ManifestLsDevsParams
	ManifestLsDevsResponse
	ManifestDevice
	Status
*/
package adminapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type ManifestAddParams struct {
	Deviceid string          `protobuf:"bytes,1,opt,name=deviceid" json:"deviceid,omitempty"`
	Metadata []*MetaKeyValue `protobuf:"bytes,2,rep,name=metadata" json:"metadata,omitempty"`
}

func (m *ManifestAddParams) Reset()                    { *m = ManifestAddParams{} }
func (m *ManifestAddParams) String() string            { return proto.CompactTextString(m) }
func (*ManifestAddParams) ProtoMessage()               {}
func (*ManifestAddParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ManifestAddParams) GetDeviceid() string {
	if m != nil {
		return m.Deviceid
	}
	return ""
}

func (m *ManifestAddParams) GetMetadata() []*MetaKeyValue {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ManifestAddResponse struct {
	Stat *Status `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
}

func (m *ManifestAddResponse) Reset()                    { *m = ManifestAddResponse{} }
func (m *ManifestAddResponse) String() string            { return proto.CompactTextString(m) }
func (*ManifestAddResponse) ProtoMessage()               {}
func (*ManifestAddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ManifestAddResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

type MetaKeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *MetaKeyValue) Reset()                    { *m = MetaKeyValue{} }
func (m *MetaKeyValue) String() string            { return proto.CompactTextString(m) }
func (*MetaKeyValue) ProtoMessage()               {}
func (*MetaKeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MetaKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MetaKeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ManifestDelParams struct {
	Deviceid string `protobuf:"bytes,1,opt,name=deviceid" json:"deviceid,omitempty"`
}

func (m *ManifestDelParams) Reset()                    { *m = ManifestDelParams{} }
func (m *ManifestDelParams) String() string            { return proto.CompactTextString(m) }
func (*ManifestDelParams) ProtoMessage()               {}
func (*ManifestDelParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ManifestDelParams) GetDeviceid() string {
	if m != nil {
		return m.Deviceid
	}
	return ""
}

type ManifestDelResponse struct {
	Stat *Status `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
}

func (m *ManifestDelResponse) Reset()                    { *m = ManifestDelResponse{} }
func (m *ManifestDelResponse) String() string            { return proto.CompactTextString(m) }
func (*ManifestDelResponse) ProtoMessage()               {}
func (*ManifestDelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ManifestDelResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

type ManifestDelPrefixParams struct {
	Deviceidprefix string `protobuf:"bytes,1,opt,name=deviceidprefix" json:"deviceidprefix,omitempty"`
}

func (m *ManifestDelPrefixParams) Reset()                    { *m = ManifestDelPrefixParams{} }
func (m *ManifestDelPrefixParams) String() string            { return proto.CompactTextString(m) }
func (*ManifestDelPrefixParams) ProtoMessage()               {}
func (*ManifestDelPrefixParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ManifestDelPrefixParams) GetDeviceidprefix() string {
	if m != nil {
		return m.Deviceidprefix
	}
	return ""
}

type ManifestDelPrefixResponse struct {
	Stat       *Status `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	Numdeleted uint32  `protobuf:"varint,2,opt,name=numdeleted" json:"numdeleted,omitempty"`
}

func (m *ManifestDelPrefixResponse) Reset()                    { *m = ManifestDelPrefixResponse{} }
func (m *ManifestDelPrefixResponse) String() string            { return proto.CompactTextString(m) }
func (*ManifestDelPrefixResponse) ProtoMessage()               {}
func (*ManifestDelPrefixResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ManifestDelPrefixResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *ManifestDelPrefixResponse) GetNumdeleted() uint32 {
	if m != nil {
		return m.Numdeleted
	}
	return 0
}

type ManifestLsDevsParams struct {
	Deviceidprefix string `protobuf:"bytes,1,opt,name=deviceidprefix" json:"deviceidprefix,omitempty"`
}

func (m *ManifestLsDevsParams) Reset()                    { *m = ManifestLsDevsParams{} }
func (m *ManifestLsDevsParams) String() string            { return proto.CompactTextString(m) }
func (*ManifestLsDevsParams) ProtoMessage()               {}
func (*ManifestLsDevsParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ManifestLsDevsParams) GetDeviceidprefix() string {
	if m != nil {
		return m.Deviceidprefix
	}
	return ""
}

type ManifestLsDevsResponse struct {
	Stat    *Status           `protobuf:"bytes,1,opt,name=stat" json:"stat,omitempty"`
	Devices []*ManifestDevice `protobuf:"bytes,2,rep,name=devices" json:"devices,omitempty"`
}

func (m *ManifestLsDevsResponse) Reset()                    { *m = ManifestLsDevsResponse{} }
func (m *ManifestLsDevsResponse) String() string            { return proto.CompactTextString(m) }
func (*ManifestLsDevsResponse) ProtoMessage()               {}
func (*ManifestLsDevsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ManifestLsDevsResponse) GetStat() *Status {
	if m != nil {
		return m.Stat
	}
	return nil
}

func (m *ManifestLsDevsResponse) GetDevices() []*ManifestDevice {
	if m != nil {
		return m.Devices
	}
	return nil
}

type ManifestDevice struct {
	Deviceid string          `protobuf:"bytes,1,opt,name=deviceid" json:"deviceid,omitempty"`
	Metadata []*MetaKeyValue `protobuf:"bytes,2,rep,name=metadata" json:"metadata,omitempty"`
}

func (m *ManifestDevice) Reset()                    { *m = ManifestDevice{} }
func (m *ManifestDevice) String() string            { return proto.CompactTextString(m) }
func (*ManifestDevice) ProtoMessage()               {}
func (*ManifestDevice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ManifestDevice) GetDeviceid() string {
	if m != nil {
		return m.Deviceid
	}
	return ""
}

func (m *ManifestDevice) GetMetadata() []*MetaKeyValue {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Status struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Status) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ManifestAddParams)(nil), "adminapi.ManifestAddParams")
	proto.RegisterType((*ManifestAddResponse)(nil), "adminapi.ManifestAddResponse")
	proto.RegisterType((*MetaKeyValue)(nil), "adminapi.MetaKeyValue")
	proto.RegisterType((*ManifestDelParams)(nil), "adminapi.ManifestDelParams")
	proto.RegisterType((*ManifestDelResponse)(nil), "adminapi.ManifestDelResponse")
	proto.RegisterType((*ManifestDelPrefixParams)(nil), "adminapi.ManifestDelPrefixParams")
	proto.RegisterType((*ManifestDelPrefixResponse)(nil), "adminapi.ManifestDelPrefixResponse")
	proto.RegisterType((*ManifestLsDevsParams)(nil), "adminapi.ManifestLsDevsParams")
	proto.RegisterType((*ManifestLsDevsResponse)(nil), "adminapi.ManifestLsDevsResponse")
	proto.RegisterType((*ManifestDevice)(nil), "adminapi.ManifestDevice")
	proto.RegisterType((*Status)(nil), "adminapi.Status")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BTrDBAdmin service

type BTrDBAdminClient interface {
	// Requires Manifest capability
	ManifestAdd(ctx context.Context, in *ManifestAddParams, opts ...grpc.CallOption) (*ManifestAddResponse, error)
	ManifestDel(ctx context.Context, in *ManifestDelParams, opts ...grpc.CallOption) (*ManifestDelResponse, error)
	ManifestDelPrefix(ctx context.Context, in *ManifestDelPrefixParams, opts ...grpc.CallOption) (*ManifestDelPrefixResponse, error)
	ManifestLsDevs(ctx context.Context, in *ManifestLsDevsParams, opts ...grpc.CallOption) (*ManifestLsDevsResponse, error)
}

type bTrDBAdminClient struct {
	cc *grpc.ClientConn
}

func NewBTrDBAdminClient(cc *grpc.ClientConn) BTrDBAdminClient {
	return &bTrDBAdminClient{cc}
}

func (c *bTrDBAdminClient) ManifestAdd(ctx context.Context, in *ManifestAddParams, opts ...grpc.CallOption) (*ManifestAddResponse, error) {
	out := new(ManifestAddResponse)
	err := grpc.Invoke(ctx, "/adminapi.BTrDBAdmin/ManifestAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBAdminClient) ManifestDel(ctx context.Context, in *ManifestDelParams, opts ...grpc.CallOption) (*ManifestDelResponse, error) {
	out := new(ManifestDelResponse)
	err := grpc.Invoke(ctx, "/adminapi.BTrDBAdmin/ManifestDel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBAdminClient) ManifestDelPrefix(ctx context.Context, in *ManifestDelPrefixParams, opts ...grpc.CallOption) (*ManifestDelPrefixResponse, error) {
	out := new(ManifestDelPrefixResponse)
	err := grpc.Invoke(ctx, "/adminapi.BTrDBAdmin/ManifestDelPrefix", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bTrDBAdminClient) ManifestLsDevs(ctx context.Context, in *ManifestLsDevsParams, opts ...grpc.CallOption) (*ManifestLsDevsResponse, error) {
	out := new(ManifestLsDevsResponse)
	err := grpc.Invoke(ctx, "/adminapi.BTrDBAdmin/ManifestLsDevs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BTrDBAdmin service

type BTrDBAdminServer interface {
	// Requires Manifest capability
	ManifestAdd(context.Context, *ManifestAddParams) (*ManifestAddResponse, error)
	ManifestDel(context.Context, *ManifestDelParams) (*ManifestDelResponse, error)
	ManifestDelPrefix(context.Context, *ManifestDelPrefixParams) (*ManifestDelPrefixResponse, error)
	ManifestLsDevs(context.Context, *ManifestLsDevsParams) (*ManifestLsDevsResponse, error)
}

func RegisterBTrDBAdminServer(s *grpc.Server, srv BTrDBAdminServer) {
	s.RegisterService(&_BTrDBAdmin_serviceDesc, srv)
}

func _BTrDBAdmin_ManifestAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestAddParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).ManifestAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/ManifestAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).ManifestAdd(ctx, req.(*ManifestAddParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDBAdmin_ManifestDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestDelParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).ManifestDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/ManifestDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).ManifestDel(ctx, req.(*ManifestDelParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDBAdmin_ManifestDelPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestDelPrefixParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).ManifestDelPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/ManifestDelPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).ManifestDelPrefix(ctx, req.(*ManifestDelPrefixParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BTrDBAdmin_ManifestLsDevs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManifestLsDevsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BTrDBAdminServer).ManifestLsDevs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/adminapi.BTrDBAdmin/ManifestLsDevs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BTrDBAdminServer).ManifestLsDevs(ctx, req.(*ManifestLsDevsParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _BTrDBAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "adminapi.BTrDBAdmin",
	HandlerType: (*BTrDBAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ManifestAdd",
			Handler:    _BTrDBAdmin_ManifestAdd_Handler,
		},
		{
			MethodName: "ManifestDel",
			Handler:    _BTrDBAdmin_ManifestDel_Handler,
		},
		{
			MethodName: "ManifestDelPrefix",
			Handler:    _BTrDBAdmin_ManifestDelPrefix_Handler,
		},
		{
			MethodName: "ManifestLsDevs",
			Handler:    _BTrDBAdmin_ManifestLsDevs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "adminapi.proto",
}

func init() { proto.RegisterFile("adminapi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x95, 0xa4, 0x94, 0x30, 0xa1, 0xa1, 0x2c, 0xa1, 0x0d, 0x86, 0x56, 0x61, 0x41, 0xa8,
	0xea, 0x21, 0x96, 0x0c, 0xe2, 0x50, 0x24, 0xa4, 0x54, 0xbe, 0x41, 0x25, 0x64, 0x10, 0x67, 0x86,
	0xec, 0x34, 0xb2, 0xf0, 0x17, 0xde, 0x8d, 0x45, 0x0f, 0x5c, 0x78, 0x05, 0x5e, 0x8b, 0x1b, 0xaf,
	0xc0, 0x83, 0x20, 0xaf, 0x3f, 0xb2, 0x49, 0x0c, 0xc8, 0x12, 0xb7, 0xdd, 0x99, 0xf1, 0xfc, 0xfe,
	0xb3, 0xfb, 0x5f, 0xc3, 0x10, 0x45, 0xe8, 0x47, 0x98, 0xf8, 0xd3, 0x24, 0x8d, 0x55, 0xcc, 0xfa,
	0xd5, 0xde, 0x7a, 0xb0, 0x88, 0xe3, 0x45, 0x40, 0x36, 0x26, 0xbe, 0x8d, 0x51, 0x14, 0x2b, 0x54,
	0x7e, 0x1c, 0xc9, 0xa2, 0x8e, 0xcf, 0xe1, 0xf6, 0x05, 0x46, 0xfe, 0x25, 0x49, 0x35, 0x13, 0xe2,
	0x0d, 0xa6, 0x18, 0x4a, 0x66, 0x41, 0x5f, 0x50, 0xe6, 0xcf, 0xc9, 0x17, 0xe3, 0xce, 0xa4, 0x73,
	0x72, 0xc3, 0xab, 0xf7, 0xcc, 0x81, 0x7e, 0x48, 0x0a, 0x05, 0x2a, 0x1c, 0x77, 0x27, 0xbd, 0x93,
	0x81, 0x73, 0x30, 0xad, 0xd9, 0x17, 0xa4, 0xf0, 0x15, 0x5d, 0xbd, 0xc7, 0x60, 0x49, 0x5e, 0x5d,
	0xc7, 0x5f, 0xc0, 0x1d, 0x03, 0xe2, 0x91, 0x4c, 0xe2, 0x48, 0x12, 0x7b, 0x0c, 0x3b, 0x52, 0xa1,
	0xd2, 0x88, 0x81, 0xb3, 0xbf, 0x6a, 0xf3, 0x56, 0xa1, 0x5a, 0x4a, 0x4f, 0x67, 0xf9, 0x73, 0xb8,
	0x69, 0xb6, 0x65, 0xfb, 0xd0, 0xfb, 0x44, 0x57, 0xa5, 0xae, 0x7c, 0xc9, 0x46, 0x70, 0x2d, 0xcb,
	0x53, 0xe3, 0xae, 0x8e, 0x15, 0x1b, 0x6e, 0xaf, 0x26, 0x73, 0x29, 0xf8, 0xf7, 0x64, 0xa6, 0x4a,
	0x97, 0x82, 0x96, 0x2a, 0x67, 0x70, 0x68, 0xd2, 0x52, 0xba, 0xf4, 0xbf, 0x94, 0xcc, 0x27, 0x30,
	0xac, 0x18, 0x89, 0x8e, 0x97, 0xe4, 0x8d, 0x28, 0x47, 0xb8, 0xb7, 0xd5, 0xa2, 0x9d, 0x0a, 0x76,
	0x0c, 0x10, 0x2d, 0x43, 0x41, 0x01, 0x29, 0x12, 0xfa, 0x38, 0xf6, 0x3c, 0x23, 0xc2, 0x5f, 0xc2,
	0xa8, 0x42, 0xbc, 0x96, 0x2e, 0x65, 0xb2, 0xa5, 0xc4, 0x14, 0x0e, 0xd6, 0xbf, 0x6f, 0xa9, 0xcf,
	0x81, 0xeb, 0x45, 0x47, 0x59, 0x7a, 0x67, 0x6c, 0x78, 0xa7, 0x9e, 0x3d, 0x2f, 0xf0, 0xaa, 0x42,
	0xfe, 0x01, 0x86, 0xeb, 0xa9, 0xff, 0x6e, 0xcf, 0x29, 0xec, 0x16, 0x2a, 0x19, 0x83, 0x9d, 0x79,
	0x2c, 0x48, 0x77, 0xdd, 0xf3, 0xf4, 0x3a, 0xf7, 0x5b, 0x28, 0x17, 0xa5, 0xb7, 0xf2, 0xa5, 0xf3,
	0xa3, 0x07, 0x70, 0xfe, 0x2e, 0x75, 0xcf, 0x67, 0x79, 0x63, 0x46, 0x30, 0x30, 0xdc, 0xcd, 0xee,
	0x6f, 0x8f, 0x54, 0xbf, 0x2c, 0xeb, 0xa8, 0x31, 0x59, 0x9d, 0x22, 0xb7, 0xbe, 0xfd, 0xfc, 0xf5,
	0xbd, 0x3b, 0xe2, 0xb7, 0xec, 0xec, 0x99, 0x1d, 0x96, 0x05, 0x28, 0xc4, 0x59, 0xe7, 0xd4, 0xc4,
	0xb8, 0x14, 0x34, 0x61, 0x6a, 0x9b, 0x37, 0x61, 0x0c, 0x4b, 0x37, 0x63, 0x04, 0x05, 0x39, 0xe6,
	0xeb, 0xfa, 0xb3, 0xd1, 0xf7, 0xce, 0x1e, 0x36, 0xc3, 0x0c, 0x97, 0x5b, 0x8f, 0xfe, 0x52, 0x52,
	0x83, 0x27, 0x1a, 0x6c, 0xf1, 0xbb, 0x1b, 0xe0, 0xc2, 0x5e, 0x39, 0xfe, 0xf3, 0xea, 0xb6, 0x0b,
	0x87, 0xb1, 0xe3, 0xed, 0xc6, 0xa6, 0x77, 0xad, 0xc9, 0x9f, 0xf2, 0x35, 0xf5, 0x48, 0x53, 0x0f,
	0x39, 0x33, 0xa9, 0x81, 0x14, 0x94, 0xc9, 0xb3, 0xce, 0xe9, 0xc7, 0x5d, 0xfd, 0x27, 0x7c, 0xfa,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xfa, 0x2d, 0xf0, 0x43, 0x05, 0x00, 0x00,
}
