package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AccountUpdateDeviceLocked(ctx context.Context, request *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.updateDeviceLocked - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("account.updateDeviceLocked - error: %v", err)
		return nil, err
	}

	err := s.PushFacade.UpdateDeviceLockedPeriod(ctx, md.UserId, md.AuthId, request.Period)
	if err != nil {
		log.Errorf("account.updateDeviceLocked - error: %v", err)
	}

	reply := mtproto.ToBool(err == nil)

	log.Debugf("account.updateDeviceLocked - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}
