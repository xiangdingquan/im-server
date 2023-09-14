package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesEditChatAdmin(ctx context.Context, request *mtproto.TLMessagesEditChatAdmin) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.editChatAdmin#a9e69f2e - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		isAdmin = mtproto.FromBool(request.GetIsAdmin())
		peer    = model.FromInputUser(md.UserId, request.UserId)
	)

	switch peer.PeerType {
	case model.PEER_USER:
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.editChatAdmin - invalid user_id, err: %v", err)
		return nil, err
	}

	chat, err := s.ChatFacade.EditChatAdmin(ctx, request.ChatId, md.UserId, peer.PeerId, isAdmin)
	if err != nil {
		log.Errorf("messages.editChatAdmin - error: ", err)
		return nil, err
	}

	log.Debugf("messages.editChatAdmin - reply: {true}")
	return model.WrapperGoFunc(mtproto.BoolTrue, func() {
		updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateChatParticipants(&mtproto.Update{
			Participants: chat.ToChatParticipants(0),
		}).To_Update())

		chat.Walk(func(userId int32, participant *model.ImmutableChatParticipant) error {
			if participant.IsChatMemberStateNormal() {
				sync_client.PushUpdates(ctx, participant.ChatParticipant.UserId, updatesHelper.ToPushUpdates(
					context.Background(),
					participant.ChatParticipant.UserId,
					s.UserFacade,
					s.ChatFacade,
					s.ChannelFacade))
			}
			return nil
		})
	}).(*mtproto.Bool), nil
}
