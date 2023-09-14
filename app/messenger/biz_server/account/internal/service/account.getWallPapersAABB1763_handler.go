package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AccountGetWallPapersAABB1763(ctx context.Context, request *mtproto.TLAccountGetWallPapersAABB1763) (*mtproto.Account_WallPapers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getWallPapers#c04cfac2 - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	wallDataList, err := s.UserFacade.GetWallPaperList(ctx)
	if err != nil {
		return nil, err
	}

	walls := mtproto.MakeTLAccountWallPapers(&mtproto.Account_WallPapers{
		Hash:       request.Hash,
		Wallpapers: make([]*mtproto.WallPaper, 0, len(wallDataList)),
	}).To_Account_WallPapers()

	log.Debugf("account.getWallPapers#c04cfac2 - reply: %s", logger.JsonDebugData(walls))
	return walls, nil
}
