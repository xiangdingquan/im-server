package service

import (
	"context"

	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.getComments#5d4ce4dd id:blogs.IdType offset_id:int limit:int = blogs.Comments;
func (s *Service) BlogsGetComments(ctx context.Context, request *mtproto.TLBlogsGetComments) (*mtproto.Blogs_Comments, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getComments - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)

	switch request.Id.PredicateName {
	case mtproto.Predicate_blogs_idTypeBlog:
	case mtproto.Predicate_blogs_idTypeComment:
	default:
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.getComments: %v", err)
		return nil, err
	}

	reply, err := s.BlogFacade.GetComments(ctx, md.GetUserId(), request.GetId(), request.GetOffsetId(), request.GetLimit())
	if err != nil {
		log.Errorf("blogs.getComments#5d4ce4dd - error: %v", err)
		return nil, err
	}

	log.Debugf("blogs.getComments#5d4ce4dd - reply: %s", reply.DebugString())
	return reply, err
}
