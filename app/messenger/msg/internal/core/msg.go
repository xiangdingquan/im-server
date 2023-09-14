package core

import (
	"context"

	"open.chat/model"
)

func (m *MsgCore) GetPeerMessageId(ctx context.Context, userId, messageId, peerId int32) int32 {
	do, _ := m.MessagesDAO.SelectPeerMessageId(ctx, peerId, userId, messageId)
	if do == nil {
		return 0
	} else {
		return do.UserMessageBoxId
	}
}

func (m *MsgCore) DeleteByMessageIdList(ctx context.Context, userId int32, idList []int32) (rowsAffected int64, err error) {
	if len(idList) == 0 {
		return 0, nil
	}
	return m.MessagesDAO.DeleteMessagesByMessageIdList(ctx, userId, idList)
}

func (m *MsgCore) GetPeerDialogMessageIdList(ctx context.Context, userId int32, idList []int32) map[int32][]int32 {
	doList, _ := m.MessagesDAO.SelectPeerDialogMessageIdList(ctx, userId, idList)
	peerMessageIdListMap := make(map[int32][]int32)

	for _, do := range doList {
		if messageIdList, ok := peerMessageIdListMap[do.UserId]; !ok {
			peerMessageIdListMap[do.UserId] = []int32{do.UserMessageBoxId}
		} else {
			peerMessageIdListMap[do.UserId] = append(messageIdList, do.UserMessageBoxId)
		}
	}

	return peerMessageIdListMap
}

func (m *MsgCore) GetMessageIdListByDialog(ctx context.Context, userId int32, peer *model.PeerUtil) []int32 {
	doList, _ := m.MessagesDAO.SelectDialogMessageIdList(ctx, userId, model.MakeDialogId(userId, peer.PeerType, peer.PeerId))
	idList := make([]int32, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		idList = append(idList, doList[i].UserMessageBoxId)
	}
	return idList
}

func (m *MsgCore) GetLastMessageAndIdListByDialog(ctx context.Context, userId int32, peer *model.PeerUtil) (lastMessage *model.MessageBox, idList []int32) {
	doList, _ := m.MessagesDAO.SelectDialogMessageIdList(ctx, userId, model.MakeDialogId(userId, peer.PeerType, peer.PeerId))
	idList = make([]int32, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		if i == 0 {
			lastMessage = makeMessageBox(&doList[i])
		}
		idList = append(idList, doList[i].UserMessageBoxId)
	}
	return
}
