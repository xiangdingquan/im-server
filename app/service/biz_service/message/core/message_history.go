package core

import (
	"context"
	"time"

	"open.chat/model"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (m *MessageCore) GetOffsetIdBackwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit, hash int32) (messages model.MessageBoxList) {
	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		doList, _ := m.MessagesDAO.SelectBackwardByOffsetIdLimit(ctx, userId, model.MakeDialogId(userId, peer.PeerType, peer.PeerId), offsetId, limit)
		log.Debugf("GetOffsetIdBackwardHistoryMessages - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	case model.PEER_CHANNEL:
		doList, _ := m.ChannelMessagesDAO.SelectBackwardByOffsetIdLimit(ctx, userId, peer.PeerId, minId, offsetId, limit)
		log.Debugf("GetOffsetIdBackwardHistoryMessages - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) GetOffsetIdForwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit, hash int32) (messages model.MessageBoxList) {
	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		doList, _ := m.MessagesDAO.SelectForwardByOffsetIdLimit(ctx, userId, model.MakeDialogId(userId, peer.PeerType, peer.PeerId), offsetId, limit)
		log.Debugf("GetOffsetIdForwardHistoryMessages - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	case model.PEER_CHANNEL:
		doList, _ := m.ChannelMessagesDAO.SelectForwardByOffsetIdLimit(ctx, userId, peer.PeerId, minId, offsetId, limit)
		log.Debugf("GetOffsetIdForwardHistoryMessages - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) GetOffsetDateBackwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetDate, minId, maxId, limit, hash int32) (messages model.MessageBoxList) {
	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		doList, _ := m.MessagesDAO.SelectBackwardByOffsetDateLimit(ctx, userId, model.MakeDialogId(userId, peer.PeerType, peer.PeerId), offsetDate, limit)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	case model.PEER_CHANNEL:
		doList, _ := m.ChannelMessagesDAO.SelectBackwardByOffsetDateLimit(ctx, userId, peer.PeerId, minId, offsetDate, limit)
		log.Debugf("GetOffsetIdForwardHistoryMessages - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) GetOffsetDateForwardHistoryMessages(ctx context.Context, userId int32, peer *model.PeerUtil, offsetDate, minId, maxId, limit, hash int32) (messages model.MessageBoxList) {
	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		doList, _ := m.MessagesDAO.SelectForwardByOffsetDateLimit(ctx, userId, model.MakeDialogId(userId, peer.PeerType, peer.PeerId), offsetDate, limit)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	case model.PEER_CHANNEL:
		doList, _ := m.ChannelMessagesDAO.SelectForwardByOffsetDateLimit(ctx, userId, peer.PeerId, minId, offsetDate, limit)
		log.Debugf("GetOffsetIdForwardHistoryMessages - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) GetMessageListByIdList(ctx context.Context, userId int32, peer *model.PeerUtil, idList []int32) (messages model.MessageBoxList) {
	messages = make([]*model.MessageBox, 0, len(idList))
	if len(idList) == 0 {
		return
	}

	switch peer.PeerType {
	case model.PEER_CHANNEL:
		doList, _ := m.ChannelMessagesDAO.SelectByMessageIdList(ctx, userId, peer.PeerId, idList)
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	default:
		doList, _ := m.MessagesDAO.SelectByMessageIdList(ctx, userId, idList)
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) GetMessageListByDataIdList(ctx context.Context, userId int32, peerType int32, dataIdList []int64) (messages model.MessageBoxList) {
	messages = make([]*model.MessageBox, 0, len(dataIdList))
	if len(dataIdList) == 0 {
		return
	}

	switch peerType {
	case model.PEER_USER_MESSAGE:
		doList, _ := m.MessagesDAO.SelectByMessageDataIdList(ctx, dataIdList)
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	case model.PEER_CHANNEL_MESSAGE:
		doList, _ := m.ChannelMessagesDAO.SelectByMessageDataIdList(ctx, dataIdList)
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	default:
		log.Errorf("invalid peer: %v", peerType)
	}
	return
}

func (m *MessageCore) GetHistoryMessagesCount(ctx context.Context, userId int32, peer *model.PeerUtil) int32 {
	var count int

	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		count = m.CommonDAO.CalcSize(ctx, "messages", map[string]interface{}{
			"user_id":   userId,
			"dialog_id": model.MakeDialogId(userId, peer.PeerType, peer.PeerId),
			"deleted":   0,
		})
	case model.PEER_CHANNEL:
		count = m.CommonDAO.CalcSize(ctx, "channel_messages", map[string]interface{}{
			"channel_id": peer.PeerId,
			"deleted":    0,
		})
	default:
		log.Errorf("invalid peer: (%d, %v)", userId, peer)
	}

	return int32(count)
}

func (m *MessageCore) GetOffsetIdBackwardUnreadMentions(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit int32) (messages model.MessageBoxList) {
	switch peer.PeerType {
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
		doList, _ := m.ChannelMessagesDAO.SelectBackwardUnreadMentions(ctx, peer.PeerId, userId, offsetId, limit)
		log.Debugf("GetOffsetIdBackwardUnreadMentions - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) GetOffsetIdForwardUnreadMentions(ctx context.Context, userId int32, peer *model.PeerUtil, offsetId, minId, maxId, limit int32) (messages model.MessageBoxList) {
	switch peer.PeerType {
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
		doList, _ := m.ChannelMessagesDAO.SelectForwardUnreadMentions(ctx, peer.PeerId, userId, offsetId, limit)
		log.Debugf("GetOffsetIdForwardHUnreadMentions - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) ReadEphemeralMsgByBetween(ctx context.Context, userId int32, peer *model.PeerUtil, minId, maxId int32) bool {
	log.Warnf("MessageCore.ReadEphemeralMsgByBetween %d,%d,%d,%d", userId, peer.PeerId, minId, maxId)
	messages := make([]*model.MessageBox, 0)
	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		doList, _ := m.MessagesDAO.SelectEphemeralByBetween(ctx, userId, model.MakeDialogId(userId, peer.PeerType, peer.PeerId), minId, maxId)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	case model.PEER_CHANNEL:
		doList, _ := m.ChannelMessagesDAO.SelectEphemeralByBetween(ctx, userId, peer.PeerId, minId, maxId)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	}

	log.Warnf("MessageCore.ReadEphemeralMsgByBetween %s", logger.JsonDebugData(messages))
	now := time.Now().Unix()
	for _, msg := range messages {
		if msg.TtlSeconds > 0 {
			if !m.Redis.AddCountDownMsg(ctx, userId, peer.PeerType, peer.PeerId, msg.MessageId, now+int64(msg.TtlSeconds)) {
				log.Error("AddCountDownMsg - fail")
			}
		}
	}
	return true
}

func (m *MessageCore) GetEphemeralExpireList(ctx context.Context) []*model.CountDownMessage {
	log.Warnf("MessageCore.GetEphemeralExpireList")
	now := time.Now().Unix()
	messages := m.Redis.GetCountDownMsg(ctx, now)
	log.Warnf("msgs %s", logger.JsonDebugData(messages))
	return messages
}

func (m *MessageCore) DelEphemeralList(ctx context.Context, messages []*model.CountDownMessage) bool {
	log.Warnf("MessageCore.DelEphemeralList %s", logger.JsonDebugData(messages))
	return m.Redis.DelCountDownMsgs(ctx, messages)
}
