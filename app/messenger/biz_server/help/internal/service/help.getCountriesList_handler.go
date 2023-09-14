package service

import (
	"context"

	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// help.getCountriesList#735787a8 lang_code:string hash:int = help.CountriesList;
func (s *Service) HelpGetCountriesList(ctx context.Context, request *mtproto.TLHelpGetCountriesList) (*mtproto.Help_CountriesList, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.getCountriesList - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Sorry: not impl HelpGetCountriesList logic
	log.Warn("help.getCountriesList - error: method HelpGetCountriesList not impl")

	return nil, mtproto.ErrMethodNotImpl
}
