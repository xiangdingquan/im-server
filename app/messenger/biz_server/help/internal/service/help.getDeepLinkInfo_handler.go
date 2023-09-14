package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetDeepLinkInfo(ctx context.Context, request *mtproto.TLHelpGetDeepLinkInfo) (*mtproto.Help_DeepLinkInfo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getDeepLinkInfo - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getDeepLinkInfo - error: %v", err)
		return nil, err
	}

	reply := mtproto.MakeTLHelpDeepLinkInfoEmpty(nil).To_Help_DeepLinkInfo()

	log.Debugf("help.getDeepLinkInfo - reply: %s", reply)
	return reply, nil
}
