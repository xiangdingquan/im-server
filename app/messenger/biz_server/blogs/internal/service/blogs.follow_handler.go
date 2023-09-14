package service

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"open.chat/model"
	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"

	sync_client "open.chat/app/messenger/sync/client"
	idgen "open.chat/app/service/idgen/client"
)

// blogs.follow#45d768c1 flags:# peer:InputPeer followed:flags.0?true = Updates;
func (s *Service) BlogsFollow(ctx context.Context, request *mtproto.TLBlogsFollow) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.follow - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peer *model.PeerUtil
		err  error
	)
	peer = model.FromInputPeer2(md.UserId, request.Peer)
	switch peer.PeerType {
	case model.PEER_USER:
		if peer.PeerId == md.UserId {
			log.Errorf("invalid peer: %v", request.Peer)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}
	default:
		log.Errorf("invalid peer: %v", request.Peer)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	if request.GetFollowed() {
		user, err := s.UserFacade.GetUserById(ctx, md.GetUserId(), peer.PeerId)
		if err != nil || user == nil || user.Deleted {
			if err == nil {
				err = mtproto.ErrUserIdInvalid
			}
			log.Errorf("invalid user: %v", request.Peer)
			return nil, err
		}
	}

	me := md.GetUserId()
	followed := request.GetFollowed()
	ok, follow, err := s.BlogFacade.Follow(ctx, me, peer.PeerId, followed)
	if err != nil {
		log.Errorf("blogs.follow#45d768c1 - error: %v", err)
		return nil, err
	}

	pushUser := peer.PeerId
	up := mtproto.MakeTLUpdateBlogFollow(&mtproto.Update{
		UserId:    follow.GetUserId(),
		TargetId:  pushUser,
		Followed:  followed,
		Date:      follow.Date,
		Pts_INT32: int32(idgen.NextBlogPtsId(ctx, me)),
		PtsCount:  1,
	}).To_Update()
	updatesMe := model.MakeUpdatesHelper(up)

	reply := updatesMe.ToReplyUpdates(ctx, md.GetUserId(), s.UserFacade, nil, nil)
	sync_client.SyncUpdatesNotMe(ctx, md.GetUserId(), md.GetAuthId(), reply)
	if ok {
		up2 := proto.Clone(up).(*mtproto.Update)
		up2.Pts_INT32 = int32(idgen.NextBlogPtsId(ctx, pushUser))
		updatesUser := model.MakeUpdatesHelper(up2)
		sync_client.PushUpdates(ctx, pushUser, updatesUser.ToPushUpdates(ctx, pushUser, s.UserFacade, nil, nil))
	}

	log.Debugf("blogs.follow#45d768c1 - reply: %s", reply.DebugString())
	return reply, nil
}
