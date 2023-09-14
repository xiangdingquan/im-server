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

func (s *Service) MessagesRequestEncryption(ctx context.Context, request *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.requestEncryption - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	DH_G_A_INVALID	g_a invalid
	// 400	USER_ID_INVALID	The provided user ID is invalid

	switch request.UserId.GetConstructor() {
	case mtproto.CRC32_inputUser:
	default:
		log.Errorf("invalid user_id: {%v}", request.UserId)
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	encryptChatData, err := s.SecretChatCore.CreateNewSecretChatData(
		ctx,
		md.UserId,
		request.UserId.UserId,
		md.AuthId,
		request.RandomId,
		request.GA)

	if err != nil {
		log.Errorf("requestEncryption error: {%v}", err)
		return nil, err
	}

	go func() {
		ctx2 := context.Background()

		pushUpdates := model.NewUpdatesLogic(encryptChatData.ParticipantId)

		chat, _ := encryptChatData.ToEncryptedChatRequested()
		updateRequestedEncryption := &mtproto.TLUpdateEncryption{Data2: &mtproto.Update{
			Date: int32(time.Now().Unix()),
			Chat: chat,
		}}

		pushUpdates.AddUpdate(updateRequestedEncryption.To_Update())

		users := s.UserFacade.GetUserListByIdList(ctx2, encryptChatData.ParticipantId, []int32{encryptChatData.AdminId, encryptChatData.ParticipantId})
		pushUpdates.AddUsers(users)
		sync_client.PushUpdates(ctx2, encryptChatData.ParticipantId, pushUpdates.ToUpdates())
	}()

	////////////////////////////////////////////////////////////////////////////////////////////////////////
	replyEncryptedChat := encryptChatData.ToEncryptedChatWaiting()

	log.Debugf("messages.requestEncryption - reply: %s", replyEncryptedChat.DebugString())
	return replyEncryptedChat, nil
}
