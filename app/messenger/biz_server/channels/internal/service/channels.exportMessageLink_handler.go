package service

import (
	"context"
	"fmt"

	"open.chat/app/pkg/env2"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsExportMessageLink(ctx context.Context, request *mtproto.TLChannelsExportMessageLink) (*mtproto.ExportedMessageLink, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.exportMessageLink - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		grouped bool
	)

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.exportMessageLink - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.exportMessageLink - error: %v", err)
		return nil, err
	}

	switch request.Constructor {
	case mtproto.CRC32_channels_exportMessageLink_e63fadeb:
		grouped = request.Grouped_FLAGBOOLEAN
	case mtproto.CRC32_channels_exportMessageLink_ceb77163:
		grouped = mtproto.FromBool(request.Grouped_BOOL)
		_ = grouped
	case mtproto.CRC32_channels_exportMessageLink_c846d22d:
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("channels.exportMessageLink - error: %v", err)
		return nil, err
	}
	channel, err := s.ChannelFacade.GetMutableChannel(ctx, request.Channel.ChannelId, md.UserId)
	if err != nil {
		log.Errorf("channels.exportMessageLink - error: %v", err)
		return nil, err
	}

	var exportedMessageLink *mtproto.ExportedMessageLink
	if channel.Channel.Username != "" {
		//
		exportedMessageLink = mtproto.MakeTLExportedMessageLink(&mtproto.ExportedMessageLink{
			Link: fmt.Sprintf("%s/c/%s/%d", env2.T_ME, channel.Channel.Username, request.Id),
			Html: "",
		}).To_ExportedMessageLink()
	} else {
		//
		exportedMessageLink = mtproto.MakeTLExportedMessageLink(&mtproto.ExportedMessageLink{
			Link: fmt.Sprintf("%s/c/%d/%d", env2.T_ME, channel.GetChannelId(), request.Id),
			Html: "",
		}).To_ExportedMessageLink()
	}

	log.Debugf("channels.exportMessageLink - reply: {%s}", exportedMessageLink.DebugString())
	return exportedMessageLink, nil
}
