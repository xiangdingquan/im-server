package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/log"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.addGroupTagMember#56dabf3a tag_id:int users:Vector<int> = Updates;
func (s *Service) BlogsAddGroupTagMember(ctx context.Context, request *mtproto.TLBlogsAddGroupTagMember) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.addGroupTagMember - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)
	if request.GetTagId() == 0 || len(request.GetUsers()) == 0 {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.addGroupTagMember: %v", err)
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

	Ids, err := s.BlogFacade.AddGroupTagMember(ctx, me, request.GetTagId(), request.GetUsers())
	if err != nil {
		log.Errorf("blogs.addGroupTagMember#56dabf3a - error: %v", err)
		return nil, err
	}

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateAddBlogGroupTagMember(&mtproto.Update{
		TagId: request.GetTagId(),
		Users: Ids,
	}).To_Update())

	ups := updatesHelper.ToReplyUpdates(ctx, me, s.UserFacade, nil, nil)
	sync_client.SyncUpdatesNotMe(ctx, me, md.GetAuthId(), ups)

	log.Debugf("blogs.addGroupTagMember#56dabf3a - reply: %s", ups.DebugString())
	return ups, nil
}
