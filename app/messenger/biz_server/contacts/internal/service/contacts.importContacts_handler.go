package service

import (
	"context"

	"strings"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/phonenumber"
)

func (s *Service) ContactsImportContacts(ctx context.Context, request *mtproto.TLContactsImportContacts) (reply *mtproto.Contacts_ImportedContacts, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.importContacts - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		region string
	)

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("contacts.importContacts - error: %v", err)
		return
	}

	reply = mtproto.MakeTLContactsImportedContacts(&mtproto.Contacts_ImportedContacts{
		Imported:       []*mtproto.ImportedContact{},
		PopularInvites: []*mtproto.PopularContact{},
		RetryContacts:  []int64{},
		Users:          []*mtproto.User{},
	}).To_Contacts_ImportedContacts()

	if len(request.Contacts) == 0 {
		log.Warn("contacts empty")
		return
	}

	region, _ = s.UserFacade.GetCountryCodeByUser(ctx, md.UserId)

	importList := make([]*mtproto.InputContact, 0, len(request.Contacts))
	for _, c := range request.Contacts {
		if c.Phone == "" || (strings.TrimSpace(c.FirstName) == "" && strings.TrimSpace(c.LastName) == "") {
			log.Warn("invalid contact: %v", c)
			continue
		}
		phoneNumber, err2 := phonenumber.MakePhoneNumberHelper(c.Phone, region)
		if err2 != nil {
			phoneNumber, err2 = phonenumber.MakePhoneNumberHelper(c.Phone, "")
			if err2 != nil {
				log.Warn("invalid phoneNumber(%s): %v", c.Phone, err2)
				continue
			}
		}
		c.Phone = phoneNumber.GetNormalizeDigits()
		importList = append(importList, c)
	}

	if len(importList) == 0 {
		log.Warn("contacts empty")
		return
	}

	go func() {
		// save phone books
		s.UserFacade.BackupPhoneBooks(context.Background(), md.AuthId, importList)
	}()

	importeds, popularInvites, updList := s.UserFacade.ImportContacts(ctx, md.UserId, importList)

	reply.Imported = importeds
	reply.PopularInvites = popularInvites

	var importedUsers []*mtproto.User

	importedIdList := make([]int32, 0, len(importeds))
	for _, i := range importeds {
		importedIdList = append(importedIdList, i.GetUserId())
	}

	importedUsers = s.UserFacade.GetUserListByIdList(ctx, md.UserId, importedIdList)
	reply.Users = importedUsers

	log.Debugf("contacts.importContacts#2c800be5 - reply: %s", reply.DebugString())
	return model.WrapperGoFunc(reply, func() {
		if len(importeds) > 0 {
			go func() {
				ctx = context.Background()

				updatePeerSettingsList := make([]*mtproto.Update, 0, len(importedUsers))
				for _, user := range importedUsers {
					updatePeerSettingsList = append(updatePeerSettingsList, mtproto.MakeTLUpdatePeerSettings(&mtproto.Update{
						Peer_PEER: model.MakePeerUser(user.GetId()),
						Settings: mtproto.MakeTLPeerSettings(&mtproto.PeerSettings{
							ReportSpam:            false,
							AddContact:            false,
							BlockContact:          false,
							ShareContact:          false,
							NeedContactsException: false,
							ReportGeo:             false,
						}).To_PeerSettings(),
					}).To_Update())
				}
				sync_client.SyncUpdatesMe(ctx,
					md.UserId,
					md.AuthId,
					md.SessionId,
					md.ServerId,
					model.MakeUpdatesByUpdatesUsers(importedUsers, updatePeerSettingsList...))

				_ = updList
			}()
		}
	}).(*mtproto.Contacts_ImportedContacts), nil
}
