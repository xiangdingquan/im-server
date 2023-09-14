package core

import (
	"context"
	"encoding/json"
	"time"

	"open.chat/app/service/biz_service/updates/internal/dao"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	PTS_UPDATE_TYPE_UNKNOWN = 0

	// pts
	PTS_UPDATE_NEW_MESSAGE           = 1
	PTS_UPDATE_DELETE_MESSAGES       = 2
	PTS_UPDATE_READ_HISTORY_OUTBOX   = 3
	PTS_UPDATE_READ_HISTORY_INBOX    = 4
	PTS_UPDATE_WEBPAGE               = 5
	PTS_UPDATE_READ_MESSAGE_CONTENTS = 6
	PTS_UPDATE_EDIT_MESSAGE          = 7

	// qts
	PTS_UPDATE_NEW_ENCRYPTED_MESSAGE = 8

	// channel pts
	PTS_UPDATE_NEW_CHANNEL_MESSAGE     = 9
	PTS_UPDATE_DELETE_CHANNEL_MESSAGES = 10
	PTS_UPDATE_EDIT_CHANNEL_MESSAGE    = 11
	PTS_UPDATE_EDIT_CHANNEL_WEBPAGE    = 12
)

type UpdatesCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *UpdatesCore {
	return &UpdatesCore{
		Dao: dao,
	}
}

func (m *UpdatesCore) GetState(ctx context.Context, authKeyId int64, userId int32) (*mtproto.Updates_State, error) {
	pts := int32(idgen.CurrentPtsId(ctx, userId))
	if pts == 0 {
		pts = int32(idgen.NextPtsId(ctx, userId))
	}

	seq := int32(idgen.CurrentSeqId(ctx, authKeyId))
	if seq == 0 {
		seq = -1
	}
	state := mtproto.MakeTLUpdatesState(&mtproto.Updates_State{
		Pts:         pts,
		Qts:         0,
		Seq:         seq,
		Date:        int32(time.Now().Unix()),
		UnreadCount: 0,
	}).To_Updates_State()

	return state, nil
}

func (m *UpdatesCore) GetDifference(ctx context.Context, authKeyId int64, userId, pts, limit int32) (*mtproto.Updates_Difference, error) {
	var (
		lastPts      = pts
		otherUpdates []*mtproto.Update
		messages     []*mtproto.Message
		userList     []*mtproto.User
		chatList     []*mtproto.Chat
		difference   *mtproto.Updates_Difference
		_            = limit
	)

	var lastSeq int32 = 0
	updateList2, lastSeq := m.GetUpdateListByGtSeq(ctx, authKeyId, userId)
	updateList := m.GetUpdateListByGtPts(ctx, userId, pts)

	if len(updateList)+len(updateList2) > 1000 {
		if len(updateList) > 0 {
			if updateList[len(updateList)-1].Pts_INT32 > lastPts {
				lastPts = updateList[len(updateList)-1].Pts_INT32
			}
		}

		difference = mtproto.MakeTLUpdatesDifferenceTooLong(&mtproto.Updates_Difference{
			Pts: lastPts,
		}).To_Updates_Difference()
	} else {
		if len(updateList) == 0 {
			state := mtproto.MakeTLUpdatesState(&mtproto.Updates_State{
				Pts:         lastPts,
				Date:        int32(time.Now().Unix()),
				UnreadCount: 0,
				Seq:         lastSeq,
			})

			difference2 := mtproto.MakeTLUpdatesDifference(&mtproto.Updates_Difference{
				NewMessages:  []*mtproto.Message{},
				OtherUpdates: updateList2,
				Users:        []*mtproto.User{},
				Chats:        []*mtproto.Chat{},
				State:        state.To_Updates_State(),
			})
			difference = difference2.To_Updates_Difference()
		} else {
			for _, update := range updateList {
				switch update.PredicateName {
				case mtproto.Predicate_updateNewMessage:
					newMessage := update.To_UpdateNewMessage()
					messages = append(messages, model.MessageUpdate(newMessage.GetMessage_MESSAGE()))
				case mtproto.Predicate_updateReadHistoryOutbox:
					readHistoryOutbox := update.To_UpdateReadHistoryOutbox()
					readHistoryOutbox.SetPtsCount(0)
					otherUpdates = append(otherUpdates, readHistoryOutbox.To_Update())
				case mtproto.Predicate_updateReadHistoryInbox:
					readHistoryInbox := update.To_UpdateReadHistoryInbox()
					readHistoryInbox.SetPtsCount(0)
					otherUpdates = append(otherUpdates, readHistoryInbox.To_Update())
				case mtproto.Predicate_updateEditMessage:
					updEditMessage := update.To_UpdateEditMessage()
					updEditMessage.SetPtsCount(0)
					otherUpdates = append(otherUpdates, updEditMessage.To_Update())
				case mtproto.Predicate_updateDeleteMessages:
					updDeleteMessage := update.To_UpdateDeleteMessages()
					updDeleteMessage.SetPtsCount(0)
					otherUpdates = append(otherUpdates, updDeleteMessage.To_Update())
				default:
					continue
				}
				if update.GetPts_INT32() > lastPts {
					lastPts = update.GetPts_INT32()
				}
			}
			if lastPts <= pts {
				lastPts = 0
			}

			state := mtproto.MakeTLUpdatesState(&mtproto.Updates_State{
				Pts:         lastPts,
				Date:        int32(time.Now().Unix()),
				UnreadCount: 0,
				Seq:         lastSeq,
			})

			difference2 := mtproto.MakeTLUpdatesDifference(&mtproto.Updates_Difference{
				NewMessages:  messages,
				OtherUpdates: append(otherUpdates, updateList...),
				Users:        userList,
				Chats:        chatList,
				State:        state.To_Updates_State(),
			})
			difference = difference2.To_Updates_Difference()
		}
	}
	return difference, nil
}

func (m *UpdatesCore) GetUpdateListByGtPts(ctx context.Context, userId, pts int32) []*mtproto.Update {
	doList, _ := m.UserPtsUpdatesDAO.SelectByGtPts(ctx, userId, pts)
	if len(doList) == 0 {
		return []*mtproto.Update{}
	}

	updates := make([]*mtproto.Update, 0, len(doList))
	for _, do := range doList {
		update := &mtproto.Update{}
		err := json.Unmarshal([]byte(do.UpdateData), update)
		if err != nil {
			log.Errorf("unmarshal pts's update(%d)error: %v", do.Id, err)
			continue
		}
		if getUpdateType(update) != do.UpdateType {
			log.Errorf("update data error.")
			continue
		}
		if update.Message_MESSAGE != nil {
			update.Message_MESSAGE = model.MessageUpdate(update.Message_MESSAGE)
		}
		updates = append(updates, update)
	}
	return updates
}

func (m *UpdatesCore) GetUpdateListByGtSeq(ctx context.Context, authKeyId int64, userId int32) ([]*mtproto.Update, int32) {
	doList, _ := m.AuthSeqUpdatesDAO.SelectByGtDate(ctx, authKeyId, userId, int32(time.Now().Unix())-30)
	if len(doList) == 0 {
		return []*mtproto.Update{}, 0
	}

	var lastSeq int32 = 0
	updates := make([]*mtproto.Update, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		log.Debugf("update: %#v", doList[i])
		update := &mtproto.Update{}
		err := json.Unmarshal([]byte(doList[i].UpdateData), update)
		if err != nil {
			log.Errorf("unmarshal seq's update(%v)error: %v", doList[i], err)
			continue
		}
		if update.Message_MESSAGE != nil {
			update.Message_MESSAGE = model.MessageUpdate(update.Message_MESSAGE)
		}
		updates = append(updates, update)
		if doList[i].Seq > lastSeq {
			lastSeq = doList[i].Seq
		}
	}
	return updates, lastSeq
}

func (m *UpdatesCore) GetChannelUpdateListByGtPts(ctx context.Context, channelId, pts int32) []*mtproto.Update {
	doList, _ := m.ChannelPtsUpdatesDAO.SelectByGtPts(ctx, channelId, pts)
	if len(doList) == 0 {
		return []*mtproto.Update{}
	}

	updates := make([]*mtproto.Update, 0, len(doList))
	for _, do := range doList {
		update := &mtproto.Update{}
		err := json.Unmarshal([]byte(do.UpdateData), update)
		if err != nil {
			log.Errorf("unmarshal pts's update(%d)error: %v", do.Id, err)
			continue
		}
		if update.Message_MESSAGE != nil {
			update.Message_MESSAGE = model.MessageUpdate(update.Message_MESSAGE)
		}
		if getUpdateType(update) != do.UpdateType {
			log.Errorf("update data error.")
			continue
		}
		updates = append(updates, update)
	}
	return updates
}

func getUpdateType(update *mtproto.Update) int8 {
	switch update.PredicateName {
	case mtproto.Predicate_updateNewMessage:
		return PTS_UPDATE_NEW_MESSAGE
	case mtproto.Predicate_updateDeleteMessages:
		return PTS_UPDATE_DELETE_MESSAGES
	case mtproto.Predicate_updateReadHistoryOutbox:
		return PTS_UPDATE_READ_HISTORY_OUTBOX
	case mtproto.Predicate_updateReadHistoryInbox:
		return PTS_UPDATE_READ_HISTORY_INBOX
	case mtproto.Predicate_updateWebPage:
		return PTS_UPDATE_WEBPAGE
	case mtproto.Predicate_updateReadMessagesContents:
		return PTS_UPDATE_READ_MESSAGE_CONTENTS
	case mtproto.Predicate_updateEditMessage:
		return PTS_UPDATE_EDIT_MESSAGE
	case mtproto.Predicate_updateNewEncryptedMessage:
		return PTS_UPDATE_NEW_ENCRYPTED_MESSAGE
	case mtproto.Predicate_updateNewChannelMessage:
		return PTS_UPDATE_NEW_CHANNEL_MESSAGE
	case mtproto.Predicate_updateDeleteChannelMessages:
		return PTS_UPDATE_DELETE_CHANNEL_MESSAGES
	case mtproto.Predicate_updateEditChannelMessage:
		return PTS_UPDATE_EDIT_CHANNEL_MESSAGE
	case mtproto.Predicate_updateChannelWebPage:
		return PTS_UPDATE_EDIT_CHANNEL_WEBPAGE
	}
	return PTS_UPDATE_TYPE_UNKNOWN
}
