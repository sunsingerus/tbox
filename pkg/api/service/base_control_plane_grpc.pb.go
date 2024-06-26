// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	common "github.com/sunsingerus/tbox/pkg/api/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ControlPlaneClient is the client API for ControlPlane service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlPlaneClient interface {
	// Bi-directional Commands stream.
	// Commands are sent from service to client and from client to server
	Tasks(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_TasksClient, error)
	// Uni-directional Metrics stream from client to server.
	Metrics(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_MetricsClient, error)
}

type controlPlaneClient struct {
	cc grpc.ClientConnInterface
}

func NewControlPlaneClient(cc grpc.ClientConnInterface) ControlPlaneClient {
	return &controlPlaneClient{cc}
}

func (c *controlPlaneClient) Tasks(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_TasksClient, error) {
	stream, err := c.cc.NewStream(ctx, &ControlPlane_ServiceDesc.Streams[0], "/api.service.ControlPlane/Tasks", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlPlaneTasksClient{stream}
	return x, nil
}

type ControlPlane_TasksClient interface {
	Send(*common.Task) error
	Recv() (*common.Task, error)
	grpc.ClientStream
}

type controlPlaneTasksClient struct {
	grpc.ClientStream
}

func (x *controlPlaneTasksClient) Send(m *common.Task) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controlPlaneTasksClient) Recv() (*common.Task, error) {
	m := new(common.Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlPlaneClient) Metrics(ctx context.Context, opts ...grpc.CallOption) (ControlPlane_MetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ControlPlane_ServiceDesc.Streams[1], "/api.service.ControlPlane/Metrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &controlPlaneMetricsClient{stream}
	return x, nil
}

type ControlPlane_MetricsClient interface {
	Send(*common.Metric) error
	CloseAndRecv() (*common.Metric, error)
	grpc.ClientStream
}

type controlPlaneMetricsClient struct {
	grpc.ClientStream
}

func (x *controlPlaneMetricsClient) Send(m *common.Metric) error {
	return x.ClientStream.SendMsg(m)
}

func (x *controlPlaneMetricsClient) CloseAndRecv() (*common.Metric, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(common.Metric)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControlPlaneServer is the server API for ControlPlane service.
// All implementations must embed UnimplementedControlPlaneServer
// for forward compatibility
type ControlPlaneServer interface {
	// Bi-directional Commands stream.
	// Commands are sent from service to client and from client to server
	Tasks(ControlPlane_TasksServer) error
	// Uni-directional Metrics stream from client to server.
	Metrics(ControlPlane_MetricsServer) error
	mustEmbedUnimplementedControlPlaneServer()
}

// UnimplementedControlPlaneServer must be embedded to have forward compatible implementations.
type UnimplementedControlPlaneServer struct {
}

func (UnimplementedControlPlaneServer) Tasks(ControlPlane_TasksServer) error {
	return status.Errorf(codes.Unimplemented, "method Tasks not implemented")
}
func (UnimplementedControlPlaneServer) Metrics(ControlPlane_MetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method Metrics not implemented")
}
func (UnimplementedControlPlaneServer) mustEmbedUnimplementedControlPlaneServer() {}

// UnsafeControlPlaneServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlPlaneServer will
// result in compilation errors.
type UnsafeControlPlaneServer interface {
	mustEmbedUnimplementedControlPlaneServer()
}

func RegisterControlPlaneServer(s grpc.ServiceRegistrar, srv ControlPlaneServer) {
	s.RegisterService(&ControlPlane_ServiceDesc, srv)
}

func _ControlPlane_Tasks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControlPlaneServer).Tasks(&controlPlaneTasksServer{stream})
}

type ControlPlane_TasksServer interface {
	Send(*common.Task) error
	Recv() (*common.Task, error)
	grpc.ServerStream
}

type controlPlaneTasksServer struct {
	grpc.ServerStream
}

func (x *controlPlaneTasksServer) Send(m *common.Task) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controlPlaneTasksServer) Recv() (*common.Task, error) {
	m := new(common.Task)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ControlPlane_Metrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ControlPlaneServer).Metrics(&controlPlaneMetricsServer{stream})
}

type ControlPlane_MetricsServer interface {
	SendAndClose(*common.Metric) error
	Recv() (*common.Metric, error)
	grpc.ServerStream
}

type controlPlaneMetricsServer struct {
	grpc.ServerStream
}

func (x *controlPlaneMetricsServer) SendAndClose(m *common.Metric) error {
	return x.ServerStream.SendMsg(m)
}

func (x *controlPlaneMetricsServer) Recv() (*common.Metric, error) {
	m := new(common.Metric)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ControlPlane_ServiceDesc is the grpc.ServiceDesc for ControlPlane service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControlPlane_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.service.ControlPlane",
	HandlerType: (*ControlPlaneServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tasks",
			Handler:       _ControlPlane_Tasks_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Metrics",
			Handler:       _ControlPlane_Metrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/service/base_control_plane.proto",
}
