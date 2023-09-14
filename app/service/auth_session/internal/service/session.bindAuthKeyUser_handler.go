package service

import (
	"context"
	"fmt"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionBindAuthKeyUser(ctx context.Context, request *authsessionpb.TLSessionBindAuthKeyUser) (*mtproto.Int64, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.bindAuthKeyUser - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	keyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, request.GetAuthKeyId())
	if err != nil {
		log.Errorf("session.bindAuthKeyUser - error: %v", err)
		return nil, err
	} else if keyData == nil {
		return nil, fmt.Errorf("not found keyId")
	}

	hash := s.AuthSessionCore.BindAuthKeyUser(ctx, keyData.PermAuthKeyId, request.GetUserId())

	log.Debugf("session.bindAuthKeyUser - reply: {%d}", hash)
	return &mtproto.Int64{V: hash}, nil
}
