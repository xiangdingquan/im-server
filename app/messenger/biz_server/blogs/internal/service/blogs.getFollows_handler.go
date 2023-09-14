package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.getFollows#b6152f28 peer:InputPeer offset:int limit:int = blogs.UserDates;
func (s *Service) BlogsGetFollows(ctx context.Context, request *mtproto.TLBlogsGetFollows) (*mtproto.Blogs_UserDates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getFollows - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peer *model.PeerUtil
		err  error
	)

	me := md.GetUserId()
	peer = model.FromInputPeer2(me, request.Peer)
	switch peer.PeerType {
	case model.PEER_SELF:
		peer.PeerType = model.PEER_USER
	case model.PEER_USER:
	default:
		log.Errorf("invalid peer: %v", request.Peer)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	if me != peer.PeerId {
		user, err := s.UserFacade.GetUserById(ctx, me, peer.PeerId)
		if err != nil || user == nil || user.Deleted {
			if err == nil {
				err = mtproto.ErrUserIdInvalid
			}
			log.Errorf("invalid user: %v", request.Peer)
			return nil, err
		}
	}

	reply, err := s.BlogFacade.GetFollows(ctx, me, peer.PeerId, request.GetOffset(), request.GetLimit())
	if err != nil {
		log.Errorf("blogs.getFollows#b6152f28 - error: %v", err)
		return nil, err
	}

	log.Debugf("blogs.getFollows#b6152f28 - reply: %s", reply.DebugString())
	return reply, err
}
