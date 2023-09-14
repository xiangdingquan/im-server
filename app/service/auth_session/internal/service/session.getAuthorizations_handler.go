package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionGetAuthorizations(ctx context.Context, request *authsessionpb.TLSessionGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.getAuthorizations - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	myKeyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, request.GetExcludeAuthKeyId())
	if err != nil {
		log.Errorf("session.getAuthorizations - error: %v", err)
		return nil, err
	}

	authorizationList := s.AuthSessionCore.GetAuthorizations(ctx, request.GetUserId(), myKeyData.PermAuthKeyId)
	reply := &mtproto.TLAccountAuthorizations{Data2: &mtproto.Account_Authorizations{
		Authorizations: authorizationList,
	}}

	log.Debugf("session.getAuthorizations - reply: {%s}", reply.DebugString())
	return reply.To_Account_Authorizations(), nil
}
