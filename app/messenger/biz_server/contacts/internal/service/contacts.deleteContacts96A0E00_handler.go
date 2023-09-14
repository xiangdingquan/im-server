package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsDeleteContacts96A0E00(ctx context.Context, request *mtproto.TLContactsDeleteContacts96A0E00) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.deleteContacts96A0E00 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	updates := model.MakeEmptyUpdates()

	for _, id := range request.GetId() {
		if id.GetPredicateName() != mtproto.Predicate_inputUser {
			continue
		}
		deleteUser, _ := s.UserFacade.GetUserById(ctx, md.UserId, id.GetUserId())
		if deleteUser == nil {
			continue
		}
		s.UserFacade.DeleteContact(ctx, md.UserId, id.GetUserId(), deleteUser.GetMutualContact())
		updates.Updates = append(updates.Updates, mtproto.MakeTLUpdatePeerSettings(&mtproto.Update{
			Peer_PEER: model.MakePeerUser(id.GetUserId()),
			Settings: mtproto.MakeTLPeerSettings(&mtproto.PeerSettings{
				ReportSpam:            false,
				AddContact:            false,
				BlockContact:          false,
				ShareContact:          false,
				NeedContactsException: false,
				ReportGeo:             false,
				Autoarchived:          false,
				GeoDistance:           nil,
			}).To_PeerSettings(),
		}).To_Update())
		updates.Users = append(updates.Users, deleteUser)
	}

	log.Debugf("contacts.deleteContacts96A0E00 - reply: %s", updates.DebugString())
	return model.WrapperGoFunc(updates, func() {
		sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, updates)
	}).(*mtproto.Updates), nil
}
