// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protoDataNode/DataNode.proto

package Proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataNodeServiceClient is the client API for DataNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNodeServiceClient interface {
	Node1(ctx context.Context, in *ReqSaveData, opts ...grpc.CallOption) (*ResSaveData, error)
	Node2(ctx context.Context, in *ReqSaveData, opts ...grpc.CallOption) (*ResSaveData, error)
	Node3(ctx context.Context, in *ReqSaveData, opts ...grpc.CallOption) (*ResSaveData, error)
	Getid(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ResData, error)
	Getid2(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ResData, error)
	Getid3(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ResData, error)
	ShutDown(ctx context.Context, in *ReqShutDown, opts ...grpc.CallOption) (*ResShutDown, error)
}

type dataNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodeServiceClient(cc grpc.ClientConnInterface) DataNodeServiceClient {
	return &dataNodeServiceClient{cc}
}

func (c *dataNodeServiceClient) Node1(ctx context.Context, in *ReqSaveData, opts ...grpc.CallOption) (*ResSaveData, error) {
	out := new(ResSaveData)
	err := c.cc.Invoke(ctx, "/grpc.DataNodeService/Node1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) Node2(ctx context.Context, in *ReqSaveData, opts ...grpc.CallOption) (*ResSaveData, error) {
	out := new(ResSaveData)
	err := c.cc.Invoke(ctx, "/grpc.DataNodeService/Node2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) Node3(ctx context.Context, in *ReqSaveData, opts ...grpc.CallOption) (*ResSaveData, error) {
	out := new(ResSaveData)
	err := c.cc.Invoke(ctx, "/grpc.DataNodeService/Node3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) Getid(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ResData, error) {
	out := new(ResData)
	err := c.cc.Invoke(ctx, "/grpc.DataNodeService/Getid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) Getid2(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ResData, error) {
	out := new(ResData)
	err := c.cc.Invoke(ctx, "/grpc.DataNodeService/Getid2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) Getid3(ctx context.Context, in *ReqData, opts ...grpc.CallOption) (*ResData, error) {
	out := new(ResData)
	err := c.cc.Invoke(ctx, "/grpc.DataNodeService/Getid3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodeServiceClient) ShutDown(ctx context.Context, in *ReqShutDown, opts ...grpc.CallOption) (*ResShutDown, error) {
	out := new(ResShutDown)
	err := c.cc.Invoke(ctx, "/grpc.DataNodeService/ShutDown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodeServiceServer is the server API for DataNodeService service.
// All implementations must embed UnimplementedDataNodeServiceServer
// for forward compatibility
type DataNodeServiceServer interface {
	Node1(context.Context, *ReqSaveData) (*ResSaveData, error)
	Node2(context.Context, *ReqSaveData) (*ResSaveData, error)
	Node3(context.Context, *ReqSaveData) (*ResSaveData, error)
	Getid(context.Context, *ReqData) (*ResData, error)
	Getid2(context.Context, *ReqData) (*ResData, error)
	Getid3(context.Context, *ReqData) (*ResData, error)
	ShutDown(context.Context, *ReqShutDown) (*ResShutDown, error)
	mustEmbedUnimplementedDataNodeServiceServer()
}

// UnimplementedDataNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataNodeServiceServer struct {
}

func (UnimplementedDataNodeServiceServer) Node1(context.Context, *ReqSaveData) (*ResSaveData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Node1 not implemented")
}
func (UnimplementedDataNodeServiceServer) Node2(context.Context, *ReqSaveData) (*ResSaveData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Node2 not implemented")
}
func (UnimplementedDataNodeServiceServer) Node3(context.Context, *ReqSaveData) (*ResSaveData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Node3 not implemented")
}
func (UnimplementedDataNodeServiceServer) Getid(context.Context, *ReqData) (*ResData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getid not implemented")
}
func (UnimplementedDataNodeServiceServer) Getid2(context.Context, *ReqData) (*ResData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getid2 not implemented")
}
func (UnimplementedDataNodeServiceServer) Getid3(context.Context, *ReqData) (*ResData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getid3 not implemented")
}
func (UnimplementedDataNodeServiceServer) ShutDown(context.Context, *ReqShutDown) (*ResShutDown, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutDown not implemented")
}
func (UnimplementedDataNodeServiceServer) mustEmbedUnimplementedDataNodeServiceServer() {}

// UnsafeDataNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodeServiceServer will
// result in compilation errors.
type UnsafeDataNodeServiceServer interface {
	mustEmbedUnimplementedDataNodeServiceServer()
}

func RegisterDataNodeServiceServer(s grpc.ServiceRegistrar, srv DataNodeServiceServer) {
	s.RegisterService(&DataNodeService_ServiceDesc, srv)
}

func _DataNodeService_Node1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSaveData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).Node1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DataNodeService/Node1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).Node1(ctx, req.(*ReqSaveData))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_Node2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSaveData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).Node2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DataNodeService/Node2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).Node2(ctx, req.(*ReqSaveData))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_Node3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSaveData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).Node3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DataNodeService/Node3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).Node3(ctx, req.(*ReqSaveData))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_Getid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).Getid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DataNodeService/Getid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).Getid(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_Getid2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).Getid2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DataNodeService/Getid2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).Getid2(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_Getid3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).Getid3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DataNodeService/Getid3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).Getid3(ctx, req.(*ReqData))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodeService_ShutDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqShutDown)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodeServiceServer).ShutDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.DataNodeService/ShutDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodeServiceServer).ShutDown(ctx, req.(*ReqShutDown))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNodeService_ServiceDesc is the grpc.ServiceDesc for DataNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.DataNodeService",
	HandlerType: (*DataNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Node1",
			Handler:    _DataNodeService_Node1_Handler,
		},
		{
			MethodName: "Node2",
			Handler:    _DataNodeService_Node2_Handler,
		},
		{
			MethodName: "Node3",
			Handler:    _DataNodeService_Node3_Handler,
		},
		{
			MethodName: "Getid",
			Handler:    _DataNodeService_Getid_Handler,
		},
		{
			MethodName: "Getid2",
			Handler:    _DataNodeService_Getid2_Handler,
		},
		{
			MethodName: "Getid3",
			Handler:    _DataNodeService_Getid3_Handler,
		},
		{
			MethodName: "ShutDown",
			Handler:    _DataNodeService_ShutDown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoDataNode/DataNode.proto",
}