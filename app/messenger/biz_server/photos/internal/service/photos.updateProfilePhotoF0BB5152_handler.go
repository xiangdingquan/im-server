package service

import (
	"context"
	"time"

	sync_client "open.chat/app/messenger/sync/client"
	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) PhotosUpdateProfilePhotoF0BB5152(ctx context.Context, request *mtproto.TLPhotosUpdateProfilePhotoF0BB5152) (*mtproto.UserProfilePhoto, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("photos.updateProfilePhoto - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("photos.updateProfilePhoto - error: %v", err)
		return nil, err
	}

	// 400	FILE_PARTS_INVALID	The number of file parts is invalid
	// 400	IMAGE_PROCESS_FAILED	Failure while processing image
	// 400	PHOTO_CROP_SIZE_SMALL	Photo is too small

	var photo *mtproto.UserProfilePhoto

	switch request.GetId().GetConstructor() {
	case mtproto.CRC32_inputPhotoEmpty:
		photos, err := s.UserFacade.GetCacheUserPhotos(ctx, md.UserId)
		if err != nil {
			log.Errorf("photos.updateProfilePhoto - error: %v", err)
			return nil, mtproto.ErrInternelServerError
		}
		photos.RemovePhotoId(photos.GetDefaultPhotoId(), func(id int64) *mtproto.Photo {
			photo := media_client.GetPhoto(id)
			return photo
		})

		err = s.UserFacade.PutCacheUserPhotos(ctx, md.UserId, photos)
		if err != nil {
			log.Errorf("photos.updateProfilePhoto - error: %v", err)
		}

		photo = photos.ToUserProfilePhoto()
	default:
		id := request.GetId().To_InputPhoto()
		sizes, _ := media_client.GetPhotoSizeList(id.GetId())
		photo = media_client.MakeUserProfilePhoto(id.GetId(), sizes)
	}

	log.Debugf("photos.updateProfilePhoto - reply: %s", photo.DebugString())
	return model.WrapperGoFunc(photo, func() {
		pushUpdates := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateUserPhoto(&mtproto.Update{
			UserId:   md.UserId,
			Date:     int32(time.Now().Unix()),
			Photo:    photo,
			Previous: mtproto.ToBool(false),
		}).To_Update())
		sync_client.PushUpdates(context.Background(), md.UserId, pushUpdates)
	}).(*mtproto.UserProfilePhoto), nil
}
