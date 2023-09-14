package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.reportPeer - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.reportPeer - error: %v", err)
		return nil, err
	}

	// Check peer invalid
	peer := model.FromInputPeer2(md.UserId, request.Peer)
	switch peer.PeerType {
	case model.PEER_USER:
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("account.reportPeer - error: %v", err)
		return nil, err
	}

	reason, text := model.FromReportReason(request.GetReason())
	s.ReportFacade.Report(ctx, md.UserId, model.ACCOUNTS_reportPeer, peer.PeerType, peer.PeerId, 0, 0, int32(reason), text)

	log.Debugf("account.reportPeer - reply: {true}")
	return mtproto.BoolTrue, nil
}
