package webpagepb

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	mtproto "open.chat/mtproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCWebPageClient is the client API for RPCWebPage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCWebPageClient interface {
	// // messages.getWebPagePreview#8b68b0cc flags:# message:string entities:flags.3?Vector<MessageEntity> = MessageMedia;
	GetPendingWebPagePreview(ctx context.Context, in *mtproto.TLMessagesGetWebPagePreview, opts ...grpc.CallOption) (*mtproto.WebPage, error)
	// messages.getWebPagePreview#8b68b0cc flags:# message:string entities:flags.3?Vector<MessageEntity> = MessageMedia;
	GetWebPagePreview(ctx context.Context, in *mtproto.TLMessagesGetWebPagePreview, opts ...grpc.CallOption) (*mtproto.WebPage, error)
	// messages.getWebPage#32ca8f91 url:string hash:int = WebPage;
	GetWebPage(ctx context.Context, in *mtproto.TLMessagesGetWebPage, opts ...grpc.CallOption) (*mtproto.WebPage, error)
}

type rPCWebPageClient struct {
	cc *grpc.ClientConn
}

func NewRPCWebPageClient(cc *grpc.ClientConn) RPCWebPageClient {
	return &rPCWebPageClient{cc}
}

func (c *rPCWebPageClient) GetPendingWebPagePreview(ctx context.Context, in *mtproto.TLMessagesGetWebPagePreview, opts ...grpc.CallOption) (*mtproto.WebPage, error) {
	out := new(mtproto.WebPage)
	err := c.cc.Invoke(ctx, "/mediapb.RPCWebPage/GetPendingWebPagePreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCWebPageClient) GetWebPagePreview(ctx context.Context, in *mtproto.TLMessagesGetWebPagePreview, opts ...grpc.CallOption) (*mtproto.WebPage, error) {
	out := new(mtproto.WebPage)
	err := c.cc.Invoke(ctx, "/mediapb.RPCWebPage/GetWebPagePreview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCWebPageClient) GetWebPage(ctx context.Context, in *mtproto.TLMessagesGetWebPage, opts ...grpc.CallOption) (*mtproto.WebPage, error) {
	out := new(mtproto.WebPage)
	err := c.cc.Invoke(ctx, "/mediapb.RPCWebPage/GetWebPage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCWebPage service

type RPCWebPageServer interface {
	// // messages.getWebPagePreview#8b68b0cc flags:# message:string entities:flags.3?Vector<MessageEntity> = MessageMedia;
	GetPendingWebPagePreview(context.Context, *mtproto.TLMessagesGetWebPagePreview) (*mtproto.WebPage, error)
	// messages.getWebPagePreview#8b68b0cc flags:# message:string entities:flags.3?Vector<MessageEntity> = MessageMedia;
	GetWebPagePreview(context.Context, *mtproto.TLMessagesGetWebPagePreview) (*mtproto.WebPage, error)
	// messages.getWebPage#32ca8f91 url:string hash:int = WebPage;
	GetWebPage(context.Context, *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error)
}

func RegisterRPCWebPageServer(s *grpc.Server, srv RPCWebPageServer) {
	s.RegisterService(&_RPCWebPage_serviceDesc, srv)
}

func _RPCWebPage_GetPendingWebPagePreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(mtproto.TLMessagesGetWebPagePreview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCWebPageServer).GetPendingWebPagePreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCWebPage/GetPendingWebPagePreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCWebPageServer).GetPendingWebPagePreview(ctx, req.(*mtproto.TLMessagesGetWebPagePreview))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCWebPage_GetWebPagePreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(mtproto.TLMessagesGetWebPagePreview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCWebPageServer).GetWebPagePreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCWebPage/GetWebPagePreview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCWebPageServer).GetWebPagePreview(ctx, req.(*mtproto.TLMessagesGetWebPagePreview))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCWebPage_GetWebPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(mtproto.TLMessagesGetWebPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCWebPageServer).GetWebPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mediapb.RPCWebPage/GetWebPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCWebPageServer).GetWebPage(ctx, req.(*mtproto.TLMessagesGetWebPage))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCWebPage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mediapb.RPCWebPage",
	HandlerType: (*RPCWebPageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPendingWebPagePreview",
			Handler:    _RPCWebPage_GetPendingWebPagePreview_Handler,
		},
		{
			MethodName: "GetWebPagePreview",
			Handler:    _RPCWebPage_GetWebPagePreview_Handler,
		},
		{
			MethodName: "GetWebPage",
			Handler:    _RPCWebPage_GetWebPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webpagepb.tl.proto",
}

func init() { proto.RegisterFile("webpagepb.tl.proto", fileDescriptor_webpagepb_tl_4acaae32f4517026) }

var fileDescriptor_webpagepb_tl_4acaae32f4517026 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xbd, 0x4a, 0x04, 0x31,
	0x14, 0x85, 0xb5, 0x51, 0x48, 0xa5, 0xc1, 0x42, 0x06, 0x41, 0xb0, 0xb0, 0x33, 0x01, 0x17, 0x1f,
	0x40, 0x2d, 0xb6, 0x11, 0x1c, 0x44, 0x14, 0x6c, 0x86, 0x24, 0x7b, 0xbc, 0x13, 0x98, 0xfc, 0x90,
	0x64, 0x76, 0xf1, 0xcd, 0x2d, 0x65, 0xb3, 0xc3, 0xca, 0x8a, 0x60, 0x61, 0x77, 0xef, 0xf9, 0xf9,
	0x9a, 0xc3, 0xf8, 0x0a, 0x3a, 0x2a, 0x42, 0xd4, 0xa2, 0x0c, 0x22, 0xa6, 0x50, 0x02, 0x3f, 0x74,
	0x58, 0x58, 0x15, 0x75, 0x73, 0x45, 0xb6, 0xf4, 0xa3, 0x16, 0x26, 0x38, 0x49, 0x81, 0x82, 0xac,
	0xbe, 0x1e, 0xdf, 0xeb, 0x57, 0x9f, 0x7a, 0x6d, 0x7a, 0xcd, 0x49, 0x36, 0x3d, 0x9c, 0x5a, 0x83,
	0xf2, 0x87, 0x37, 0x93, 0x7a, 0xb6, 0xab, 0x76, 0x19, 0x69, 0x69, 0x0d, 0x36, 0xee, 0xf5, 0xe7,
	0x3e, 0x63, 0x4f, 0xed, 0xfd, 0x2b, 0x74, 0xab, 0x08, 0xfc, 0x85, 0x9d, 0xce, 0x51, 0x5a, 0xf8,
	0x85, 0xf5, 0x34, 0x89, 0x6d, 0xc2, 0xd2, 0x62, 0xc5, 0x2f, 0x85, 0x2b, 0xb5, 0x24, 0x9e, 0x1f,
	0x3a, 0x87, 0x9c, 0x15, 0x21, 0x77, 0x84, 0xb2, 0x9b, 0x6b, 0x8e, 0xb6, 0xb9, 0xc9, 0xb8, 0xd8,
	0xe3, 0x8f, 0xec, 0x78, 0xfe, 0x33, 0xf8, 0x2f, 0xe0, 0x2d, 0x63, 0xdf, 0x40, 0x7e, 0xfe, 0x07,
	0xe9, 0x37, 0xc4, 0xdd, 0xcd, 0xdb, 0xcc, 0x43, 0x8f, 0x83, 0x12, 0xa6, 0x57, 0x45, 0x9a, 0x61,
	0xcc, 0x05, 0x49, 0xaa, 0x18, 0xe5, 0x1a, 0x01, 0x4f, 0x48, 0x72, 0x9a, 0x47, 0x6e, 0x67, 0xd2,
	0x07, 0x95, 0x33, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x40, 0x3a, 0x32, 0x8c, 0xba, 0x01, 0x00,
	0x00,
}
