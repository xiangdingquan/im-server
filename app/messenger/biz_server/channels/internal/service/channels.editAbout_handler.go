package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsEditAbout(ctx context.Context, request *mtproto.TLChannelsEditAbout) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.editAbout - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.editAbout - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.editAbout - error: %v", err)
		return nil, err
	}

	log.Debugf("editAbout: {channel_id: %d, editId: %d, about: %s}", request.Channel.ChannelId, md.UserId, request.About)
	_, err := s.ChannelFacade.EditAbout(ctx, request.Channel.ChannelId, md.UserId, request.About)
	if err != nil {
		log.Errorf("channels.editAbout - error: %v", err)
		return nil, err
	}

	log.Debugf("channels.editAbout - reply: {true}")
	return mtproto.ToBool(true), nil
}
