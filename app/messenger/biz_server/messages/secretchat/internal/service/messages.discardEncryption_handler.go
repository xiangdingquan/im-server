package service

import (
	"context"
	"time"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesDiscardEncryption(ctx context.Context, request *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.discardEncryption - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	encryptChatData, err := s.SecretChatCore.MakeSecretChatData(ctx, request.ChatId, 0)
	if err != nil {
		log.Errorf("discardEncryption error: {%v}", err)
		return nil, err
	}

	var participantId int32
	if md.UserId == encryptChatData.AdminId {
		participantId = encryptChatData.ParticipantId
	} else {
		participantId = encryptChatData.AdminId
	}

	err = s.SecretChatCore.AddClosedRequest(ctx, encryptChatData.Id, md.UserId, participantId)
	if err != nil {
		log.Errorf("discardEncryption error: {%v}", err)
		return nil, err
	}

	go func() {
		ctx2 := context.Background()

		pushUpdates := model.NewUpdatesLogic(encryptChatData.AdminId)
		updateRequestedEncryption := &mtproto.TLUpdateEncryption{Data2: &mtproto.Update{
			Date: int32(time.Now().Unix()),
			Chat: &mtproto.EncryptedChat{
				Constructor: mtproto.CRC32_encryptedChatDiscarded,
				Id:          encryptChatData.Id,
			},
		}}
		pushUpdates.AddUpdate(updateRequestedEncryption.To_Update())

		var pushUserId int32
		if md.UserId == encryptChatData.AdminId {
			pushUserId = encryptChatData.ParticipantId
		} else {
			pushUserId = encryptChatData.AdminId
		}

		users := s.UserFacade.GetUserListByIdList(ctx2, pushUserId, []int32{encryptChatData.AdminId, encryptChatData.ParticipantId})
		pushUpdates.AddUsers(users)
		sync_client.PushUpdates(ctx2, pushUserId, pushUpdates.ToUpdates())
	}()

	log.Debugf("messages.discardEncryption - reply: {true}")
	return mtproto.ToBool(true), nil
}
