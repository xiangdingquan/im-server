package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetGlobalPrivacySettings(ctx context.Context, request *mtproto.TLAccountGetGlobalPrivacySettings) (*mtproto.GlobalPrivacySettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getGlobalPrivacySettings - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	gps := mtproto.MakeTLGlobalPrivacySettings(&mtproto.GlobalPrivacySettings{
		ArchiveAndMuteNewNoncontactPeers: mtproto.ToBool(true),
	}).To_GlobalPrivacySettings()
	log.Error(("account.getGlobalPrivacySettings - not imp AccountGetGlobalPrivacySettings"))
	return gps, nil
}
