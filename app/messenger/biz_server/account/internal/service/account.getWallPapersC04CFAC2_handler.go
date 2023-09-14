package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AccountGetWallPapersC04CFAC2(ctx context.Context, request *mtproto.TLAccountGetWallPapersC04CFAC2) (*mtproto.Vector_WallPaper, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getWallPapers#c04cfac2 - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	wallDataList, err := s.UserFacade.GetWallPaperList(ctx)
	if err != nil {
		return nil, err
	}

	walls := &mtproto.Vector_WallPaper{
		Datas: make([]*mtproto.WallPaper, 0, len(wallDataList)),
	}

	log.Debugf("account.getWallPapers#c04cfac2 - reply: %s", logger.JsonDebugData(walls))
	return walls, nil
}
