package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountSaveAutoDownloadSettings(ctx context.Context, request *mtproto.TLAccountSaveAutoDownloadSettings) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getAutoDownloadSettings - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	reply := mtproto.ToBool(true)
	log.Debugf("account.getAutoDownloadSettings#56da0b3f - reply: %s\n", reply)

	return reply, nil
}
