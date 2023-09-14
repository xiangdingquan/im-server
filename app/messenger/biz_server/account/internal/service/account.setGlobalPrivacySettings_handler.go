package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountSetGlobalPrivacySettings(ctx context.Context, request *mtproto.TLAccountSetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.setGlobalPrivacySettings - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	gps := mtproto.MakeTLGlobalPrivacySettings(request.Settings).To_GlobalPrivacySettings()
	log.Error(("account.setGlobalPrivacySettings - not imp AccountSetGlobalPrivacySettings"))
	return gps, nil
}
