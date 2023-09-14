package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.readHistory#a3c20ca8 flags:# pts:flags.0?int read_id:flags.1?blogs.IdType = Bool;
func (s *Service) BlogsReadHistory(ctx context.Context, request *mtproto.TLBlogsReadHistory) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.readHistory - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		maxReadId *mtproto.Blogs_IdType
		err       error
	)

	maxReadId, err = s.BlogFacade.ReadHistory(ctx, md.GetUserId(), request.GetPts().GetValue(), request.GetReadId())
	if err != nil {
		log.Errorf("blogs.readHistory#a3c20ca8 - error: %v", err)
		return nil, err
	}

	if maxReadId != nil {
		updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateBlogReadHistory(&mtproto.Update{
			MaxReadIdType: maxReadId,
		}).To_Update())
		sync_client.SyncUpdatesNotMe(ctx, md.GetUserId(), md.GetAuthId(), updatesHelper.ToReplyUpdates(ctx, md.GetUserId(), s.UserFacade, nil, nil))
	}

	reply := mtproto.BoolTrue
	log.Debugf("blogs.readHistory#a3c20ca8 - reply: %s", reply.DebugString())
	return reply, nil
}
