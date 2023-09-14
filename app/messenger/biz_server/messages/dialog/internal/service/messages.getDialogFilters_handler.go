package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetDialogFilters(ctx context.Context, request *mtproto.TLMessagesGetDialogFilters) (*mtproto.Vector_DialogFilter, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getDialogFilters - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getDialogFilters - error: %v", err)
		return nil, err
	}

	dialogFilterList := s.PrivateFacade.GetDialogFilters(ctx, md.UserId)

	dialogFilters := &mtproto.Vector_DialogFilter{
		Datas: make([]*mtproto.DialogFilter, 0, len(dialogFilterList)),
	}
	for _, df := range dialogFilterList {
		dialogFilters.Datas = append(dialogFilters.Datas, df.DialogFilter)
	}

	log.Debugf("messages.getDialogFilters - reply: %s", dialogFilters.DebugString())
	return dialogFilters, nil
}
