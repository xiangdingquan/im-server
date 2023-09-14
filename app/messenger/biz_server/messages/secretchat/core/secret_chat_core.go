package core

import (
	"context"
	"encoding/json"
	"open.chat/app/messenger/biz_server/messages/secretchat/dal/dataobject"
	"open.chat/app/messenger/biz_server/messages/secretchat/dao"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

type SecretChatCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *SecretChatCore {
	return &SecretChatCore{dao}
}

func (m *SecretChatCore) GetDifference(ctx context.Context, authKeyId int64, qts int32) (int32, []int32, []*mtproto.EncryptedMessage, error) {
	updDOList, err := m.SecretChatQtsUpdatesDAO.SelectGtQts(ctx, authKeyId, qts)
	if err != nil {
		return -1, nil, nil, err
	}

	if len(updDOList) == 0 {
		return -1, []int32{}, []*mtproto.EncryptedMessage{}, nil
	}

	idList := make([]int64, 0, len(updDOList))
	for i := 0; i < len(updDOList); i++ {
		idList = append(idList, updDOList[i].ChatMessageId)
	}

	msgDOList, err := m.SecretChatMessagesDAO.SelectList(ctx, idList)
	if err != nil {
		return -1, nil, nil, err
	}

	if len(msgDOList) == 0 {
		return -1, []int32{}, []*mtproto.EncryptedMessage{}, nil
	}

	lastQts := int32(-1)
	userIdList := make([]int32, len(updDOList)*2)

	appendIdF := func(idList []int32, id int32) []int32 {
		for _, i := range idList {
			if id != 0 && i == id {
				return idList
			}
		}
		idList = append(idList, id)
		return idList
	}

	newEncryptedMessages := make([]*mtproto.EncryptedMessage, 0, len(updDOList))
	for i := 0; i < len(updDOList); i++ {
		for j := 0; j < len(msgDOList); j++ {
			if updDOList[i].ChatMessageId == msgDOList[j].Id {
				encryptedMessage := &mtproto.EncryptedMessage{}
				err := json.Unmarshal(hack.Bytes(msgDOList[j].MessageData), encryptedMessage)
				if err != nil {
					log.Errorf("json unmarsh error - %v", err)
					continue
				}
				newEncryptedMessages = append(newEncryptedMessages, encryptedMessage)
				userIdList = appendIdF(userIdList, updDOList[i].UserId)
				userIdList = appendIdF(userIdList, msgDOList[j].SenderUserId)
				break
			}
		}
		lastQts = updDOList[i].Qts
	}

	return lastQts, userIdList, newEncryptedMessages, nil
}

func (m *SecretChatCore) GetRequested(ctx context.Context, selfId int32) ([]*SecretChatData, error) {
	l, err := m.SelectRequested(ctx, selfId)
	if err != nil {
		return nil, err
	}
	out := make([]*SecretChatData, len(l))
	for i, do := range l {
		out[i] = &SecretChatData{
			SecretChatsDO:  do,
			SecretChatCore: m,
		}
	}
	return out, nil
}

func (m *SecretChatCore) AddClosedRequest(ctx context.Context, secretChatId, selfId, participantId int32) error {
	do, err := m.SecretChatsCloseRequestsDAO.SelectByUserAndChat(ctx, selfId, secretChatId)
	if err != nil {
		return err
	}

	if do != nil {
		_, err = m.SecretChatsCloseRequestsDAO.UpdateClosed(ctx, selfId, secretChatId)
		return err
	}

	do = &dataobject.SecretChatCloseRequestsDo{
		SecretChatId: secretChatId,
		FromUID:      selfId,
		ToUID:        participantId,
	}
	_, _, err = m.SecretChatsCloseRequestsDAO.Insert(ctx, do)
	return err
}

func (m *SecretChatCore) GetPendingClosed(ctx context.Context, uid int32) ([]int32, error) {
	requests, err := m.SecretChatsCloseRequestsDAO.SelectByUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	reqIds := make([]int32, len(requests))
	for i, v := range requests {
		reqIds[i] = v.SecretChatId
	}

	return reqIds, nil
}
