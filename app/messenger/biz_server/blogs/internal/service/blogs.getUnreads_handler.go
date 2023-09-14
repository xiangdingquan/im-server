package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.getUnreads#f00410ba pts:int limit:int = blogs.Unreads;
func (s *Service) BlogsGetUnreads(ctx context.Context, request *mtproto.TLBlogsGetUnreads) (*mtproto.Blogs_Unreads, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getUnreads - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	reply, err := s.BlogFacade.GetUnreads(ctx, md.GetUserId(), request.GetPts(), request.GetLimit())
	if err != nil {
		log.Errorf("blogs.getUnreads#f00410ba - error: %v", err)
		return nil, err
	}

	var uids model.IDList
	for _, update := range reply.NewUpdates {
		switch update.PredicateName {
		case mtproto.Predicate_updateNewBlog:
			uids.AddIfNot(update.GetBlog().GetUserId())
		case mtproto.Predicate_updateBlogFollow:
			uids.AddIfNot(update.GetUserId())
			uids.AddIfNot(update.GetTargetId())
		case mtproto.Predicate_updateBlogComment:
			uids.AddIfNot(update.GetUserId())
		case mtproto.Predicate_updateBlogLike:
			uids.AddIfNot(update.GetUserId())
		}
	}

	mutableUsers := s.UserFacade.GetMutableUsers(ctx, uids...)
	for _, id := range uids {
		user, _ := mutableUsers.ToUnsafeUser(md.GetUserId(), id)
		if user != nil {
			reply.Users = append(reply.Users, user)
		}
	}

	log.Debugf("blogs.getUnreads#f00410ba - reply: %s", reply.DebugString())
	return reply, err
}
