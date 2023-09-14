package service

import (
	"context"

	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.getLikes#2490ae31 id:blogs.IdType offset:int limit:int = blogs.UserDates;
func (s *Service) BlogsGetLikes(ctx context.Context, request *mtproto.TLBlogsGetLikes) (*mtproto.Blogs_UserDates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getLikes - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)

	switch request.Id.PredicateName {
	case mtproto.Predicate_blogs_idTypeBlog:
	case mtproto.Predicate_blogs_idTypeComment:
	default:
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.getLikes: %v", err)
		return nil, err
	}

	reply, err := s.BlogFacade.GetLikes(ctx, md.GetUserId(), request.GetId(), request.GetOffsetId(), request.GetLimit())
	if err != nil {
		log.Errorf("blogs.getLikes#2490ae31 - error: %v", err)
		return nil, err
	}

	log.Debugf("blogs.getLikes#2490ae31 - reply: %s", reply.DebugString())
	return reply, err
}
