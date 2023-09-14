package service

import (
	"context"

	"open.chat/model"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetPeerSettings(ctx context.Context, request *mtproto.TLMessagesGetPeerSettings) (*mtproto.PeerSettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getPeerSettings#3672e09c - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	peer := model.FromInputPeer2(md.UserId, request.Peer)
	// 400	CHANNEL_INVALID	The provided channel is invalid
	// 400	PEER_ID_INVALID	The provided peer id is invalid
	peerSettings, err := s.UserFacade.GetPeerSettings(ctx, md.UserId, peer)
	if err != nil {
		log.Errorf("messages.getPeerSettings - error: %v", err)
		peerSettings = mtproto.MakeTLPeerSettings(nil).To_PeerSettings()
	}

	log.Debugf("messages.getPeerSettings#3672e09c - reply: %s", peerSettings.DebugString())
	return peerSettings, nil
}
