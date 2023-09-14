package model

import "open.chat/mtproto"

type UserPhotos struct {
	Photo  *mtproto.Photo `json:"photo,omitempty"`
	IdList []int64        `json:"id_list,omitempty"`
}

func MakeUserPhotos() *UserPhotos {
	return &UserPhotos{
		Photo:  mtproto.MakeTLPhotoEmpty(nil).To_Photo(),
		IdList: []int64{},
	}
}

func (m *UserPhotos) GetDefaultPhotoId() int64 {
	return m.Photo.GetId()
}

func (m *UserPhotos) AddPhotoId(id int64, cb func(id int64) *mtproto.Photo) {
	m.Photo = cb(id)
	m.IdList = AddID64ListFrontIfNot(m.IdList, id)
}

func (m *UserPhotos) RemovePhotoId(id int64, cb func(id int64) *mtproto.Photo) {
	if len(m.IdList) <= 1 {
		m.IdList = []int64{}
		m.Photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
	} else {
		if id == m.Photo.GetId() {
			id = m.IdList[1]
			m.IdList = m.IdList[1:]
		} else {
			for i, j := range m.IdList {
				if j == id {
					m.IdList = append(m.IdList[:i], m.IdList[i+1:]...)
				}
			}
		}
		m.Photo = cb(id)
	}
}

func (m *UserPhotos) ToUserProfilePhoto() (photo *mtproto.UserProfilePhoto) {
	return MakeUserProfilePhotoByPhoto(m.Photo)
}

func MakeChatPhotoByPhoto(photo *mtproto.Photo) (chatPhoto *mtproto.ChatPhoto) {
	switch photo.GetPredicateName() {
	case mtproto.Predicate_photo:
		sizes := photo.GetSizes()
		if len(sizes) == 0 {
			chatPhoto = mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto()
		} else {
			chatPhoto = mtproto.MakeTLChatPhoto(&mtproto.ChatPhoto{
				PhotoSmall: sizes[0].GetLocation(),
				PhotoBig:   sizes[len(sizes)-1].GetLocation(),
				DcId:       2,
			}).To_ChatPhoto()
		}
	default:
		chatPhoto = mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto()
	}
	return
}

func MakeUserProfilePhotoByPhoto(photo *mtproto.Photo) (userProfilePhoto *mtproto.UserProfilePhoto) {
	switch photo.GetPredicateName() {
	case mtproto.Predicate_photo:
		sizes := photo.GetSizes()
		if len(sizes) == 0 {
			userProfilePhoto = mtproto.MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
		} else {
			userProfilePhoto = mtproto.MakeTLUserProfilePhoto(&mtproto.UserProfilePhoto{
				PhotoId:    photo.GetId(),
				PhotoSmall: sizes[0].GetLocation(),
				PhotoBig:   sizes[len(sizes)-1].GetLocation(),
				DcId:       2,
			}).To_UserProfilePhoto()
		}
	default:
		userProfilePhoto = mtproto.MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	}
	return
}
