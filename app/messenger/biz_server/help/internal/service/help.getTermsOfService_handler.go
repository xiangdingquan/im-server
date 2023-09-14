package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetTermsOfService(ctx context.Context, request *mtproto.TLHelpGetTermsOfService) (*mtproto.Help_TermsOfService, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getTermsOfService - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	err := mtproto.ErrMethodNotImpl

	log.Warnf("help.getTermsOfService - not impl help.getTermsOfService#350170f3: %v", err)

	return nil, err
}
