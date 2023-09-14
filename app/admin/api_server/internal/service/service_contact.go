package service

import (
	"context"
	"strconv"

	"open.chat/app/admin/api_server/api"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) AddContact(ctx context.Context, i *api.AddContact) (r *mtproto.PeerSettings, err error) {
	log.Debugf("addContact - request: %s", i.DebugString())

	var (
		rUpdates *mtproto.Updates
		id       int64
	)

	if id, err = strconv.ParseInt(i.Id, 10, 64); err != nil {
		err = mtproto.ErrInternelServerError
		return
	}

	if rUpdates, err = s.addContact(ctx, i); err != nil {
		log.Errorf("addContact - error: %v", err)
		return
	}

	model.VisitUpdates(int32(id), rUpdates, map[string]model.UpdateVisitedFunc{
		mtproto.Predicate_updatePeerSettings: func(userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			r = update.GetSettings()
		},
	})

	if r == nil {
		log.Errorf("addContact - error: invalid Bool type")
		err = mtproto.ErrInternelServerError
		return
	}

	go func() {
		var myUser *mtproto.User
		for _, user := range rUpdates.Users {
			if int64(user.Id) == id {
				myUser = user
				break
			}
		}
		if myUser == nil {
			log.Errorf("addContact - not found id: %d", id)
			return
		}

		i2 := &api.AddContact{
			Id:                       i.ContactId,
			AddPhonePrivacyException: i.AddPhonePrivacyException,
			ContactId:                i.Id,
			FirstName:                myUser.GetFirstName().GetValue(),
			LastName:                 myUser.GetLastName().GetValue(),
			Phone:                    myUser.GetPhone().GetValue(),
		}

		if _, err2 := s.addContact(context.Background(), i2); err2 != nil {
			log.Errorf("addContact error: %v", err2)
		}
	}()

	log.Debugf("addContact - reply: %s", r.DebugString())
	return
}

func (s *Service) addContact(ctx context.Context, i *api.AddContact) (r *mtproto.Updates, err error) {
	var (
		result        mtproto.TLObject
		ok            bool
		id, contactId int64
	)

	if id, err = strconv.ParseInt(i.Id, 10, 64); err != nil {
		err = mtproto.ErrInternelServerError
		return
	}

	if contactId, err = strconv.ParseInt(i.ContactId, 10, 64); err != nil {
		err = mtproto.ErrInternelServerError
		return
	}

	req := &mtproto.TLContactsAddContact{
		AddPhonePrivacyException: i.AddPhonePrivacyException,
		Id:                       mtproto.MakeTLInputUser(&mtproto.InputUser{UserId: int32(contactId)}).To_InputUser(),
		FirstName:                i.FirstName,
		LastName:                 i.LastName,
		Phone:                    i.Phone,
	}

	result, err = s.Invoke(ctx, int32(id), req)
	if err != nil {
		log.Errorf("addContact - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.Updates); !ok {
		log.Errorf("addContact - error: invalid Bool type")
		err = mtproto.ErrInternelServerError
		return
	}

	return
}
