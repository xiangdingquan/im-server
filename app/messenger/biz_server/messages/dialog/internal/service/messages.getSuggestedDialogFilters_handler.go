package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetSuggestedDialogFilters(ctx context.Context, request *mtproto.TLMessagesGetSuggestedDialogFilters) (*mtproto.Vector_DialogFilterSuggested, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getSuggestedDialogFilters - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getSuggestedDialogFilters - error: %v", err)
		return nil, err
	}

	dialogFilters := &mtproto.Vector_DialogFilterSuggested{
		Datas: []*mtproto.DialogFilterSuggested{},
	}

	log.Debugf("messages.getSuggestedDialogFilters - reply: %s", dialogFilters.DebugString())
	return dialogFilters, nil
}
