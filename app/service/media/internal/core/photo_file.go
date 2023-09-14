package core

import (
	"context"

	"open.chat/app/service/dfs/dfspb"
	"open.chat/app/service/media/internal/dal/dataobject"
	"open.chat/mtproto"
)

const (
	kPhotoSizeOriginalType = "0"
	kPhotoSizeSmallType    = "s"
	kPhotoSizeMediumType   = "m"
	kPhotoSizeXLargeType   = "x"
	kPhotoSizeYLargeType   = "y"
	kPhotoSizeAType        = "a"
	kPhotoSizeBType        = "b"
	kPhotoSizeCType        = "c"

	kPhotoSizeOriginalSize = 0
	kPhotoSizeSmallSize    = 90
	kPhotoSizeMediumSize   = 320
	kPhotoSizeXLargeSize   = 800
	kPhotoSizeYLargeSize   = 1280
	kPhotoSizeASize        = 160
	kPhotoSizeBSize        = 320
	kPhotoSizeCSize        = 640

	kPhotoSizeAIndex = 4
)

var sizeList = []int{
	kPhotoSizeOriginalSize,
	kPhotoSizeSmallSize,
	kPhotoSizeMediumSize,
	kPhotoSizeXLargeSize,
	kPhotoSizeYLargeSize,
	kPhotoSizeASize,
	kPhotoSizeBSize,
	kPhotoSizeCSize,
}

func getSizeType(idx int) string {
	switch idx {
	case 0:
		return kPhotoSizeOriginalType
	case 1:
		return kPhotoSizeSmallType
	case 2:
		return kPhotoSizeMediumType
	case 3:
		return kPhotoSizeXLargeType
	case 4:
		return kPhotoSizeYLargeType
	case 5:
		return kPhotoSizeAType
	case 6:
		return kPhotoSizeBType
	case 7:
		return kPhotoSizeCType
	}

	return ""
}

type resizeInfo struct {
	isWidth bool
	size    int
}

func (m *MediaCore) GetPhotoSizeList(ctx context.Context, photoId int64) (sizes []*mtproto.PhotoSize) {
	doList, _ := m.PhotoDatasDAO.SelectListByPhotoId(ctx, photoId)
	sizes = make([]*mtproto.PhotoSize, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		sizeData := &mtproto.PhotoSize{
			PredicateName: mtproto.Predicate_photoSize,
			Type:          getSizeType(int(doList[i].LocalId)),
			W:             doList[i].Width,
			H:             doList[i].Height,
			Size2:         doList[i].FileSize,
			Location: &mtproto.FileLocation{
				PredicateName: mtproto.Predicate_fileLocation,
				VolumeId:      doList[i].VolumeId,
				LocalId:       int32(doList[i].LocalId),
				Secret:        doList[i].AccessHash,
				DcId:          doList[i].DcId,
			},
		}

		sizes = append(sizes, sizeData)
	}
	return
}

func (m *MediaCore) UploadPhotoFile2(ctx context.Context, fileMDList []*dfspb.PhotoFileMetadata) (photoId, accessHsh int64, sizeList []*mtproto.PhotoSize, err error) {
	sizeList = make([]*mtproto.PhotoSize, 0, 4)

	for i, fileMD := range fileMDList {
		photoDatasDO := &dataobject.PhotoDatasDO{
			PhotoId:    fileMD.PhotoId,
			PhotoType:  fileMD.PhotoType,
			DcId:       fileMD.DcId,
			VolumeId:   fileMD.VolumeId,
			LocalId:    fileMD.LocalId,
			AccessHash: fileMD.SecretId,
			Width:      fileMD.Width,
			Height:     fileMD.Height,
			FileSize:   fileMD.FileSize,
			FilePath:   fileMD.FilePath,
			Ext:        fileMD.Ext,
		}

		m.PhotoDatasDAO.Insert(ctx, photoDatasDO)

		photoSizeData := &mtproto.PhotoSize{
			PredicateName: mtproto.Predicate_photoSize,
			Type:          getSizeType(int(fileMD.LocalId)),
			W:             photoDatasDO.Width,
			H:             photoDatasDO.Height,
			Size2:         photoDatasDO.FileSize,
			Location: &mtproto.FileLocation{
				PredicateName: mtproto.Predicate_fileLocation,
				VolumeId:      photoDatasDO.VolumeId,
				LocalId:       fileMD.LocalId,
				Secret:        photoDatasDO.AccessHash,
				DcId:          photoDatasDO.DcId,
			},
		}

		if i == 0 {
			photoId = fileMD.PhotoId
			accessHsh = fileMD.SecretId
		}
		sizeList = append(sizeList, photoSizeData)
	}
	return
}
