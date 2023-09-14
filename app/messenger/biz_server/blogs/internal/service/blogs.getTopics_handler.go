package service

import (
	"context"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) BlogsGetTopics(ctx context.Context, request *mtproto.TLBlogsGetTopics) (*mtproto.Blogs_Topics, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getTopics - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return s.BlogFacade.GetTopics(ctx, request.GetFromTopicId(), request.GetLimit())
}
