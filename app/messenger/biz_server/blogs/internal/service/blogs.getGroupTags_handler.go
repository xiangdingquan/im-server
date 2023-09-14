package service

import (
	"context"

	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.getGroupTags#c636efb6 = blogs.GroupTags;
func (s *Service) BlogsGetGroupTags(ctx context.Context, request *mtproto.TLBlogsGetGroupTags) (*mtproto.Blogs_GroupTags, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getGroupTags - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)
	reply, err := s.BlogFacade.GetGroupTags(ctx, md.GetUserId())
	if err != nil {
		log.Errorf("blogs.getGroupTags#c636efb6 - error: %v", err)
		return nil, err
	}

	log.Debugf("blogs.getGroupTags#c636efb6 - reply: %s", reply.DebugString())
	return reply, err
}
