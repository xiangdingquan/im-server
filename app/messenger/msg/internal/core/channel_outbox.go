package core

import (
	"context"
	"fmt"
	"time"

	"open.chat/app/messenger/msg/internal/dal/dataobject"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (m *MsgCore) SendChannelMessage(ctx context.Context, sendUserId int32, channelId int32, randomId int64, message *mtproto.Message, dmUsers []int32) (*model.MessageBox, error) {
	var (
		outBoxMsgId       = int32(idgen.NextChannelMessageBoxId(ctx, channelId))
		pts               = int32(idgen.NextChannelPtsId(ctx, channelId))
		ptsCount    int32 = 1
	)

	if pts == 0 || outBoxMsgId == 0 {
		return nil, mtproto.ErrInternelServerError
	}

	if pts == 1 {
		pts = int32(idgen.NextChannelPtsId(ctx, channelId))
	}

	message.Out = true
	message.PeerId = model.MakePeerChannel(channelId)
	message.Id = outBoxMsgId
	mType, mData := model.EncodeMessage(message)

	message.Mentioned = false
	message.MediaUnread = true

	outBoxMsg := &model.MessageBox{
		SelfUserId:        sendUserId,
		SendUserId:        sendUserId,
		MessageId:         int32(outBoxMsgId),
		DialogId:          model.MakeDialogId(sendUserId, model.PEER_CHANNEL, channelId),
		DialogMessageId:   int32(outBoxMsgId),
		MessageDataId:     int64(channelId)<<32 | int64(outBoxMsgId),
		RandomId:          randomId,
		Pts:               pts,
		PtsCount:          ptsCount,
		MessageFilterType: int8(model.GetMediaType(message)),
		MessageBoxType:    model.MESSAGE_BOX_TYPE_CHANNEL,
		MessageType:       int8(mType),
		Message:           message,
		Views:             0,
		ReplyOwnerId:      0,
		TtlSeconds:        message.GetTtlSeconds().GetValue(),
		DmUsers:           dmUsers,
	}

	if message.Views != nil {
		outBoxMsg.Views = message.Views.Value
	} else if message.Forwards != nil {
		outBoxMsg.Views = message.Forwards.Value
	} else {
		if message.ReplyTo != nil && message.ReplyTo.GetReplyToMsgId() != 0 {
			replyId, _ := m.ChannelMessagesDAO.SelectByMessageId(ctx, channelId, message.ReplyTo.GetReplyToMsgId())
			if replyId != nil {
				outBoxMsg.ReplyOwnerId = replyId.SenderUserId
			}
		}
	}

	log.Debugf("channelMsg: %v", outBoxMsg)

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		channelMessageDO := &dataobject.ChannelMessagesDO{
			Id:               0,
			ChannelId:        channelId,
			ChannelMessageId: outBoxMsg.MessageId,
			SenderUserId:     outBoxMsg.SelfUserId,
			RandomId:         outBoxMsg.RandomId,
			Pts:              outBoxMsg.Pts,
			MessageDataId:    outBoxMsg.MessageDataId,
			MessageType:      outBoxMsg.MessageType,
			MessageData:      hack.String(mData),
			MediaType:        int8(model.GetMediaType(message)),
			Message:          message.Message,
			HasMediaUnread:   0,
			EditMessage:      "",
			EditDate:         0,
			Views:            outBoxMsg.Views,
			Date:             outBoxMsg.Message.Date,
			TtlSeconds:       outBoxMsg.TtlSeconds,
			HasDM:            util.BoolToInt8(len(outBoxMsg.DmUsers) > 0),
		}

		lastInsertId, rowsAffected, err := m.ChannelMessagesDAO.InsertOrGetIdTx(tx, channelMessageDO)
		log.Debugf("last_inser_id = %d, rows_affected = %d, err = %v", lastInsertId, rowsAffected, err)
		if err != nil {
			result.Err = err
			return
		}

		if rowsAffected == 0 {
			if lastInsertId > 0 {
				result.Data = lastInsertId
				return
			} else {
				result.Err = fmt.Errorf("insert error")
				return
			}
		}

		_, err = m.ChannelsDAO.UpdateTopMessageTx(tx, outBoxMsg.MessageId, outBoxMsg.Pts, outBoxMsg.Message.Date, channelId)
		if err != nil {
			result.Err = err
			return
		}

		_, err = m.ChannelParticipantsDAO.UpdateReadInboxMaxIdTx(tx, outBoxMsg.MessageId, sendUserId, channelId)
		if err != nil {
			result.Err = err
			return
		}

		ptsUpdatesDO := &dataobject.ChannelPtsUpdatesDO{
			ChannelId:    channelId,
			Pts:          pts,
			PtsCount:     ptsCount,
			UpdateType:   9,
			NewMessageId: outBoxMsg.MessageId,
			Date2:        int32(time.Now().Unix()),
		}
		_, _, err = m.ChannelPtsUpdatesDAO.InsertTx(tx, ptsUpdatesDO)
		if err != nil {
			result.Err = err
		}

		var visiblesDos []*dataobject.ChannelMessageVisiblesDO
		for _, userId := range outBoxMsg.DmUsers {
			visiblesDos = append(visiblesDos, &dataobject.ChannelMessageVisiblesDO{
				UserId:    userId,
				ChannelId: channelId,
				MessageId: outBoxMsg.MessageId,
			})
		}
		if len(visiblesDos) > 0 {
			_, _, err = m.ChannelMessageVisiblesDAO.InsertBulkTx(tx, visiblesDos)
			if err != nil {
				result.Err = err
			}
		}
	})

	if tR.Err != nil {
		return nil, tR.Err
	}

	switch tR.Data.(type) {
	case int64:
		if tR.Data.(int64) <= 0 {
			log.Error("unknown error")
			return nil, fmt.Errorf("fatal unknown error")
		}

		do, err := m.ChannelMessagesDAO.SelectById(ctx, tR.Data.(int64))
		if err != nil {
			return nil, err
		}
		if do != nil {
			outBoxMsg = &model.MessageBox{
				SelfUserId:        sendUserId,
				SendUserId:        sendUserId,
				MessageId:         do.ChannelMessageId,
				DialogId:          model.MakeDialogId(sendUserId, model.PEER_CHANNEL, channelId),
				DialogMessageId:   do.ChannelMessageId,
				MessageDataId:     int64(channelId)<<32 | int64(do.ChannelMessageId),
				RandomId:          randomId,
				Pts:               do.Pts,
				PtsCount:          ptsCount,
				MessageFilterType: do.MediaType,
				MessageBoxType:    model.MESSAGE_BOX_TYPE_CHANNEL,
				MessageType:       int8(mType),
				Message:           nil,
				ReplyOwnerId:      outBoxMsg.ReplyOwnerId,
				TtlSeconds:        outBoxMsg.TtlSeconds,
			}
			outBoxMsg.Message, _ = model.DecodeMessage(int(do.MessageType), hack.Bytes(do.MessageData))
		} else {
			log.Error("unknown error")
			return nil, fmt.Errorf("fatal unknown error")
		}
	default:
	}

	return outBoxMsg, nil
}

func (m *MsgCore) EditChannelOutboxMessage(ctx context.Context, fromId int32, toId int32, message *mtproto.Message) (box *model.MessageBox, err error) {
	mType, mData := model.EncodeMessage(message)
	if _, err = m.ChannelMessagesDAO.Update(ctx, map[string]interface{}{
		"message_data": hack.String(mData),
		"message":      message.Message,
	}, toId, message.Id); err != nil {
		return
	}
	box = &model.MessageBox{
		SelfUserId:        fromId,
		SendUserId:        fromId,
		MessageId:         int32(message.Id),
		DialogId:          model.MakeDialogId(fromId, model.PEER_CHANNEL, toId),
		DialogMessageId:   int32(message.Id),
		MessageDataId:     int64(toId)<<32 | int64(message.Id),
		RandomId:          0,
		Pts:               int32(idgen.NextChannelPtsId(ctx, toId)),
		PtsCount:          1,
		MessageFilterType: int8(model.GetMediaType(message)),
		MessageBoxType:    model.MESSAGE_BOX_TYPE_CHANNEL,
		MessageType:       int8(mType),
		Message:           message,
		Views:             0,
		ReplyOwnerId:      0,
	}
	return
}

func (m *MsgCore) DeleteMessagesJustSelf(ctx context.Context, user_id int32, channel_id int32, msg_ids []int32) (int32, error) {
	pts := int32(idgen.NextChannelNPtsId(ctx, channel_id, len(msg_ids)))
	do := &dataobject.ChannelMessagesDeleteDO{
		UserId:    user_id,
		ChannelId: channel_id,
	}
	tR := sqlx.TxWrapper(ctx, m.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		for _, msg_id := range msg_ids {
			do.MessageId = msg_id
			_, _, result.Err = m.ChannelMessagesDeleteDAO.InsertOrGetIdTx(tx, do)
			if result.Err != nil {
				return
			}
		}
		m.ChannelMessagesDAO.RemoveMessagesTx(tx, channel_id, msg_ids)
		_, result.Err = m.ChannelsDAO.UpdateTopMessagePtsTx(tx, pts, int32(time.Now().Unix()), channel_id)
		if result.Err != nil {
			return
		}
	})

	if tR.Err != nil {
		return 0, tR.Err
	}
	return pts, nil
}

func (m *MsgCore) DeleteChannelMessages(ctx context.Context, channelId int32, msgIds []int32) (int32, error) {
	lastMessageId, err := m.ChannelMessagesDAO.SelectLastMessageNotIdList(ctx, channelId, msgIds)
	if err != nil {
		return 0, err
	}

	pts := int32(idgen.NextChannelNPtsId(ctx, channelId, len(msgIds)))
	tR := sqlx.TxWrapper(ctx, m.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, result.Err = m.ChannelMessagesDAO.DeleteMessagesTx(tx, channelId, msgIds)
		if result.Err != nil {
			return
		}

		_, err = m.ChannelsDAO.UpdateTopMessageTx(tx, lastMessageId, pts, int32(time.Now().Unix()), channelId)
		if err != nil {
			result.Err = err
			return
		}
	})
	if tR.Err != nil {
		return 0, tR.Err
	}
	return pts, nil
}

func (m *MsgCore) GetChannelMessageIdListBySenderUserId(ctx context.Context, channelId, senderId int32) []int32 {
	idList, _ := m.ChannelMessagesDAO.SelectMessageIdListBySenderUserId(ctx, channelId, senderId)
	return idList
}
