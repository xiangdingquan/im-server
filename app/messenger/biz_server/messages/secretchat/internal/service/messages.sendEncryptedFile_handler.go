package service

import (
	"context"
	"fmt"
	"time"

	sync_client "open.chat/app/messenger/sync/client"
	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func makeEncryptedFile(authKeyId int64, file *mtproto.InputEncryptedFile) (*mtproto.EncryptedFile, error) {
	switch file.GetConstructor() {
	case mtproto.CRC32_inputEncryptedFileEmpty:
		return mtproto.MakeTLEncryptedFileEmpty(nil).To_EncryptedFile(), nil
	case mtproto.CRC32_inputEncryptedFileUploaded:
		return media_client.UploadEncryptedFile(authKeyId, file)
	case mtproto.CRC32_inputEncryptedFile:
		return media_client.GetEncryptedFile(file.Id, file.AccessHash)
	case mtproto.CRC32_inputEncryptedFileBigUploaded:
		return media_client.UploadEncryptedFile(authKeyId, file)
	default:
	}
	return nil, fmt.Errorf("invalid constructor")
}

func (s *Service) MessagesSendEncryptedFile(ctx context.Context, request *mtproto.TLMessagesSendEncryptedFile) (*mtproto.Messages_SentEncryptedMessage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.sendEncryptedFile - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if request.Peer == nil {
		log.Errorf("peer is nil")
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	encryptChatData, err := s.SecretChatCore.MakeSecretChatData(ctx, request.Peer.ChatId, request.Peer.AccessHash)
	if err != nil {
		log.Errorf("sendEncrypted error: {%v}", err)
		return nil, err
	}

	encryptedFile, err := makeEncryptedFile(md.AuthId, request.GetFile())
	if err != nil {
		log.Errorf("requestEncryption error: {%v}", err)
		return nil, err
	}

	date := int32(time.Now().Unix())
	encryptedMessage := &mtproto.EncryptedMessage{
		Constructor: mtproto.CRC32_encryptedMessage,
		RandomId:    request.GetRandomId(),
		ChatId:      encryptChatData.Id,
		Date:        date,
		Bytes:       request.Data,
		File:        encryptedFile,
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
		File: encryptedFile,
	}}

	log.Debugf("messages.sendEncryptedFile - reply: %s", sentEncryptedMessage.DebugString())
	return sentEncryptedMessage.To_Messages_SentEncryptedMessage(), nil
}
