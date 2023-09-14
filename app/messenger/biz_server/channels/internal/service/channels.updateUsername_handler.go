package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsUpdateUsername(ctx context.Context, request *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.updateUsername - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.updateUsername - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.updateUsername - error: %v", err)
		return nil, err
	}

	channelId := request.Channel.ChannelId

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, channelId, md.UserId)
	if err != nil {
		log.Errorf("channels.updateUsername - error: %v", err)
		return nil, err
	}
	me := channel.GetImmutableChannelParticipant(md.UserId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("channels.updateUsername - error: %v", err)
		return nil, err
	}

	if !me.CanAdminInviteUsers() {
		err = mtproto.ErrChatAdminRequired
		log.Errorf("channels.updateUsername - error: %v", err)
		return nil, err
	}

	if request.Username == "" {
		s.UsernameFacade.DeleteUsername(ctx, channel.Channel.Username)
	} else {
		if !model.CheckUsernameInvalid(request.Username) {
			log.Errorf("channels.updateUsername - format error: %v", request.Username)
			return nil, mtproto.ErrUsernameInvalid
		}
		s.UsernameFacade.UpdateUsername(ctx, model.PEER_CHANNEL, channelId, request.Username)
	}
	s.ChannelFacade.UpdateUsername(ctx, channelId, md.UserId, request.Username)

	log.Debugf("channels.updateUsername - reply: {true}")
	return mtproto.ToBool(true), nil
}
