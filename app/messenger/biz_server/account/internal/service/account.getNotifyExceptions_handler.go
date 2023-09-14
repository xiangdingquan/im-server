package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetNotifyExceptions(ctx context.Context, request *mtproto.TLAccountGetNotifyExceptions) (reply *mtproto.Updates, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getNotifyExceptions - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("account.getNotifyExceptions - error: %v", err)
		return nil, err
	}

	// handle compare_sound

	settings, err := s.UserFacade.GetAllNotifySettings(ctx, md.UserId)
	if err != nil {
		log.Errorf("account.getNotifyExceptions - error: %v", err)
		return nil, err
	}

	var (
		userIdList    []int32
		chatIdList    []int32
		channelIdList []int32
		resultUpdates = model.NewUpdatesLogic(md.UserId)
	)
	for k, v := range settings {
		peer := &model.PeerUtil{
			PeerType: int32(k >> 32),
			PeerId:   int32(k & 0xffffffff),
		}

		switch peer.PeerType {
		case model.PEER_USER:
			userIdList = append(userIdList, peer.PeerId)
		case model.PEER_CHAT:
			chatIdList = append(chatIdList, peer.PeerId)
		case model.PEER_CHANNEL:
			channelIdList = append(channelIdList, peer.PeerId)
		default:
			continue
		}

		updateNotifySettings := mtproto.MakeTLUpdateNotifySettings(&mtproto.Update{
			Peer_NOTIFYPEER: peer.ToNotifyPeer(),
			NotifySettings:  v,
		}).To_Update()
		resultUpdates.AddUpdate(updateNotifySettings)
	}

	userList := s.UserFacade.GetUserListByIdList(ctx, md.UserId, userIdList)
	chatList := s.ChatFacade.GetChatListByIdList(ctx, md.UserId, chatIdList)
	channelList := s.ChannelFacade.GetChannelListByIdList(ctx, md.UserId, channelIdList...)

	resultUpdates.AddUsers(userList)
	resultUpdates.AddChats(chatList)
	resultUpdates.AddChats(channelList)

	reply = resultUpdates.ToUpdates()

	log.Debugf("account.getNotifyExceptions - reply %s", reply.DebugString())
	return reply, nil
}
