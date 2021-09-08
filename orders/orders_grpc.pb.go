// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package orders_management

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

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementClient interface {
	CreateNewOrder(ctx context.Context, in *NewOrder, opts ...grpc.CallOption) (*Order, error)
}

type orderManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementClient(cc grpc.ClientConnInterface) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) CreateNewOrder(ctx context.Context, in *NewOrder, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/orders.OrderManagement/CreateNewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderManagementServer is the server API for OrderManagement service.
// All implementations must embed UnimplementedOrderManagementServer
// for forward compatibility
type OrderManagementServer interface {
	CreateNewOrder(context.Context, *NewOrder) (*Order, error)
	mustEmbedUnimplementedOrderManagementServer()
}

// UnimplementedOrderManagementServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (UnimplementedOrderManagementServer) CreateNewOrder(context.Context, *NewOrder) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewOrder not implemented")
}
func (UnimplementedOrderManagementServer) mustEmbedUnimplementedOrderManagementServer() {}

// UnsafeOrderManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServer will
// result in compilation errors.
type UnsafeOrderManagementServer interface {
	mustEmbedUnimplementedOrderManagementServer()
}

func RegisterOrderManagementServer(s grpc.ServiceRegistrar, srv OrderManagementServer) {
	s.RegisterService(&OrderManagement_ServiceDesc, srv)
}

func _OrderManagement_CreateNewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderManagementServer).CreateNewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.OrderManagement/CreateNewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderManagementServer).CreateNewOrder(ctx, req.(*NewOrder))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderManagement_ServiceDesc is the grpc.ServiceDesc for OrderManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewOrder",
			Handler:    _OrderManagement_CreateNewOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/orders.proto",
}
