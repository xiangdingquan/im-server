package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) SessionResetAuthorization(ctx context.Context, request *authsessionpb.TLSessionResetAuthorization) (*authsessionpb.VectorLong, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.resetAuthorization - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		excludeKeyId = request.AuthKeyId
	)

	if excludeKeyId != 0 {
		myKeyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, request.AuthKeyId)
		if err != nil {
			log.Errorf("session.getAuthorizations - error: %v", err)
			return nil, err
		} else if myKeyData == nil {
			log.Errorf("session.getAuthorizations - error: %v", err)
			err = mtproto.ErrAuthKeyInvalid
			return nil, err
		} else {
			excludeKeyId = myKeyData.PermAuthKeyId
		}
	}

	keyIdList := s.AuthSessionCore.ResetAuthorization(ctx, request.UserId, excludeKeyId, request.Hash)
	log.Debugf("keyIdList: %v", keyIdList)
	reply := &authsessionpb.VectorLong{
		Datas: make([]int64, 0, len(keyIdList)),
	}
	for _, keyId := range keyIdList {
		keyData, _ := s.Dao.GetAuthKey(ctx, keyId)
		if keyData != nil {
			if keyData.TempAuthKeyId != 0 {
				reply.Datas = append(reply.Datas, keyData.TempAuthKeyId)
			} else {
				reply.Datas = append(reply.Datas, keyId)
			}
		}
	}

	log.Debugf("session.resetAuthorization - reply: {%s}", logger.JsonDebugData(reply))
	return reply, nil
}
