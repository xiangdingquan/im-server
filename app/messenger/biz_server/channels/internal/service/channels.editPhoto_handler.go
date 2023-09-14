package service

import (
	"context"
	"math/rand"
	"time"

	"open.chat/app/messenger/msg/msgpb"
	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) makeEditPhotoAction(
	ctx context.Context,
	authKeyId int64,
	chatPhoto *mtproto.InputChatPhoto) (*mtproto.MessageAction, int64, error) {

	var (
		action  *mtproto.MessageAction
		photoId int64
	)

	switch chatPhoto.GetPredicateName() {
	case mtproto.Predicate_inputChatPhotoEmpty:
		photoId = 0
		action = model.MakeMessageActionChatDeletePhoto()
	case mtproto.Predicate_inputChatUploadedPhoto:
		result, err := media_client.UploadPhotoFile(authKeyId, chatPhoto.File)
		if err != nil {
			log.Errorf("UploadPhoto error: %v", err)
			return nil, 0, err
		}
		action = model.MakeMessageActionChatEditPhoto(mtproto.MakeTLPhoto(&mtproto.Photo{
			Id:          result.PhotoId,
			HasStickers: false,
			AccessHash:  result.AccessHash,
			Date:        int32(time.Now().Unix()),
			Sizes:       result.SizeList,
			DcId:        2,
		}).To_Photo())
		photoId = result.PhotoId
	case mtproto.Predicate_inputChatPhoto:
		log.Warnf("not impl inputChatPhoto")

		photoId = 0
		action = model.MakeMessageActionChatDeletePhoto()
	default:
		log.Errorf("invalid classId: %d", chatPhoto.GetPredicateName())

		photoId = 0
		action = model.MakeMessageActionChatDeletePhoto()
	}

	return action, photoId, nil
}

func (s *Service) ChannelsEditPhoto(ctx context.Context, request *mtproto.TLChannelsEditPhoto) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.editPhoto#f12e57c9 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.editPhoto - error: %v", err)
		return nil, err
	}

	chatPhotoAction, photoId, err := s.makeEditPhotoAction(ctx, md.AuthId, request.Photo)
	if err != nil {
		log.Errorf("channels.editPhoto - error: %v", err)
		return nil, err
	}

	var chatPhoto *mtproto.Photo
	if photoId == 0 {
		chatPhoto = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	} else {
		chatPhoto = chatPhotoAction.Photo
	}
	channel, err := s.ChannelFacade.EditPhoto(ctx, request.Channel.ChannelId, md.UserId, chatPhoto)
	if err != nil {
		log.Errorf("channels.editPhoto - error: %v", err)
		return nil, err
	}

	result, err := s.MsgFacade.SendMessage(
		ctx,
		md.UserId,
		md.AuthId,
		model.MakeChannelPeerUtil(request.Channel.ChannelId),
		&msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      channel.MakeMessageService(md.UserId, false, 0, chatPhotoAction),
			ScheduleDate: nil,
		})

	if err != nil {
		log.Errorf("channels.editPhoto - error: %v", err)
		return nil, err
	}

	log.Debugf("channels.editPhoto - reply: {%s}", result.DebugString())
	return result, nil
}
