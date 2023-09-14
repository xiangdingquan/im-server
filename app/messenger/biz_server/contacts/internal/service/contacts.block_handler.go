package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsBlock(ctx context.Context, request *mtproto.TLContactsBlock) (reply *mtproto.Bool, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.block#332b49fc - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		blockId int32
		id      = request.Id_INPUTPEER
	)

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("contacts.block - error: %v", err)
		return
	}
	if request.Id_INPUTUSER != nil {
		id = model.FromInputUser(md.UserId, request.Id_INPUTUSER).ToInputPeer()
	}

	switch id.PredicateName {
	case mtproto.Predicate_inputPeerUser:
		blockId = id.UserId
	default:
		err = mtproto.ErrContactIdInvalid
		log.Errorf("contacts.block - error: %v", err)
		return
	}

	blocked := s.UserFacade.BlockUser(ctx, md.UserId, blockId)

	if blocked {
		go func() {
			updateUserBlocked := mtproto.MakeTLUpdateUserBlocked(&mtproto.Update{
				UserId:  blockId,
				Blocked: mtproto.ToBool(true),
			}).To_Update()

			blockedUpdates := model.NewUpdatesLogic(md.UserId)
			blockedUpdates.AddUpdate(updateUserBlocked)
			user, _ := s.UserFacade.GetUserById(context.Background(), md.UserId, blockId)
			blockedUpdates.AddUser(user)

			sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, blockedUpdates.ToUpdates())
		}()
	}

	// Blocked会影响收件箱
	reply = mtproto.ToBool(blocked)
	log.Debugf("contacts.block#332b49fc - reply: {%v}", reply.DebugString())
	return
}
