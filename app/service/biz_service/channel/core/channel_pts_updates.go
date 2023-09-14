package core

import (
	"context"
	"encoding/json"
	"time"

	"open.chat/app/service/biz_service/channel/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

const (
	ptsUnknown              = 0
	ptsNewChannelMessage    = 9
	ptsDeleteChannelMessage = 10
	ptsEditChannelMessage   = 11
	ptsEditChannelWebpage   = 12
)

func getUpdateType(update *mtproto.Update) int8 {
	switch update.PredicateName {
	case mtproto.Predicate_updateNewChannelMessage:
		return ptsNewChannelMessage
	case mtproto.Predicate_updateDeleteChannelMessages:
		return ptsDeleteChannelMessage
	case mtproto.Predicate_updateEditChannelMessage:
		return ptsEditChannelMessage
	case mtproto.Predicate_updateChannelWebPage:
		return ptsEditChannelWebpage
	}
	return ptsUnknown
}

func (m *ChannelCore) GetChannelUpdateListByGtPts(ctx context.Context, channelId, userId, pts, limit int32) []*mtproto.Update {
	doList, _ := m.ChannelPtsUpdatesDAO.SelectByGtPts(ctx, channelId, pts, limit)
	if len(doList) == 0 {
		return []*mtproto.Update{}
	}

	newMessageIdList := make([]int32, 0, len(doList))
	updates := make([]*mtproto.Update, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		switch doList[i].UpdateType {
		case ptsNewChannelMessage:
			updates = append(updates, mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
				Pts_INT32: doList[i].Pts,
				PtsCount:  doList[i].PtsCount,
				Message_MESSAGE: mtproto.MakeTLMessageEmpty(&mtproto.Message{
					Id: doList[i].NewMessageId,
				}).To_Message(),
			}).To_Update())
			newMessageIdList = append(newMessageIdList, doList[i].NewMessageId)
		default:
			update := &mtproto.Update{}
			err := json.Unmarshal([]byte(doList[i].UpdateData), update)
			if err != nil {
				log.Errorf("unmarshal pts's update(%d) error: %v", doList[i].Id, err)
				continue
			}
			if update.Message_MESSAGE != nil {
				update.Message_MESSAGE = model.MessageUpdate(update.Message_MESSAGE)
			}
			if getUpdateType(update) != doList[i].UpdateType {
				log.Errorf("update data error.")
				continue
			}
			updates = append(updates, update)
		}
	}

	if len(newMessageIdList) > 0 {
		messageList, _ := m.GetMessageListByIdList(ctx, userId, channelId, newMessageIdList)
		for _, upd := range updates {
			if upd.PredicateName == mtproto.Predicate_updateNewChannelMessage {
				for _, m2 := range messageList {
					if m2.Id == upd.Message_MESSAGE.Id {
						upd.Message_MESSAGE = model.MessageUpdate(m2)
						break
					}
				}
			}
		}
	}

	return updates
}

func (m *ChannelCore) AddOtherUpdateToPtsQueue(ctx context.Context, channelId, pts, ptsCount int32, update *mtproto.Update) error {
	updateData, err := json.Marshal(update)

	if err == nil {
		do := &dataobject.ChannelPtsUpdatesDO{
			ChannelId:  channelId,
			Pts:        pts,
			PtsCount:   ptsCount,
			UpdateType: getUpdateType(update),
			UpdateData: string(updateData),
			Date2:      int32(time.Now().Unix()),
		}
		_, _, err = m.ChannelPtsUpdatesDAO.Insert(ctx, do)
	}
	return err
}

func (m *ChannelCore) AddMessageToPtsQueue(tx *sqlx.Tx, channelId, pts, ptsCount, newMessageId int32) error {
	do := &dataobject.ChannelPtsUpdatesDO{
		ChannelId:    channelId,
		Pts:          pts,
		PtsCount:     ptsCount,
		UpdateType:   ptsNewChannelMessage,
		NewMessageId: newMessageId,
		Date2:        int32(time.Now().Unix()),
	}
	_, _, err := m.ChannelPtsUpdatesDAO.InsertTx(tx, do)
	return err
}

func (m *ChannelCore) GetDifference(ctx context.Context, userId, date int32) (int32, []*mtproto.Update, error) {
	channelIdList, err := m.ChannelParticipantsDAO.SelectByGTOffsetDate2(ctx, userId, date)
	if err != nil {
		return 0, nil, err
	}

	if len(channelIdList) == 0 {
		return date, []*mtproto.Update{}, nil
	}

	doList, _ := m.ChannelPtsUpdatesDAO.SelectByGtDate2(ctx, channelIdList, date-1)
	if len(doList) == 0 {
		return date, []*mtproto.Update{}, nil
	} else if len(doList) > 100 {
		//
	}

	lastDate := date

	newMessagesIdList := make(map[int32][]int32)
	updates := make([]*mtproto.Update, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		cpu := &doList[i]
		lastDate = cpu.Date2
		switch cpu.UpdateType {
		case ptsNewChannelMessage:
			updates = append(updates, mtproto.MakeTLUpdateNewChannelMessage(&mtproto.Update{
				Pts_INT32: cpu.Pts,
				PtsCount:  cpu.PtsCount,
				Message_MESSAGE: mtproto.MakeTLMessageEmpty(&mtproto.Message{
					Id: cpu.NewMessageId,
				}).To_Message(),
			}).To_Update())

			if v, ok := newMessagesIdList[cpu.ChannelId]; ok {
				newMessagesIdList[cpu.ChannelId] = append(v, cpu.NewMessageId)
			} else {
				newMessagesIdList[cpu.ChannelId] = []int32{cpu.NewMessageId}
			}
		default:
			update := &mtproto.Update{}
			err := json.Unmarshal([]byte(cpu.UpdateData), update)
			if err != nil {
				log.Errorf("unmarshal pts's update(%d) error: %v", cpu.Id, err)
				continue
			}
			if update.Message_MESSAGE != nil {
				update.Message_MESSAGE = model.MessageUpdate(update.Message_MESSAGE)
			}
			if getUpdateType(update) != cpu.UpdateType {
				log.Errorf("update data error.")
				continue
			}
			updates = append(updates, update)
		}
	}

	for k, v := range newMessagesIdList {
		if len(v) == 0 {
			continue
		}

		channel, err := m.GetMutableChannel(ctx, k, userId)
		if err != nil {
			log.Errorf("m.GetMutableChannel - error: %v", err)
			continue
		}

		participant := channel.GetImmutableChannelParticipant(userId)
		if participant != nil && participant.IsKicked() {
			log.Errorf("channel.GetImmutableChannelParticipant - error: invalid participant(%d)", userId)
			continue
		}

		messageList, _ := m.GetMessageListByIdList(ctx, userId, k, v)
		for _, upd := range updates {
			if upd.PredicateName == mtproto.Predicate_updateNewChannelMessage {
				for _, m2 := range messageList {
					if m2.Id == upd.Message_MESSAGE.Id {
						newMessage := model.MessageUpdate(m2)
						if newMessage.PredicateName == mtproto.Predicate_message {
							newMessage.Mentioned = model.CheckHasMention(newMessage.Entities, userId)
						}
						if newMessage.GetId() > participant.ReadInboxMaxId {
							newMessage.MediaUnread = true
						}
						upd.Message_MESSAGE = newMessage
						break
					}
				}
			}
		}
	}

	return lastDate, updates, nil
}
