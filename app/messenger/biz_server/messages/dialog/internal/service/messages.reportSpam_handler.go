package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReportSpam(ctx context.Context, request *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.reportSpam#cf1592db - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.reportSpam - error: %v", err)
		return nil, err
	}

	peer := model.FromInputPeer2(md.UserId, request.GetPeer())
	if peer.PeerType == model.PEER_EMPTY {
		return mtproto.ToBool(false), nil
	}

	if peer.PeerType == model.PEER_USER || peer.PeerType == model.PEER_CHAT {
	}

	s.ReportFacade.Report(ctx, md.UserId, model.MESSAGES_reportSpam, peer.PeerType, peer.PeerId, 0, 0, int32(model.REASON_SPAM), "")

	log.Debugf("messages.reportSpam#cf1592db - reply: {true}")
	return mtproto.ToBool(true), nil
}
