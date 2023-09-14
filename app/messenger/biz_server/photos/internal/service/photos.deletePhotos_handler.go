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

func (s *Service) PhotosDeletePhotos(ctx context.Context, request *mtproto.TLPhotosDeletePhotos) (*mtproto.Vector_Long, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("photos.deletePhotos - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("photos.deletePhotos - error: %v", err)
		return nil, err
	}

	reply := &mtproto.Vector_Long{
		Datas: make([]int64, 0, len(request.Id)),
	}

	if len(request.Id) < 1 {
		log.Debugf("photos.deletePhotos#87cf7f2f - reply: %s", reply.DebugString())
		return reply, nil
	}

	cachePhotos, err := s.UserFacade.GetCacheUserPhotos(ctx, md.UserId)
	if err != nil {
		log.Errorf("photos.updateProfilePhoto - error: %v", err)
		return nil, err
	}

	for _, id := range request.Id {
		cachePhotos.RemovePhotoId(id.GetId(), func(id int64) *mtproto.Photo {
			return nil
		})
		reply.Datas = append(reply.Datas, id.GetId())
	}

	if len(cachePhotos.IdList) == 0 {
		cachePhotos.Photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	} else {
		cachePhotos.Photo = media_client.GetPhoto(cachePhotos.IdList[0])
	}

	err = s.UserFacade.PutCacheUserPhotos(ctx, md.UserId, cachePhotos)
	if err != nil {
		log.Errorf("photos.updateProfilePhoto - error: %v", err)
	}

	log.Debugf("photos.deletePhotos#87cf7f2f - reply: %v", reply)
	return model.WrapperGoFunc(reply, func() {
		pushUpdates := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateUserPhoto(&mtproto.Update{
			UserId:   md.UserId,
			Date:     int32(time.Now().Unix()),
			Photo:    cachePhotos.ToUserProfilePhoto(),
			Previous: mtproto.ToBool(false),
		}).To_Update())
		sync_client.PushUpdates(context.Background(), md.UserId, pushUpdates)
	}).(*mtproto.Vector_Long), nil
}
