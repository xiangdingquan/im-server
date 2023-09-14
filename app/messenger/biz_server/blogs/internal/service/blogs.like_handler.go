package service

import (
	"context"
	"github.com/pkg/errors"

	"github.com/gogo/protobuf/proto"
	"open.chat/model"
	"open.chat/pkg/log"

	sync_client "open.chat/app/messenger/sync/client"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.like#4c835b23 flags:# id:blogs.IdType liked:flags.0?true = Updates;
func (s *Service) BlogsLike(ctx context.Context, request *mtproto.TLBlogsLike) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.like - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err        error
		blogIdType *mtproto.Blogs_IdType
	)

	isBanned, err := s.BlogFacade.IsBannedUser(ctx, md.UserId)
	if err != nil {
		log.Errorf("blogs.sendComment - check banned user, %v", err)
		return nil, err
	}
	if isBanned {
		log.Errorf("blogs.sendComment - %d banned", md.UserId)
		return nil, errors.New("banned")
	}

	switch request.GetId().GetPredicateName() {
	case mtproto.Predicate_blogs_idTypeBlog:
		blogIdType = mtproto.MakeTLBlogsIdTypeBlog(&mtproto.Blogs_IdType{
			BlogId: request.GetId().GetBlogId(),
		}).To_Blogs_IdType()
	case mtproto.Predicate_blogs_idTypeComment:
		blogIdType = mtproto.MakeTLBlogsIdTypeComment(&mtproto.Blogs_IdType{
			CommentId: request.GetId().GetCommentId(),
		}).To_Blogs_IdType()
	default:
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.like: %v", err)
		return nil, err
	}

	me := md.GetUserId()
	liked := request.GetLiked()
	pushUser, like, err := s.BlogFacade.Like(ctx, me, blogIdType, liked)
	if err != nil {
		log.Errorf("blogs.like#4c835b23 - error: %v", err)
		return nil, err
	}

	up := mtproto.MakeTLUpdateBlogLike(&mtproto.Update{
		BlogIdType: blogIdType,
		UserId:     like.GetUserId(),
		Liked:      liked,
		Date:       like.GetDate(),
		Pts_INT32:  int32(idgen.NextBlogPtsId(ctx, me)),
		PtsCount:   1,
	}).To_Update()
	updatesMe := model.MakeUpdatesHelper(up)

	reply := updatesMe.ToReplyUpdates(ctx, me, s.UserFacade, nil, nil)
	sync_client.SyncUpdatesNotMe(ctx, me, md.GetAuthId(), reply)
	if liked && pushUser != 0 && pushUser != me {
		up2 := proto.Clone(up).(*mtproto.Update)
		up2.Pts_INT32 = int32(idgen.NextBlogPtsId(ctx, pushUser))
		updatesUser := model.MakeUpdatesHelper(up2)
		sync_client.PushUpdates(ctx, pushUser, updatesUser.ToPushUpdates(ctx, pushUser, s.UserFacade, nil, nil))
	}

	log.Debugf("blogs.like#4c835b23 - reply: %s", reply.DebugString())
	return reply, err
}
