package service

import (
	"context"
	"math/rand"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesAddChatUser(ctx context.Context, request *mtproto.TLMessagesAddChatUser) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.addChatUser#f9a0aa09 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err  error
		peer = model.FromInputUser(md.UserId, request.UserId)
		chat *model.MutableChat
	)

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.addChatUser - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_USER:
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("messages.addChatUser - error: %v", err)
		return nil, err
	}

	// 400	USERS_TOO_MUCH	The maximum number of users has been exceeded (to create a chat, for example)
	// 400	USER_ALREADY_PARTICIPANT	The user is already in the group
	// 400	USER_ID_INVALID	The provided user ID is invalid
	// 403	USER_NOT_MUTUAL_CONTACT	The provided user is not a mutual contact
	// 403	USER_PRIVACY_RESTRICTED	The user's privacy settings do not allow you to do this
	isContact, _ := s.UserFacade.GetContactAndMutual(ctx, peer.PeerId, md.UserId)
	allowAddChat := s.UserFacade.CheckPrivacy(ctx, int(model.CHAT_INVITE), peer.PeerId, md.UserId, isContact)
	if !allowAddChat {
		err = mtproto.ErrUserPrivacyRestricted
		log.Errorf("not allow addChat: %v", err)
		return nil, err
	}

	chat, err = s.ChatFacade.AddChatUser(ctx, request.ChatId, md.UserId, peer.PeerId)
	if err != nil {
		log.Errorf("addChatUser error: %v", err)
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
			Message:      chat.MakeMessageService(md.UserId, model.MakeMessageActionChatAddUser(peer.PeerId)),
			ScheduleDate: nil,
		})

	if err != nil {
		log.Errorf("addChatUser error: %v", err)
		return nil, err
	}

	log.Debugf("messages.addChatUser#f9a0aa09 - reply: {%v}", replyUpdates.DebugString())
	return replyUpdates, nil
}
