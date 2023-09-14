package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsExportInvite(ctx context.Context, request *mtproto.TLChannelsExportInvite) (*mtproto.ExportedChatInvite, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.exportInvite - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.exportInvite - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.exportInvite - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.ExportChannelInvite(ctx, request.Channel.ChannelId, md.UserId)
	if err != nil {
		log.Errorf("channels.exportInvite - error: %v", err)
		return nil, err
	}

	exportedChatInvite := mtproto.MakeTLChatInviteExported(&mtproto.ExportedChatInvite{
		Link: channel.Channel.Link,
	}).To_ExportedChatInvite()

	log.Debugf("channels.exportInvite - reply: {%s}", exportedChatInvite.DebugString())
	return exportedChatInvite, nil
}
