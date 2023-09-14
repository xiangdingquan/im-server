package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsEditLocation(ctx context.Context, request *mtproto.TLChannelsEditLocation) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.editLocation - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.editCreator - error: %v", err)
		return nil, err
	}

	res, err := s.ChannelFacade.EditLocation(ctx, request.Channel.ChannelId, md.UserId, request.GetGeoPoint(), request.GetAddress())
	if err != nil {
		log.Errorf("ChannelsEditLocation error - %v", err)
		return nil, err
	}
	var reply *mtproto.Bool
	if res {
		reply = mtproto.BoolTrue
	} else {
		reply = mtproto.BoolFalse
	}

	log.Debugf("channels.editLocation - reply: {%s}", reply.DebugString())

	return reply, nil
}
