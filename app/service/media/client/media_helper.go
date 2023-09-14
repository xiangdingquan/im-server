package media_client

import (
	"time"

	"open.chat/mtproto"
)

func GetUserProfilePhoto(photoId int64) (photo *mtproto.UserProfilePhoto) {
	if photoId == 0 {
		photo = mtproto.MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	} else {
		sizes, _ := GetPhotoSizeList(photoId)
		photo = MakeUserProfilePhoto(photoId, sizes)
	}

	return
}

func GetChatPhoto(photoId int64) (photo *mtproto.ChatPhoto) {
	if photoId == 0 {
		photo = mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto()
	} else {
		sizes, _ := GetPhotoSizeList(photoId)
		photo = MakeChatPhoto(sizes)
	}

	return
}

func GetPhoto(photoId int64) (photo *mtproto.Photo) {
	if photoId == 0 {
		photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	} else {
		sizes, _ := GetPhotoSizeList(photoId)
		if len(sizes) == 0 {
			photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
		} else {
			photo2 := mtproto.MakeTLPhoto(&mtproto.Photo{
				Id:          photoId,
				HasStickers: false,
				AccessHash:  photoId,
				Date:        int32(time.Now().Unix()),
				Sizes:       sizes,
				DcId:        2,
			})
			photo = photo2.To_Photo()
		}
	}
	return
}

func MakeUserProfilePhoto(photoId int64, sizes []*mtproto.PhotoSize) *mtproto.UserProfilePhoto {
	if len(sizes) == 0 {
		return mtproto.MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	}

	photo := mtproto.MakeTLUserProfilePhoto(&mtproto.UserProfilePhoto{
		PhotoId:    photoId,
		PhotoSmall: sizes[0].GetLocation(),
		PhotoBig:   sizes[len(sizes)-1].GetLocation(),
		DcId:       2,
	})

	return photo.To_UserProfilePhoto()
}

func MakeChatPhoto(sizes []*mtproto.PhotoSize) *mtproto.ChatPhoto {
	if len(sizes) == 0 {
		return mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto()
	}

	photo := mtproto.MakeTLChatPhoto(&mtproto.ChatPhoto{
		PhotoSmall: sizes[0].GetLocation(),
		PhotoBig:   sizes[len(sizes)-1].GetLocation(),
		DcId:       2,
	})

	return photo.To_ChatPhoto()
}
