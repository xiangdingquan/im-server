package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsDeleteUserHistory(ctx context.Context, request *mtproto.TLChannelsDeleteUserHistory) (*mtproto.Messages_AffectedHistory, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Infof("channels.deleteUserHistory - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.deleteUserHistory - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.deleteUserHistory - error: %v", err)
		return nil, err
	}

	peer := model.FromInputUser(md.UserId, request.UserId)
	switch peer.PeerType {
	case model.PEER_USER:
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("channels.deleteUserHistory - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, request.Channel.ChannelId, md.UserId, peer.PeerId)
	if err != nil {
		log.Errorf("channels.deleteUserHistory - error: %v", err)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(md.UserId)
	if me == nil || !me.CanAdminDeleteMessages() {
		err = mtproto.ErrChatAdminRequired
		log.Errorf("channels.deleteUserHistory - error: %v", err)
		return nil, err
	}

	affectedHistory, err := s.MsgFacade.DeleteChannelUserHistory(ctx,
		md.UserId,
		md.AuthId,
		request.Channel.ChannelId,
		peer)

	if err != nil {
		log.Errorf("channels.deleteUserHistory - error: %v", err)
		return nil, err
	}

	log.Debugf("channels.deleteUserHistory - reply: %s", affectedHistory.DebugString())
	return affectedHistory, nil
}
