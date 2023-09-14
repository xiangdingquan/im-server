package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReportEncryptedSpam(ctx context.Context, request *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.reportEncryptedSpam - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.reportEncryptedSpam - error: %v", err)
		return nil, err
	}

	// CHAT_ID_INVALID
	if request.Peer == nil {
		log.Errorf("messages.reportEncryptedSpam - peer is nil")
		return mtproto.ToBool(false), nil
	}

	s.ReportFacade.Report(ctx, md.UserId, model.MESSAGES_reportEncryptedSpam, model.PEER_ENCRYPTED_CHAT, request.Peer.ChatId, 0, 0, int32(model.REASON_SPAM), "")

	log.Debugf("messages.reportEncryptedSpam - reply: {true}")
	return mtproto.ToBool(true), nil
}
