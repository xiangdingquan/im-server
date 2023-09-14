package service

import (
	"context"
	"time"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetPromoData(ctx context.Context, request *mtproto.TLHelpGetPromoData) (*mtproto.Help_PromoData, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getPromoData - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getProxyData - error: %v", err)
		return nil, err
	}

	reply := mtproto.MakeTLHelpPromoDataEmpty(&mtproto.Help_PromoData{
		Expires: int32(time.Now().Unix() + 60*60),
	}).To_Help_PromoData()

	log.Debugf("help.getPromoData - reply: %s", reply.DebugString())
	return reply, nil
}
