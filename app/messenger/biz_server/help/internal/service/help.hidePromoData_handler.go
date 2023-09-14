package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpHidePromoData(ctx context.Context, request *mtproto.TLHelpHidePromoData) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.hidePromoData - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.BoolFalse, nil
}
