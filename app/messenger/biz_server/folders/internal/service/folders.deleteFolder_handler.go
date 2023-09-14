package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) FoldersDeleteFolder(ctx context.Context, request *mtproto.TLFoldersDeleteFolder) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("folders.deleteFolder - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("folders.deleteFolder - error: %v", err)
		return nil, err
	}

	return model.MakeEmptyUpdates(), nil
}
