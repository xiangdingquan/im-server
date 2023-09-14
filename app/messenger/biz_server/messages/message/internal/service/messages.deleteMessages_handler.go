package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesDeleteMessages(ctx context.Context, request *mtproto.TLMessagesDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.deleteMessages#e58e95d2 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	affectedMessages, err := s.MsgFacade.DeleteMessages(ctx,
		md.UserId,
		md.AuthId,
		&model.PeerUtil{PeerType: model.PEER_EMPTY},
		request.Revoke,
		request.Id)

	if err != nil {
		log.Errorf("messages.deleteMessages#e58e95d2 - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.deleteMessages#e58e95d2 - reply: %s", affectedMessages.DebugString())
	return affectedMessages, nil
}
