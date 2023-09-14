package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/log"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.deleteGroupTag#21dffb22 tag_ids:Vector<int> = Updates;
func (s *Service) BlogsDeleteGroupTag(ctx context.Context, request *mtproto.TLBlogsDeleteGroupTag) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.deleteGroupTag - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)

	if len(request.GetTagIds()) == 0 {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.createGroupTag: %v", err)
		return nil, err
	}

	me := md.GetUserId()
	Ids, err := s.BlogFacade.DeleteGroupTag(ctx, me, request.GetTagIds())
	if err != nil {
		log.Errorf("blogs.deleteGroupTag#21dffb22 - error: %v", err)
		return nil, err
	}

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateDeleteBlogGroupTag(&mtproto.Update{
		TagIds: Ids,
	}).To_Update())

	ups := updatesHelper.ToReplyUpdates(ctx, me, s.UserFacade, nil, nil)
	sync_client.SyncUpdatesNotMe(ctx, me, md.GetAuthId(), ups)

	log.Debugf("blogs.deleteGroupTag#21dffb22 - reply: %s", ups.DebugString())
	return ups, nil
}
