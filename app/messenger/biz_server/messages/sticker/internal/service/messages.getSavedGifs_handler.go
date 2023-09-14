package service

import (
	"context"

	media_client "open.chat/app/service/media/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetSavedGifs(ctx context.Context, request *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getSavedGifs - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	idList := s.StickerCore.GetSavedGifs(ctx, md.UserId)
	documents, _ := media_client.GetDocumentByIdList(idList)

	savedGifs := mtproto.MakeTLMessagesSavedGifs(&mtproto.Messages_SavedGifs{
		Hash: request.Hash,
		Gifs: documents,
	}).To_Messages_SavedGifs()

	log.Debugf("messages.getSavedGifs - reply: %s", savedGifs.DebugString())
	return savedGifs, nil
}
