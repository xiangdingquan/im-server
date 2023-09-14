package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AuthResetAuthorizations(ctx context.Context, request *mtproto.TLAuthResetAuthorizations) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.resetAuthorizations#9fab0d1a - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	tKeyIdList, _ := s.AuthSessionRpcClient.SessionResetAuthorization(ctx, &authsessionpb.TLSessionResetAuthorization{
		UserId:    md.UserId,
		AuthKeyId: md.AuthId,
		Hash:      0,
	})

	for _, id := range tKeyIdList.Datas {
		upds := mtproto.MakeTLUpdateAccountResetAuthorization(&mtproto.Updates{
			UserId:    md.UserId,
			AuthKeyId: id,
		}).To_Updates()
		sync_client.SyncUpdatesMe(ctx, md.UserId, id, 0, "", upds)
	}

	log.Debugf("auth.resetAuthorizations#9fab0d1a - reply: {%v}", tKeyIdList)
	return mtproto.ToBool(true), nil
}
