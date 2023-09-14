package service

import (
	"context"

	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) PhotosGetUserPhotos(ctx context.Context, request *mtproto.TLPhotosGetUserPhotos) (*mtproto.Photos_Photos, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("photos.getUserPhotos - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	// 400	MAX_ID_INVALID	The provided max ID is invalid
	// 400	USER_ID_INVALID	The provided user ID is invalid
	userId := model.FromInputUser(md.UserId, request.UserId)
	switch userId.PeerType {
	case model.PEER_SELF, model.PEER_USER:
	default:
		err := mtproto.ErrUserIdInvalid
		log.Errorf("photos.getUserPhotos - error: %v", err)
		return nil, err
	}

	cachePhotos, err := s.UserFacade.GetCacheUserPhotos(ctx, userId.PeerId)
	if err != nil {
		log.Errorf("photos.getUserPhotos - error: %v", err)
		return nil, err
	}

	photos := mtproto.MakeTLPhotosPhotos(&mtproto.Photos_Photos{
		Photos: make([]*mtproto.Photo, 0, len(cachePhotos.IdList)),
		Users:  []*mtproto.User{},
	}).To_Photos_Photos()

	for _, id := range cachePhotos.IdList {
		photos.Photos = append(photos.Photos, media_client.GetPhoto(id))
	}

	log.Debugf("photos.getUserPhotos - reply: %s", photos.DebugString())
	return photos, nil
}
