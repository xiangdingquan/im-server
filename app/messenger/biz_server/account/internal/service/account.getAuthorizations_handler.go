package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AccountGetAuthorizations(ctx context.Context, request *mtproto.TLAccountGetAuthorizations) (*mtproto.Account_Authorizations, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getAuthorizations - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.getAuthorizations - error: %v", err)
		return nil, err
	}

	authorizations, err := s.RPCSessionClient.SessionGetAuthorizations(ctx, &authsessionpb.TLSessionGetAuthorizations{
		UserId:           md.UserId,
		ExcludeAuthKeyId: md.AuthId,
	})
	if err == nil {
		log.Debugf("account.getAuthorizations - reply: {%s}", logger.JsonDebugData(authorizations))
	} else {
		log.Errorf("account.getAuthorizations - error: %v", err)
	}

	log.Debugf("account.getAuthorizations - reply: %v", authorizations)
	return authorizations, err
}
