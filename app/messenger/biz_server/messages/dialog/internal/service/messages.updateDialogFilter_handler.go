package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesUpdateDialogFilter(ctx context.Context, request *mtproto.TLMessagesUpdateDialogFilter) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.updateDialogFilter - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.updateDialogFilter - error: %v", err)
		return nil, err
	}

	if request.GetFilter() == nil {
		s.PrivateFacade.DeleteDialogFilter(ctx, md.UserId, request.Id)
	} else {
		s.PrivateFacade.InsertOrUpdateDialogFilter(ctx, md.UserId, request.Id, request.GetFilter())
	}

	log.Debugf("messages.updateDialogFilter - reply: {true}")
	return mtproto.BoolTrue, nil
}
