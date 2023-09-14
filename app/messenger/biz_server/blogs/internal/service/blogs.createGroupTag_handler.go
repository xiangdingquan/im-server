package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/log"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.createGroupTag#64b2a32e title:string users:Vector<int> = Updates;
func (s *Service) BlogsCreateGroupTag(ctx context.Context, request *mtproto.TLBlogsCreateGroupTag) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.createGroupTag - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)
	/*
		groupTag := mtproto.MakeTLBlogsGroupTag(&mtproto.Blogs_GroupTag{
			TagId: 100001,
			Title: request.GetTitle(),
			Users: request.GetUsers(),
		}).To_Blogs_GroupTag()
		updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateNewBlogGroupTag(&mtproto.Update{
			GroupTag: groupTag,
		}).To_Update())
		ups := updatesHelper.ToPushUpdates(ctx, md.GetUserId(), nil, nil, nil)
		sync_client.PushUpdates(ctx, md.GetUserId(), ups)

		groupTag.TagId = 100002
		updatesHelper = model.MakeUpdatesHelper(mtproto.MakeTLUpdateNewBlogGroupTag(&mtproto.Update{
			GroupTag: groupTag,
		}).To_Update())
		ups = updatesHelper.ToReplyUpdates(ctx, md.GetUserId(), nil, nil, nil)
		return ups, nil
	*/

	if len(request.GetUsers()) == 0 {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.createGroupTag: %v", err)
		return nil, err
	}

	me := md.GetUserId()
	users := s.UserFacade.GetUserListByIdList(ctx, me, request.GetUsers())
	userIds := make([]int32, 0, len(users))
	for _, user := range users {
		if !user.Deleted && user.GetId() != me {
			userIds = append(userIds, user.GetId())
		}
	}

	if len(userIds) == 0 {
		err = mtproto.ErrUserIdInvalid
		log.Errorf("blogs.addGroupTagMember#56dabf3a - error: %s\n", err.Error())
		return nil, err
	}

	groupTag, err := s.BlogFacade.CreateGroupTag(ctx, me, request.GetTitle(), request.GetUsers())
	if err != nil {
		log.Errorf("blogs.createGroupTag#64b2a32e - error: %v", err)
		return nil, err
	}

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateNewBlogGroupTag(&mtproto.Update{
		GroupTag: groupTag,
	}).To_Update())

	ups := updatesHelper.ToReplyUpdates(ctx, me, s.UserFacade, nil, nil)
	sync_client.SyncUpdatesNotMe(ctx, me, md.GetAuthId(), ups)

	log.Debugf("blogs.createGroupTag#64b2a32e - reply: %s", ups.DebugString())
	return ups, nil
}
