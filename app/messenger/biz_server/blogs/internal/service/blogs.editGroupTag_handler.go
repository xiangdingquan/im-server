package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/log"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.editGroupTag#12dd4287 tag_id:int title:string = Bool;
func (s *Service) BlogsEditGroupTag(ctx context.Context, request *mtproto.TLBlogsEditGroupTag) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.editGroupTag - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)
	if request.GetTagId() == 0 {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.createGroupTag: %v", err)
		return nil, err
	}

	me := md.GetUserId()
	_, err = s.BlogFacade.EditGroupTag(ctx, me, request.GetTagId(), request.GetTitle())
	if err != nil {
		log.Errorf("blogs.editGroupTag#12dd4287 - error: %v", err)
		return nil, err
	}

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateEditBlogGroupTag(&mtproto.Update{
		TagId: request.GetTagId(),
		Title: request.GetTitle(),
	}).To_Update())

	sync_client.SyncUpdatesNotMe(ctx, me, md.GetAuthId(), updatesHelper.ToReplyUpdates(ctx, me, s.UserFacade, nil, nil))

	reply := mtproto.BoolTrue
	log.Debugf("blogs.editGroupTag#12dd4287 - reply: %s", reply.DebugString())
	return reply, nil
}
