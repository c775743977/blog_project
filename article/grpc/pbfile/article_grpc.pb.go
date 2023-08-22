// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: article.proto

package pbfile

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

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	SaveArticleInfo(ctx context.Context, in *Article, opts ...grpc.CallOption) (*ArticleRes, error)
	GetArticlesInfo(ctx context.Context, in *ArticlesReq, opts ...grpc.CallOption) (*ArticlesRes, error)
	GetArticleInfo(ctx context.Context, in *ArticleReq, opts ...grpc.CallOption) (*ArticlesRes, error)
	CheckArtileExist(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckRes, error)
	GetArtilesByAuthor(ctx context.Context, in *AuthorReq, opts ...grpc.CallOption) (*ArticlesRes, error)
	DelArticle(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*IDReq, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) SaveArticleInfo(ctx context.Context, in *Article, opts ...grpc.CallOption) (*ArticleRes, error) {
	out := new(ArticleRes)
	err := c.cc.Invoke(ctx, "/column.ArticleService/SaveArticleInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticlesInfo(ctx context.Context, in *ArticlesReq, opts ...grpc.CallOption) (*ArticlesRes, error) {
	out := new(ArticlesRes)
	err := c.cc.Invoke(ctx, "/column.ArticleService/GetArticlesInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticleInfo(ctx context.Context, in *ArticleReq, opts ...grpc.CallOption) (*ArticlesRes, error) {
	out := new(ArticlesRes)
	err := c.cc.Invoke(ctx, "/column.ArticleService/GetArticleInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) CheckArtileExist(ctx context.Context, in *CheckReq, opts ...grpc.CallOption) (*CheckRes, error) {
	out := new(CheckRes)
	err := c.cc.Invoke(ctx, "/column.ArticleService/CheckArtileExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArtilesByAuthor(ctx context.Context, in *AuthorReq, opts ...grpc.CallOption) (*ArticlesRes, error) {
	out := new(ArticlesRes)
	err := c.cc.Invoke(ctx, "/column.ArticleService/GetArtilesByAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) DelArticle(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*IDReq, error) {
	out := new(IDReq)
	err := c.cc.Invoke(ctx, "/column.ArticleService/DelArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations must embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	SaveArticleInfo(context.Context, *Article) (*ArticleRes, error)
	GetArticlesInfo(context.Context, *ArticlesReq) (*ArticlesRes, error)
	GetArticleInfo(context.Context, *ArticleReq) (*ArticlesRes, error)
	CheckArtileExist(context.Context, *CheckReq) (*CheckRes, error)
	GetArtilesByAuthor(context.Context, *AuthorReq) (*ArticlesRes, error)
	DelArticle(context.Context, *IDReq) (*IDReq, error)
	mustEmbedUnimplementedArticleServiceServer()
}

// UnimplementedArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) SaveArticleInfo(context.Context, *Article) (*ArticleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveArticleInfo not implemented")
}
func (UnimplementedArticleServiceServer) GetArticlesInfo(context.Context, *ArticlesReq) (*ArticlesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesInfo not implemented")
}
func (UnimplementedArticleServiceServer) GetArticleInfo(context.Context, *ArticleReq) (*ArticlesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleInfo not implemented")
}
func (UnimplementedArticleServiceServer) CheckArtileExist(context.Context, *CheckReq) (*CheckRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckArtileExist not implemented")
}
func (UnimplementedArticleServiceServer) GetArtilesByAuthor(context.Context, *AuthorReq) (*ArticlesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtilesByAuthor not implemented")
}
func (UnimplementedArticleServiceServer) DelArticle(context.Context, *IDReq) (*IDReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelArticle not implemented")
}
func (UnimplementedArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_SaveArticleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).SaveArticleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/column.ArticleService/SaveArticleInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).SaveArticleInfo(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticlesInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticlesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticlesInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/column.ArticleService/GetArticlesInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticlesInfo(ctx, req.(*ArticlesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticleInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticleInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/column.ArticleService/GetArticleInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticleInfo(ctx, req.(*ArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_CheckArtileExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CheckArtileExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/column.ArticleService/CheckArtileExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CheckArtileExist(ctx, req.(*CheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArtilesByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArtilesByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/column.ArticleService/GetArtilesByAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArtilesByAuthor(ctx, req.(*AuthorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_DelArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).DelArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/column.ArticleService/DelArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).DelArticle(ctx, req.(*IDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "column.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveArticleInfo",
			Handler:    _ArticleService_SaveArticleInfo_Handler,
		},
		{
			MethodName: "GetArticlesInfo",
			Handler:    _ArticleService_GetArticlesInfo_Handler,
		},
		{
			MethodName: "GetArticleInfo",
			Handler:    _ArticleService_GetArticleInfo_Handler,
		},
		{
			MethodName: "CheckArtileExist",
			Handler:    _ArticleService_CheckArtileExist_Handler,
		},
		{
			MethodName: "GetArtilesByAuthor",
			Handler:    _ArticleService_GetArtilesByAuthor_Handler,
		},
		{
			MethodName: "DelArticle",
			Handler:    _ArticleService_DelArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}