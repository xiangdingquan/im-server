package service

import (
	"context"

	"math/rand"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsCreateChannel(ctx context.Context, request *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.createChannel#f4893d7f - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.createChannel - error: %v", err)
		return nil, err
	}

	if request.Broadcast && request.Megagroup {
		log.Errorf("broadcast and megagroup == true")
		return nil, mtproto.ErrInputRequestInvalid
	} else if !request.Broadcast && !request.Megagroup {
		log.Errorf("broadcast and megagroup == false")
		return nil, mtproto.ErrInputRequestInvalid
	}

	if request.Title == "" {
		log.Errorf("title empty")
		return nil, mtproto.ErrChatTitleEmpty
	}

	key := crypto.CreateAuthKey()
	_, err := s.RPCSessionClient.SessionSetAuthKey(ctx, &authsessionpb.TLSessionSetAuthKey{
		AuthKey: &authsessionpb.AuthKeyInfo{
			AuthKeyId:          key.AuthKeyId(),
			AuthKey:            key.AuthKey(),
			AuthKeyType:        model.AuthKeyTypePerm,
			PermAuthKeyId:      key.AuthKeyId(),
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		},
		FutureSalt: nil,
	})
	if err != nil {
		log.Errorf("create channel secret key error")
		return nil, err
	}

	channel, err := s.ChannelFacade.CreateChannel(
		ctx,
		md.UserId,
		key.AuthKeyId(),
		request.Broadcast,
		request.Title,
		request.About,
		request.Notice,
		request.GetGeoPoint(),
		request.GetAddress(),
		md.ClientMsgId)

	if err != nil {
		log.Errorf("createChannel error - %v", err)
		return nil, err
	}

	resultUpdates, err := s.MsgFacade.SendMessage(
		ctx,
		md.UserId,
		md.AuthId,
		model.MakeChannelPeerUtil(channel.GetChannelId()),
		&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      channel.MakeMessageService(md.UserId, false, 0, model.MakeMessageActionChannelCreate(request.Title)),
			ScheduleDate: nil,
		},
	)

	if err != nil {
		log.Errorf("channels.createChannel#f4893d7f - error: %v", err)
		return nil, err
	}

	log.Debugf("channels.createChannel#f4893d7f - reply: {%s}", resultUpdates.DebugString())
	return resultUpdates, nil
}
