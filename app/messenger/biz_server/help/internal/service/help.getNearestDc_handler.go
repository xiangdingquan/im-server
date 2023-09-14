package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetNearestDc(ctx context.Context, request *mtproto.TLHelpGetNearestDc) (*mtproto.NearestDc, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getNearestDc - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("help.getNearestDc - error: %v", err)
		return nil, err
	}

	dc := mtproto.MakeTLNearestDc(&mtproto.NearestDc{
		Country:   "US",
		ThisDc:    2,
		NearestDc: 2,
	}).To_NearestDc()

	log.Debugf("help.getNearestDc#1fb33026 - reply: %s", dc.DebugString())
	return dc, nil
}
