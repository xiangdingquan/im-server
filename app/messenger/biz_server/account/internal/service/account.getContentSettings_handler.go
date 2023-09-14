package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetContentSettings(ctx context.Context, request *mtproto.TLAccountGetContentSettings) (*mtproto.Account_ContentSettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getContentSettings - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	//  add
	contentSettings := mtproto.MakeTLAccountContentSettings(&mtproto.Account_ContentSettings{
		SensitiveEnabled:   false,
		SensitiveCanChange: false,
	}).To_Account_ContentSettings()

	log.Debugf("account.getContentSettings - reply: %s", contentSettings.DebugString())
	return contentSettings, nil
}
