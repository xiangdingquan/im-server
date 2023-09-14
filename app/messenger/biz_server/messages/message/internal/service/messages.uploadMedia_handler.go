package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesUploadMedia(ctx context.Context, request *mtproto.TLMessagesUploadMedia) (*mtproto.MessageMedia, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.uploadMedia#519bc2b1 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	messageMedia, _ := s.makeMediaByInputMedia(ctx, md.UserId, md.AuthId, nil, request.GetMedia())

	log.Debugf("messages.uploadMedia#519bc2b1 - reply: %s", messageMedia.DebugString())
	return messageMedia, nil
}
