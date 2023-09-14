package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetUserInfo(ctx context.Context, request *mtproto.TLHelpGetUserInfo) (*mtproto.Help_UserInfo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getUserInfo - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getUserInfo - error: %v", err)
		return nil, err
	}

	reply := mtproto.MakeTLHelpUserInfoEmpty(&mtproto.Help_UserInfo{}).To_Help_UserInfo()

	log.Debugf("help.getUserInfo - reply: %s", reply.DebugString())
	return reply, nil
}
