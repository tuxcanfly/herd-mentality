// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: hm/hm/query.proto

package hm

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
	Query_Params_FullMethodName      = "/hm.hm.Query/Params"
	Query_Question_FullMethodName    = "/hm.hm.Query/Question"
	Query_QuestionAll_FullMethodName = "/hm.hm.Query/QuestionAll"
	Query_Answer_FullMethodName      = "/hm.hm.Query/Answer"
	Query_AnswerAll_FullMethodName   = "/hm.hm.Query/AnswerAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Question items.
	Question(ctx context.Context, in *QueryGetQuestionRequest, opts ...grpc.CallOption) (*QueryGetQuestionResponse, error)
	QuestionAll(ctx context.Context, in *QueryAllQuestionRequest, opts ...grpc.CallOption) (*QueryAllQuestionResponse, error)
	// Queries a list of Answer items.
	Answer(ctx context.Context, in *QueryGetAnswerRequest, opts ...grpc.CallOption) (*QueryGetAnswerResponse, error)
	AnswerAll(ctx context.Context, in *QueryAllAnswerRequest, opts ...grpc.CallOption) (*QueryAllAnswerResponse, error)
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

func (c *queryClient) Question(ctx context.Context, in *QueryGetQuestionRequest, opts ...grpc.CallOption) (*QueryGetQuestionResponse, error) {
	out := new(QueryGetQuestionResponse)
	err := c.cc.Invoke(ctx, Query_Question_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QuestionAll(ctx context.Context, in *QueryAllQuestionRequest, opts ...grpc.CallOption) (*QueryAllQuestionResponse, error) {
	out := new(QueryAllQuestionResponse)
	err := c.cc.Invoke(ctx, Query_QuestionAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Answer(ctx context.Context, in *QueryGetAnswerRequest, opts ...grpc.CallOption) (*QueryGetAnswerResponse, error) {
	out := new(QueryGetAnswerResponse)
	err := c.cc.Invoke(ctx, Query_Answer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AnswerAll(ctx context.Context, in *QueryAllAnswerRequest, opts ...grpc.CallOption) (*QueryAllAnswerResponse, error) {
	out := new(QueryAllAnswerResponse)
	err := c.cc.Invoke(ctx, Query_AnswerAll_FullMethodName, in, out, opts...)
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
	// Queries a list of Question items.
	Question(context.Context, *QueryGetQuestionRequest) (*QueryGetQuestionResponse, error)
	QuestionAll(context.Context, *QueryAllQuestionRequest) (*QueryAllQuestionResponse, error)
	// Queries a list of Answer items.
	Answer(context.Context, *QueryGetAnswerRequest) (*QueryGetAnswerResponse, error)
	AnswerAll(context.Context, *QueryAllAnswerRequest) (*QueryAllAnswerResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Question(context.Context, *QueryGetQuestionRequest) (*QueryGetQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Question not implemented")
}
func (UnimplementedQueryServer) QuestionAll(context.Context, *QueryAllQuestionRequest) (*QueryAllQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuestionAll not implemented")
}
func (UnimplementedQueryServer) Answer(context.Context, *QueryGetAnswerRequest) (*QueryGetAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Answer not implemented")
}
func (UnimplementedQueryServer) AnswerAll(context.Context, *QueryAllAnswerRequest) (*QueryAllAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerAll not implemented")
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

func _Query_Question_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Question(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Question_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Question(ctx, req.(*QueryGetQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QuestionAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QuestionAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_QuestionAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QuestionAll(ctx, req.(*QueryAllQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Answer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Answer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Answer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Answer(ctx, req.(*QueryGetAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AnswerAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AnswerAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AnswerAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AnswerAll(ctx, req.(*QueryAllAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hm.hm.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Question",
			Handler:    _Query_Question_Handler,
		},
		{
			MethodName: "QuestionAll",
			Handler:    _Query_QuestionAll_Handler,
		},
		{
			MethodName: "Answer",
			Handler:    _Query_Answer_Handler,
		},
		{
			MethodName: "AnswerAll",
			Handler:    _Query_AnswerAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hm/hm/query.proto",
}
