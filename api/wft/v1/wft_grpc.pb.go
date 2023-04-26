// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: wft/v1/wft.proto

package v1

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

const (
	Wft_GetWarframeMarketItem_FullMethodName    = "/wft.v1.Wft/GetWarframeMarketItem"
	Wft_GetWarframeOfficalItem_FullMethodName   = "/wft.v1.Wft/GetWarframeOfficalItem"
	Wft_GetWarframeMarketAuction_FullMethodName = "/wft.v1.Wft/GetWarframeMarketAuction"
)

// WftClient is the client API for Wft service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WftClient interface {
	GetWarframeMarketItem(ctx context.Context, in *NameSearchReq, opts ...grpc.CallOption) (*WmItemResp, error)
	GetWarframeOfficalItem(ctx context.Context, in *NameSearchReq, opts ...grpc.CallOption) (*WarframeItemResp, error)
	GetWarframeMarketAuction(ctx context.Context, in *NameSearchReq, opts ...grpc.CallOption) (*WmItemResp, error)
}

type wftClient struct {
	cc grpc.ClientConnInterface
}

func NewWftClient(cc grpc.ClientConnInterface) WftClient {
	return &wftClient{cc}
}

func (c *wftClient) GetWarframeMarketItem(ctx context.Context, in *NameSearchReq, opts ...grpc.CallOption) (*WmItemResp, error) {
	out := new(WmItemResp)
	err := c.cc.Invoke(ctx, Wft_GetWarframeMarketItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wftClient) GetWarframeOfficalItem(ctx context.Context, in *NameSearchReq, opts ...grpc.CallOption) (*WarframeItemResp, error) {
	out := new(WarframeItemResp)
	err := c.cc.Invoke(ctx, Wft_GetWarframeOfficalItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wftClient) GetWarframeMarketAuction(ctx context.Context, in *NameSearchReq, opts ...grpc.CallOption) (*WmItemResp, error) {
	out := new(WmItemResp)
	err := c.cc.Invoke(ctx, Wft_GetWarframeMarketAuction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WftServer is the server API for Wft service.
// All implementations must embed UnimplementedWftServer
// for forward compatibility
type WftServer interface {
	GetWarframeMarketItem(context.Context, *NameSearchReq) (*WmItemResp, error)
	GetWarframeOfficalItem(context.Context, *NameSearchReq) (*WarframeItemResp, error)
	GetWarframeMarketAuction(context.Context, *NameSearchReq) (*WmItemResp, error)
	mustEmbedUnimplementedWftServer()
}

// UnimplementedWftServer must be embedded to have forward compatible implementations.
type UnimplementedWftServer struct {
}

func (UnimplementedWftServer) GetWarframeMarketItem(context.Context, *NameSearchReq) (*WmItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarframeMarketItem not implemented")
}
func (UnimplementedWftServer) GetWarframeOfficalItem(context.Context, *NameSearchReq) (*WarframeItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarframeOfficalItem not implemented")
}
func (UnimplementedWftServer) GetWarframeMarketAuction(context.Context, *NameSearchReq) (*WmItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarframeMarketAuction not implemented")
}
func (UnimplementedWftServer) mustEmbedUnimplementedWftServer() {}

// UnsafeWftServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WftServer will
// result in compilation errors.
type UnsafeWftServer interface {
	mustEmbedUnimplementedWftServer()
}

func RegisterWftServer(s grpc.ServiceRegistrar, srv WftServer) {
	s.RegisterService(&Wft_ServiceDesc, srv)
}

func _Wft_GetWarframeMarketItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WftServer).GetWarframeMarketItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wft_GetWarframeMarketItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WftServer).GetWarframeMarketItem(ctx, req.(*NameSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wft_GetWarframeOfficalItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WftServer).GetWarframeOfficalItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wft_GetWarframeOfficalItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WftServer).GetWarframeOfficalItem(ctx, req.(*NameSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wft_GetWarframeMarketAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WftServer).GetWarframeMarketAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wft_GetWarframeMarketAuction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WftServer).GetWarframeMarketAuction(ctx, req.(*NameSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Wft_ServiceDesc is the grpc.ServiceDesc for Wft service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wft_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wft.v1.Wft",
	HandlerType: (*WftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWarframeMarketItem",
			Handler:    _Wft_GetWarframeMarketItem_Handler,
		},
		{
			MethodName: "GetWarframeOfficalItem",
			Handler:    _Wft_GetWarframeOfficalItem_Handler,
		},
		{
			MethodName: "GetWarframeMarketAuction",
			Handler:    _Wft_GetWarframeMarketAuction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wft/v1/wft.proto",
}
