package core

import (
	"context"
	"encoding/json"

	"open.chat/model"
	"open.chat/pkg/hack"
)

func (m *UserCore) GetCacheUserPhotos(ctx context.Context, userId int32) (*model.UserPhotos, error) {
	do, err := m.UsersDAO.SelectProfilePhotos(ctx, userId)
	if err != nil {
		return nil, err
	}

	photos := model.MakeUserPhotos()
	if do != nil {
		photos = new(model.UserPhotos)
		json.Unmarshal(hack.Bytes(do.Photos), photos)
	}

	return photos, nil
}

func (m *UserCore) PutCacheUserPhotos(ctx context.Context, userId int32, photos *model.UserPhotos) error {
	pData, _ := json.Marshal(photos)
	_, err := m.UsersDAO.UpdateProfilePhotos(ctx, hack.String(pData), userId)
	return err
}
