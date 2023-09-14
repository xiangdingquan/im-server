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

func (s *Service) MessagesSendEncryptedService(ctx context.Context, request *mtproto.TLMessagesSendEncryptedService) (*mtproto.Messages_SentEncryptedMessage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.sendEncryptedService - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if request.Peer == nil {
		log.Errorf("peer is nil")
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	encryptChatData, err := s.SecretChatCore.MakeSecretChatData(ctx, request.Peer.ChatId, request.Peer.AccessHash)
	if err != nil {
		log.Errorf("sendEncrypted error: {%v}", err)
		return nil, err
	}

	date := int32(time.Now().Unix())
	encryptedMessage := &mtproto.EncryptedMessage{
		Constructor: mtproto.CRC32_encryptedMessageService,
		RandomId:    request.GetRandomId(),
		ChatId:      encryptChatData.Id,
		Date:        date,
		Bytes:       request.GetData(),
	}

	qts, err := encryptChatData.SendEncryptedMessage(ctx, md.UserId, md.AuthId, encryptedMessage)
	if err != nil {
		log.Errorf("sendEncrypted error: {%v}", err)
		return nil, err
	}

	go func() {
		ctx2 := context.Background()

		peerId, _ := encryptChatData.GetSecretChatPeerId(md.UserId)

		pushUpdates := model.NewUpdatesLogic(peerId)
		updateNewEncryptedMessage := &mtproto.TLUpdateNewEncryptedMessage{Data2: &mtproto.Update{
			Message_ENCRYPTEDMESSAGE: encryptedMessage,
			Qts:                      qts,
		}}
		pushUpdates.AddUpdate(updateNewEncryptedMessage.To_Update())

		users := s.UserFacade.GetUserListByIdList(ctx2, peerId, []int32{encryptChatData.AdminId, encryptChatData.ParticipantId})
		pushUpdates.AddUsers(users)
		sync_client.PushUpdates(ctx2, peerId, pushUpdates.ToUpdates())

	}()

	sentEncryptedMessage := &mtproto.TLMessagesSentEncryptedMessage{Data2: &mtproto.Messages_SentEncryptedMessage{
		Date: date,
	}}

	log.Debugf("messages.sendEncryptedService - reply: %s", sentEncryptedMessage.DebugString())
	return sentEncryptedMessage.To_Messages_SentEncryptedMessage(), nil
}
