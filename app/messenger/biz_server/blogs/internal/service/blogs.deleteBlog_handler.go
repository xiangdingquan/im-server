package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/log"

	sync_client "open.chat/app/messenger/sync/client"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.deleteBlog#c3e55a42 blogs:Vector<int> = Updates;
func (s *Service) BlogsDeleteBlog(ctx context.Context, request *mtproto.TLBlogsDeleteBlog) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.deleteBlog - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err error
	)
	if len(request.GetBlogs()) == 0 {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.deleteBlog: %v", err)
		return nil, err
	}

	me := md.GetUserId()
	deleteIds, shieldIds, err := s.BlogFacade.DeleteBlog(ctx, me, request.GetBlogs())
	if err != nil {
		log.Errorf("blogs.deleteBlog#c3e55a42 - error: %v", err)
		return nil, err
	}

	totalIds := append(deleteIds, shieldIds...)
	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateDeleteBlog(&mtproto.Update{
		Blogs:     totalIds,
		Pts_INT32: int32(idgen.NextBlogNPtsId(ctx, me, len(totalIds))),
		PtsCount:  int32(len(totalIds)),
	}).To_Update())

	ups := updatesHelper.ToReplyUpdates(ctx, me, s.UserFacade, nil, nil)
	sync_client.SyncUpdatesNotMe(ctx, me, md.GetAuthId(), ups)

	log.Debugf("blogs.deleteBlog#c3e55a42 - reply: %s", ups.DebugString())
	return ups, nil
}
