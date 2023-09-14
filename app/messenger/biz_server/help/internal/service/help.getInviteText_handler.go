package service

import (
	"context"
	"fmt"

	"open.chat/app/pkg/env2"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetInviteText(ctx context.Context, request *mtproto.TLHelpGetInviteText) (*mtproto.Help_InviteText, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getInviteText - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getAppConfig - error: %v", err)
		return nil, err
	}

	inviteText := mtproto.MakeTLHelpInviteText(&mtproto.Help_InviteText{
		Message: fmt.Sprintf("Hey, I'm using %s to chat. Join me. Download it here: https://%s/dl",
			env2.MY_APP_NAME,
			env2.MY_WEB_SITE),
	})

	log.Debugf("help.getInviteText - reply: %s", inviteText.DebugString())
	return inviteText.To_Help_InviteText(), nil
}
