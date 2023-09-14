package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) ContactsDeleteContact(ctx context.Context, request *mtproto.TLContactsDeleteContact) (*mtproto.Contacts_Link, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.deleteContact#8e953744 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		deleteId int32
		id       = request.Id
	)

	switch id.GetConstructor() {
	case mtproto.CRC32_inputUserSelf:
		deleteId = md.UserId
	case mtproto.CRC32_inputUser:
		if ok := s.UserFacade.CheckUserAccessHash(ctx, id.GetUserId(), id.GetAccessHash()); !ok {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
			log.Errorf("%v: is access_hash error", err)
			return nil, err
		}
		deleteId = id.GetUserId()
	default:
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		log.Errorf("%v: is inputUserEmpty", err)
		return nil, err
	}

	deleteUser, _ := s.UserFacade.GetUserById(ctx, md.UserId, deleteId)

	needUpdate := s.UserFacade.DeleteContact(ctx, md.UserId, deleteId, deleteUser.MutualContact)

	selfUpdates := model.NewUpdatesLogic(md.UserId)
	myLink, foreignLink := s.UserFacade.GetContactLink(ctx, md.UserId, deleteId)
	contactLink := &mtproto.TLUpdateContactLink{Data2: &mtproto.Update{
		UserId:      deleteId,
		MyLink:      myLink,
		ForeignLink: foreignLink,
	}}
	selfUpdates.AddUpdate(contactLink.To_Update())
	selfUpdates.AddUser(deleteUser)

	sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, selfUpdates.ToUpdates())

	if needUpdate {
		contactUpdates := model.NewUpdatesLogic(deleteUser.Id)
		myLink, foreignLink := s.UserFacade.GetContactLink(ctx, deleteId, md.UserId)
		contactLink2 := &mtproto.TLUpdateContactLink{Data2: &mtproto.Update{
			UserId:      md.UserId,
			MyLink:      myLink,
			ForeignLink: foreignLink,
		}}
		contactUpdates.AddUpdate(contactLink2.To_Update())

		selfUser, _ := s.UserFacade.GetUserById(ctx, deleteId, md.UserId)
		contactUpdates.AddUser(selfUser)
		sync_client.PushUpdates(ctx, deleteId, contactUpdates.ToUpdates())
	}

	////////////////////////////////////////////////////////////////////////////////////////
	contactsLink := &mtproto.TLContactsLink{Data2: &mtproto.Contacts_Link{
		MyLink:      contactLink.Data2.MyLink,
		ForeignLink: contactLink.Data2.ForeignLink,
		User:        deleteUser,
	}}

	log.Debugf("contacts.deleteContact#8e953744 - reply: %s", logger.JsonDebugData(contactsLink))
	return contactsLink.To_Contacts_Link(), nil
}
