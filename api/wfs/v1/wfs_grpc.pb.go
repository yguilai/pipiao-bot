// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: wfs/v1/wfs.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Wfs_GetCycles_FullMethodName              = "/wfs.v1.Wfs/GetCycles"
	Wfs_GetAlerts_FullMethodName              = "/wfs.v1.Wfs/GetAlerts"
	Wfs_GetSorties_FullMethodName             = "/wfs.v1.Wfs/GetSorties"
	Wfs_GetVoidTrader_FullMethodName          = "/wfs.v1.Wfs/GetVoidTrader"
	Wfs_GetWarframeMarket_FullMethodName      = "/wfs.v1.Wfs/GetWarframeMarket"
	Wfs_GetWarframeMarketRiven_FullMethodName = "/wfs.v1.Wfs/GetWarframeMarketRiven"
	Wfs_GetWarframeMarketLich_FullMethodName  = "/wfs.v1.Wfs/GetWarframeMarketLich"
	Wfs_GetInvasions_FullMethodName           = "/wfs.v1.Wfs/GetInvasions"
	Wfs_GetNightwaves_FullMethodName          = "/wfs.v1.Wfs/GetNightwaves"
	Wfs_GetKuva_FullMethodName                = "/wfs.v1.Wfs/GetKuva"
	Wfs_GetVoidStorm_FullMethodName           = "/wfs.v1.Wfs/GetVoidStorm"
	Wfs_GetDailyDeal_FullMethodName           = "/wfs.v1.Wfs/GetDailyDeal"
)

// WfsClient is the client API for Wfs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WfsClient interface {
	GetCycles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CyclesResp, error)
	GetAlerts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AlertsResp, error)
	GetSorties(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SortiesResp, error)
	GetVoidTrader(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VoidTraderResp, error)
	GetWarframeMarket(ctx context.Context, in *WarframeMarketReq, opts ...grpc.CallOption) (*WarframeMarketResp, error)
	GetWarframeMarketRiven(ctx context.Context, in *WarframeMarketRivenReq, opts ...grpc.CallOption) (*WarframeMarketRivenResp, error)
	GetWarframeMarketLich(ctx context.Context, in *WarframeMarketLichReq, opts ...grpc.CallOption) (*WarframeMarketLichResp, error)
	GetInvasions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InvasionsResp, error)
	GetNightwaves(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NightwavesResp, error)
	GetKuva(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*KuvaResp, error)
	GetVoidStorm(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VoidStormResp, error)
	GetDailyDeal(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DailyDealsResp, error)
}

type wfsClient struct {
	cc grpc.ClientConnInterface
}

func NewWfsClient(cc grpc.ClientConnInterface) WfsClient {
	return &wfsClient{cc}
}

func (c *wfsClient) GetCycles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CyclesResp, error) {
	out := new(CyclesResp)
	err := c.cc.Invoke(ctx, Wfs_GetCycles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetAlerts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AlertsResp, error) {
	out := new(AlertsResp)
	err := c.cc.Invoke(ctx, Wfs_GetAlerts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetSorties(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SortiesResp, error) {
	out := new(SortiesResp)
	err := c.cc.Invoke(ctx, Wfs_GetSorties_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetVoidTrader(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VoidTraderResp, error) {
	out := new(VoidTraderResp)
	err := c.cc.Invoke(ctx, Wfs_GetVoidTrader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetWarframeMarket(ctx context.Context, in *WarframeMarketReq, opts ...grpc.CallOption) (*WarframeMarketResp, error) {
	out := new(WarframeMarketResp)
	err := c.cc.Invoke(ctx, Wfs_GetWarframeMarket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetWarframeMarketRiven(ctx context.Context, in *WarframeMarketRivenReq, opts ...grpc.CallOption) (*WarframeMarketRivenResp, error) {
	out := new(WarframeMarketRivenResp)
	err := c.cc.Invoke(ctx, Wfs_GetWarframeMarketRiven_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetWarframeMarketLich(ctx context.Context, in *WarframeMarketLichReq, opts ...grpc.CallOption) (*WarframeMarketLichResp, error) {
	out := new(WarframeMarketLichResp)
	err := c.cc.Invoke(ctx, Wfs_GetWarframeMarketLich_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetInvasions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InvasionsResp, error) {
	out := new(InvasionsResp)
	err := c.cc.Invoke(ctx, Wfs_GetInvasions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetNightwaves(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NightwavesResp, error) {
	out := new(NightwavesResp)
	err := c.cc.Invoke(ctx, Wfs_GetNightwaves_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetKuva(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*KuvaResp, error) {
	out := new(KuvaResp)
	err := c.cc.Invoke(ctx, Wfs_GetKuva_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetVoidStorm(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VoidStormResp, error) {
	out := new(VoidStormResp)
	err := c.cc.Invoke(ctx, Wfs_GetVoidStorm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wfsClient) GetDailyDeal(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DailyDealsResp, error) {
	out := new(DailyDealsResp)
	err := c.cc.Invoke(ctx, Wfs_GetDailyDeal_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WfsServer is the server API for Wfs service.
// All implementations must embed UnimplementedWfsServer
// for forward compatibility
type WfsServer interface {
	GetCycles(context.Context, *emptypb.Empty) (*CyclesResp, error)
	GetAlerts(context.Context, *emptypb.Empty) (*AlertsResp, error)
	GetSorties(context.Context, *emptypb.Empty) (*SortiesResp, error)
	GetVoidTrader(context.Context, *emptypb.Empty) (*VoidTraderResp, error)
	GetWarframeMarket(context.Context, *WarframeMarketReq) (*WarframeMarketResp, error)
	GetWarframeMarketRiven(context.Context, *WarframeMarketRivenReq) (*WarframeMarketRivenResp, error)
	GetWarframeMarketLich(context.Context, *WarframeMarketLichReq) (*WarframeMarketLichResp, error)
	GetInvasions(context.Context, *emptypb.Empty) (*InvasionsResp, error)
	GetNightwaves(context.Context, *emptypb.Empty) (*NightwavesResp, error)
	GetKuva(context.Context, *emptypb.Empty) (*KuvaResp, error)
	GetVoidStorm(context.Context, *emptypb.Empty) (*VoidStormResp, error)
	GetDailyDeal(context.Context, *emptypb.Empty) (*DailyDealsResp, error)
	mustEmbedUnimplementedWfsServer()
}

// UnimplementedWfsServer must be embedded to have forward compatible implementations.
type UnimplementedWfsServer struct {
}

func (UnimplementedWfsServer) GetCycles(context.Context, *emptypb.Empty) (*CyclesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCycles not implemented")
}
func (UnimplementedWfsServer) GetAlerts(context.Context, *emptypb.Empty) (*AlertsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlerts not implemented")
}
func (UnimplementedWfsServer) GetSorties(context.Context, *emptypb.Empty) (*SortiesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSorties not implemented")
}
func (UnimplementedWfsServer) GetVoidTrader(context.Context, *emptypb.Empty) (*VoidTraderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoidTrader not implemented")
}
func (UnimplementedWfsServer) GetWarframeMarket(context.Context, *WarframeMarketReq) (*WarframeMarketResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarframeMarket not implemented")
}
func (UnimplementedWfsServer) GetWarframeMarketRiven(context.Context, *WarframeMarketRivenReq) (*WarframeMarketRivenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarframeMarketRiven not implemented")
}
func (UnimplementedWfsServer) GetWarframeMarketLich(context.Context, *WarframeMarketLichReq) (*WarframeMarketLichResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWarframeMarketLich not implemented")
}
func (UnimplementedWfsServer) GetInvasions(context.Context, *emptypb.Empty) (*InvasionsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvasions not implemented")
}
func (UnimplementedWfsServer) GetNightwaves(context.Context, *emptypb.Empty) (*NightwavesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNightwaves not implemented")
}
func (UnimplementedWfsServer) GetKuva(context.Context, *emptypb.Empty) (*KuvaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKuva not implemented")
}
func (UnimplementedWfsServer) GetVoidStorm(context.Context, *emptypb.Empty) (*VoidStormResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoidStorm not implemented")
}
func (UnimplementedWfsServer) GetDailyDeal(context.Context, *emptypb.Empty) (*DailyDealsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDailyDeal not implemented")
}
func (UnimplementedWfsServer) mustEmbedUnimplementedWfsServer() {}

// UnsafeWfsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WfsServer will
// result in compilation errors.
type UnsafeWfsServer interface {
	mustEmbedUnimplementedWfsServer()
}

func RegisterWfsServer(s grpc.ServiceRegistrar, srv WfsServer) {
	s.RegisterService(&Wfs_ServiceDesc, srv)
}

func _Wfs_GetCycles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetCycles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetCycles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetCycles(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetAlerts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetSorties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetSorties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetSorties_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetSorties(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetVoidTrader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetVoidTrader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetVoidTrader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetVoidTrader(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetWarframeMarket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WarframeMarketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetWarframeMarket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetWarframeMarket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetWarframeMarket(ctx, req.(*WarframeMarketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetWarframeMarketRiven_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WarframeMarketRivenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetWarframeMarketRiven(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetWarframeMarketRiven_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetWarframeMarketRiven(ctx, req.(*WarframeMarketRivenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetWarframeMarketLich_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WarframeMarketLichReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetWarframeMarketLich(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetWarframeMarketLich_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetWarframeMarketLich(ctx, req.(*WarframeMarketLichReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetInvasions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetInvasions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetInvasions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetInvasions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetNightwaves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetNightwaves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetNightwaves_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetNightwaves(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetKuva_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetKuva(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetKuva_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetKuva(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetVoidStorm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetVoidStorm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetVoidStorm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetVoidStorm(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wfs_GetDailyDeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WfsServer).GetDailyDeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wfs_GetDailyDeal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WfsServer).GetDailyDeal(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Wfs_ServiceDesc is the grpc.ServiceDesc for Wfs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wfs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wfs.v1.Wfs",
	HandlerType: (*WfsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCycles",
			Handler:    _Wfs_GetCycles_Handler,
		},
		{
			MethodName: "GetAlerts",
			Handler:    _Wfs_GetAlerts_Handler,
		},
		{
			MethodName: "GetSorties",
			Handler:    _Wfs_GetSorties_Handler,
		},
		{
			MethodName: "GetVoidTrader",
			Handler:    _Wfs_GetVoidTrader_Handler,
		},
		{
			MethodName: "GetWarframeMarket",
			Handler:    _Wfs_GetWarframeMarket_Handler,
		},
		{
			MethodName: "GetWarframeMarketRiven",
			Handler:    _Wfs_GetWarframeMarketRiven_Handler,
		},
		{
			MethodName: "GetWarframeMarketLich",
			Handler:    _Wfs_GetWarframeMarketLich_Handler,
		},
		{
			MethodName: "GetInvasions",
			Handler:    _Wfs_GetInvasions_Handler,
		},
		{
			MethodName: "GetNightwaves",
			Handler:    _Wfs_GetNightwaves_Handler,
		},
		{
			MethodName: "GetKuva",
			Handler:    _Wfs_GetKuva_Handler,
		},
		{
			MethodName: "GetVoidStorm",
			Handler:    _Wfs_GetVoidStorm_Handler,
		},
		{
			MethodName: "GetDailyDeal",
			Handler:    _Wfs_GetDailyDeal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wfs/v1/wfs.proto",
}
