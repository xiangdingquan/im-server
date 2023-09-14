package service

import (
	"context"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) BlogsGetHotTopics(ctx context.Context, request *mtproto.TLBlogsGetHotTopics) (*mtproto.Blogs_Topics, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getHotTopics - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	topics, err := s.BlogFacade.GetHotTopics(ctx, request.GetLimit())
	if err != nil {
		log.Errorf("blogs.getHotTopics, error: %v", err)
		return nil, err
	}

	log.Debugf("blogs.getHotTopics - reply: %s", topics.DebugString())
	return topics, nil
}
