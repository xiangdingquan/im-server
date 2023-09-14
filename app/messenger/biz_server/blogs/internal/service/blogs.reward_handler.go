package service

import (
	"context"

	"github.com/gogo/protobuf/proto"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.reward#a1a73acd flags:# user:int blog:int payment_password:flags.0?string amount:double = Updates;
func (s *Service) BlogsReward(ctx context.Context, request *mtproto.TLBlogsReward) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.reward - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	var (
		ok     bool
		err    error
		blog   *mtproto.MicroBlogs
		wallet *mtproto.Wallet_Info
	)

	me := md.GetUserId()
	if me == request.GetUser() {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	blog, err = s.BlogFacade.GetBlogs(ctx, me, []int32{request.GetBlog()})
	if err != nil || blog == nil {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	blogs := blog.GetBlogs()
	if len(blogs) == 0 || blogs[0] == nil {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	if blogs[0].GetId() != request.GetBlog() || blogs[0].GetUserId() != request.GetUser() {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	wallet, err = s.WalletFacade.GetInfo(ctx, me)
	if err != nil {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	if wallet == nil {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	if wallet.HasPassword {
		ok, err = s.WalletFacade.CheckPassword(ctx, me, request.GetPaymentPassword().GetValue())
		if err != nil || !ok {
			err = mtproto.ErrButtonTypeInvalid
			log.Errorf("blogs.reward: %v", err)
			return nil, err
		}
	}

	if wallet.Balance < request.GetAmount() {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	ok, err = s.WalletFacade.RewardBlog(ctx, me, request.GetUser(), request.GetBlog(), request.GetAmount())
	if err != nil {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	pushUser := request.GetUser()
	up := mtproto.MakeTLUpdateBlogReward(&mtproto.Update{
		UserId:   me,
		TargetId: pushUser,
		BlogId:   request.GetBlog(),
		Amount:   request.Amount,
	}).To_Update()
	updatesMe := model.MakeUpdatesHelper(up)

	reply := updatesMe.ToReplyUpdates(ctx, md.GetUserId(), s.UserFacade, nil, nil)
	sync_client.SyncUpdatesNotMe(ctx, md.GetUserId(), md.GetAuthId(), reply)
	if ok {
		up2 := proto.Clone(up).(*mtproto.Update)
		updatesUser := model.MakeUpdatesHelper(up2)
		sync_client.PushUpdates(ctx, pushUser, updatesUser.ToPushUpdates(ctx, pushUser, s.UserFacade, nil, nil))
	}

	log.Debugf("blogs.reward#a1a73acd - reply: %s", reply.DebugString())
	return reply, nil
}
