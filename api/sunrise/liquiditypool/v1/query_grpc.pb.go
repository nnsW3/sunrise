// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sunrise/liquiditypool/v1/query.proto

package liquiditypoolv1

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
	Query_Params_FullMethodName  = "/sunrise.liquiditypool.v1.Query/Params"
	Query_Pair_FullMethodName    = "/sunrise.liquiditypool.v1.Query/Pair"
	Query_PairAll_FullMethodName = "/sunrise.liquiditypool.v1.Query/PairAll"
	Query_Pool_FullMethodName    = "/sunrise.liquiditypool.v1.Query/Pool"
	Query_PoolAll_FullMethodName = "/sunrise.liquiditypool.v1.Query/PoolAll"
	Query_Twap_FullMethodName    = "/sunrise.liquiditypool.v1.Query/Twap"
	Query_TwapAll_FullMethodName = "/sunrise.liquiditypool.v1.Query/TwapAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Pair items.
	Pair(ctx context.Context, in *QueryGetPairRequest, opts ...grpc.CallOption) (*QueryGetPairResponse, error)
	PairAll(ctx context.Context, in *QueryAllPairRequest, opts ...grpc.CallOption) (*QueryAllPairResponse, error)
	// Queries a list of Pool items.
	Pool(ctx context.Context, in *QueryGetPoolRequest, opts ...grpc.CallOption) (*QueryGetPoolResponse, error)
	PoolAll(ctx context.Context, in *QueryAllPoolRequest, opts ...grpc.CallOption) (*QueryAllPoolResponse, error)
	// Queries a list of Twap items.
	Twap(ctx context.Context, in *QueryGetTwapRequest, opts ...grpc.CallOption) (*QueryGetTwapResponse, error)
	TwapAll(ctx context.Context, in *QueryAllTwapRequest, opts ...grpc.CallOption) (*QueryAllTwapResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pair(ctx context.Context, in *QueryGetPairRequest, opts ...grpc.CallOption) (*QueryGetPairResponse, error) {
	out := new(QueryGetPairResponse)
	err := c.cc.Invoke(ctx, Query_Pair_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PairAll(ctx context.Context, in *QueryAllPairRequest, opts ...grpc.CallOption) (*QueryAllPairResponse, error) {
	out := new(QueryAllPairResponse)
	err := c.cc.Invoke(ctx, Query_PairAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pool(ctx context.Context, in *QueryGetPoolRequest, opts ...grpc.CallOption) (*QueryGetPoolResponse, error) {
	out := new(QueryGetPoolResponse)
	err := c.cc.Invoke(ctx, Query_Pool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PoolAll(ctx context.Context, in *QueryAllPoolRequest, opts ...grpc.CallOption) (*QueryAllPoolResponse, error) {
	out := new(QueryAllPoolResponse)
	err := c.cc.Invoke(ctx, Query_PoolAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Twap(ctx context.Context, in *QueryGetTwapRequest, opts ...grpc.CallOption) (*QueryGetTwapResponse, error) {
	out := new(QueryGetTwapResponse)
	err := c.cc.Invoke(ctx, Query_Twap_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TwapAll(ctx context.Context, in *QueryAllTwapRequest, opts ...grpc.CallOption) (*QueryAllTwapResponse, error) {
	out := new(QueryAllTwapResponse)
	err := c.cc.Invoke(ctx, Query_TwapAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Pair items.
	Pair(context.Context, *QueryGetPairRequest) (*QueryGetPairResponse, error)
	PairAll(context.Context, *QueryAllPairRequest) (*QueryAllPairResponse, error)
	// Queries a list of Pool items.
	Pool(context.Context, *QueryGetPoolRequest) (*QueryGetPoolResponse, error)
	PoolAll(context.Context, *QueryAllPoolRequest) (*QueryAllPoolResponse, error)
	// Queries a list of Twap items.
	Twap(context.Context, *QueryGetTwapRequest) (*QueryGetTwapResponse, error)
	TwapAll(context.Context, *QueryAllTwapRequest) (*QueryAllTwapResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Pair(context.Context, *QueryGetPairRequest) (*QueryGetPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pair not implemented")
}
func (UnimplementedQueryServer) PairAll(context.Context, *QueryAllPairRequest) (*QueryAllPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PairAll not implemented")
}
func (UnimplementedQueryServer) Pool(context.Context, *QueryGetPoolRequest) (*QueryGetPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pool not implemented")
}
func (UnimplementedQueryServer) PoolAll(context.Context, *QueryAllPoolRequest) (*QueryAllPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PoolAll not implemented")
}
func (UnimplementedQueryServer) Twap(context.Context, *QueryGetTwapRequest) (*QueryGetTwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Twap not implemented")
}
func (UnimplementedQueryServer) TwapAll(context.Context, *QueryAllTwapRequest) (*QueryAllTwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TwapAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Pair_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pair(ctx, req.(*QueryGetPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PairAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PairAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PairAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PairAll(ctx, req.(*QueryAllPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Pool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pool(ctx, req.(*QueryGetPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PoolAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PoolAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PoolAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PoolAll(ctx, req.(*QueryAllPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Twap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Twap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Twap_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Twap(ctx, req.(*QueryGetTwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TwapAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TwapAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TwapAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TwapAll(ctx, req.(*QueryAllTwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sunrise.liquiditypool.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Pair",
			Handler:    _Query_Pair_Handler,
		},
		{
			MethodName: "PairAll",
			Handler:    _Query_PairAll_Handler,
		},
		{
			MethodName: "Pool",
			Handler:    _Query_Pool_Handler,
		},
		{
			MethodName: "PoolAll",
			Handler:    _Query_PoolAll_Handler,
		},
		{
			MethodName: "Twap",
			Handler:    _Query_Twap_Handler,
		},
		{
			MethodName: "TwapAll",
			Handler:    _Query_TwapAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sunrise/liquiditypool/v1/query.proto",
}