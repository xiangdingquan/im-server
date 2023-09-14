package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsGetContacts(ctx context.Context, request *mtproto.TLContactsGetContacts) (*mtproto.Contacts_Contacts, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.getContacts#c023849f - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		contacts *mtproto.Contacts_Contacts
	)

	contactList := s.UserFacade.GetContactList(ctx, md.UserId)
	if len(contactList) > 0 {
		idList := make([]int32, 0, len(contactList))
		for _, c := range contactList {
			idList = append(idList, c.UserId)
		}

		log.Debugf("contactIdList - {%v}", idList)

		users := s.UserFacade.GetUserListByIdList(ctx, md.UserId, idList)
		contacts = mtproto.MakeTLContactsContacts(&mtproto.Contacts_Contacts{
			Contacts:   contactList,
			SavedCount: 0,
			Users:      users,
		}).To_Contacts_Contacts()
	} else {
		contacts = mtproto.MakeTLContactsContacts(nil).To_Contacts_Contacts()
	}

	log.Debugf("contacts.getContacts#c023849f - reply: %s\n", contacts.DebugString())
	return contacts, nil
}
