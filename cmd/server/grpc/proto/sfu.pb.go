// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/server/grpc/proto/sfu.proto

package proto

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

type SignalRequest struct {
	// Types that are valid to be assigned to Payload:
	//	*SignalRequest_Join
	//	*SignalRequest_Negotiate
	//	*SignalRequest_Trickle
	Payload              isSignalRequest_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SignalRequest) Reset()         { *m = SignalRequest{} }
func (m *SignalRequest) String() string { return proto.CompactTextString(m) }
func (*SignalRequest) ProtoMessage()    {}
func (*SignalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca80ff2c9b7a4e60, []int{0}
}

func (m *SignalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalRequest.Unmarshal(m, b)
}
func (m *SignalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalRequest.Marshal(b, m, deterministic)
}
func (m *SignalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalRequest.Merge(m, src)
}
func (m *SignalRequest) XXX_Size() int {
	return xxx_messageInfo_SignalRequest.Size(m)
}
func (m *SignalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignalRequest proto.InternalMessageInfo

type isSignalRequest_Payload interface {
	isSignalRequest_Payload()
}

type SignalRequest_Join struct {
	Join *JoinRequest `protobuf:"bytes,1,opt,name=join,proto3,oneof"`
}

type SignalRequest_Negotiate struct {
	Negotiate *SessionDescription `protobuf:"bytes,2,opt,name=negotiate,proto3,oneof"`
}

type SignalRequest_Trickle struct {
	Trickle *Trickle `protobuf:"bytes,3,opt,name=trickle,proto3,oneof"`
}

func (*SignalRequest_Join) isSignalRequest_Payload() {}

func (*SignalRequest_Negotiate) isSignalRequest_Payload() {}

func (*SignalRequest_Trickle) isSignalRequest_Payload() {}

func (m *SignalRequest) GetPayload() isSignalRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SignalRequest) GetJoin() *JoinRequest {
	if x, ok := m.GetPayload().(*SignalRequest_Join); ok {
		return x.Join
	}
	return nil
}

func (m *SignalRequest) GetNegotiate() *SessionDescription {
	if x, ok := m.GetPayload().(*SignalRequest_Negotiate); ok {
		return x.Negotiate
	}
	return nil
}

func (m *SignalRequest) GetTrickle() *Trickle {
	if x, ok := m.GetPayload().(*SignalRequest_Trickle); ok {
		return x.Trickle
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignalRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignalRequest_Join)(nil),
		(*SignalRequest_Negotiate)(nil),
		(*SignalRequest_Trickle)(nil),
	}
}

type SignalReply struct {
	// Types that are valid to be assigned to Payload:
	//	*SignalReply_Join
	//	*SignalReply_Negotiate
	//	*SignalReply_Trickle
	Payload              isSignalReply_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SignalReply) Reset()         { *m = SignalReply{} }
func (m *SignalReply) String() string { return proto.CompactTextString(m) }
func (*SignalReply) ProtoMessage()    {}
func (*SignalReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca80ff2c9b7a4e60, []int{1}
}

func (m *SignalReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalReply.Unmarshal(m, b)
}
func (m *SignalReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalReply.Marshal(b, m, deterministic)
}
func (m *SignalReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalReply.Merge(m, src)
}
func (m *SignalReply) XXX_Size() int {
	return xxx_messageInfo_SignalReply.Size(m)
}
func (m *SignalReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalReply.DiscardUnknown(m)
}

var xxx_messageInfo_SignalReply proto.InternalMessageInfo

type isSignalReply_Payload interface {
	isSignalReply_Payload()
}

type SignalReply_Join struct {
	Join *JoinReply `protobuf:"bytes,1,opt,name=join,proto3,oneof"`
}

type SignalReply_Negotiate struct {
	Negotiate *SessionDescription `protobuf:"bytes,2,opt,name=negotiate,proto3,oneof"`
}

type SignalReply_Trickle struct {
	Trickle *Trickle `protobuf:"bytes,3,opt,name=trickle,proto3,oneof"`
}

func (*SignalReply_Join) isSignalReply_Payload() {}

func (*SignalReply_Negotiate) isSignalReply_Payload() {}

func (*SignalReply_Trickle) isSignalReply_Payload() {}

func (m *SignalReply) GetPayload() isSignalReply_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SignalReply) GetJoin() *JoinReply {
	if x, ok := m.GetPayload().(*SignalReply_Join); ok {
		return x.Join
	}
	return nil
}

func (m *SignalReply) GetNegotiate() *SessionDescription {
	if x, ok := m.GetPayload().(*SignalReply_Negotiate); ok {
		return x.Negotiate
	}
	return nil
}

func (m *SignalReply) GetTrickle() *Trickle {
	if x, ok := m.GetPayload().(*SignalReply_Trickle); ok {
		return x.Trickle
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignalReply) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignalReply_Join)(nil),
		(*SignalReply_Negotiate)(nil),
		(*SignalReply_Trickle)(nil),
	}
}

type JoinRequest struct {
	Rid                  string              `protobuf:"bytes,1,opt,name=rid,proto3" json:"rid,omitempty"`
	Offer                *SessionDescription `protobuf:"bytes,2,opt,name=offer,proto3" json:"offer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca80ff2c9b7a4e60, []int{2}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetRid() string {
	if m != nil {
		return m.Rid
	}
	return ""
}

func (m *JoinRequest) GetOffer() *SessionDescription {
	if m != nil {
		return m.Offer
	}
	return nil
}

type JoinReply struct {
	Pid                  string              `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Answer               *SessionDescription `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *JoinReply) Reset()         { *m = JoinReply{} }
func (m *JoinReply) String() string { return proto.CompactTextString(m) }
func (*JoinReply) ProtoMessage()    {}
func (*JoinReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca80ff2c9b7a4e60, []int{3}
}

func (m *JoinReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinReply.Unmarshal(m, b)
}
func (m *JoinReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinReply.Marshal(b, m, deterministic)
}
func (m *JoinReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinReply.Merge(m, src)
}
func (m *JoinReply) XXX_Size() int {
	return xxx_messageInfo_JoinReply.Size(m)
}
func (m *JoinReply) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinReply.DiscardUnknown(m)
}

var xxx_messageInfo_JoinReply proto.InternalMessageInfo

func (m *JoinReply) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *JoinReply) GetAnswer() *SessionDescription {
	if m != nil {
		return m.Answer
	}
	return nil
}

type Trickle struct {
	Init                 string   `protobuf:"bytes,1,opt,name=init,proto3" json:"init,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trickle) Reset()         { *m = Trickle{} }
func (m *Trickle) String() string { return proto.CompactTextString(m) }
func (*Trickle) ProtoMessage()    {}
func (*Trickle) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca80ff2c9b7a4e60, []int{4}
}

func (m *Trickle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trickle.Unmarshal(m, b)
}
func (m *Trickle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trickle.Marshal(b, m, deterministic)
}
func (m *Trickle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trickle.Merge(m, src)
}
func (m *Trickle) XXX_Size() int {
	return xxx_messageInfo_Trickle.Size(m)
}
func (m *Trickle) XXX_DiscardUnknown() {
	xxx_messageInfo_Trickle.DiscardUnknown(m)
}

var xxx_messageInfo_Trickle proto.InternalMessageInfo

func (m *Trickle) GetInit() string {
	if m != nil {
		return m.Init
	}
	return ""
}

type SessionDescription struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Sdp                  []byte   `protobuf:"bytes,2,opt,name=sdp,proto3" json:"sdp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionDescription) Reset()         { *m = SessionDescription{} }
func (m *SessionDescription) String() string { return proto.CompactTextString(m) }
func (*SessionDescription) ProtoMessage()    {}
func (*SessionDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca80ff2c9b7a4e60, []int{5}
}

func (m *SessionDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionDescription.Unmarshal(m, b)
}
func (m *SessionDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionDescription.Marshal(b, m, deterministic)
}
func (m *SessionDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionDescription.Merge(m, src)
}
func (m *SessionDescription) XXX_Size() int {
	return xxx_messageInfo_SessionDescription.Size(m)
}
func (m *SessionDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionDescription.DiscardUnknown(m)
}

var xxx_messageInfo_SessionDescription proto.InternalMessageInfo

func (m *SessionDescription) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *SessionDescription) GetSdp() []byte {
	if m != nil {
		return m.Sdp
	}
	return nil
}

func init() {
	proto.RegisterType((*SignalRequest)(nil), "sfu.SignalRequest")
	proto.RegisterType((*SignalReply)(nil), "sfu.SignalReply")
	proto.RegisterType((*JoinRequest)(nil), "sfu.JoinRequest")
	proto.RegisterType((*JoinReply)(nil), "sfu.JoinReply")
	proto.RegisterType((*Trickle)(nil), "sfu.Trickle")
	proto.RegisterType((*SessionDescription)(nil), "sfu.SessionDescription")
}

func init() { proto.RegisterFile("cmd/server/grpc/proto/sfu.proto", fileDescriptor_ca80ff2c9b7a4e60) }

var fileDescriptor_ca80ff2c9b7a4e60 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xc1, 0x4b, 0xe3, 0x40,
	0x14, 0xc6, 0x93, 0x4d, 0xb7, 0x25, 0xaf, 0xdd, 0xa5, 0xcc, 0x65, 0xcb, 0xc2, 0xb2, 0x12, 0x44,
	0x7a, 0x69, 0x47, 0xaa, 0x20, 0xe8, 0xad, 0x88, 0x14, 0x0f, 0x3d, 0xa4, 0x7a, 0xf1, 0x96, 0x26,
	0x93, 0xf8, 0x34, 0x9d, 0x19, 0x67, 0x26, 0x4a, 0xfe, 0x1c, 0x0f, 0xfe, 0x9f, 0x92, 0x49, 0x6a,
	0x5b, 0x14, 0x7a, 0xf4, 0x94, 0x8f, 0x79, 0xdf, 0xf7, 0xe5, 0x37, 0xc9, 0x83, 0xff, 0xf1, 0x2a,
	0xa1, 0x9a, 0xa9, 0x67, 0xa6, 0x68, 0xa6, 0x64, 0x4c, 0xa5, 0x12, 0x46, 0x50, 0x9d, 0x16, 0x63,
	0xab, 0x88, 0xa7, 0xd3, 0x22, 0x78, 0x73, 0xe1, 0xd7, 0x02, 0x33, 0x1e, 0xe5, 0x21, 0x7b, 0x2a,
	0x98, 0x36, 0xe4, 0x08, 0x5a, 0x0f, 0x02, 0xf9, 0xc0, 0x3d, 0x70, 0x87, 0xdd, 0x49, 0x7f, 0x5c,
	0x05, 0xae, 0x05, 0xf2, 0x66, 0x3e, 0x73, 0x42, 0x3b, 0x27, 0x67, 0xe0, 0x73, 0x96, 0x09, 0x83,
	0x91, 0x61, 0x83, 0x1f, 0xd6, 0xfc, 0xc7, 0x9a, 0x17, 0x4c, 0x6b, 0x14, 0xfc, 0x92, 0xe9, 0x58,
	0xa1, 0x34, 0x28, 0xf8, 0xcc, 0x09, 0x37, 0x5e, 0x32, 0x84, 0x8e, 0x51, 0x18, 0x3f, 0xe6, 0x6c,
	0xe0, 0xd9, 0x58, 0xcf, 0xc6, 0x6e, 0xea, 0xb3, 0x99, 0x13, 0xae, 0xc7, 0x53, 0x1f, 0x3a, 0x32,
	0x2a, 0x73, 0x11, 0x25, 0xc1, 0xab, 0x0b, 0xdd, 0x35, 0xa7, 0xcc, 0x4b, 0x72, 0xb8, 0x43, 0xf9,
	0x7b, 0x8b, 0x52, 0xe6, 0xe5, 0x37, 0x31, 0xce, 0xa1, 0xbb, 0xf5, 0xa1, 0x48, 0x1f, 0x3c, 0x85,
	0x89, 0x25, 0xf4, 0xc3, 0x4a, 0x92, 0x11, 0xfc, 0x14, 0x69, 0xca, 0xd4, 0x1e, 0x94, 0xb0, 0x76,
	0x05, 0x73, 0xf0, 0x3f, 0xae, 0x54, 0xb5, 0xc9, 0x4d, 0x9b, 0xc4, 0x84, 0x50, 0x68, 0x47, 0x5c,
	0xbf, 0xec, 0xaf, 0x6b, 0x6c, 0xc1, 0x3f, 0xe8, 0x34, 0x17, 0x20, 0x04, 0x5a, 0xc8, 0xd1, 0x34,
	0x75, 0x56, 0x07, 0xe7, 0x40, 0x3e, 0x87, 0x2b, 0xa7, 0x29, 0x25, 0x5b, 0x3b, 0x2b, 0x5d, 0xb1,
	0xe8, 0x44, 0xda, 0xd7, 0xf6, 0xc2, 0x4a, 0x4e, 0x2e, 0xc0, 0x5b, 0x5c, 0xdd, 0x92, 0x53, 0x68,
	0xd7, 0x3f, 0x89, 0x90, 0x1a, 0x66, 0x7b, 0xb3, 0xfe, 0xf6, 0x77, 0xce, 0x64, 0x5e, 0x06, 0xce,
	0xd0, 0x3d, 0x76, 0xa7, 0xf4, 0x6e, 0x94, 0xa1, 0xb9, 0x2f, 0x96, 0xe3, 0x58, 0xac, 0xa8, 0x44,
	0xc1, 0x29, 0x0a, 0x3e, 0xd2, 0x69, 0x41, 0xbf, 0xdc, 0xe1, 0x65, 0xdb, 0x3e, 0x4e, 0xde, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x68, 0xed, 0x3f, 0x46, 0xe3, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SFUClient is the client API for SFU service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SFUClient interface {
	Signal(ctx context.Context, opts ...grpc.CallOption) (SFU_SignalClient, error)
}

type sFUClient struct {
	cc *grpc.ClientConn
}

func NewSFUClient(cc *grpc.ClientConn) SFUClient {
	return &sFUClient{cc}
}

func (c *sFUClient) Signal(ctx context.Context, opts ...grpc.CallOption) (SFU_SignalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SFU_serviceDesc.Streams[0], "/sfu.SFU/Signal", opts...)
	if err != nil {
		return nil, err
	}
	x := &sFUSignalClient{stream}
	return x, nil
}

type SFU_SignalClient interface {
	Send(*SignalRequest) error
	Recv() (*SignalReply, error)
	grpc.ClientStream
}

type sFUSignalClient struct {
	grpc.ClientStream
}

func (x *sFUSignalClient) Send(m *SignalRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sFUSignalClient) Recv() (*SignalReply, error) {
	m := new(SignalReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SFUServer is the server API for SFU service.
type SFUServer interface {
	Signal(SFU_SignalServer) error
}

// UnimplementedSFUServer can be embedded to have forward compatible implementations.
type UnimplementedSFUServer struct {
}

func (*UnimplementedSFUServer) Signal(srv SFU_SignalServer) error {
	return status.Errorf(codes.Unimplemented, "method Signal not implemented")
}

func RegisterSFUServer(s *grpc.Server, srv SFUServer) {
	s.RegisterService(&_SFU_serviceDesc, srv)
}

func _SFU_Signal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SFUServer).Signal(&sFUSignalServer{stream})
}

type SFU_SignalServer interface {
	Send(*SignalReply) error
	Recv() (*SignalRequest, error)
	grpc.ServerStream
}

type sFUSignalServer struct {
	grpc.ServerStream
}

func (x *sFUSignalServer) Send(m *SignalReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sFUSignalServer) Recv() (*SignalRequest, error) {
	m := new(SignalRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SFU_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sfu.SFU",
	HandlerType: (*SFUServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Signal",
			Handler:       _SFU_Signal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cmd/server/grpc/proto/sfu.proto",
}
