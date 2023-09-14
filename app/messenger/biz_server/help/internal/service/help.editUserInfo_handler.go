package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpEditUserInfo(ctx context.Context, request *mtproto.TLHelpEditUserInfo) (*mtproto.Help_UserInfo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.editUserInfo - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.editUserInfo - error: %v", err)
		return nil, err
	}

	reply := mtproto.MakeTLHelpUserInfoEmpty(&mtproto.Help_UserInfo{}).To_Help_UserInfo()

	log.Debugf("help.editUserInfo - reply: %s", reply.DebugString())
	return reply, nil
}
