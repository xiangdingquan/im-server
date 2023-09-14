package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetCdnConfig(ctx context.Context, request *mtproto.TLHelpGetCdnConfig) (*mtproto.CdnConfig, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getCdnConfig#52029342 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.MakeTLCdnConfig(&mtproto.CdnConfig{
		PublicKeys: []*mtproto.CdnPublicKey{},
	}).To_CdnConfig(), nil
}
