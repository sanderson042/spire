// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hostservices

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsServiceClient interface {
	SetGauge(ctx context.Context, in *SetGaugeRequest, opts ...grpc.CallOption) (*SetGaugeResponse, error)
	EmitKey(ctx context.Context, in *EmitKeyRequest, opts ...grpc.CallOption) (*EmitKeyResponse, error)
	IncrCounter(ctx context.Context, in *IncrCounterRequest, opts ...grpc.CallOption) (*IncrCounterResponse, error)
	AddSample(ctx context.Context, in *AddSampleRequest, opts ...grpc.CallOption) (*AddSampleResponse, error)
	MeasureSince(ctx context.Context, in *MeasureSinceRequest, opts ...grpc.CallOption) (*MeasureSinceResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) SetGauge(ctx context.Context, in *SetGaugeRequest, opts ...grpc.CallOption) (*SetGaugeResponse, error) {
	out := new(SetGaugeResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/SetGauge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) EmitKey(ctx context.Context, in *EmitKeyRequest, opts ...grpc.CallOption) (*EmitKeyResponse, error) {
	out := new(EmitKeyResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/EmitKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) IncrCounter(ctx context.Context, in *IncrCounterRequest, opts ...grpc.CallOption) (*IncrCounterResponse, error) {
	out := new(IncrCounterResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/IncrCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) AddSample(ctx context.Context, in *AddSampleRequest, opts ...grpc.CallOption) (*AddSampleResponse, error) {
	out := new(AddSampleResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/AddSample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) MeasureSince(ctx context.Context, in *MeasureSinceRequest, opts ...grpc.CallOption) (*MeasureSinceResponse, error) {
	out := new(MeasureSinceResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/MeasureSince", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
// All implementations must embed UnimplementedMetricsServiceServer
// for forward compatibility
type MetricsServiceServer interface {
	SetGauge(context.Context, *SetGaugeRequest) (*SetGaugeResponse, error)
	EmitKey(context.Context, *EmitKeyRequest) (*EmitKeyResponse, error)
	IncrCounter(context.Context, *IncrCounterRequest) (*IncrCounterResponse, error)
	AddSample(context.Context, *AddSampleRequest) (*AddSampleResponse, error)
	MeasureSince(context.Context, *MeasureSinceRequest) (*MeasureSinceResponse, error)
	mustEmbedUnimplementedMetricsServiceServer()
}

// UnimplementedMetricsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (UnimplementedMetricsServiceServer) SetGauge(context.Context, *SetGaugeRequest) (*SetGaugeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGauge not implemented")
}
func (UnimplementedMetricsServiceServer) EmitKey(context.Context, *EmitKeyRequest) (*EmitKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitKey not implemented")
}
func (UnimplementedMetricsServiceServer) IncrCounter(context.Context, *IncrCounterRequest) (*IncrCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrCounter not implemented")
}
func (UnimplementedMetricsServiceServer) AddSample(context.Context, *AddSampleRequest) (*AddSampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSample not implemented")
}
func (UnimplementedMetricsServiceServer) MeasureSince(context.Context, *MeasureSinceRequest) (*MeasureSinceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MeasureSince not implemented")
}
func (UnimplementedMetricsServiceServer) mustEmbedUnimplementedMetricsServiceServer() {}

// UnsafeMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServiceServer will
// result in compilation errors.
type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_SetGauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGaugeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).SetGauge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/SetGauge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).SetGauge(ctx, req.(*SetGaugeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_EmitKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).EmitKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/EmitKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).EmitKey(ctx, req.(*EmitKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_IncrCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).IncrCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/IncrCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).IncrCounter(ctx, req.(*IncrCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_AddSample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).AddSample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/AddSample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).AddSample(ctx, req.(*AddSampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_MeasureSince_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasureSinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).MeasureSince(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/MeasureSince",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).MeasureSince(ctx, req.(*MeasureSinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.common.hostservices.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGauge",
			Handler:    _MetricsService_SetGauge_Handler,
		},
		{
			MethodName: "EmitKey",
			Handler:    _MetricsService_EmitKey_Handler,
		},
		{
			MethodName: "IncrCounter",
			Handler:    _MetricsService_IncrCounter_Handler,
		},
		{
			MethodName: "AddSample",
			Handler:    _MetricsService_AddSample_Handler,
		},
		{
			MethodName: "MeasureSince",
			Handler:    _MetricsService_MeasureSince_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/common/hostservices/metricsservice.proto",
}
