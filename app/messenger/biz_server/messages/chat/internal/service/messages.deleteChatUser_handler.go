package service

import (
	"context"

	"math/rand"

	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesDeleteChatUser(ctx context.Context, request *mtproto.TLMessagesDeleteChatUser) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.deleteChatUser#e0611f16 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	peer := model.FromInputUser(md.UserId, request.UserId)

	switch peer.PeerType {
	case model.PEER_USER:
	case model.PEER_SELF:
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.deleteChatUser - invalid peer", err)
		return nil, err
	}

	chat, err := s.ChatFacade.DeleteChatUser(ctx, request.ChatId, md.UserId, peer.PeerId)
	if err != nil {
		log.Errorf("messages.deleteChatUser - error: %v", err)
		return nil, err
	}
	replyUpdates, err := s.MsgFacade.SendMessage(ctx,
		md.UserId,
		md.AuthId,
		model.MakeChatPeerUtil(request.ChatId),
		&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      chat.MakeMessageService(md.UserId, model.MakeMessageActionChatDeleteUser(peer.PeerId)),
			ScheduleDate: nil,
		})

	if err != nil {
		log.Errorf("messages.deleteChatUser - error: %v", err)
		return nil, err
	}

	updateChatParticipants := mtproto.MakeTLUpdateChatParticipants(&mtproto.Update{
		Participants: chat.ToChatParticipants(0),
	}).To_Update()
	if peer.PeerType == model.PEER_USER {
		replyUpdates.Updates = append(replyUpdates.Updates, updateChatParticipants)
	}

	log.Debugf("messages.deleteChatUser#e0611f16 - reply: {%s}", replyUpdates.DebugString())
	return model.WrapperGoFunc(replyUpdates, func() {
		updatesHelper := model.MakeUpdatesHelper(updateChatParticipants)

		chat.Walk(func(userId int32, participant *model.ImmutableChatParticipant) error {
			if userId != md.UserId && userId != peer.PeerId {
				sync_client.PushUpdates(ctx, userId,
					updatesHelper.ToPushUpdates(context.Background(), userId, s.UserFacade, s.ChatFacade, s.ChannelFacade))
			}
			return nil
		})
	}).(*mtproto.Updates), nil
}
