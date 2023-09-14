package core

import (
	"context"
	"time"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
	"open.chat/pkg/phonenumber"
	"open.chat/pkg/util"
)

func (m *UserCore) makePhoneNumber(ctx context.Context, selfId int32, Phone string) string {
	region, _ := m.GetCountryCodeByUser(ctx, selfId)
	phoneNumber, err2 := phonenumber.MakePhoneNumberHelper(Phone, region)
	if err2 != nil {
		phoneNumber, err2 = phonenumber.MakePhoneNumberHelper(Phone, "")
		if err2 != nil {
			return ""
		}
	}
	return phoneNumber.GetNormalizeDigits()
}

func (m *UserCore) getContactsById(ctx context.Context, userId int32, idList ...int32) (contacts map[int32]*model.Contact, err error) {
	var contactsDOList []dataobject.UserContactsDO
	if contactsDOList, err = m.UserContactsDAO.SelectListByIdList(ctx, userId, idList); err != nil {
		return
	}
	contacts = make(map[int32]*model.Contact, len(contactsDOList))
	cacheContacts := make([]*model.Contact, 0, len(contactsDOList))
	for i := 0; i < len(contactsDOList); i++ {
		c := &model.Contact{
			SelfUserId:    userId,
			ContactUserId: contactsDOList[i].ContactUserId,
			PhoneNumber:   contactsDOList[i].ContactPhone,
			FirstName:     contactsDOList[i].ContactFirstName,
			LastName:      contactsDOList[i].ContactLastName,
		}
		contacts[contactsDOList[i].ContactUserId] = c
		cacheContacts = append(cacheContacts, c)
	}
	return
}

func (m *UserCore) GetContactUserIdList(ctx context.Context, userId int32) []int32 {
	contactsDOList, _ := m.UserContactsDAO.SelectUserContacts(ctx, userId)
	idList := make([]int32, 0, len(contactsDOList))

	for _, do := range contactsDOList {
		idList = append(idList, do.ContactUserId)
	}
	return idList
}

func (m *UserCore) GetMutualContactUserIdList(ctx context.Context, mutual bool, userId int32) []int32 {
	contactsDOList, _ := m.UserContactsDAO.SelectUserContacts(ctx, userId)
	idList := make([]int32, 0, len(contactsDOList))
	for _, do := range contactsDOList {
		if util.Int8ToBool(do.Mutual) == mutual {
			idList = append(idList, do.ContactUserId)
		}
	}
	return idList
}

func (m *UserCore) GetContactLink(ctx context.Context, userId, contactId int32) (myLink, foreignLink *mtproto.ContactLink) {
	if userId == contactId {
		myLink = mtproto.MakeTLContactLinkContact(nil).To_ContactLink()
		foreignLink = mtproto.MakeTLContactLinkContact(nil).To_ContactLink()
	} else {
		myContactDO, _ := m.UserContactsDAO.SelectByContactId(ctx, userId, contactId)
		foreignContactDO, _ := m.UserContactsDAO.SelectByContactId(ctx, contactId, userId)

		if myContactDO == nil || myContactDO.IsDeleted == 1 {
			if myContactDO == nil {
				myLink = mtproto.MakeTLContactLinkNone(nil).To_ContactLink()
			} else {
				myLink = mtproto.MakeTLContactLinkHasPhone(nil).To_ContactLink()
			}

			if foreignContactDO == nil {
				foreignLink = mtproto.MakeTLContactLinkUnknown(nil).To_ContactLink()
			} else {
				foreignLink = mtproto.MakeTLContactLinkHasPhone(nil).To_ContactLink()
			}
		} else {
			myLink = mtproto.MakeTLContactLinkContact(nil).To_ContactLink()
			if foreignContactDO == nil {
				foreignLink = mtproto.MakeTLContactLinkNone(nil).To_ContactLink()
			} else {
				if foreignContactDO.IsDeleted == 1 {
					foreignLink = mtproto.MakeTLContactLinkHasPhone(nil).To_ContactLink()
				} else {
					foreignLink = mtproto.MakeTLContactLinkContact(nil).To_ContactLink()
				}
			}
		}
	}

	return
}

func (m *UserCore) GetAllContactList(ctx context.Context, userId int32) []*mtproto.Contact {
	doList, _ := m.UserContactsDAO.SelectAllUserContacts(ctx, userId)
	contactList := make([]*mtproto.Contact, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		c2 := &mtproto.Contact{
			PredicateName: mtproto.Predicate_contact,
			Constructor:   mtproto.CRC32_contact,
			UserId:        doList[i].ContactUserId,
			Mutual:        mtproto.ToBool(doList[i].Mutual == 1),
		}
		contactList = append(contactList, c2)
	}
	return contactList
}

func (m *UserCore) GetContactList(ctx context.Context, userId int32) []*mtproto.Contact {
	doList, _ := m.UserContactsDAO.SelectUserContacts(ctx, userId)
	contactList := make([]*mtproto.Contact, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		c2 := mtproto.MakeTLContact(&mtproto.Contact{
			UserId: doList[i].ContactUserId,
			Mutual: mtproto.ToBool(doList[i].Mutual == 1),
		}).To_Contact()
		contactList = append(contactList, c2)
	}
	contactList = append(contactList, mtproto.MakeTLContact(&mtproto.Contact{
		UserId: userId,
		Mutual: mtproto.ToBool(true),
	}).To_Contact())
	return contactList
}

func (m *UserCore) DeleteContact(ctx context.Context, userId, deleteId int32, mutual bool) bool {
	var needUpdate = false
	m.UserContactsDAO.DeleteContacts(ctx, userId, []int32{deleteId})
	if deleteId != userId && mutual {
		m.UserContactsDAO.UpdateMutual(ctx, 0, deleteId, userId)
		needUpdate = true
	}

	return needUpdate
}

func (m *UserCore) SearchContacts(ctx context.Context, userId int32, q string, limit int32) ([]int32, []int32) {
	contactList := m.GetContactList(ctx, userId)
	idList := make([]int32, 0, len(contactList)+1)
	idList = append(idList, userId)
	for _, c2 := range contactList {
		idList = append(idList, c2.UserId)
	}
	var (
		userIdList    []int32
		channelIdList []int32
	)

	q2 := q + "%"
	doList, _ := m.UsernameDAO.SearchByQueryNotIdList(ctx, q2, idList, limit)
	for i := 0; i < len(doList); i++ {
		switch doList[i].PeerType {
		case model.PEER_USER:
			isContact, _ := m.GetContactAndMutual(ctx, doList[i].PeerId, userId)
			if m.CheckPrivacy(ctx, model.ADDED_BY_USERNAME, doList[i].PeerId, userId, isContact) {
				userIdList = append(userIdList, doList[i].PeerId)
			}
		case model.PEER_CHANNEL:
			channelIdList = append(channelIdList, doList[i].PeerId)
		}
	}

	phone := m.makePhoneNumber(ctx, userId, q)
	if phone != "" && !phonenumber.IsNotPhoneUser(phone) && !phonenumber.IsVirtualUser(phone) {
		usersDOList, _ := m.UsersDAO.SelectUsersByPhoneList(ctx, []string{phone})
		for i := 0; i < len(usersDOList); i++ {
			isContact, _ := m.GetContactAndMutual(ctx, usersDOList[i].Id, userId)
			if !m.CheckPrivacy(ctx, model.ADDED_BY_PHONE, usersDOList[i].Id, userId, isContact) {
				continue
			}
			userIdList = append(userIdList, usersDOList[i].Id)
		}
	}

	return userIdList, channelIdList
}

// /////////////////////////////////////////////////////////////////////////////////////
type contactItem struct {
	c               *mtproto.InputContact
	unregistered    bool  // 未注册
	userId          int32 // 已经注册的用户ID
	contactId       int32 // 已经注册是我的联系人
	importContactId int32 // 已经注册的反向联系人
}

func (m *UserCore) ImportContacts(ctx context.Context, userId int32, contacts []*mtproto.InputContact) ([]*mtproto.ImportedContact, []*mtproto.PopularContact, []int32) {
	var (
		importedContacts  = make([]*mtproto.ImportedContact, 0, len(contacts))
		popularContactMap = make(map[string]*mtproto.TLPopularContact, len(contacts))
		updList           = make([]int32, 0, len(contacts))
	)

	importContacts := make(map[string]*contactItem)
	// 1. 整理
	phoneList := make([]string, 0, len(contacts))
	for _, c2 := range contacts {
		phoneList = append(phoneList, c2.Phone)
		importContacts[c2.Phone] = &contactItem{unregistered: true, c: c2}
	}

	// 2. 已注册
	registeredContacts, _ := m.UsersDAO.SelectUsersByPhoneList(ctx, phoneList)
	var contactIdList []int32
	for i := 0; i < len(registeredContacts); i++ {
		if c2, ok := importContacts[registeredContacts[i].Phone]; ok {
			c2.unregistered = false
			c2.userId = registeredContacts[i].Id
			phoneList = append(phoneList, registeredContacts[i].Phone)
			contactIdList = append(contactIdList, registeredContacts[i].Id)
		} else {
			c2.unregistered = true
		}
	}

	if len(contactIdList) > 0 {
		// 3. 我的联系人
		myContacts, _ := m.UserContactsDAO.SelectListByIdList(ctx, userId, contactIdList)
		log.Infof("myContacts - %v", myContacts)
		for i := 0; i < len(myContacts); i++ {
			if c2, ok := importContacts[myContacts[i].ContactPhone]; ok {
				c2.contactId = myContacts[i].ContactUserId
			}
		}
	}

	if len(contactIdList) > 0 {
		// 4. 反向联系人
		importedMyContacts, _ := m.ImportedContactsDAO.SelectListByImportedList(ctx, userId, contactIdList)
		log.Infof("importedMyContacts - %v", importedMyContacts)
		for i := 0; i < len(importedMyContacts); i++ {
			for _, c2 := range importContacts {
				if c2.userId == importedMyContacts[i].ImportedUserId {
					c2.importContactId = c2.userId
					break
				}
			}
		}
	}

	// clear phoneList
	phoneList = phoneList[0:0]
	for _, c2 := range importContacts {
		if c2.unregistered {
			go func() {
				// 1. 未注册 - popular inviter
				unregisteredContactsDO := &dataobject.UnregisteredContactsDO{
					Phone:           c2.c.Phone,
					ImporterUserId:  userId,
					ImportFirstName: c2.c.FirstName,
					ImportLastName:  c2.c.LastName,
				}
				m.UnregisteredContactsDAO.InsertOrUpdate(context.Background(), unregisteredContactsDO)
			}()
			phoneList = append(phoneList, c2.c.Phone)
			popularContact := mtproto.MakeTLPopularContact(&mtproto.PopularContact{
				ClientId:  c2.c.ClientId,
				Importers: 1,
			})
			popularContactMap[c2.c.Phone] = popularContact
		} else {
			// 已经注册
			if !m.CheckPrivacy(ctx, model.ADDED_BY_PHONE, c2.userId, userId, c2.importContactId != 0) {
				continue
			}

			userContactsDO := &dataobject.UserContactsDO{
				OwnerUserId:      userId,
				ContactUserId:    c2.userId,
				ContactPhone:     c2.c.Phone,
				ContactFirstName: c2.c.FirstName,
				ContactLastName:  c2.c.LastName,
				Date2:            int32(time.Now().Unix()),
			}

			if c2.contactId > 0 {
				if c2.importContactId > 0 {
					updList = append(updList, c2.importContactId)
				}

				// 联系人已经存在，刷新first_name, last_name
				m.UserContactsDAO.UpdateContactName(ctx, userContactsDO.ContactFirstName,
					userContactsDO.ContactLastName,
					userContactsDO.OwnerUserId,
					userContactsDO.ContactUserId)
			} else {
				userContactsDO.IsDeleted = 0
				if c2.importContactId > 0 {
					userContactsDO.Mutual = 1
					updList = append(updList, c2.importContactId)

					m.UserContactsDAO.UpdateMutual(ctx, 1, userContactsDO.ContactUserId, userContactsDO.OwnerUserId)
				} else {
					importedContactsDO := &dataobject.ImportedContactsDO{
						UserId:         userContactsDO.ContactUserId,
						ImportedUserId: userContactsDO.OwnerUserId,
					}
					m.ImportedContactsDAO.InsertOrUpdate(ctx, importedContactsDO)
				}
				m.UserContactsDAO.InsertOrUpdate(ctx, userContactsDO)
			}

			log.Infof("userContactsDO - %v", userContactsDO)
			log.Infof("c2 - %v", c2)

			importedContact := mtproto.MakeTLImportedContact(&mtproto.ImportedContact{
				UserId:   userContactsDO.ContactUserId,
				ClientId: c2.c.ClientId,
			})
			importedContacts = append(importedContacts, importedContact.To_ImportedContact())
		}
	}

	//
	popularContacts := make([]*mtproto.PopularContact, 0, len(phoneList))
	if len(phoneList) > 0 {
		popularDOList, _ := m.PopularContactsDAO.SelectImportersList(ctx, phoneList)
		for i := 0; i < len(popularDOList); i++ {
			if c2, ok := popularContactMap[popularDOList[i].Phone]; ok {
				c2.SetImporters(popularDOList[i].Importers + 1)
			}
		}

		for _, c2 := range popularContactMap {
			popularContacts = append(popularContacts, c2.To_PopularContact())
		}
	}

	return importedContacts, popularContacts, updList
}

func (m *UserCore) CheckContactAndMutualByUserId(ctx context.Context, selfId, contactId int32) (bool, bool) {
	do, _ := m.UserContactsDAO.SelectContact(ctx, selfId, contactId)
	if do == nil {
		return false, false
	} else {
		return true, do.Mutual == 1
	}
}

func (m *UserCore) GetContactAndMutual(ctx context.Context, selfUserId, id int32) (bool, bool) {
	return m.CheckContactAndMutualByUserId(ctx, selfUserId, id)
}

func (m *UserCore) BackupPhoneBooks(ctx context.Context, authKeyId int64, contacts []*mtproto.InputContact) {
	do := &dataobject.PhoneBooksDO{
		AuthKeyId: authKeyId,
	}
	for _, c := range contacts {
		if c.GetClientId() == 0 {
			continue
		}

		do.ClientId = c.GetClientId()
		do.Phone = c.GetPhone()
		do.FirstName = c.GetFirstName()
		do.LastName = c.GetLastName()
		m.PhoneBooksDAO.InsertOrUpdate(ctx, do)
	}
}

func (m *UserCore) AddContact(ctx context.Context, selfUserId, contactId int32, addPhonePrivacyException bool, contactFirstName, contactLastName string) (bool, error) {
	// 1. Check mutal_contact
	meDO, err := m.Dao.UserContactsDAO.SelectByContactId(ctx, selfUserId, contactId)
	if err != nil {
		return false, err
	}

	var (
		needCheckMutual = false
		changeMutual    = false
	)

	if meDO == nil || meDO.IsDeleted == 1 {
		needCheckMutual = true
	}

	if meDO == nil {
		meDO = &dataobject.UserContactsDO{
			OwnerUserId:      selfUserId,
			ContactUserId:    contactId,
			ContactPhone:     "",
			ContactFirstName: contactFirstName,
			ContactLastName:  contactLastName,
			Mutual:           0,
			IsDeleted:        0,
			Date2:            int32(time.Now().Unix()),
		}
	} else {
		meDO.ContactFirstName = contactFirstName
		meDO.ContactLastName = contactLastName
		meDO.IsDeleted = 0
	}

	// not contact
	if needCheckMutual {
		contactDO, _ := m.Dao.UserContactsDAO.SelectByContactId(ctx, contactId, selfUserId)
		if contactDO != nil && contactDO.IsDeleted != 1 {
			meDO.Mutual = 1
			changeMutual = true
		}
	}

	tR := sqlx.TxWrapper(ctx, m.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if changeMutual {
			m.UserContactsDAO.UpdateMutualTx(tx, 1, contactId, selfUserId)
		}
		_, _, err = m.Dao.UserContactsDAO.InsertOrUpdateTx(tx, meDO)
		if err != nil {
			result.Err = err
			return
		}
	})

	return changeMutual, tR.Err
}

func (m *UserCore) GetImportersByPhone(ctx context.Context, phone string) []*mtproto.InputContact {
	doList, _ := m.UnregisteredContactsDAO.SelectImportersByPhone(ctx, phone)
	if len(doList) == 0 {
		return []*mtproto.InputContact{}
	}

	contacts := make([]*mtproto.InputContact, len(doList))
	for i := 0; i < len(doList); i++ {
		contacts[i] = mtproto.MakeTLInputPhoneContact(&mtproto.InputContact{
			ClientId:  int64(doList[i].ImporterUserId),
			Phone:     "",
			FirstName: doList[i].ImportFirstName,
			LastName:  doList[i].ImportLastName,
		}).To_InputContact()
	}

	return contacts
}

func (m *UserCore) DeleteImportersByPhone(ctx context.Context, phone string) {
	m.UnregisteredContactsDAO.DeleteImportersByPhone(ctx, phone)
}
