package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionSetClientSessionInfo(ctx context.Context, request *authsessionpb.TLSessionSetClientSessionInfo) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.setClientSessionInfo - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	clientSession := request.GetSession()
	if clientSession == nil {
		err := mtproto.ErrInputRequestInvalid
		log.Errorf("session.setClientSessionInfo - error: %v", err)
		return nil, err
	}

	keyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, request.GetSession().GetAuthKeyId())
	if err != nil {
		log.Errorf("session.setClientSessionInfo - error: %v", err)
		return nil, err
	}

	clientSession.AuthKeyId = keyData.PermAuthKeyId
	r := s.AuthSessionCore.SetClientSessionInfo(ctx, clientSession)

	log.Debugf("session.setClientSessionInfo - reply: {%v}", r)
	return mtproto.ToBool(r), nil
}
