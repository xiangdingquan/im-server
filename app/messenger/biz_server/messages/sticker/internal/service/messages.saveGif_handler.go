package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSaveGif(ctx context.Context, request *mtproto.TLMessagesSaveGif) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.saveGif - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	unsave := mtproto.FromBool(request.Unsave)
	if !unsave {
		s.StickerCore.SaveGif(ctx, md.UserId, request.GetId().GetId())
	} else {
		s.StickerCore.DeleteSavedGif(ctx, md.UserId, request.GetId().GetId())
	}

	log.Debugf("messages.saveGif - reply: {true}")
	return mtproto.ToBool(true), nil
}
