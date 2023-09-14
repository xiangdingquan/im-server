package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsReportSpam(ctx context.Context, request *mtproto.TLChannelsReportSpam) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Infof("channels.reportSpam - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.reportSpam - error: %v", err)
		return nil, err
	}

	if len(request.Id) == 0 {
		log.Errorf("channels.reportSpam error - id empty")
		return mtproto.BoolFalse, nil
	}

	if request.UserId == nil {
		log.Errorf("channels.reportSpam error - user_id is nil")
		return mtproto.BoolFalse, nil
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.readHistory - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, request.Channel.ChannelId, md.UserId)
	if err != nil {
		log.Errorf("channels.reportSpam error - %v")
		return nil, err
	}

	s.ReportFacade.ReportIdList(ctx,
		md.UserId,
		model.CHANNELS_reportSpam,
		model.PEER_CHANNEL,
		channel.GetChannelId(),
		request.UserId.UserId,
		request.Id,
		int32(model.REASON_SPAM),
		"")

	log.Debugf("channels.reportSpam - reply: {true}")
	return mtproto.ToBool(true), nil
}
