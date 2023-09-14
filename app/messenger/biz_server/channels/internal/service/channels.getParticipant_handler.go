package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetParticipant(ctx context.Context, request *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getParticipant - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err         error
		channelId   int32
		userId      = model.FromInputUser(md.UserId, request.UserId)
		participant *mtproto.ChannelParticipant
	)

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.getParticipant - error: %v", err)
		return nil, err
	} else {
		channelId = request.Channel.ChannelId
	}

	switch userId.PeerType {
	case model.PEER_SELF:
	case model.PEER_USER:
	default:
		err := mtproto.ErrUserIdInvalid
		log.Errorf("channels.getParticipant - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, channelId, userId.PeerId)
	if err != nil {
		log.Errorf("channels.getParticipant - error: %v", err)
		return nil, err
	}

	p := channel.GetImmutableChannelParticipant(userId.PeerId)
	if p == nil || p.IsLeft() {
		err := mtproto.ErrUserNotParticipant
		log.Errorf("channels.getParticipant - error: %v", err)
		return nil, err
	}

	participant, err = p.TryGetChannelParticipantSelf(md.UserId)
	if err != nil {
		log.Errorf("channels.getParticipant - error: %v", err)
		return nil, err
	}

	channelParticipants := mtproto.MakeTLChannelsChannelParticipant(&mtproto.Channels_ChannelParticipant{
		Participant: participant,
		Users: s.UserFacade.GetUserListByIdList(
			ctx,
			md.UserId,
			[]int32{participant.UserId, participant.InviterId_INT32}),
	}).To_Channels_ChannelParticipant()

	log.Debugf("channels.getParticipant - reply: {%s}", channelParticipants.DebugString())
	return channelParticipants, nil
}
