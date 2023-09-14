package core

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/gogo/protobuf/proto"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (m *MsgCore) makeMessageInBox(fromId int32, peer *model.PeerUtil, toUserId int32, clientRandomId int64, dialogMessageId int32, messageDataId int64, message *mtproto.Message) *model.MessageBox {
	mentioned := model.CheckHasMention(message.Entities, toUserId)
	log.Infof("insert to inbox: %#v, message: {%#v}", mentioned, message)

	did := model.MakeDialogId(fromId, peer.PeerType, peer.PeerId)
	return &model.MessageBox{
		SelfUserId:        fromId,
		MessageId:         0,
		DialogId:          did,
		DialogMessageId:   dialogMessageId,
		MessageDataId:     messageDataId,
		RandomId:          clientRandomId,
		MessageFilterType: int8(model.GetMediaType(message)),
	}
}

func (m *MsgCore) sendMessageToInbox(ctx context.Context, fromId int32, peer *model.PeerUtil, toUserId int32, dialogMessageId int32, messageDataId, clientRandomId int64, message2 *mtproto.Message) (*model.MessageBox, error) {
	var (
		inBoxMsgId = int32(idgen.NextMessageBoxId(ctx, toUserId))
		dialogId   = model.MakeDialogId(fromId, peer.PeerType, peer.PeerId)
		date       = int32(time.Now().Unix())
		message    = proto.Clone(message2).(*mtproto.Message)
	)

	message.Out = false
	message.PeerId = peer.ToPeer()
	if peer.PeerType == model.PEER_USER {
		message.PeerId.UserId = fromId
	}
	message.Id = inBoxMsgId
	var replyOwnerId int32 = 0
	if message.ReplyTo != nil {
		if replyId, _ := m.MessagesDAO.SelectPeerMessageId(ctx, toUserId, fromId, message.ReplyTo.ReplyToMsgId); replyId != nil {
			replyOwnerId = replyId.SenderUserId
			message.ReplyTo.ReplyToMsgId = replyId.UserMessageBoxId
			if peer.PeerType == model.PEER_CHAT && replyId.UserId == toUserId {
				message.Mentioned = true
				message.MediaUnread = true
			}
		} else {
			message.ReplyTo = nil
		}
	}

	if peer.PeerType == model.PEER_CHAT {
		if !message.Mentioned {
			message.Mentioned = model.CheckHasMention(message.Entities, toUserId)
			if message.Mentioned {
				message.MediaUnread = true
			}
		}
	}

	if !message.MediaUnread {
		message.MediaUnread = model.CheckHasMediaUnread(message)
	}

	mType, mData := model.EncodeMessage(message)
	inBox := &model.MessageBox{
		SelfUserId:        toUserId,
		SendUserId:        fromId,
		MessageId:         inBoxMsgId,
		DialogId:          dialogId,
		DialogMessageId:   dialogMessageId,
		MessageDataId:     messageDataId,
		RandomId:          clientRandomId,
		Pts:               0,
		PtsCount:          1,
		MessageFilterType: int8(model.GetMediaType(message)),
		MessageBoxType:    model.MESSAGE_BOX_TYPE_INCOMING,
		MessageType:       int8(mType),
		Message:           message,
		Views:             0,
		ReplyOwnerId:      replyOwnerId,
		TtlSeconds:        message.GetTtlSeconds().GetValue(),
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		inBoxDO := &dataobject.MessagesDO{
			UserId:           inBox.SelfUserId,
			UserMessageBoxId: inBox.MessageId,
			DialogId:         inBox.DialogId,
			DialogMessageId:  inBox.DialogMessageId,
			SenderUserId:     fromId,
			PeerType:         int8(peer.PeerType),
			PeerId:           peer.PeerId,
			RandomId:         inBox.RandomId,
			MessageType:      int8(mType),
			MessageData:      hack.String(mData),
			MessageDataId:    inBox.MessageDataId,
			MessageDataType:  int32(inBox.MessageFilterType),
			Message:          message.Message,
			Pts:              0,
			PtsCount:         0,
			MessageBoxType:   inBox.MessageBoxType,
			ReplyToMsgId:     0,
			Mentioned:        util.BoolToInt8(message.Mentioned),
			MediaUnread:      util.BoolToInt8(message.MediaUnread),
			HasMediaUnread:   0,
			Date2:            date,
			Deleted:          0,
			TtlSeconds:       inBox.TtlSeconds,
		}

		_, _, result.Err = m.MessagesDAO.InsertOrReturnIdTx(tx, inBoxDO)
		if result.Err != nil {
			return
		}

		switch peer.PeerType {
		case model.PEER_USER:
			dialogDO := &dataobject.ConversationsDO{
				UserId:      inBox.SelfUserId,
				PeerId:      fromId,
				TopMessage:  int32(inBoxMsgId),
				UnreadCount: 1,
				Date2:       date,
				Deleted:     0,
			}

			_, _, result.Err = m.ConversationsDAO.InsertOrUpdateTx(tx, dialogDO)
		case model.PEER_CHAT:
			if inBox.Message.GetAction().GetPredicateName() == mtproto.Predicate_messageActionChatMigrateTo {
				m.ChatParticipantsDAO.UpdateInboxDialogTx(tx, map[string]interface{}{
					"read_inbox_max_id":     inBoxMsgId,
					"read_outbox_max_id":    inBoxMsgId,
					"unread_count":          0,
					"unread_mentions_count": 0,
					"draft_type":            0,
					"draft_message_data":    "",
				}, inBox.SelfUserId, peer.PeerId)
			} else {
				cMap := map[string]interface{}{
					"top_message": inBoxMsgId,
					"date2":       date,
					"unread_mark": 0,
				}
				_, result.Err = m.ChatParticipantsDAO.UpdateInboxDialogTx(tx, cMap, inBox.SelfUserId, peer.PeerId)
			}
			if result.Err != nil {
				return
			}
		default:
			result.Err = fmt.Errorf("fatal error - invalid peer_type: %v", peer)
		}
	})

	if tR.Err != nil {
		return nil, tR.Err
	}

	inBox.Pts = int32(idgen.NextPtsId(ctx, toUserId))
	inBox.PtsCount = 1

	return inBox, nil
}

func (m *MsgCore) SendUserMessageToInbox(ctx context.Context, fromId, toId, dialogMessageId int32, messageDataId, clientRandomId int64, message *mtproto.Message) (*model.MessageBox, error) {
	peer := model.MakeUserPeerUtil(toId)
	return m.sendMessageToInbox(ctx, fromId, peer, toId, dialogMessageId, messageDataId, clientRandomId, message)
}

func (m *MsgCore) SendChatMessageToInbox(ctx context.Context, fromId, chatId, toId, dialogMessageId int32, messageDataId, clientRandomId int64, message *mtproto.Message) (*model.MessageBox, error) {
	peer := model.MakeChatPeerUtil(chatId)
	return m.sendMessageToInbox(ctx, fromId, peer, toId, dialogMessageId, messageDataId, clientRandomId, message)
}

func (m *MsgCore) SendUserMultiMessageToInbox(ctx context.Context, fromId, toId int32, dialogMessageIdList []int32, messageDataIdList, clientRandomIdList []int64, messageList []*mtproto.Message) ([]*model.MessageBox, error) {
	var inBoxList = make([]*model.MessageBox, 0, len(messageList))
	for i, message := range messageList {
		peer := model.MakeUserPeerUtil(toId)
		inBox, _ := m.sendMessageToInbox(ctx, fromId, peer, toId, dialogMessageIdList[i], messageDataIdList[i], clientRandomIdList[i], message)
		inBoxList = append(inBoxList, inBox)
	}
	return inBoxList, nil
}

func (m *MsgCore) SendChatMultiMessageToInbox(ctx context.Context, fromId, chatId, toId int32, dialogMessageIdList []int32, messageDataIdList, clientRandomIdList []int64, messageList []*mtproto.Message) ([]*model.MessageBox, error) {
	var inBoxList = make([]*model.MessageBox, 0, len(messageList))
	for i, message := range messageList {
		peer := model.MakeChatPeerUtil(chatId)
		inBox, _ := m.sendMessageToInbox(ctx, fromId, peer, toId, dialogMessageIdList[i], messageDataIdList[i], clientRandomIdList[i], message)
		inBoxList = append(inBoxList, inBox)
	}
	return inBoxList, nil
}

func (m *MsgCore) DeleteInboxMessages(ctx context.Context, deleteUserId int32, deleteMsgDataIds []int64, cb func(ctx context.Context, userId int32, idList []int32)) error {
	var (
		deletedDialogsMap = map[int32][]*dataobject.MessagesDO{}
	)

	mDOList, err := m.MessagesDAO.SelectByMessageDataIdList(ctx, deleteMsgDataIds)
	if err != nil {
		return err
	}

	for i := 0; i < len(mDOList); i++ {
		if mDOList[i].UserId == deleteUserId {
			continue
		}

		if v, ok := deletedDialogsMap[mDOList[i].UserId]; !ok {
			deletedDialogsMap[mDOList[i].UserId] = []*dataobject.MessagesDO{&mDOList[i]}
		} else {
			deletedDialogsMap[mDOList[i].UserId] = append(v, &mDOList[i])
		}
	}
	for userId, msgDOList := range deletedDialogsMap {

		var (
			topMessageIndex int32
			dialogId        int64
			msgIds          []int32
		)

		for i := 0; i < len(msgDOList); i++ {
			if dialogId == 0 {
				dialogId = msgDOList[i].DialogId
			}

			if dialogId != msgDOList[i].DialogId {
				err = mtproto.ErrMessageIdInvalid
				return err
			}
			msgIds = append(msgIds, msgDOList[i].UserMessageBoxId)
		}
		topMessageDOList, err := m.MessagesDAO.SelectDialogLastMessageList(ctx, userId, dialogId, int32(len(msgIds)+1))
		if err != nil {
			return err
		} else if len(topMessageDOList) == 0 {

		} else {
			topMessageIndex = math.MaxInt32
		}

		getLastTopMessage := func(topMessage2 int32) int32 {
			for i := 0; i < len(topMessageDOList); i++ {
				if topMessageDOList[i].UserMessageBoxId >= topMessage2 {
					continue
				} else {
					return topMessageDOList[i].UserMessageBoxId
				}
			}
			return 0
		}

		_, err = m.ConversationsDAO.SelectByPeer(ctx, userId, model.GetPeerIdByDialogId(userId, dialogId))
		if err != nil {
			return err
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
			return tR.Err
		}

		if cb != nil {
			cb(ctx, userId, msgIds)
		}
	}
	return nil
}

func (m *MsgCore) EditUserInboxMessage(ctx context.Context, fromId int32, peerId int32, message *mtproto.Message) (box *model.MessageBox, err error) {
	var peerMsgDO *dataobject.MessagesDO

	peerMsgDO, err = m.MessagesDAO.SelectPeerMessageId(ctx, peerId, fromId, message.Id)
	if err != nil {
		return
	} else if peerMsgDO == nil {
		return
	}

	message.Out = false
	if message.PeerId != nil && message.PeerId.PredicateName == mtproto.Predicate_peerUser {
		message.PeerId.UserId = fromId
	}
	message.Id = peerMsgDO.UserMessageBoxId
	var replyOwnerId int32 = 0
	if message.ReplyTo != nil {
		peerMessage, _ := model.DecodeMessage(int(peerMsgDO.MessageType), hack.Bytes(peerMsgDO.MessageData))
		message.ReplyTo = peerMessage.ReplyTo
		if message.ReplyTo != nil {
			if replyId, _ := m.MessagesDAO.SelectPeerMessageId(ctx, peerId, fromId, message.ReplyTo.ReplyToMsgId); replyId != nil {
				replyOwnerId = replyId.SenderUserId
			}
		}
	}

	mType, mData := model.EncodeMessage(message)
	if _, err = m.MessagesDAO.UpdateEditMessage(ctx, int8(mType), hack.String(mData), message.Message, peerId, message.Id); err != nil {
		return
	}

	box = &model.MessageBox{
		SelfUserId:        peerId,
		SendUserId:        0,
		MessageId:         message.Id,
		DialogId:          0,
		DialogMessageId:   0,
		RandomId:          0,
		Pts:               int32(idgen.NextPtsId(ctx, peerId)),
		PtsCount:          1,
		MessageFilterType: 0,
		MessageBoxType:    0,
		MessageType:       0,
		Message:           message,
		Views:             0,
		ReplyOwnerId:      replyOwnerId,
	}
	return
}

func (m *MsgCore) EditChatInboxMessage(ctx context.Context, fromId int32, peer *model.PeerUtil, toId int32, message *mtproto.Message) (box *model.MessageBox, err error) {
	var peerMsgDO *dataobject.MessagesDO

	peerMsgDO, err = m.MessagesDAO.SelectPeerMessageId(ctx, toId, fromId, message.Id)
	if err != nil {
		return
	} else if peerMsgDO == nil {
		return
	}

	message.Out = false
	if message.PeerId != nil && message.PeerId.PredicateName == mtproto.Predicate_peerUser {
		message.PeerId.UserId = fromId
	}
	message.Id = peerMsgDO.UserMessageBoxId
	var replyOwnerId int32 = 0
	if message.ReplyTo != nil {
		peerMessage, _ := model.DecodeMessage(int(peerMsgDO.MessageType), hack.Bytes(peerMsgDO.MessageData))
		message.ReplyTo = peerMessage.ReplyTo
		if message.ReplyTo != nil {
			if replyId, _ := m.MessagesDAO.SelectPeerMessageId(ctx, toId, fromId, message.ReplyTo.ReplyToMsgId); replyId != nil {
				replyOwnerId = replyId.SenderUserId
			}
		}
	}

	mType, mData := model.EncodeMessage(message)
	if _, err = m.MessagesDAO.UpdateEditMessage(ctx, int8(mType), hack.String(mData), message.Message, toId, message.Id); err != nil {
		return
	}

	box = &model.MessageBox{
		SelfUserId:        toId,
		SendUserId:        0,
		MessageId:         message.Id,
		DialogId:          0,
		DialogMessageId:   0,
		RandomId:          0,
		Pts:               int32(idgen.NextPtsId(ctx, toId)),
		PtsCount:          1,
		MessageFilterType: 0,
		MessageBoxType:    0,
		MessageType:       0,
		Message:           message,
		Views:             0,
		ReplyOwnerId:      replyOwnerId,
	}
	return
}
