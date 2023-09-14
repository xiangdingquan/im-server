package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSearchGifs(ctx context.Context, request *mtproto.TLMessagesSearchGifs) (*mtproto.Messages_FoundGifs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.searchGifs#bf9a776b - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("not Impl messages.searchGifs#bf9a776b")

	foundGif := mtproto.MakeTLMessagesFoundGifs(&mtproto.Messages_FoundGifs{
		NextOffset: 0,
		Results:    []*mtproto.FoundGif{},
	}).To_Messages_FoundGifs()

	return foundGif, nil
}
