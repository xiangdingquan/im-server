package core

import (
	"context"
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/random2"
	"open.chat/pkg/util"
)

func makeUserDataByDO(userDO *dataobject.UsersDO, lastSeenAt int32) *model.UserData {
	uData := &model.UserData{
		Id:                userDO.Id,
		UserType:          model.UserTypeRegular,
		AccessHash:        userDO.AccessHash,
		SecretKeyId:       userDO.SecretKeyId,
		FirstName:         userDO.FirstName,
		LastName:          userDO.LastName,
		Username:          userDO.Username,
		Phone:             userDO.Phone,
		CountryCode:       userDO.CountryCode,
		Country:           userDO.Country,
		Province:          userDO.Province,
		City:              userDO.City,
		CityCode:          userDO.CityCode,
		Gender:            int32(userDO.Gender),
		Birth:             userDO.Birth,
		Verified:          util.Int8ToBool(userDO.Verified),
		About:             userDO.About,
		State:             userDO.State,
		IsBot:             util.Int8ToBool(userDO.IsBot),
		IsVirtualUser:     util.Int8ToBool(userDO.IsVirtualUser),
		IsInternal:        util.Int8ToBool(userDO.IsInternal),
		AccountDaysTtl:    userDO.AccountDaysTtl,
		Photo:             nil,
		ProfilePhoto:      nil,
		Min:               util.Int8ToBool(userDO.Min),
		Restricted:        util.Int8ToBool(userDO.Restricted),
		RestrictionReason: userDO.RestrictionReason,
		Deleted:           util.Int8ToBool(userDO.Deleted),
		DeleteReason:      userDO.DeleteReason,
		LastSeenAt:        lastSeenAt,
	}

	if userDO.Photos == "" {
		uData.Photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
		uData.ProfilePhoto = mtproto.MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	} else {
		cachePhotos := new(model.UserPhotos)
		err := json.Unmarshal(hack.Bytes(userDO.Photos), cachePhotos)
		if err != nil {
			uData.Photo = mtproto.MakeTLPhotoEmpty(nil).To_Photo()
			uData.ProfilePhoto = mtproto.MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
		} else {
			uData.Photo = cachePhotos.Photo
			uData.ProfilePhoto = cachePhotos.ToUserProfilePhoto()
		}
	}
	return uData
}

func makeBotDataByDO(botDO *dataobject.BotsDO) (bData *model.BotData) {
	if botDO != nil {
		bData = &model.BotData{
			Id:                   botDO.Id,
			BotId:                botDO.BotId,
			BotType:              int(botDO.BotType),
			CreatorUserId:        botDO.CreatorUserId,
			Token:                botDO.Token,
			Description:          botDO.Description,
			BotChatHistory:       util.Int8ToBool(botDO.BotChatHistory),
			BotNochats:           util.Int8ToBool(botDO.BotNochats),
			Verified:             util.Int8ToBool(botDO.Verified),
			BotInlineGeo:         util.Int8ToBool(botDO.BotInlineGeo),
			BotInfoVersion:       botDO.BotInfoVersion,
			BotInlinePlaceholder: botDO.BotInlinePlaceholder,
		}
	}
	return bData
}

func makeUserContactByDO(c *dataobject.UserContactsDO) *model.Contact {
	return &model.Contact{
		SelfUserId:    c.OwnerUserId,
		ContactUserId: c.ContactUserId,
		PhoneNumber:   c.ContactPhone,
		FirstName:     c.ContactFirstName,
		LastName:      c.ContactLastName,
	}

}
func (m *UserCore) CreateNewUser(ctx context.Context, keyId int64, phoneNumber, countryCode, firstName, lastName string) (*mtproto.User, error) {
	var (
		err          error
		userDO       *dataobject.UsersDO
		now          = time.Now().Unix()
		defaultRules = []*mtproto.PrivacyRule{
			mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule(),
		}

		phoneNumberRules = []*mtproto.PrivacyRule{
			mtproto.MakeTLPrivacyValueDisallowAll(nil).To_PrivacyRule(),
		}
	)

	isVirtualUser := strings.HasPrefix(phoneNumber, "86101")
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		userDO = &dataobject.UsersDO{
			UserType:      0,
			AccessHash:    rand.Int63(),
			Phone:         phoneNumber,
			SecretKeyId:   keyId,
			FirstName:     firstName,
			LastName:      lastName,
			CountryCode:   countryCode,
			IsVirtualUser: util.BoolToInt8(isVirtualUser),
		}
		if lastInsertId, _, err := m.UsersDAO.InsertTx(tx, userDO); err != nil {
			if sqlx.IsDuplicate(err) {
				result.Err = mtproto.ErrPhoneNumberOccupied
				return
			}
			result.Err = err
			return
		} else {
			userDO.Id = int32(lastInsertId)
		}

		// presences
		presencesDO := &dataobject.UserPresencesDO{
			UserId:     userDO.Id,
			LastSeenAt: now,
		}

		if _, _, err := m.UserPresencesDAO.InsertTx(tx, presencesDO); err != nil {
			result.Err = err
			return
		}

		if _, _, err := m.UserWalletDAO.InsertTx(tx, &dataobject.UserWalletDO{
			UID:     userDO.Id,
			Address: random2.RandomAlphanumeric(uint(rand.Intn(10) + 55)),
			Data:    int32(now),
		}); err != nil {
			result.Err = err
			return
		}

		if _, _, err := m.UserBlogsDAO.InsertTx(tx, &dataobject.UserBlogsDO{
			UserId: userDO.Id,
			Data:   int32(now),
		}); err != nil {
			result.Err = err
			return
		}

		// privacy
		bData, _ := json.Marshal(defaultRules)
		bData2, _ := json.Marshal(phoneNumberRules)
		doList := make([]*dataobject.UserPrivaciesDO, 0, model.MAX_KEY_TYPE)
		for i := model.STATUS_TIMESTAMP; i <= model.MAX_KEY_TYPE; i++ {
			if i == model.PHONE_NUMBER {
				doList = append(doList, &dataobject.UserPrivaciesDO{
					Id:      1,
					UserId:  userDO.Id,
					KeyType: int8(i),
					Rules:   hack.String(bData2),
				})
			} else {
				doList = append(doList, &dataobject.UserPrivaciesDO{
					Id:      1,
					UserId:  userDO.Id,
					KeyType: int8(i),
					Rules:   hack.String(bData),
				})
			}
		}

		log.Debugf("doList - %v", doList)
		_, _, err = m.UserPrivaciesDAO.InsertBulk(ctx, doList)
		if err != nil {
			result.Err = err
			return
		}
	})

	if tR.Err != nil {
		log.Errorf("createNewUser2 error: %v", tR.Err)
		return nil, tR.Err
	}

	// put privacy to cache
	var privacyList = make(map[int][]*mtproto.PrivacyRule, model.MAX_KEY_TYPE)
	for i := model.STATUS_TIMESTAMP; i <= model.MAX_KEY_TYPE; i++ {
		if i == model.PHONE_NUMBER {
			privacyList[i] = phoneNumberRules
		} else {
			privacyList[i] = defaultRules
		}
	}

	user := model.MakeImmutableUser(
		makeUserDataByDO(userDO, int32(now)),
		nil,
		nil,
		privacyList)

	// put to cache
	err = m.Redis.PutCacheUser2(ctx, user)
	err = m.Redis.SetPrivacyList(ctx, userDO.Id, privacyList)
	m.Redis.SetContactList(ctx, userDO.Id)

	return user.ToImmutableUser(user), nil
}

func (m *UserCore) getImmutableUser(ctx context.Context, id int32) (user *model.ImmutableUser, err error) {
	var (
		usersDO     *dataobject.UsersDO
		presencesDO *dataobject.UserPresencesDO
		botsDO      *dataobject.BotsDO
		lastSeenAt  int64
	)

	usersDO, err = m.UsersDAO.SelectById(ctx, id)
	if err != nil {
		return
	} else if usersDO == nil {
		err = mtproto.ErrUserIdInvalid
		return
	}
	if presencesDO, _ = m.UserPresencesDAO.Select(ctx, id); presencesDO != nil {
		lastSeenAt = presencesDO.LastSeenAt
	}

	if usersDO.IsVirtualUser != 0 {
		lastSeenAt = time.Now().Unix()
	}

	if usersDO.IsBot == 1 {
		botsDO, _ = m.BotsDAO.Select(ctx, id)
	}

	user = model.MakeImmutableUser(
		makeUserDataByDO(usersDO, int32(lastSeenAt)),
		makeBotDataByDO(botsDO),
		nil,
		nil)
	return
}

func (m *UserCore) GetMutableUsers(ctx context.Context, idList ...int32) (users model.MutableUsers) {
	return m.getMutableUsersBySelfId(ctx, 0, idList...)
}

func (m *UserCore) getMutableUsersBySelfId(ctx context.Context, selfId int32, idList ...int32) (users model.MutableUsers) {
	users = model.NewMutableUsers()

	if selfId == 0 && len(idList) == 0 {
		return
	}

	if selfId != 0 {
		idList = append(idList, selfId)
	}

	doList, _ := m.UsersDAO.SelectUsersByIdList(ctx, idList)
	if len(doList) == 0 {
		log.Warnf("getContactsByIdr(%v) - not found", idList)
		return
	}

	var (
		botIdList []int32
		idList2   []int32
		now       = time.Now().Unix()
	)

	for i := 0; i < len(doList); i++ {
		if doList[i].Deleted == 1 {
			users[doList[i].Id] = model.MakeImmutableUser(&model.UserData{
				Deleted:    true,
				Id:         doList[i].Id,
				AccessHash: doList[i].AccessHash,
			}, nil, nil, nil)
		} else {
			if doList[i].IsBot == 1 {
				botIdList = append(botIdList, doList[i].Id)
			} else {
				idList2 = append(idList2, doList[i].Id)
			}
			users[doList[i].Id] = model.MakeImmutableUser(
				makeUserDataByDO(&doList[i], 0),
				nil,
				nil,
				nil)
		}
	}

	// 1. bot
	if len(botIdList) > 0 {
		botDOList, _ := m.BotsDAO.SelectByIdList(ctx, botIdList)
		for i := 0; i < len(botIdList); i++ {
			users[botDOList[i].BotId].Bot = makeBotDataByDO(&botDOList[i])
		}
	}

	// 2. lastSeen
	if len(idList2) > 0 {
		presenceDOList, _ := m.UserPresencesDAO.SelectList(ctx, idList2)
		for i := 0; i < len(presenceDOList); i++ {
			users[presenceDOList[i].UserId].User.LastSeenAt = int32(presenceDOList[i].LastSeenAt)
			if users[presenceDOList[i].UserId].User.IsVirtualUser {
				users[presenceDOList[i].UserId].User.LastSeenAt = int32(now)
			}
		}
	}

	// contacts
	if len(idList2) > 1 {
		var contactDOList []dataobject.UserContactsDO
		if selfId != 0 {
			contactDOList, _ = m.UserContactsDAO.SelectListByIdList(ctx, selfId, idList2)
		} else {
			contactDOList, _ = m.UserContactsDAO.SelectListByOwnerListAndContactList(ctx, idList2, idList2)
		}
		for i := 0; i < len(contactDOList); i++ {
			if users[contactDOList[i].OwnerUserId].Contacts == nil {
				users[contactDOList[i].OwnerUserId].Contacts = map[int32]*model.Contact{
					contactDOList[i].ContactUserId: makeUserContactByDO(&contactDOList[i]),
				}
			} else {
				users[contactDOList[i].OwnerUserId].Contacts[contactDOList[i].ContactUserId] = makeUserContactByDO(&contactDOList[i])
			}
		}
	}

	// privacy
	if len(idList2) > 1 {
		privacyDOList, _ := m.UserPrivaciesDAO.SelectUsersPrivacyList(ctx, idList2, []int32{
			model.STATUS_TIMESTAMP,
			model.PROFILE_PHOTO,
			model.PHONE_NUMBER,
		})

		for i := 0; i < len(privacyDOList); i++ {
			if users[privacyDOList[i].UserId].PrivacyRules == nil {
				users[privacyDOList[i].UserId].PrivacyRules = map[int][]*mtproto.PrivacyRule{
					int(privacyDOList[i].KeyType): makePrivacyRulesByDO(&privacyDOList[i]),
				}
			} else {
				users[privacyDOList[i].UserId].PrivacyRules[int(privacyDOList[i].KeyType)] = makePrivacyRulesByDO(&privacyDOList[i])
			}
		}
	}

	return
}
