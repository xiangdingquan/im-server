package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetSupportName(ctx context.Context, request *mtproto.TLHelpGetSupportName) (*mtproto.Help_SupportName, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getSupportName - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getSupportName - error: %v", err)
		return nil, err
	}

	// Get localized name of the telegram support user
	reply := mtproto.MakeTLHelpSupportName(&mtproto.Help_SupportName{
		Name: supportName,
	}).To_Help_SupportName()

	log.Debugf("help.getSupportName - reply: %s", reply.DebugString())
	return reply, nil
}
