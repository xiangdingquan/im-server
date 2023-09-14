package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) PhotosUpdateProfilePhoto72D4742C(ctx context.Context, request *mtproto.TLPhotosUpdateProfilePhoto72D4742C) (*mtproto.Photos_Photo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("photos.updateProfilePhoto#72d4742c - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("photos.updateProfilePhoto - error: %v", err)
		return nil, err
	}

	log.Warnf("not impl photos.updateProfilePhoto#72d4742c")

	return nil, mtproto.ErrMethodNotImpl
}
