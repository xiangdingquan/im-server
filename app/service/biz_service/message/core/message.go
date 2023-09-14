package core

import (
	"context"
	"fmt"
	"math"

	"open.chat/app/service/biz_service/message/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (m *MessageCore) makeMessageBox(ctx context.Context, selfUserId int32, do *dataobject.MessagesDO) (box *model.MessageBox) {
	box = &model.MessageBox{
		SelfUserId:        do.UserId,
		SendUserId:        do.SenderUserId,
		MessageId:         do.UserMessageBoxId,
		DialogId:          do.DialogId,
		DialogMessageId:   do.DialogMessageId,
		MessageDataId:     do.MessageDataId,
		RandomId:          do.RandomId,
		Pts:               do.Pts,
		PtsCount:          do.PtsCount,
		MessageFilterType: 0,
		MessageBoxType:    do.MessageBoxType,
		MessageType:       do.MessageType,
		Message:           nil,
		TtlSeconds:        do.TtlSeconds,
	}
	box.Message, _ = model.DecodeMessage(int(do.MessageType), hack.Bytes(do.MessageData))
	box.Message.MediaUnread = util.Int8ToBool(do.MediaUnread)
	pollId, _ := model.GetPollIdByMessage(box.Message.Media)
	if pollId != 0 {
		poll, _ := m.PollFacade.GetMediaPoll(ctx, selfUserId, pollId)
		if poll != nil {
			box.Message.Media = poll.ToMessageMedia()
		}
	}
	return
}

func (m *MessageCore) GetMessageBox(ctx context.Context, peerType, ownerId, messageId int32) (*model.MessageBox, error) {
	var (
		mBox *model.MessageBox
	)

	switch peerType {
	case model.PEER_USER, model.PEER_CHAT:
		boxDO, err := m.MessagesDAO.SelectByMessageId(ctx, ownerId, messageId)
		if err != nil {
			return nil, err
		}
		if boxDO == nil {
			return nil, fmt.Errorf("query is empty: getMessageBox(%d, %d, %d)", peerType, ownerId, messageId)
		}

		mBox = m.makeMessageBox(ctx, ownerId, boxDO)
	case model.PEER_CHANNEL:
		log.Warn("blocked, License key from https://dong.chat required to unlock enterprise features.")
	default:
		return nil, fmt.Errorf("invalid peer_id")
	}

	return mBox, nil
}

func (m *MessageCore) GetPeerUserMessageId(ctx context.Context, userId, messageId, peerUserId int32) int32 {
	pDO, _ := m.MessagesDAO.SelectPeerUserMessage(ctx, peerUserId, userId, messageId)
	if pDO == nil {
		return 0
	}

	return pDO.UserMessageBoxId
}

func (m *MessageCore) GetPeerUserMessage(ctx context.Context, userId, messageId, peerUserId int32) (*model.MessageBox, error) {
	pDO, err := m.MessagesDAO.SelectPeerUserMessage(ctx, peerUserId, userId, messageId)
	if err != nil {
		return nil, err
	} else if pDO == nil {
		return nil, mtproto.ErrMsgIdInvalid
	}
	return m.makeMessageBox(ctx, peerUserId, pDO), nil
}

func (m *MessageCore) GetPeerChatMessageList(ctx context.Context, userId, messageId, peerChatId int32) (peerMsgs map[int32]*model.MessageBox) {
	peerMsgs = make(map[int32]*model.MessageBox)
	myDO, err := m.MessagesDAO.SelectByMessageId(ctx, userId, messageId)
	if err != nil || myDO == nil {
		return
	}
	pDOList, err := m.MessagesDAO.SelectByMessageDataIdList(ctx, []int64{myDO.MessageDataId})
	if err != nil || len(pDOList) == 0 {
		return
	}

	for i := 0; i < len(pDOList); i++ {
		peerMsgs[pDOList[i].UserId] = m.makeMessageBox(ctx, pDOList[i].UserId, &pDOList[i])
	}
	return
}

func (m *MessageCore) GetUserMessage(ctx context.Context, userId int32, id int32) (*model.MessageBox, error) {
	myDO, err := m.MessagesDAO.SelectByMessageId(ctx, userId, id)
	if err != nil {
		return nil, err
	} else if myDO == nil {
		return nil, mtproto.ErrMessageIdInvalid
	}
	return m.makeMessageBox(ctx, userId, myDO), nil
}

func (m *MessageCore) GetUserMessageList(ctx context.Context, userId int32, idList []int32) (boxList []*model.MessageBox) {
	boxList = make([]*model.MessageBox, 0, len(idList))
	if len(idList) == 0 {
		return
	}

	doList, err := m.MessagesDAO.SelectByMessageIdList(ctx, userId, idList)
	if err != nil {
		log.Errorf("getUserMessageList - error: %v", err)
	} else {
		for i := 0; i < len(doList); i++ {
			boxList = append(boxList, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) GetUserMessageListByDataIdList(ctx context.Context, userId int32, idList []int64) (boxList []*model.MessageBox) {
	boxList = make([]*model.MessageBox, 0, len(idList))
	if len(idList) == 0 {
		return
	}

	doList, err := m.MessagesDAO.SelectByMessageDataIdList(ctx, idList)
	if err != nil {
		log.Errorf("getUserMessageListByDataIdList - error: %v", err)
	} else {
		for i := 0; i < len(doList); i++ {
			boxList = append(boxList, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) UpdateMediaUnread(ctx context.Context, userId int32, id int32) {
	m.MessagesDAO.UpdateMediaUnread(ctx, userId, id)
}

func (m *MessageCore) SearchByMediaType(ctx context.Context, userId int32, peer *model.PeerUtil, mediaType model.MediaType, minId, offset, limit int32) (boxList []*model.MessageBox) {
	if mediaType == model.MEDIA_PHONE_CALL {
		return m.SearchByPhoneCall(ctx, userId, offset, limit)
	} else {
		return m.searchByMediaType(ctx, userId, peer, mediaType, minId, offset, limit)
	}
}

func (m *MessageCore) searchByMediaType(ctx context.Context, userId int32, peer *model.PeerUtil, mediaType model.MediaType, minId, offset, limit int32) (boxList []*model.MessageBox) {
	if peer.PeerType == model.PEER_CHANNEL {
		doList, err := m.ChannelMessagesDAO.SelectByMediaType(ctx, userId, peer.PeerId, minId, int32(mediaType), offset, limit)
		if err != nil {
			log.Errorf("searchByMediaType - error: %v", err)
		} else {
			for i := 0; i < len(doList); i++ {
				boxList = append(boxList, m.makeChannelMessageBox(ctx, userId, &doList[i]))
			}
		}
	} else {
		doList, err := m.MessagesDAO.SelectByMediaType(ctx, userId, model.MakeDialogId(userId, peer.PeerType, peer.PeerId), int32(mediaType), offset, limit)
		if err != nil {
			log.Errorf("searchByMediaType - error: %v", err)
		} else {
			for i := 0; i < len(doList); i++ {
				boxList = append(boxList, m.makeMessageBox(ctx, userId, &doList[i]))
			}
		}
	}
	return
}

func (m *MessageCore) SearchByPhoneCall(ctx context.Context, userId int32, offset, limit int32) (boxList []*model.MessageBox) {
	doList, err := m.MessagesDAO.SelectPhoneCallList(ctx, userId, int32(model.MEDIA_PHONE_CALL), offset, limit)
	if err != nil {
		log.Errorf("searchByPhoneCall - error: %v", err)
	} else {
		for i := 0; i < len(doList); i++ {
			boxList = append(boxList, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) Search(ctx context.Context, userId int32, peer *model.PeerUtil, q string, minId, offset, limit int32) (messages model.MessageBoxList) {
	if offset == 0 {
		offset = math.MaxInt32
	}
	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		doList, _ := m.MessagesDAO.Search(ctx, userId, model.MakeDialogId(userId, peer.PeerType, peer.PeerId), "%"+q+"%", offset, limit)
		log.Debugf("Search - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeMessageBox(ctx, userId, &doList[i]))
		}
	case model.PEER_CHANNEL:
		doList, _ := m.ChannelMessagesDAO.Search(ctx, userId, peer.PeerId, minId, "%"+q+"%", offset, limit)
		log.Debugf("Search - %#v", doList)
		messages = make([]*model.MessageBox, 0, len(doList))
		for i := 0; i < len(doList); i++ {
			messages = append(messages, m.makeChannelMessageBox(ctx, userId, &doList[i]))
		}
	}
	return
}

func (m *MessageCore) SearchGlobal(ctx context.Context, userId int32, q string, offset, limit int32) (messages model.MessageBoxList) {
	if offset == 0 {
		offset = math.MaxInt32
	}
	doList, _ := m.MessagesDAO.SearchGlobal(ctx, userId, "%"+q+"%", offset, limit)
	log.Debugf("Search - %#v", doList)
	messages = make([]*model.MessageBox, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		messages = append(messages, m.makeMessageBox(ctx, userId, &doList[i]))
	}

	return
}
