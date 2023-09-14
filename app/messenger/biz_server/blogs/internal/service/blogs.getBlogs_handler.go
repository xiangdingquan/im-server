package service

import (
	"context"

	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.getBlogs#7f44d787 blogs:Vector<int> = MicroBlogs;
func (s *Service) BlogsGetBlogs(ctx context.Context, request *mtproto.TLBlogsGetBlogs) (*mtproto.MicroBlogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getBlogs - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)

	if len(request.GetBlogs()) == 0 {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.getBlogs: %v", err)
		return nil, err
	}

	reply, err := s.BlogFacade.GetBlogs(ctx, md.GetUserId(), request.GetBlogs())
	if err != nil {
		log.Errorf("blogs.getBlogs#7f44d787 - error: %v", err)
		return nil, err
	}

	log.Debugf("blogs.getBlogs#7f44d787 - reply: %s", reply.DebugString())
	return reply, err
}
