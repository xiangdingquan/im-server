package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesClearAllDrafts(ctx context.Context, request *mtproto.TLMessagesClearAllDrafts) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.clearAllDrafts - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peerUsers    []int32
		peerChats    []int32
		peerChannels []int32
	)

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.clearAllDrafts - error: %v", err)
		return nil, err
	}

	if peerUsers, _, _ = s.PrivateFacade.GetAllDrafts(ctx, md.UserId); len(peerUsers) > 0 {
		s.PrivateFacade.ClearAllDrafts(ctx, md.UserId)
	}

	if peerChats, _, _ = s.ChatFacade.GetAllDrafts(ctx, md.UserId); len(peerChats) > 0 {
		s.ChatFacade.ClearAllDrafts(ctx, md.UserId)
	}

	if peerChannels, _, _ = s.ChannelFacade.GetAllDrafts(ctx, md.UserId); len(peerChats) > 0 {
		s.ChannelFacade.ClearAllDrafts(ctx, md.UserId)
	}

	reply := mtproto.ToBool(true)
	log.Debugf("messages.clearAllDrafts - reply: %s", reply.DebugString())

	if len(peerUsers) == 0 && len(peerChats) == 0 && len(peerChannels) == 0 {
		return reply, nil
	} else {
		return model.WrapperGoFunc(reply, func() {
			ctx := context.Background()
			syncUpdatesHelper := model.MakeUpdatesHelper()
			draft := model.MakeDraftMessageEmpty(0)

			for _, id := range peerUsers {
				updateDraftMessage := mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
					Peer_PEER: model.MakePeerUser(id),
					Draft:     draft,
				})
				syncUpdatesHelper.PushBackUpdate(updateDraftMessage.To_Update())
			}

			for _, id := range peerChats {
				updateDraftMessage := mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
					Peer_PEER: model.MakePeerChat(id),
					Draft:     draft,
				})
				syncUpdatesHelper.PushBackUpdate(updateDraftMessage.To_Update())
			}

			for _, id := range peerChannels {
				updateDraftMessage := mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
					Peer_PEER: model.MakePeerChannel(id),
					Draft:     draft,
				})
				syncUpdatesHelper.PushBackUpdate(updateDraftMessage.To_Update())
			}
			syncUpdates := syncUpdatesHelper.ToSyncNotMeUpdates(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)

			sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, syncUpdates)

		}).(*mtproto.Bool), nil
	}
}
