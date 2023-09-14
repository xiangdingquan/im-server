package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) UpdatesGetState(ctx context.Context, request *mtproto.TLUpdatesGetState) (*mtproto.Updates_State, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("updates.getState#edd4882a  - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	state, err := s.UpdatesFacade.GetState(ctx, md.AuthId, md.UserId)

	log.Debugf("updates.getState#edd4882a  - reply: %s", logger.JsonDebugData(state))
	return state, err
}
