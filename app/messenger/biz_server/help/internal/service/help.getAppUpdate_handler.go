package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetAppUpdate(ctx context.Context, request *mtproto.TLHelpGetAppUpdate) (*mtproto.Help_AppUpdate, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getAppUpdate - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getAppUpdate - error: %v", err)
		return nil, err
	}

	reply := mtproto.MakeTLHelpNoAppUpdate(nil).To_Help_AppUpdate()

	log.Debugf("help.getAppUpdate - reply: %s\n", reply.DebugString())
	return reply, nil
}
