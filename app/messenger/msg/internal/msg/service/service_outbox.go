package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) sendUserMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, randomId int64,
	message *mtproto.Message, cb func(did int32, mid int64, inboxMsg *mtproto.Message) error) (*mtproto.Updates, error) {
	sendMe := fromUserId == toUserId
	if !sendMe {
	}

	hasDuplicateMessage, err := s.MsgCore.HasDuplicateMessage(ctx, fromUserId, randomId)
	if err != nil {
		log.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := s.MsgCore.GetDuplicateMessage(ctx, fromUserId, randomId)
		if err != nil {
			log.Errorf("checkDuplicateMessage error - %v", err)
			return nil, err
		} else if upd != nil {
			return upd, nil
		}
	}

	outBox, err := s.MsgCore.SendUserMessage(ctx, fromUserId, toUserId, randomId, message)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if !hasDuplicateMessage && cb != nil {
		err = cb(outBox.DialogMessageId, outBox.MessageDataId, outBox.ToMessage(fromUserId))
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	updatesHelper := model.MakeUpdatesHelper()
	updatesHelper.PushBackUpdate(mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Pts_INT32:       outBox.Pts,
		PtsCount:        outBox.PtsCount,
		RandomId:        outBox.RandomId,
		Message_MESSAGE: model.MessageUpdate(outBox.Message),
	}).To_Update())
	syncNotMe := updatesHelper.ToSyncNotMeUpdates(ctx, fromUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	err = sync_client.SyncUpdatesNotMe(ctx, fromUserId, fromAuthKeyId, syncNotMe)
	if err != nil {
		return nil, err
	}

	replyUpdates := updatesHelper.ToReplyUpdates(ctx, fromUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	s.MsgCore.PutDuplicateMessage(ctx, fromUserId, randomId, replyUpdates)

	return replyUpdates, nil
}

func (s *Service) pushUserMessage(ctx context.Context, fromUserId int32, toUserId int32, randomId int64,
	message *mtproto.Message, pushType int32, cb func(did int32, mid int64, inboxMsg *mtproto.Message) error) error {
	sendMe := fromUserId == toUserId
	if !sendMe {
	}

	hasDuplicateMessage, err := s.MsgCore.HasDuplicateMessage(ctx, fromUserId, randomId)
	if err != nil {
		log.Errorf("checkDuplicateMessage error - %v", err)
		return err
	} else if hasDuplicateMessage {
		upd, err := s.MsgCore.GetDuplicateMessage(ctx, fromUserId, randomId)
		if err != nil {
			log.Errorf("checkDuplicateMessage error - %v", err)
			return err
		} else if upd != nil {
			return nil
		}
	}

	outBox, err := s.MsgCore.SendUserMessage(ctx, fromUserId, toUserId, randomId, message)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	if !hasDuplicateMessage && cb != nil {
		err = cb(outBox.DialogMessageId, outBox.MessageDataId, outBox.ToMessage(fromUserId))
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}
	updatesHelper := model.MakeUpdatesHelper()
	updatesHelper.PushBackUpdate(mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Pts_INT32:       outBox.Pts,
		PtsCount:        outBox.PtsCount,
		RandomId:        outBox.RandomId,
		Message_MESSAGE: model.MessageUpdate(outBox.Message),
	}).To_Update())
	pushUpdates := updatesHelper.ToPushUpdates(ctx, fromUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	err = sync_client.PushUpdates(ctx, fromUserId, pushUpdates)

	return err
}

func (s *Service) sendChatMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, chatId int32,
	randomId int64, message *mtproto.Message, cb func(did int32, mid int64, inboxMsg *mtproto.Message) error) (*mtproto.Updates, error) {
	hasDuplicateMessage, err := s.MsgCore.HasDuplicateMessage(ctx, fromUserId, randomId)
	if err != nil {
		log.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := s.MsgCore.GetDuplicateMessage(ctx, fromUserId, randomId)
		if err != nil {
			log.Errorf("checkDuplicateMessage error - %v", err)
			return nil, err
		} else if upd != nil {
			return upd, nil
		}
	}

	outBox, err := s.MsgCore.SendChatMessage(ctx, fromUserId, chatId, randomId, message)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if !hasDuplicateMessage && cb != nil {
		err = cb(outBox.DialogMessageId, outBox.MessageDataId, outBox.ToMessage(fromUserId))
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
		Pts_INT32:       outBox.Pts,
		PtsCount:        outBox.PtsCount,
		RandomId:        outBox.RandomId,
		Message_MESSAGE: model.MessageUpdate(outBox.Message),
	}).To_Update())
	syncNotMeUpdates := updatesHelper.ToSyncNotMeUpdates(ctx, fromUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	err = sync_client.SyncUpdatesNotMe(ctx, fromUserId, fromAuthKeyId, syncNotMeUpdates)
	if err != nil {
		return nil, err
	}

	replyUpdates := updatesHelper.ToReplyUpdates(ctx, fromUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	s.MsgCore.PutDuplicateMessage(ctx, fromUserId, randomId, replyUpdates)

	return replyUpdates, nil
}

func (s *Service) sendUserMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, toUserId int32, randomIdList []int64,
	messages []*mtproto.Message, cb func(didList []int32, midList []int64, inboxMsgList []*mtproto.Message) error) (*mtproto.Updates, error) {
	sendMe := fromUserId == toUserId
	if !sendMe {
	}

	outBoxList, err := s.MsgCore.SendUserMultiMessage(ctx, fromUserId, toUserId, randomIdList, messages)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if cb != nil {
		var (
			didList       = make([]int32, 0, len(outBoxList))
			midList       = make([]int64, 0, len(outBoxList))
			outBoxMsgList = make([]*mtproto.Message, 0, len(outBoxList))
		)

		for _, outBoxMsg := range outBoxList {
			didList = append(didList, outBoxMsg.DialogMessageId)
			midList = append(midList, outBoxMsg.MessageDataId)
			outBoxMsgList = append(outBoxMsgList, outBoxMsg.ToMessage(fromUserId))

		}
		err = cb(didList, midList, outBoxMsgList)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	updatesHelper := model.MakeUpdatesHelper()
	for _, outBox := range outBoxList {
		updatesHelper.PushBackUpdate(mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
			Pts_INT32:       outBox.Pts,
			PtsCount:        outBox.PtsCount,
			RandomId:        outBox.RandomId,
			Message_MESSAGE: model.MessageUpdate(outBox.Message),
		}).To_Update())
	}

	syncNotMeUpdates := updatesHelper.ToSyncNotMeUpdates(ctx, fromUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	err = sync_client.SyncUpdatesNotMe(ctx, fromUserId, fromAuthKeyId, syncNotMeUpdates)
	if err != nil {
		return nil, err
	}

	replyUpdates := updatesHelper.ToReplyUpdates(ctx, fromUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	return replyUpdates, nil
}

func (s *Service) sendChatMultiMessage(ctx context.Context, fromUserId int32, fromAuthKeyId int64, chatId int32, randomIds []int64,
	messages []*mtproto.Message, cb func(didList []int32, midList []int64, inboxMsgList []*mtproto.Message) error) (*mtproto.Updates, error) {
	outBoxList, err := s.MsgCore.SendChatMultiMessage(ctx, fromUserId, chatId, randomIds, messages)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if cb != nil {
		var (
			didList       = make([]int32, 0, len(outBoxList))
			midList       = make([]int64, 0, len(outBoxList))
			outBoxMsgList = make([]*mtproto.Message, 0, len(outBoxList))
		)

		for _, outBoxMsg := range outBoxList {
			didList = append(didList, outBoxMsg.DialogMessageId)
			midList = append(midList, outBoxMsg.MessageDataId)
			outBoxMsgList = append(outBoxMsgList, outBoxMsg.ToMessage(fromUserId))

		}
		err = cb(didList, midList, outBoxMsgList)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}

	updatesHelper := model.MakeUpdatesHelper()
	for _, outBox := range outBoxList {
		updatesHelper.PushBackUpdate(mtproto.MakeTLUpdateNewMessage(&mtproto.Update{
			Pts_INT32:       outBox.Pts,
			PtsCount:        outBox.PtsCount,
			RandomId:        outBox.RandomId,
			Message_MESSAGE: model.MessageUpdate(outBox.Message),
		}).To_Update())
	}

	syncNotMeUpdates := updatesHelper.ToSyncNotMeUpdates(ctx, fromUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	err = sync_client.SyncUpdatesNotMe(ctx, fromUserId, fromAuthKeyId, syncNotMeUpdates)
	if err != nil {
		return nil, err
	}

	replyUpdates := updatesHelper.ToReplyUpdates(ctx, fromUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	return replyUpdates, nil
}
