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

func (s *Service) PhotosUploadProfilePhoto(ctx context.Context, request *mtproto.TLPhotosUploadProfilePhoto) (*mtproto.Photos_Photo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("photos.uploadProfilePhoto - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("photos.uploadProfilePhoto - error: %v", err)
		return nil, err
	}

	file := request.GetFile()
	result, err := media_client.UploadProfilePhotoFile(md.AuthId, file)
	if err != nil {
		log.Errorf("UploadPhoto error: %v", err)
		return nil, err
	}

	cachePhotos, err := s.UserFacade.GetCacheUserPhotos(ctx, md.UserId)
	if err != nil {
		log.Errorf("UploadPhoto error: %v", err)
		return nil, err
	}

	cachePhotos.AddPhotoId(result.PhotoId, func(id int64) *mtproto.Photo {
		return mtproto.MakeTLPhoto(&mtproto.Photo{
			Id:          result.PhotoId,
			HasStickers: false,
			AccessHash:  result.AccessHash,
			Date:        int32(time.Now().Unix()),
			Sizes:       result.SizeList,
			DcId:        2,
		}).To_Photo()
	})

	err = s.UserFacade.PutCacheUserPhotos(ctx, md.UserId, cachePhotos)
	if err != nil {
		log.Errorf("UploadPhoto error: %v", err)
	}

	photos := mtproto.MakeTLPhotosPhoto(&mtproto.Photos_Photo{
		Photo: cachePhotos.Photo,
		Users: []*mtproto.User{},
	}).To_Photos_Photo()

	log.Debugf("photos.uploadProfilePhoto - reply: %s", photos.DebugString())
	return model.WrapperGoFunc(photos, func() {
		pushUpdates := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateUserPhoto(&mtproto.Update{
			UserId:   md.UserId,
			Date:     int32(time.Now().Unix()),
			Photo:    media_client.MakeUserProfilePhoto(result.PhotoId, result.SizeList),
			Previous: mtproto.ToBool(false),
		}).To_Update())
		sync_client.PushUpdates(context.Background(), md.UserId, pushUpdates)
	}).(*mtproto.Photos_Photo), nil
}
