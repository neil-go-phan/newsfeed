// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: crawler.proto

package serverproto

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
	CrawlerService_TestRSSCrawler_FullMethodName    = "/crawlerproto.CrawlerService/TestRSSCrawler"
	CrawlerService_TestCustomCrawler_FullMethodName = "/crawlerproto.CrawlerService/TestCustomCrawler"
)

// CrawlerServiceClient is the client API for CrawlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrawlerServiceClient interface {
	TestRSSCrawler(ctx context.Context, in *Crawler, opts ...grpc.CallOption) (*TestResult, error)
	TestCustomCrawler(ctx context.Context, in *Crawler, opts ...grpc.CallOption) (*TestResult, error)
}

type crawlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrawlerServiceClient(cc grpc.ClientConnInterface) CrawlerServiceClient {
	return &crawlerServiceClient{cc}
}

func (c *crawlerServiceClient) TestRSSCrawler(ctx context.Context, in *Crawler, opts ...grpc.CallOption) (*TestResult, error) {
	out := new(TestResult)
	err := c.cc.Invoke(ctx, CrawlerService_TestRSSCrawler_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlerServiceClient) TestCustomCrawler(ctx context.Context, in *Crawler, opts ...grpc.CallOption) (*TestResult, error) {
	out := new(TestResult)
	err := c.cc.Invoke(ctx, CrawlerService_TestCustomCrawler_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrawlerServiceServer is the server API for CrawlerService service.
// All implementations must embed UnimplementedCrawlerServiceServer
// for forward compatibility
type CrawlerServiceServer interface {
	TestRSSCrawler(context.Context, *Crawler) (*TestResult, error)
	TestCustomCrawler(context.Context, *Crawler) (*TestResult, error)
	mustEmbedUnimplementedCrawlerServiceServer()
}

// UnimplementedCrawlerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCrawlerServiceServer struct {
}

func (UnimplementedCrawlerServiceServer) TestRSSCrawler(context.Context, *Crawler) (*TestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRSSCrawler not implemented")
}
func (UnimplementedCrawlerServiceServer) TestCustomCrawler(context.Context, *Crawler) (*TestResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestCustomCrawler not implemented")
}
func (UnimplementedCrawlerServiceServer) mustEmbedUnimplementedCrawlerServiceServer() {}

// UnsafeCrawlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrawlerServiceServer will
// result in compilation errors.
type UnsafeCrawlerServiceServer interface {
	mustEmbedUnimplementedCrawlerServiceServer()
}

func RegisterCrawlerServiceServer(s grpc.ServiceRegistrar, srv CrawlerServiceServer) {
	s.RegisterService(&CrawlerService_ServiceDesc, srv)
}

func _CrawlerService_TestRSSCrawler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Crawler)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServiceServer).TestRSSCrawler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrawlerService_TestRSSCrawler_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServiceServer).TestRSSCrawler(ctx, req.(*Crawler))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrawlerService_TestCustomCrawler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Crawler)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlerServiceServer).TestCustomCrawler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CrawlerService_TestCustomCrawler_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlerServiceServer).TestCustomCrawler(ctx, req.(*Crawler))
	}
	return interceptor(ctx, in, info, handler)
}

// CrawlerService_ServiceDesc is the grpc.ServiceDesc for CrawlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrawlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crawlerproto.CrawlerService",
	HandlerType: (*CrawlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestRSSCrawler",
			Handler:    _CrawlerService_TestRSSCrawler_Handler,
		},
		{
			MethodName: "TestCustomCrawler",
			Handler:    _CrawlerService_TestCustomCrawler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crawler.proto",
}
