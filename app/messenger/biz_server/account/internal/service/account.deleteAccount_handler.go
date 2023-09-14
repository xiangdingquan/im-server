package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AccountDeleteAccount(ctx context.Context, request *mtproto.TLAccountDeleteAccount) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("AccountDeleteAccount - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	deletedOk, _ := s.UserFacade.DeleteUser(ctx, md.UserId, request.GetReason())

	log.Debugf("AccountDeleteAccount - reply: {%v}", deletedOk)
	return mtproto.ToBool(deletedOk), nil
}
