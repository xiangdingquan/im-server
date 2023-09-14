package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetAdminLog(ctx context.Context, request *mtproto.TLChannelsGetAdminLog) (*mtproto.Channels_AdminLogResults, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getAdminLog - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.getAdminLog - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.getAdminLog - error: %v", err)
		return nil, err
	}

	var (
		limit = request.Limit
	)

	if limit > 50 {
		limit = 50
	}

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, request.Channel.ChannelId, md.UserId)
	if err != nil {
		log.Errorf("channels.exportMessageLink - error: %v", err)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(md.UserId)
	if me == nil || !me.IsCreatorOrAdmin() {
		err := mtproto.ErrChatAdminRequired
		log.Errorf("channels.getAdminLog - error: %v", err)
		return nil, err
	}

	admins := make([]int32, 0, len(request.GetAdmins()))
	for _, admin := range request.GetAdmins() {
		admins = append(admins, model.FromInputUser(md.UserId, admin).PeerId)
	}
	adminLogEvents := s.ChannelFacade.GetAdminLogs(ctx,
		request.Channel.ChannelId,
		request.Q,
		int32(model.FromChannelAdminLogEventsFilter(request.GetEventsFilter())), // request.EventsFilter,
		admins, // request.Admins,
		request.MaxId,
		request.MinId,
		limit)

	adminLogResults := mtproto.MakeTLChannelsAdminLogResults(&mtproto.Channels_AdminLogResults{
		Events: adminLogEvents,
		Users:  []*mtproto.User{},
		Chats:  []*mtproto.Chat{},
	}).To_Channels_AdminLogResults()

	log.Debugf("channels.getAdminLog - reply: {%s}", adminLogResults.DebugString())
	return adminLogResults, nil
}
