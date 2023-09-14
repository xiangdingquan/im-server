package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsUnblock(ctx context.Context, request *mtproto.TLContactsUnblock) (reply *mtproto.Bool, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.unblock - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		blockedId int32
		id        = request.Id_INPUTPEER
	)

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("contacts.unblock - error: %v", err)
		return
	}

	if request.Id_INPUTUSER != nil {
		id = model.FromInputUser(md.UserId, request.Id_INPUTUSER).ToInputPeer()
	}

	switch id.PredicateName {
	case mtproto.Predicate_inputPeerUser:
		blockedId = id.UserId
	default:
		err = mtproto.ErrContactIdInvalid
		log.Errorf("contacts.unblock - error: %v", err)
		return
	}

	unBlocked := s.UserFacade.UnBlockUser(ctx, md.UserId, blockedId)
	if unBlocked {
		go func() {
			// Sync unblocked: updateUserBlocked
			updateUserUnBlocked := mtproto.MakeTLUpdateUserBlocked(&mtproto.Update{
				UserId:  blockedId,
				Blocked: mtproto.ToBool(false),
			}).To_Update()

			unBlockedUpdates := model.NewUpdatesLogic(md.UserId)
			unBlockedUpdates.AddUpdate(updateUserUnBlocked)
			user, _ := s.UserFacade.GetUserById(context.Background(), md.UserId, blockedId)
			unBlockedUpdates.AddUser(user)
			sync_client.PushUpdates(context.Background(), blockedId, unBlockedUpdates.ToUpdates())
		}()
	}

	log.Debugf("contacts.unblock#e54100bd - reply: {%v}", unBlocked)
	return mtproto.ToBool(unBlocked), nil
}
