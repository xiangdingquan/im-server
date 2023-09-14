package service

import (
	"context"
	"time"

	"math/rand"

	"open.chat/app/messenger/msg/msgpb"
	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesEditChatPhoto(ctx context.Context, request *mtproto.TLMessagesEditChatPhoto) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.editChatPhoto#ca4c79d8 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		photoId int64 = 0
		action  *mtproto.MessageAction
	)

	chatPhoto := request.GetPhoto()
	photo := mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	switch chatPhoto.GetConstructor() {
	case mtproto.CRC32_inputChatPhotoEmpty:
		photoId = 0
		action = mtproto.MakeTLMessageActionChatDeletePhoto(nil).To_MessageAction()
	case mtproto.CRC32_inputChatUploadedPhoto_927c55b4:
		file := chatPhoto.GetFile()
		result, err := media_client.UploadPhotoFile(md.AuthId, file)
		if err != nil {
			log.Errorf("UploadPhoto error: %v", err)
			return nil, err
		}
		photoId = result.PhotoId
		photo = mtproto.MakeTLPhoto(&mtproto.Photo{
			Id:          photoId,
			HasStickers: false,
			AccessHash:  result.AccessHash,
			Date:        int32(time.Now().Unix()),
			Sizes:       result.SizeList,
			DcId:        2,
		}).To_Photo()
		action = model.MakeMessageActionChatEditPhoto(photo)
	case mtproto.CRC32_inputChatUploadedPhoto_c642724e:
		log.Warnf("not impl CRC32_inputChatUploadedPhoto_c642724e")
		return nil, mtproto.ErrInputRequestInvalid
	case mtproto.CRC32_inputChatPhoto:
	default:
	}

	chat, err := s.ChatFacade.EditChatPhoto(ctx, request.ChatId, md.UserId, photo)
	if err != nil {
		log.Errorf("messages.editChatDefaultBannedRights - error: %v", err)
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
			Message:      chat.MakeMessageService(md.UserId, action),
			ScheduleDate: nil,
		})

	if err != nil {
		log.Errorf("messages.editChatTitle - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.editChatPhoto#ca4c79d8 - reply: {%s}", replyUpdates.DebugString())
	return replyUpdates, nil
}
