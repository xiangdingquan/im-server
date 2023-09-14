package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesHideReportSpam(ctx context.Context, request *mtproto.TLMessagesHideReportSpam) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.hideReportSpam#a8f1709b - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	peer := model.FromInputPeer2(md.UserId, request.GetPeer())
	if peer.PeerType == model.PEER_USER || peer.PeerType == model.PEER_CHAT {

	}

	log.Debugf("messages.hideReportSpam#a8f1709b - reply: {true}")
	return mtproto.ToBool(true), nil
}
