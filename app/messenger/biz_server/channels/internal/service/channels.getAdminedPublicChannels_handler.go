package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetAdminedPublicChannels(ctx context.Context, request *mtproto.TLChannelsGetAdminedPublicChannels) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getAdminedPublicChannels - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		byLocation bool
		checkLimit bool
	)

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.getAdminedPublicChannels - error: %v", err)
		return nil, err
	}

	switch request.Constructor {
	case mtproto.CRC32_channels_getAdminedPublicChannels_f8b036af:
	case mtproto.CRC32_channels_getAdminedPublicChannels_8d8d82d7:
		byLocation = request.ByLocation
		checkLimit = request.CheckLimit
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("channels.getAdminedPublicChannels - error: %v", err)
		return nil, err
	}

	chats, err := s.ChannelFacade.GetAdminedPublicChannels(ctx, md.UserId, byLocation, checkLimit)
	if err != nil {
		log.Errorf("channels.getAdminedPublicChannels - error: %v", err)
		return nil, err
	}

	messagesChats := mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Count: int32(len(chats)),
		Chats: chats.ToChats(md.UserId),
	}).To_Messages_Chats()

	log.Debugf("channels.getAdminedPublicChannels - reply: {%s}", messagesChats.DebugString())
	return messagesChats, nil
}
