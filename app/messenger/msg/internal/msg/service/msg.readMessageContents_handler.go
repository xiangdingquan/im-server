package service

import (
	"context"
	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) ReadMessageContents(ctx context.Context, r *msgpb.ReadMessageContentsRequest) (reply *mtproto.Messages_AffectedMessages, err error) {
	log.Debugf("ReadMessageContents - request: %s", r.GoString())

	if r.From == nil {
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("DeleteMessages - %v", err)
		return
	}

	var (
		pts, ptsCount int32
	)

	switch r.PeerType {
	case model.PEER_USER:
		id := make([]int32, 0, len(r.Id))
		for _, m := range r.Id {
			s.MsgCore.UpdateMediaUnread(ctx, r.From.Id, m.Id)
			id = append(id, m.Id)
		}

		if r.From.Id != r.PeerId {
			s.inboxClient.ReadUserMediaUnreadToInbox(ctx, &msgpb.InboxUserReadMediaUnread{
				From: r.From,
				Id:   id,
			})
		}

		if len(r.Id) > 0 {
			ptsCount = int32(len(id))
			pts = int32(idgen.NextNPtsId(ctx, r.From.Id, len(r.Id))) - ptsCount + 1
		} else {
			ptsCount = 0
			pts = int32(idgen.CurrentPtsId(ctx, r.From.Id))
		}
	case model.PEER_CHAT:
		id := make([]int32, 0, len(r.Id))
		for _, m := range r.Id {
			if m.IsMentioned {
				s.MsgCore.UpdateMentioned(ctx, r.From.Id, m.Id)
			} else {
				s.MsgCore.UpdateMediaUnread(ctx, r.From.Id, m.Id)
				id = append(id, m.Id)
			}
		}
		if len(id) > 0 {
			s.inboxClient.ReadChatMediaUnreadToInbox(ctx, &msgpb.InboxChatReadMediaUnread{
				From:       r.From,
				PeerChatId: r.PeerId,
				Id:         id,
			})
		}

		if len(r.Id) > 0 {
			ptsCount = int32(len(id))
			pts = int32(idgen.NextNPtsId(ctx, r.From.Id, len(r.Id))) - ptsCount + 1
		} else {
			ptsCount = 0
			pts = int32(idgen.CurrentPtsId(ctx, r.From.Id))
		}
	case model.PEER_CHANNEL:
		id := make([]int32, 0, len(r.Id))
		for _, m := range r.Id {
			if m.IsMentioned {
				s.MsgCore.UpdateChannelUnreadReadMention(ctx, r.From.Id, r.PeerId, m.Id)
			} else {
				s.MsgCore.UpdateChannelMediaUnread(ctx, r.PeerId, m.Id)
				id = append(id, m.Id)
			}
		}
		if len(id) > 0 {
			sync_client.SyncUpdatesNotMe(ctx,
				r.From.Id,
				r.From.AuthKeyId,
				model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateChannelReadMessagesContents(&mtproto.Update{
					ChannelId: r.PeerId,
					Messages:  id,
				}).To_Update()))
		}
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("DeleteMessages - error: %v", err)
		return
	}

	reply = mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	}).To_Messages_AffectedMessages()

	log.Debugf("ReadMessageContents - reply: %s", reply.DebugString())
	return
}
