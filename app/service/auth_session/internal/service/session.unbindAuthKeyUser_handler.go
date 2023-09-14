package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionUnbindAuthKeyUser(ctx context.Context, request *authsessionpb.TLSessionUnbindAuthKeyUser) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.unbindAuthKeyUser - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		unBindKeyId = request.AuthKeyId
	)

	if unBindKeyId != 0 {
		keyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, unBindKeyId)
		if err != nil {
			log.Errorf("session.unbindAuthKeyUser - error: %v", err)
			return nil, err
		} else if keyData == nil {
			err = mtproto.ErrAuthKeyInvalid
			return nil, err
		} else {
			unBindKeyId = keyData.PermAuthKeyId
		}
	}

	s.AuthSessionCore.UnbindAuthUser(ctx, unBindKeyId, request.UserId)

	log.Debugf("session.unbindAuthKeyUser - reply: {true}")
	return mtproto.BoolTrue, nil
}
