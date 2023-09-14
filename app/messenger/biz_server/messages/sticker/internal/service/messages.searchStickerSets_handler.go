package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSearchStickerSets(ctx context.Context, request *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.searchStickerSets#c2b7d08b - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Warn("not Impl messages.searchStickerSets#c2b7d08b")

	foundStickerSets := &mtproto.TLMessagesFoundStickerSets{Data2: &mtproto.Messages_FoundStickerSets{
		Hash: 0,
		Sets: []*mtproto.StickerSetCovered{},
	}}
	return foundStickerSets.To_Messages_FoundStickerSets(), nil
}
