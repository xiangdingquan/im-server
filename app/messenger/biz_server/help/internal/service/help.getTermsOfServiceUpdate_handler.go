package service

import (
	"context"
	"time"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetTermsOfServiceUpdate(ctx context.Context, request *mtproto.TLHelpGetTermsOfServiceUpdate) (*mtproto.Help_TermsOfServiceUpdate, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.HelpGetTermsOfServiceUpdate - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.HelpGetTermsOfServiceUpdate - error: %v", err)
		return nil, err
	}

	termsOfServiceUpdate := mtproto.MakeTLHelpTermsOfServiceUpdateEmpty(&mtproto.Help_TermsOfServiceUpdate{
		Expires: int32(time.Now().Unix() + 3600),
	}).To_Help_TermsOfServiceUpdate()

	log.Debugf("help.HelpGetTermsOfServiceUpdate - reply: %s", termsOfServiceUpdate)
	return termsOfServiceUpdate, nil
}
