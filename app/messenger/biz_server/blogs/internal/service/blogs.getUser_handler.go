package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.getUser#dd7cabf8 peer:InputPeer = blogs.User;
func (s *Service) BlogsGetUser(ctx context.Context, request *mtproto.TLBlogsGetUser) (*mtproto.Blogs_User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getUser - metadata: %s, request: %s", md.DebugString(), request.DebugString())

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

	reply, err := s.BlogFacade.GetUser(ctx, me, peer.PeerId)
	if err != nil {
		log.Errorf("blogs.getUser#dd7cabf8 - error: %v", err)
		return nil, err
	}

	log.Debugf("blogs.getUser#dd7cabf8 - reply: %s", reply.DebugString())
	return reply, err

}
