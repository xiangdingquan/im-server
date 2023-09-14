package service

import (
	"context"
	"open.chat/app/interface/botway/botapi"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) SendSticker(ctx context.Context, token string, req *botapi.SendSticker2) (*botapi.Message, error) {
	log.Warnf("sendSticker - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) GetStickerSet(ctx context.Context, token string, req *botapi.GetStickerSet2) (*botapi.StickerSet, error) {
	log.Warnf("getStickerSet - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) UploadStickerFile(ctx context.Context, token string, req *botapi.UploadStickerFile2) (*botapi.File, error) {
	log.Warnf("uploadStickerFile - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) CreateNewStickerSet(ctx context.Context, token string, req *botapi.CreateNewStickerSet2) (bool, error) {
	log.Warnf("createNewStickerSet - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) AddStickerToSet(ctx context.Context, token string, req *botapi.AddStickerToSet2) (bool, error) {
	log.Warnf("addStickerToSet - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) SetStickerPositionInSet(ctx context.Context, token string, req *botapi.SetStickerPositionInSet2) (bool, error) {
	log.Warnf("stopPoll - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) DeleteStickerFromSet(ctx context.Context, token string, req *botapi.DeleteStickerFromSet2) (bool, error) {
	log.Warnf("deleteStickerFromSet - method not impl")
	return false, mtproto.ErrMethodNotImpl
}
