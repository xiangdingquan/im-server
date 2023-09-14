package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) ContactsGetBlocked(ctx context.Context, request *mtproto.TLContactsGetBlocked) (*mtproto.Contacts_Blocked, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.getBlocked#f57c350f - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		limit = request.Limit
	)
	if limit > 50 {
		limit = 50
	}

	blockeds := s.UserFacade.GetBlockedList(ctx, md.UserId, request.Offset, limit)

	blocked := &mtproto.Contacts_Blocked{
		Blocked_VECTORCONTACTBLOCKED: blockeds,
		Blocked_VECTORPEERBLOCKED:    make([]*mtproto.PeerBlocked, len(blockeds)),
		Users:                        make([]*mtproto.User, 0, len(blockeds)),
	}

	var uIds model.IDList
	for i, b := range blockeds {
		blocked.Blocked_VECTORPEERBLOCKED[i] = mtproto.MakeTLPeerBlocked(&mtproto.PeerBlocked{
			PeerId: model.MakePeerUser(b.UserId),
			Date:   b.Date,
		}).To_PeerBlocked()
		uIds.AddIfNot(b.UserId)
	}

	blocked.Users = s.UserFacade.GetUserListByIdList(ctx, md.UserId, uIds)
	contactsBlocked := mtproto.MakeTLContactsBlocked(blocked).To_Contacts_Blocked()
	log.Debugf("contacts.getBlocked#f57c350f - reply: %s\n", logger.JsonDebugData(contactsBlocked))
	return contactsBlocked, nil
}
