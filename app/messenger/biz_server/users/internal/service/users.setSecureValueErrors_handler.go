package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) UsersSetSecureValueErrors(ctx context.Context, request *mtproto.TLUsersSetSecureValueErrors) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("users.setSecureValueErrors - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	result := mtproto.BoolFalse
	log.Debugf("users.setSecureValueErrors - reply: {false}")

	return result, nil
}
