package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthBindTempAuthKey(ctx context.Context, request *mtproto.TLAuthBindTempAuthKey) (reply *mtproto.Bool, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.bindTempAuthKey - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	err = mtproto.ErrMethodNotImpl
	log.Warn("auth.bindTempAuthKey - method not impl")

	return mtproto.BoolTrue, nil
}
