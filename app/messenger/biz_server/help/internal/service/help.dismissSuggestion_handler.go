package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) HelpDismissSuggestion(ctx context.Context, request *mtproto.TLHelpDismissSuggestion) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("help.dismissSuggestion - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("help.dismissSuggestion - not imp HelpDismissSuggestion")
}
