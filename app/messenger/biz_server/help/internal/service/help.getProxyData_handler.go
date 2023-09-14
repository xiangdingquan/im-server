package service

import (
	"context"
	"time"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetProxyData(ctx context.Context, request *mtproto.TLHelpGetProxyData) (*mtproto.Help_ProxyData, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getProxyData - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getProxyData - error: %v", err)
		return nil, err
	}

	reply := mtproto.MakeTLHelpProxyDataEmpty(&mtproto.Help_ProxyData{
		Expires: int32(time.Now().Unix() + 60*60),
	}).To_Help_ProxyData()

	log.Debugf("help.getProxyData - reply: %s", reply.DebugString())
	return reply, nil
}
