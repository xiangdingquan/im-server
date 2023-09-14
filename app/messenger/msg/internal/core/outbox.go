package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	"math"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func makeMessageBoxByDO(boxDO *dataobject.MessagesDO) *model.MessageBox {
	message, _ := model.DecodeMessage(int(boxDO.MessageType), hack.Bytes(boxDO.MessageData))
	return &model.MessageBox{
		SelfUserId:        boxDO.UserId,
		SendUserId:        boxDO.SenderUserId,
		MessageId:         boxDO.UserMessageBoxId,
		DialogId:          boxDO.DialogId,
		DialogMessageId:   boxDO.DialogMessageId,
		MessageDataId:     boxDO.MessageDataId,
		RandomId:          boxDO.RandomId,
		Pts:               0,
		PtsCount:          0,
		MessageFilterType: int8(boxDO.MessageDataType),
		MessageBoxType:    boxDO.MessageBoxType,
		MessageType:       boxDO.MessageType,
		Message:           message,
		Views:             0,
		ReplyOwnerId:      0,
		TtlSeconds:        boxDO.TtlSeconds,
	}
}

func (m *MsgCore) sendMessageToOutbox(ctx context.Context, fromId int32, peer *model.PeerUtil, clientRandomId int64, message *mtproto.Message) (*model.MessageBox, error) {
	var (
		dialogId        = model.MakeDialogId(fromId, peer.PeerType, peer.PeerId)
		dialogMessageId = int32(idgen.NextMessageDataId(ctx, dialogId))
		outBoxMsgId     = int32(idgen.NextMessageBoxId(ctx, fromId))
		err             error
	)

	if dialogMessageId == 0 || outBoxMsgId == 0 {
		err = mtproto.ErrInternelServerError
		return nil, err
	}

	message.Out = true
	message.PeerId = peer.ToPeer()
	message.Id = outBoxMsgId
	message.MediaUnread = model.CheckHasMediaUnread(message)
	messageBoxType := int8(model.MESSAGE_BOX_TYPE_OUTGOING)
	if peer.PeerType == model.PEER_USER && fromId == peer.PeerId {
		message.Out = false
		messageBoxType = model.MESSAGE_BOX_TYPE_INCOMING
	}
	mType, mData := model.EncodeMessage(message)

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		outMsgBox := &model.MessageBox{
			SelfUserId:        fromId,
			SendUserId:        fromId,
			MessageId:         outBoxMsgId,
			DialogId:          dialogId,
			DialogMessageId:   dialogMessageId,
			MessageDataId:     idgen.GetUUID(),
			RandomId:          clientRandomId,
			Pts:               0,
			PtsCount:          0,
			MessageFilterType: int8(model.GetMediaType(message)),
			MessageBoxType:    messageBoxType,
			MessageType:       int8(mType),
			Message:           message,
			TtlSeconds:        message.GetTtlSeconds().GetValue(),
		}

		outBoxDO := &dataobject.MessagesDO{
			UserId:           outMsgBox.SelfUserId,
			UserMessageBoxId: outMsgBox.MessageId,
			DialogId:         outMsgBox.DialogId,
			DialogMessageId:  outMsgBox.DialogMessageId,
			SenderUserId:     outMsgBox.SelfUserId,
			PeerType:         int8(peer.PeerType),
			PeerId:           peer.PeerId,
			RandomId:         outMsgBox.RandomId,
			MessageType:      outMsgBox.MessageType,
			MessageData:      hack.String(mData),
			MessageDataId:    outMsgBox.MessageDataId,
			MessageDataType:  int32(outMsgBox.MessageFilterType),
			Message:          message.Message,
			Pts:              0,
			PtsCount:         0,
			MessageBoxType:   outMsgBox.MessageBoxType,
			ReplyToMsgId:     0,
			Mentioned:        0,
			MediaUnread:      util.BoolToInt8(message.MediaUnread),
			HasMediaUnread:   0,
			TtlSeconds:       outMsgBox.TtlSeconds,
			Date2:            int32(time.Now().Unix()),
			Deleted:          0,
		}

		lastInsertId, rowsAffected, err := m.MessagesDAO.InsertOrReturnIdTx(tx, outBoxDO)
		if err != nil {
			result.Err = err
			return
		}

		if rowsAffected == 0 {
			if lastInsertId > 0 {
				result.Data = lastInsertId
				return
			} else {
				result.Err = errors.New("insert error")
				return
			}
		}

		outMsgBox.Pts = int32(idgen.NextPtsId(ctx, fromId))
		outMsgBox.PtsCount = 1
		log.Debugf("sendMessage - (pts: %d, pts_count: %d)", outMsgBox.Pts, outMsgBox.PtsCount)

		switch peer.PeerType {
		case model.PEER_USER:
			dialogDO := &dataobject.ConversationsDO{
				UserId:     fromId,
				PeerId:     peer.PeerId,
				TopMessage: int32(outBoxMsgId),
				Date2:      int32(time.Now().Unix()),
			}
			if dialogMessageId > 1 {
				cMap := map[string]interface{}{
					"top_message": dialogDO.TopMessage,
					"date2":       dialogDO.Date2,
					"unread_mark": 0,
				}

				if true {
					cMap["draft_message_data"] = ""
					cMap["draft_type"] = 0
				}

				var rowsAffected int64
				rowsAffected, result.Err = m.ConversationsDAO.UpdateOutboxDialogTx(tx, cMap, fromId, peer.PeerId)
				if result.Err != nil {
					return
				}
				if rowsAffected == 0 {
					_, _, result.Err = m.ConversationsDAO.InsertIgnoreTx(tx, dialogDO)
				}
			} else {
				_, _, result.Err = m.ConversationsDAO.InsertIgnoreTx(tx, dialogDO)
			}

			result.Data = outMsgBox
			return
		case model.PEER_CHAT:
			cMap := map[string]interface{}{
				"top_message": int32(outBoxMsgId),
				"date2":       int32(time.Now().Unix()),
				"unread_mark": 0,
			}
			if true {
				cMap["draft_message_data"] = ""
				cMap["draft_type"] = 0
			}
			_, result.Err = m.ChatParticipantsDAO.UpdateOutboxDialogTx(tx, cMap, fromId, peer.PeerId)
			if result.Err != nil {
				log.Errorf("%v", result)
				return
			}
			result.Data = outMsgBox
			return
		default:
			result.Err = fmt.Errorf("fatal error - invalid peer_type: %v", peer)
			log.Errorf("%v", result)
		}
	})

	if tR.Err != nil {
		return nil, tR.Err
	}

	var outBox *model.MessageBox

	switch tR.Data.(type) {
	case *model.MessageBox:
		outBox = tR.Data.(*model.MessageBox)

	case int64:
		if tR.Data.(int64) <= 0 {
			log.Error("unknown error")
			return nil, errors.New("fatal unknown error")
		}

		do, err := m.MessagesDAO.SelectByRandomId(ctx, fromId, clientRandomId)
		if err != nil {
			return nil, err
		}
		if do != nil {
			outBox = makeMessageBoxByDO(do)
			outBox.Pts = int32(idgen.CurrentPtsId(ctx, outBox.SelfUserId))
			outBox.PtsCount = 1
		} else {
			log.Error("unknown error")
			return nil, errors.New("fatal unknown error")
		}
	default:
		log.Error("unknown error")
		return nil, errors.New("fatal unknown error")
	}

	return outBox, nil
}

func (m *MsgCore) SendUserMessage(ctx context.Context, fromId int32, toId int32, clientRandomId int64, message *mtproto.Message) (*model.MessageBox, error) {
	peer := model.MakeUserPeerUtil(toId)
	return m.sendMessageToOutbox(ctx, fromId, peer, clientRandomId, message)
}

func (m *MsgCore) SendUserMultiMessage(ctx context.Context, fromId int32, toId int32, randomIdList []int64, messageList []*mtproto.Message) ([]*model.MessageBox, error) {
	var (
		outBoxList []*model.MessageBox
	)

	for i, msg := range messageList {
		peer := model.MakeUserPeerUtil(toId)
		outBox, _ := m.sendMessageToOutbox(ctx, fromId, peer, randomIdList[i], msg)
		outBoxList = append(outBoxList, outBox)
	}

	return outBoxList, nil
}

func (m *MsgCore) SendChatMessage(ctx context.Context, fromId int32, chatId int32, clientRandomId int64, message *mtproto.Message) (*model.MessageBox, error) {
	peer := model.MakeChatPeerUtil(chatId)
	return m.sendMessageToOutbox(ctx, fromId, peer, clientRandomId, message)
}

func (m *MsgCore) SendChatMultiMessage(ctx context.Context, fromId int32, chatId int32, randomIdList []int64, messageList []*mtproto.Message) ([]*model.MessageBox, error) {
	var (
		outBoxList []*model.MessageBox
	)

	for i, msg := range messageList {
		peer := model.MakeChatPeerUtil(chatId)
		outBox, _ := m.sendMessageToOutbox(ctx, fromId, peer, randomIdList[i], msg)
		outBoxList = append(outBoxList, outBox)
	}

	return outBoxList, nil
}

func (m *MsgCore) DeleteMessages(ctx context.Context, userId int32, msgIds []int32) ([]int64, error) {
	msgDOList, err := m.MessagesDAO.SelectByMessageIdList(ctx, userId, msgIds)
	if err != nil {
		return nil, err
	}

	if len(msgDOList) == 0 {
		return []int64{}, nil
	} else {

	}

	var (
		topMessageIndex      int32
		dialogId             int64
		deletedMsgDataIdList = make([]int64, 0, len(msgDOList))
	)

	for i := 0; i < len(msgDOList); i++ {
		if dialogId == 0 {
			dialogId = msgDOList[i].DialogId
		}

		if dialogId != msgDOList[i].DialogId {
			err = mtproto.ErrMessageIdInvalid
			return nil, err
		}
	}

	topMessageDOList, err := m.MessagesDAO.SelectDialogLastMessageList(ctx, userId, dialogId, int32(len(msgIds)+1))
	if err != nil {
		return nil, err
	} else if len(topMessageDOList) == 0 {

	} else {
		topMessageIndex = math.MaxInt32
	}

	getLastTopMessage := func(topMessage2 int32) int32 {
		for i := 0; i < len(topMessageDOList); i++ {
			if topMessageDOList[i].UserMessageBoxId >= topMessage2 {
				continue
			}
			return topMessageDOList[i].UserMessageBoxId
		}
		return 0
	}

	_, err = m.ConversationsDAO.SelectByPeer(ctx, userId, model.GetPeerIdByDialogId(userId, dialogId))
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(msgDOList); i++ {
		topMessage := getLastTopMessage(topMessageIndex)
		if topMessage == msgDOList[i].UserMessageBoxId {
			topMessageIndex = topMessage
		}
	}

	tR := sqlx.TxWrapper(ctx, m.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, result.Err = m.MessagesDAO.DeleteMessagesByMessageIdList(ctx, userId, msgIds)
		if result.Err != nil {
			return
		}
		_, result.Err = m.ConversationsDAO.UpdateCustomMap(ctx, map[string]interface{}{
			"top_message": getLastTopMessage(topMessageIndex),
		}, userId, model.GetPeerIdByDialogId(userId, dialogId))
	})
	if tR.Err != nil {
		return nil, tR.Err
	}

	// collect deletedMsgDataIdList
	for i := 0; i < len(msgDOList); i++ {
		deletedMsgDataIdList = append(deletedMsgDataIdList, msgDOList[i].MessageDataId)
	}
	return deletedMsgDataIdList, nil
}

func (m *MsgCore) editOutboxMessage(ctx context.Context, fromId int32, peer *model.PeerUtil, toId int32, message *mtproto.Message) (box *model.MessageBox, err error) {
	mType, mData := model.EncodeMessage(message)
	if _, err = m.MessagesDAO.UpdateEditMessage(ctx, int8(mType), hack.String(mData), message.Message, fromId, message.Id); err != nil {
		return
	}
	box = &model.MessageBox{
		SelfUserId:        fromId,
		MessageId:         message.Id,
		DialogId:          model.MakeDialogId(fromId, peer.PeerType, peer.PeerId),
		DialogMessageId:   0,
		RandomId:          0,
		Pts:               int32(idgen.NextPtsId(ctx, fromId)),
		PtsCount:          1,
		MessageFilterType: 0,
		MessageBoxType:    0,
		MessageType:       int8(mType),
		Message:           message,
	}
	return
}

func (m *MsgCore) EditUserOutboxMessage(ctx context.Context, fromId int32, toId int32, message *mtproto.Message) (*model.MessageBox, error) {
	peer := &model.PeerUtil{PeerType: model.PEER_USER, PeerId: toId}
	return m.editOutboxMessage(ctx, fromId, peer, toId, message)
}

func (m *MsgCore) EditChatOutboxMessage(ctx context.Context, fromId int32, toId int32, message *mtproto.Message) (*model.MessageBox, error) {
	peer := &model.PeerUtil{PeerType: model.PEER_CHAT, PeerId: toId}
	return m.editOutboxMessage(ctx, fromId, peer, toId, message)
}

func (m *MsgCore) UpdateMediaUnread(ctx context.Context, userId int32, id int32) error {
	_, err := m.MessagesDAO.UpdateMediaUnread(ctx, userId, id)
	return err
}

func (m *MsgCore) UpdateMentioned(ctx context.Context, userId int32, id int32) error {
	_, err := m.MessagesDAO.UpdateMentionedAndMediaUnread(ctx, userId, id)
	return err
}
