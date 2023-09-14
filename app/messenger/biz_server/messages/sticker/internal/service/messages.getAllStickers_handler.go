package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetAllStickers(ctx context.Context, request *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getAllStickers#1c9618b1 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	stickers := &mtproto.TLMessagesAllStickers{Data2: &mtproto.Messages_AllStickers{
		Hash: 0, // hash规则
		Sets: s.StickerCore.GetStickerSetList(ctx, request.Hash),
	}}

	var acc uint32 = 0
	sets := stickers.GetSets()
	for _, set := range sets {
		if set.GetArchived() {
			continue
		}
		acc = acc*0x4F25 + uint32(set.GetHash())
	}
	hash := int32(acc & 0x7FFFFFFF)
	stickers.SetHash(hash)

	log.Debugf("messages.getAllStickers#1c9618b1 - reply: %s", stickers.DebugString())
	return stickers.To_Messages_AllStickers(), nil
}
