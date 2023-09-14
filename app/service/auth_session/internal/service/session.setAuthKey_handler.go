package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionSetAuthKey(ctx context.Context, request *authsessionpb.TLSessionSetAuthKey) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.setAuthKey - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		keyInfo = request.GetAuthKey()
		salt    *mtproto.TLFutureSalt
	)

	if request.FutureSalt != nil {
		salt = request.FutureSalt.To_FutureSalt()
	}
	err := s.AuthSessionCore.InsertAuthKey(
		ctx,
		&model.AuthKeyData{
			AuthKeyId:          keyInfo.AuthKeyId,
			AuthKeyType:        int(keyInfo.AuthKeyType),
			AuthKey:            keyInfo.AuthKey,
			PermAuthKeyId:      keyInfo.PermAuthKeyId,
			TempAuthKeyId:      keyInfo.TempAuthKeyId,
			MediaTempAuthKeyId: keyInfo.MediaTempAuthKeyId,
		},
		salt)
	if err != nil {
		log.Errorf("session.setAuthKey - error: %v", err)
		return mtproto.ToBool(false), nil
	}

	log.Debugf("session.setAuthKey - reply: {true}")
	return mtproto.ToBool(true), nil
}
