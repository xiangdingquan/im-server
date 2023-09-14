package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpTest(ctx context.Context, request *mtproto.TLHelpTest) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.test#c0e202f7 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	reply := mtproto.ToBool(false)

	log.Debugf("help.test#c0e202f7 - reply: {false}")
	return reply, nil
}
