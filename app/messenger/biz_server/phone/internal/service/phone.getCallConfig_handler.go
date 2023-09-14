package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) PhoneGetCallConfig(ctx context.Context, request *mtproto.TLPhoneGetCallConfig) (*mtproto.DataJSON, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("phone.getCallConfig - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("phone.getCallConfig - error: %v", err)
		return nil, err
	}

	reply := mtproto.MakeTLDataJSON(&mtproto.DataJSON{
		Data: model.GetCallConfigDataJson(),
	}).To_DataJSON()

	log.Debugf("phone.getCallConfig - reply %s", reply.DebugString())
	return reply, nil
}
