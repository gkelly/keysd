// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fido2.proto

package service

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type CredentialType int32

const (
	UnknownCredential CredentialType = 0
	ES256             CredentialType = -7
	EDDSA             CredentialType = -8
	RS256             CredentialType = -257
)

var CredentialType_name = map[int32]string{
	0:    "UNKNOWN_CREDENTIAL",
	-7:   "ES256",
	-8:   "EDDSA",
	-257: "RS256",
}

var CredentialType_value = map[string]int32{
	"UNKNOWN_CREDENTIAL": 0,
	"ES256":              -7,
	"EDDSA":              -8,
	"RS256":              -257,
}

func (x CredentialType) String() string {
	return proto.EnumName(CredentialType_name, int32(x))
}

func (CredentialType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_05bb79244b7f5be6, []int{0}
}

type DeviceInfo struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	productID            int32    `protobuf:"varint,2,opt,name=productId,proto3" json:"productId,omitempty"`
	vendorID             int32    `protobuf:"varint,3,opt,name=vendorId,proto3" json:"vendorId,omitempty"`
	Manufacturer         string   `protobuf:"bytes,4,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Product              string   `protobuf:"bytes,5,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceInfo) Reset()         { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()    {}
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_05bb79244b7f5be6, []int{0}
}
func (m *DeviceInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeviceInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceInfo.Merge(m, src)
}
func (m *DeviceInfo) XXX_Size() int {
	return m.Size()
}
func (m *DeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceInfo proto.InternalMessageInfo

// Credential ...
type Credential struct {
	ID                   []byte         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthData             []byte         `protobuf:"bytes,2,opt,name=authData,proto3" json:"authData,omitempty"`
	ClientDataHash       []byte         `protobuf:"bytes,3,opt,name=clientDataHash,proto3" json:"clientDataHash,omitempty"`
	Type                 CredentialType `protobuf:"varint,4,opt,name=type,proto3,enum=service.CredentialType" json:"type,omitempty"`
	PubKey               []byte         `protobuf:"bytes,5,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Cert                 []byte         `protobuf:"bytes,6,opt,name=cert,proto3" json:"cert,omitempty"`
	Sig                  []byte         `protobuf:"bytes,7,opt,name=sig,proto3" json:"sig,omitempty"`
	Format               string         `protobuf:"bytes,8,opt,name=format,proto3" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Credential) Reset()         { *m = Credential{} }
func (m *Credential) String() string { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()    {}
func (*Credential) Descriptor() ([]byte, []int) {
	return fileDescriptor_05bb79244b7f5be6, []int{1}
}
func (m *Credential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credential.Merge(m, src)
}
func (m *Credential) XXX_Size() int {
	return m.Size()
}
func (m *Credential) XXX_DiscardUnknown() {
	xxx_messageInfo_Credential.DiscardUnknown(m)
}

var xxx_messageInfo_Credential proto.InternalMessageInfo

type DetectDevicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetectDevicesRequest) Reset()         { *m = DetectDevicesRequest{} }
func (m *DetectDevicesRequest) String() string { return proto.CompactTextString(m) }
func (*DetectDevicesRequest) ProtoMessage()    {}
func (*DetectDevicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_05bb79244b7f5be6, []int{2}
}
func (m *DetectDevicesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DetectDevicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DetectDevicesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DetectDevicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectDevicesRequest.Merge(m, src)
}
func (m *DetectDevicesRequest) XXX_Size() int {
	return m.Size()
}
func (m *DetectDevicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectDevicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetectDevicesRequest proto.InternalMessageInfo

type DetectDevicesResponse struct {
	Detected             []*DeviceInfo `protobuf:"bytes,1,rep,name=detected,proto3" json:"detected,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DetectDevicesResponse) Reset()         { *m = DetectDevicesResponse{} }
func (m *DetectDevicesResponse) String() string { return proto.CompactTextString(m) }
func (*DetectDevicesResponse) ProtoMessage()    {}
func (*DetectDevicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05bb79244b7f5be6, []int{3}
}
func (m *DetectDevicesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DetectDevicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DetectDevicesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DetectDevicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectDevicesResponse.Merge(m, src)
}
func (m *DetectDevicesResponse) XXX_Size() int {
	return m.Size()
}
func (m *DetectDevicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectDevicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetectDevicesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("service.CredentialType", CredentialType_name, CredentialType_value)
	proto.RegisterType((*DeviceInfo)(nil), "service.DeviceInfo")
	proto.RegisterType((*Credential)(nil), "service.Credential")
	proto.RegisterType((*DetectDevicesRequest)(nil), "service.DetectDevicesRequest")
	proto.RegisterType((*DetectDevicesResponse)(nil), "service.DetectDevicesResponse")
}

func init() { proto.RegisterFile("fido2.proto", fileDescriptor_05bb79244b7f5be6) }

var fileDescriptor_05bb79244b7f5be6 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0xf9, 0x9f, 0xa9, 0x1b, 0x85, 0xa5, 0x0d, 0x96, 0x45, 0x9d, 0xc8, 0x07, 0x14,
	0x51, 0x91, 0x4a, 0x41, 0x70, 0x4f, 0xea, 0x48, 0x8d, 0x8a, 0x02, 0xda, 0xb6, 0xe2, 0x88, 0x1c,
	0x7b, 0x93, 0x58, 0xb4, 0x5e, 0x63, 0xaf, 0x8b, 0xf2, 0x06, 0xa8, 0xef, 0xd0, 0x13, 0x3c, 0x02,
	0x5c, 0x78, 0x82, 0x1e, 0x39, 0x72, 0x8a, 0x88, 0x4f, 0x48, 0x5c, 0x38, 0xc2, 0x01, 0x15, 0xed,
	0x26, 0x75, 0x9a, 0x20, 0x72, 0x9a, 0xf9, 0x7d, 0xb3, 0xd9, 0x6f, 0xbf, 0x91, 0x61, 0x63, 0xe8,
	0x3a, 0xac, 0xd5, 0xf4, 0x03, 0xc6, 0x19, 0x2e, 0x84, 0x34, 0x38, 0x77, 0x6d, 0xaa, 0x6d, 0x8d,
	0xd8, 0x88, 0x49, 0xb6, 0x27, 0xaa, 0xb9, 0x6c, 0x7c, 0x42, 0x00, 0x26, 0x15, 0x03, 0x3d, 0x6f,
	0xc8, 0x30, 0x86, 0xac, 0x6f, 0xf1, 0xb1, 0x8a, 0xea, 0xa8, 0x51, 0x22, 0xb2, 0xc6, 0xbb, 0x50,
	0xf2, 0x03, 0xe6, 0x44, 0x36, 0xef, 0x39, 0x6a, 0xba, 0x8e, 0x1a, 0xb9, 0xce, 0x66, 0x3c, 0xad,
	0x25, 0xd0, 0x24, 0x4b, 0x1d, 0x37, 0xa0, 0x78, 0x4e, 0x3d, 0x87, 0x05, 0x3d, 0x47, 0xcd, 0xc8,
	0x59, 0x25, 0x9e, 0xd6, 0x6e, 0x98, 0x49, 0x12, 0x15, 0x1b, 0xa0, 0x9c, 0x59, 0x5e, 0x34, 0xb4,
	0x6c, 0x1e, 0x05, 0x34, 0x50, 0xb3, 0xf2, 0xca, 0x15, 0x86, 0x55, 0x28, 0x2c, 0xfe, 0x5a, 0xcd,
	0x49, 0xf9, 0xa6, 0x35, 0x7e, 0x20, 0x80, 0xfd, 0x80, 0x3a, 0xd4, 0xe3, 0xae, 0x75, 0x8a, 0xab,
	0x90, 0x76, 0x1d, 0xe9, 0x5a, 0xe9, 0xe4, 0xe3, 0x69, 0x2d, 0xdd, 0x33, 0x49, 0xda, 0x75, 0xb0,
	0x06, 0x45, 0x2b, 0xe2, 0x63, 0xd3, 0xe2, 0x96, 0xb4, 0xae, 0x90, 0xa4, 0xc7, 0x0f, 0xa0, 0x6c,
	0x9f, 0xba, 0xd4, 0xe3, 0xa2, 0x3b, 0xb0, 0xc2, 0xb1, 0x34, 0xac, 0x90, 0x35, 0x8a, 0x77, 0x21,
	0xcb, 0x27, 0x3e, 0x95, 0x06, 0xcb, 0xad, 0x7b, 0xcd, 0x45, 0xa0, 0xcd, 0xe5, 0xf5, 0xc7, 0x13,
	0x9f, 0x12, 0x39, 0x84, 0xab, 0x90, 0xf7, 0xa3, 0xc1, 0x21, 0x9d, 0x48, 0xc3, 0x0a, 0x59, 0x74,
	0x22, 0x58, 0x9b, 0x06, 0x5c, 0xcd, 0x4b, 0x2a, 0x6b, 0x5c, 0x81, 0x4c, 0xe8, 0x8e, 0xd4, 0x82,
	0x44, 0xa2, 0x14, 0xa7, 0x87, 0x2c, 0x38, 0xb3, 0xb8, 0x5a, 0x94, 0xcf, 0x5d, 0x74, 0x46, 0x15,
	0xb6, 0x4c, 0xca, 0xa9, 0xcd, 0xe7, 0xab, 0x0a, 0x09, 0x7d, 0x13, 0xd1, 0x90, 0x1b, 0x07, 0xb0,
	0xbd, 0xc6, 0x43, 0x9f, 0x79, 0x21, 0xc5, 0x7b, 0x50, 0x74, 0xa4, 0x40, 0x45, 0x2a, 0x99, 0xc6,
	0x46, 0xeb, 0x6e, 0xe2, 0x7b, 0xb9, 0x6e, 0x92, 0x0c, 0x3d, 0xfc, 0x88, 0xa0, 0xbc, 0xfa, 0x20,
	0xfc, 0x08, 0xf0, 0x49, 0xff, 0xb0, 0xff, 0xfc, 0x65, 0xff, 0xd5, 0x3e, 0xe9, 0x9a, 0xdd, 0xfe,
	0x71, 0xaf, 0xfd, 0xac, 0x92, 0xd2, 0xb6, 0x2f, 0x2e, 0xeb, 0x77, 0x4e, 0xbc, 0xd7, 0x1e, 0x7b,
	0xeb, 0xdd, 0x5a, 0xc1, 0x0e, 0xe4, 0xba, 0x47, 0xad, 0x27, 0x4f, 0x2b, 0xbf, 0xaf, 0x17, 0x3f,
	0xa4, 0x95, 0x2e, 0x2e, 0xeb, 0x73, 0x2a, 0x65, 0xd3, 0x3c, 0x6a, 0x57, 0x7e, 0xad, 0xcb, 0x82,
	0x0a, 0x99, 0xc8, 0xd3, 0xd7, 0x7f, 0x56, 0x65, 0x49, 0xb5, 0xea, 0xbb, 0xf7, 0x7a, 0xea, 0xf3,
	0x07, 0x7d, 0xcd, 0x63, 0x6b, 0x00, 0xe5, 0x76, 0xc4, 0xc7, 0x82, 0xd8, 0x16, 0x67, 0x41, 0x88,
	0x5f, 0xc0, 0xe6, 0x4a, 0x24, 0x78, 0xe7, 0xd6, 0xc3, 0xff, 0x8d, 0x50, 0xd3, 0xff, 0x27, 0xcf,
	0x93, 0x34, 0x52, 0x9d, 0xfb, 0x57, 0x33, 0x3d, 0xf5, 0x75, 0xa6, 0xa3, 0x9f, 0x33, 0x1d, 0x5d,
	0xc5, 0x3a, 0xfa, 0x12, 0xeb, 0xe8, 0x5b, 0xac, 0xa3, 0xef, 0xb1, 0x8e, 0x06, 0x79, 0xf9, 0x1d,
	0x3d, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xac, 0xd8, 0xdd, 0x75, 0x03, 0x00, 0x00,
}

func (this *DeviceInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&service.DeviceInfo{")
	s = append(s, "Path: "+fmt.Sprintf("%#v", this.Path)+",\n")
	s = append(s, "productID: "+fmt.Sprintf("%#v", this.productID)+",\n")
	s = append(s, "vendorID: "+fmt.Sprintf("%#v", this.vendorID)+",\n")
	s = append(s, "Manufacturer: "+fmt.Sprintf("%#v", this.Manufacturer)+",\n")
	s = append(s, "Product: "+fmt.Sprintf("%#v", this.Product)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Credential) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&service.Credential{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "AuthData: "+fmt.Sprintf("%#v", this.AuthData)+",\n")
	s = append(s, "ClientDataHash: "+fmt.Sprintf("%#v", this.ClientDataHash)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "PubKey: "+fmt.Sprintf("%#v", this.PubKey)+",\n")
	s = append(s, "Cert: "+fmt.Sprintf("%#v", this.Cert)+",\n")
	s = append(s, "Sig: "+fmt.Sprintf("%#v", this.Sig)+",\n")
	s = append(s, "Format: "+fmt.Sprintf("%#v", this.Format)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DetectDevicesRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&service.DetectDevicesRequest{")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DetectDevicesResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&service.DetectDevicesResponse{")
	if this.Detected != nil {
		s = append(s, "Detected: "+fmt.Sprintf("%#v", this.Detected)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFido2(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticatorsClient is the client API for Authenticators service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticatorsClient interface {
	DetectDevices(ctx context.Context, in *DetectDevicesRequest, opts ...grpc.CallOption) (*DetectDevicesResponse, error)
}

type authenticatorsClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticatorsClient(cc *grpc.ClientConn) AuthenticatorsClient {
	return &authenticatorsClient{cc}
}

func (c *authenticatorsClient) DetectDevices(ctx context.Context, in *DetectDevicesRequest, opts ...grpc.CallOption) (*DetectDevicesResponse, error) {
	out := new(DetectDevicesResponse)
	err := c.cc.Invoke(ctx, "/service.Authenticators/DetectDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticatorsServer is the server API for Authenticators service.
type AuthenticatorsServer interface {
	DetectDevices(context.Context, *DetectDevicesRequest) (*DetectDevicesResponse, error)
}

// UnimplementedAuthenticatorsServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticatorsServer struct {
}

func (*UnimplementedAuthenticatorsServer) DetectDevices(ctx context.Context, req *DetectDevicesRequest) (*DetectDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectDevices not implemented")
}

func RegisterAuthenticatorsServer(s *grpc.Server, srv AuthenticatorsServer) {
	s.RegisterService(&_Authenticators_serviceDesc, srv)
}

func _Authenticators_DetectDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorsServer).DetectDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Authenticators/DetectDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorsServer).DetectDevices(ctx, req.(*DetectDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authenticators_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Authenticators",
	HandlerType: (*AuthenticatorsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DetectDevices",
			Handler:    _Authenticators_DetectDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fido2.proto",
}

func (m *DeviceInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeviceInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeviceInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Product) > 0 {
		i -= len(m.Product)
		copy(dAtA[i:], m.Product)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.Product)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Manufacturer) > 0 {
		i -= len(m.Manufacturer)
		copy(dAtA[i:], m.Manufacturer)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.Manufacturer)))
		i--
		dAtA[i] = 0x22
	}
	if m.vendorID != 0 {
		i = encodeVarintFido2(dAtA, i, uint64(m.vendorID))
		i--
		dAtA[i] = 0x18
	}
	if m.productID != 0 {
		i = encodeVarintFido2(dAtA, i, uint64(m.productID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Credential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Credential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Format) > 0 {
		i -= len(m.Format)
		copy(dAtA[i:], m.Format)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.Format)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Cert) > 0 {
		i -= len(m.Cert)
		copy(dAtA[i:], m.Cert)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.Cert)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Type != 0 {
		i = encodeVarintFido2(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ClientDataHash) > 0 {
		i -= len(m.ClientDataHash)
		copy(dAtA[i:], m.ClientDataHash)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.ClientDataHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AuthData) > 0 {
		i -= len(m.AuthData)
		copy(dAtA[i:], m.AuthData)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.AuthData)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintFido2(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DetectDevicesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DetectDevicesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DetectDevicesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *DetectDevicesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DetectDevicesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DetectDevicesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Detected) > 0 {
		for iNdEx := len(m.Detected) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Detected[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFido2(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintFido2(dAtA []byte, offset int, v uint64) int {
	offset -= sovFido2(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeviceInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	if m.productID != 0 {
		n += 1 + sovFido2(uint64(m.productID))
	}
	if m.vendorID != 0 {
		n += 1 + sovFido2(uint64(m.vendorID))
	}
	l = len(m.Manufacturer)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	l = len(m.Product)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Credential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	l = len(m.AuthData)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	l = len(m.ClientDataHash)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovFido2(uint64(m.Type))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	l = len(m.Cert)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	l = len(m.Format)
	if l > 0 {
		n += 1 + l + sovFido2(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DetectDevicesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DetectDevicesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Detected) > 0 {
		for _, e := range m.Detected {
			l = e.Size()
			n += 1 + l + sovFido2(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFido2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFido2(x uint64) (n int) {
	return sovFido2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeviceInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFido2
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeviceInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeviceInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field productID", wireType)
			}
			m.productID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.productID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field vendorID", wireType)
			}
			m.vendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.vendorID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Manufacturer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Manufacturer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Product", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Product = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFido2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFido2
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFido2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Credential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFido2
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Credential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthData = append(m.AuthData[:0], dAtA[iNdEx:postIndex]...)
			if m.AuthData == nil {
				m.AuthData = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientDataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientDataHash = append(m.ClientDataHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ClientDataHash == nil {
				m.ClientDataHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= CredentialType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cert = append(m.Cert[:0], dAtA[iNdEx:postIndex]...)
			if m.Cert == nil {
				m.Cert = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Format = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFido2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFido2
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFido2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DetectDevicesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFido2
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DetectDevicesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DetectDevicesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFido2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFido2
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFido2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DetectDevicesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFido2
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DetectDevicesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DetectDevicesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Detected", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFido2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFido2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Detected = append(m.Detected, &DeviceInfo{})
			if err := m.Detected[len(m.Detected)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFido2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFido2
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFido2
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFido2(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFido2
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFido2
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFido2
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFido2
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFido2
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFido2        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFido2          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFido2 = fmt.Errorf("proto: unexpected end of group")
)
