package service

import (
	"context"

	"open.chat/app/messenger/sync/client"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountResetAuthorization(ctx context.Context, request *mtproto.TLAccountResetAuthorization) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.resetAuthorization#df77f3bc - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if request.Hash == 0 {
		log.Debugf("account.resetAuthorization#df77f3bc - hash is 0")
		return mtproto.ToBool(false), nil
	}

	tKeyIdList, err := s.RPCSessionClient.SessionResetAuthorization(ctx, &authsessionpb.TLSessionResetAuthorization{
		UserId:    md.UserId,
		AuthKeyId: md.AuthId,
		Hash:      request.Hash,
	})

	if err != nil {
		log.Errorf("account.resetAuthorization#df77f3bc - error: %v", err)
		return nil, err
	}

	for _, id := range tKeyIdList.Datas {
		// notify kill session
		upds := mtproto.MakeTLUpdateAccountResetAuthorization(&mtproto.Updates{
			UserId:    md.UserId,
			AuthKeyId: id,
		}).To_Updates()
		sync_client.SyncUpdatesMe(ctx, md.UserId, id, 0, "", upds)
	}

	log.Debugf("account.resetAuthorization#df77f3bc - reply: {%d}", tKeyIdList)
	return mtproto.ToBool(true), nil
}
