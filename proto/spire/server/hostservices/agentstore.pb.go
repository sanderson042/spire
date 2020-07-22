// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire/server/hostservices/agentstore.proto

package hostservices

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

type AgentInfo struct {
	AgentId              string   `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentInfo) Reset()         { *m = AgentInfo{} }
func (m *AgentInfo) String() string { return proto.CompactTextString(m) }
func (*AgentInfo) ProtoMessage()    {}
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f94b0b7e4f5b3ba, []int{0}
}

func (m *AgentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentInfo.Unmarshal(m, b)
}
func (m *AgentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentInfo.Marshal(b, m, deterministic)
}
func (m *AgentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentInfo.Merge(m, src)
}
func (m *AgentInfo) XXX_Size() int {
	return xxx_messageInfo_AgentInfo.Size(m)
}
func (m *AgentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AgentInfo proto.InternalMessageInfo

func (m *AgentInfo) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

type GetAgentInfoRequest struct {
	AgentId              string   `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgentInfoRequest) Reset()         { *m = GetAgentInfoRequest{} }
func (m *GetAgentInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetAgentInfoRequest) ProtoMessage()    {}
func (*GetAgentInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f94b0b7e4f5b3ba, []int{1}
}

func (m *GetAgentInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentInfoRequest.Unmarshal(m, b)
}
func (m *GetAgentInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetAgentInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentInfoRequest.Merge(m, src)
}
func (m *GetAgentInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetAgentInfoRequest.Size(m)
}
func (m *GetAgentInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentInfoRequest proto.InternalMessageInfo

func (m *GetAgentInfoRequest) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

type GetAgentInfoResponse struct {
	Info                 *AgentInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAgentInfoResponse) Reset()         { *m = GetAgentInfoResponse{} }
func (m *GetAgentInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetAgentInfoResponse) ProtoMessage()    {}
func (*GetAgentInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f94b0b7e4f5b3ba, []int{2}
}

func (m *GetAgentInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentInfoResponse.Unmarshal(m, b)
}
func (m *GetAgentInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetAgentInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentInfoResponse.Merge(m, src)
}
func (m *GetAgentInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetAgentInfoResponse.Size(m)
}
func (m *GetAgentInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentInfoResponse proto.InternalMessageInfo

func (m *GetAgentInfoResponse) GetInfo() *AgentInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentInfo)(nil), "spire.server.hostservices.AgentInfo")
	proto.RegisterType((*GetAgentInfoRequest)(nil), "spire.server.hostservices.GetAgentInfoRequest")
	proto.RegisterType((*GetAgentInfoResponse)(nil), "spire.server.hostservices.GetAgentInfoResponse")
}

func init() {
	proto.RegisterFile("spire/server/hostservices/agentstore.proto", fileDescriptor_6f94b0b7e4f5b3ba)
}

var fileDescriptor_6f94b0b7e4f5b3ba = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4b, 0xc5, 0x30,
	0x10, 0xc6, 0x79, 0x20, 0xea, 0x3b, 0x9d, 0xa2, 0x83, 0xcf, 0x49, 0x8a, 0x88, 0x38, 0x24, 0x52,
	0x07, 0x75, 0xd4, 0x45, 0xba, 0x49, 0xdd, 0x5c, 0xc4, 0xb6, 0x97, 0x36, 0x83, 0xb9, 0x98, 0x4b,
	0xdd, 0xfc, 0xdf, 0xa5, 0x57, 0xd0, 0x0a, 0xb6, 0xb8, 0x25, 0xe1, 0xf7, 0x3b, 0xbe, 0xdc, 0x07,
	0x17, 0x1c, 0x5c, 0x44, 0xc3, 0x18, 0x3f, 0x30, 0x9a, 0x8e, 0x38, 0x0d, 0x47, 0x57, 0x23, 0x9b,
	0xd7, 0x16, 0x7d, 0xe2, 0x44, 0x11, 0x75, 0x88, 0x94, 0x48, 0x6d, 0x84, 0xd5, 0x23, 0xab, 0xa7,
	0x6c, 0x76, 0x06, 0xeb, 0xbb, 0x01, 0x2f, 0xbc, 0x25, 0xb5, 0x81, 0x5d, 0x71, 0x5f, 0x5c, 0x73,
	0xb4, 0x3a, 0x59, 0x9d, 0xaf, 0xcb, 0x1d, 0xb9, 0x17, 0x4d, 0x76, 0x09, 0x07, 0x0f, 0x98, 0xbe,
	0xd1, 0x12, 0xdf, 0x7b, 0xe4, 0xb4, 0x64, 0x3c, 0xc2, 0xe1, 0x6f, 0x83, 0x03, 0x79, 0x46, 0x75,
	0x03, 0x5b, 0xce, 0x5b, 0x12, 0x7c, 0x2f, 0x3f, 0xd5, 0xb3, 0xd9, 0xf4, 0x8f, 0x2b, 0x46, 0xfe,
	0x09, 0x20, 0x4f, 0x4f, 0xc3, 0xd7, 0x14, 0xc1, 0xfe, 0x74, 0xbe, 0xd2, 0x0b, 0x93, 0xfe, 0x88,
	0x7e, 0x6c, 0xfe, 0xcd, 0x8f, 0xc1, 0xef, 0x6f, 0x9f, 0xaf, 0x5b, 0x97, 0xba, 0xbe, 0xd2, 0x35,
	0xbd, 0x19, 0x0e, 0xce, 0x5a, 0x34, 0x63, 0x0b, 0xb2, 0x66, 0x33, 0xdb, 0x48, 0xb5, 0x2d, 0xc0,
	0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x9d, 0x80, 0x2c, 0xb5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentStoreClient is the client API for AgentStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentStoreClient interface {
	GetAgentInfo(ctx context.Context, in *GetAgentInfoRequest, opts ...grpc.CallOption) (*GetAgentInfoResponse, error)
}

type agentStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentStoreClient(cc grpc.ClientConnInterface) AgentStoreClient {
	return &agentStoreClient{cc}
}

func (c *agentStoreClient) GetAgentInfo(ctx context.Context, in *GetAgentInfoRequest, opts ...grpc.CallOption) (*GetAgentInfoResponse, error) {
	out := new(GetAgentInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.server.hostservices.AgentStore/GetAgentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentStoreServer is the server API for AgentStore service.
type AgentStoreServer interface {
	GetAgentInfo(context.Context, *GetAgentInfoRequest) (*GetAgentInfoResponse, error)
}

// UnimplementedAgentStoreServer can be embedded to have forward compatible implementations.
type UnimplementedAgentStoreServer struct {
}

func (*UnimplementedAgentStoreServer) GetAgentInfo(ctx context.Context, req *GetAgentInfoRequest) (*GetAgentInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentInfo not implemented")
}

func RegisterAgentStoreServer(s *grpc.Server, srv AgentStoreServer) {
	s.RegisterService(&_AgentStore_serviceDesc, srv)
}

func _AgentStore_GetAgentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentStoreServer).GetAgentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.hostservices.AgentStore/GetAgentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentStoreServer).GetAgentInfo(ctx, req.(*GetAgentInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.server.hostservices.AgentStore",
	HandlerType: (*AgentStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgentInfo",
			Handler:    _AgentStore_GetAgentInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/server/hostservices/agentstore.proto",
}
