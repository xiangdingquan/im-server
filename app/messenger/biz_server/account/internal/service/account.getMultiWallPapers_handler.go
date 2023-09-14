package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetMultiWallPapers(ctx context.Context, request *mtproto.TLAccountGetMultiWallPapers) (*mtproto.Vector_WallPaper, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getMultiWallPapers - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return &mtproto.Vector_WallPaper{
		Datas: []*mtproto.WallPaper{},
	}, nil
}
