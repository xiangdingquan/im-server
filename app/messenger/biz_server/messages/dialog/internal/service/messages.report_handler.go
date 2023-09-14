package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReport(ctx context.Context, request *mtproto.TLMessagesReport) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.report#bd82b658 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.report - error: %v", err)
		return nil, err
	}

	if len(request.Id) == 0 {
		log.Errorf("messages.report#bd82b658 error - id empty")
		return mtproto.ToBool(false), nil
	}

	peer := model.FromInputPeer2(md.UserId, request.Peer)
	if peer.IsEmpty() || peer.IsEmpty() {
		log.Errorf("messages.report#bd82b658 error - peer is empty or self")
		return mtproto.ToBool(false), nil
	}

	r, text := model.FromReportReason(request.Reason)
	s.ReportFacade.ReportIdList(ctx, md.UserId, model.MESSAGES_report, peer.PeerType, peer.PeerId, 0, request.Id, int32(r), text)

	log.Debugf("messages.report#bd82b658 - reply: {true}")
	return mtproto.ToBool(true), nil
}
