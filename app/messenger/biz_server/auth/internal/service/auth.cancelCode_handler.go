package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
	"open.chat/pkg/phonenumber"
)

func (s *Service) AuthCancelCode(ctx context.Context, request *mtproto.TLAuthCancelCode) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.cancelCode - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("auth.cancelCode - error: %v", err)
		return nil, err
	}

	phoneNumber, err := phonenumber.CheckAndGetPhoneNumber(request.GetPhoneNumber())
	if err != nil {
		log.Errorf("check phone_number error - %v", err)
		err = mtproto.ErrPhoneNumberInvalid
		return nil, err
	}
	_ = phoneNumber

	canceled := mtproto.ToBool(true)

	log.Debugf("auth.cancelCode -  - reply: %s", logger.JsonDebugData(canceled))
	return canceled, nil
}
