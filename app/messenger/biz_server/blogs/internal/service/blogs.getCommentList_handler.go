package service

import (
	"context"

	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.getCommentList#6af5634f comments:Vector<int> = blogs.Comments;
func (s *Service) BlogsGetCommentList(ctx context.Context, request *mtproto.TLBlogsGetCommentList) (*mtproto.Blogs_Comments, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getCommentList - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)

	if len(request.GetComments()) == 0 {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.getComments: %v", err)
		return nil, err
	}

	reply, err := s.BlogFacade.GetCommentList(ctx, md.GetUserId(), request.GetComments())
	if err != nil {
		log.Errorf("blogs.getCommentList#6af5634f - error: %v", err)
		return nil, err
	}

	log.Debugf("blogs.getCommentList#6af5634f - reply: %s", reply.DebugString())
	return reply, err
}
