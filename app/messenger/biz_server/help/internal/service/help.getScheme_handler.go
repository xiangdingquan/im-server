package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpGetScheme(ctx context.Context, request *mtproto.TLHelpGetScheme) (*mtproto.Scheme, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getScheme#dbb69a9e - metadata: %s, request: %s", request.DebugString(), md.DebugString())

	scheme := mtproto.MakeTLScheme(&mtproto.Scheme{
		SchemeRaw: "",
		Version:   1,
	}).To_Scheme()

	log.Debugf("help.getScheme#dbb69a9e - reply: %s", scheme.DebugString())
	return scheme, nil
}
