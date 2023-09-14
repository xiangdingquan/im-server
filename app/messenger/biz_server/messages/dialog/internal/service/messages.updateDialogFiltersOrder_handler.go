package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesUpdateDialogFiltersOrder(ctx context.Context, request *mtproto.TLMessagesUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.updateDialogFiltersOrder - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.updateDialogFiltersOrder - error: %v", err)
		return nil, err
	}

	s.PrivateFacade.UpdateDialogFiltersOrder(ctx, md.UserId, request.Order)

	log.Debugf("messages.updateDialogFiltersOrder - reply: {true}")
	return mtproto.BoolTrue, nil
}
