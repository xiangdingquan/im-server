package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsAddContact(ctx context.Context, request *mtproto.TLContactsAddContact) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.addContact - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	// 400	CONTACT_ID_INVALID	The provided contact ID is invalid
	id := model.FromInputUser(md.UserId, request.Id)
	if id.PeerType != model.PEER_USER {
		err := mtproto.ErrContactIdInvalid
		log.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	users := s.UserFacade.GetMutableUsers(ctx, md.UserId, id.PeerId)
	u, ok := users.GetImmutableUser(id.PeerId)
	if !ok {
		err := mtproto.ErrContactIdInvalid
		log.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	if request.FirstName == "" {
		request.FirstName = u.User.FirstName
		request.LastName = u.User.LastName
		//err := mtproto.ErrFirstNameInvalid
		//log.Errorf("contacts.addContact - error: %v", err)
		//return nil, err
	}

	_, err := s.UserFacade.AddContact(ctx, md.UserId, id.PeerId, request.AddPhonePrivacyException, request.FirstName, request.LastName)
	if err != nil {
		log.Errorf("contacts.addContact - error: %v", err)
		return nil, err
	}

	meUpdates := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdatePeerSettings(&mtproto.Update{
		Peer_PEER: id.ToPeer(),
		Settings: mtproto.MakeTLPeerSettings(&mtproto.PeerSettings{
			ReportSpam:            false,
			AddContact:            false,
			BlockContact:          false,
			ShareContact:          false,
			NeedContactsException: false,
			ReportGeo:             false,
		}).To_PeerSettings(),
	}).To_Update())
	meUpdates.Users = s.UserFacade.GetUserListByIdList(ctx, md.UserId, []int32{md.UserId, id.PeerId})

	log.Debugf("contacts.addContact - reply: %s", meUpdates.DebugString())
	return meUpdates, nil
}
