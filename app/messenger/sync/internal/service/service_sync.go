package service

import (
	"context"
	"fmt"

	"open.chat/app/bots/botpb"
	"open.chat/app/interface/session/sessionpb"
	"open.chat/app/messenger/push/pushpb"
	"open.chat/app/messenger/sync/syncpb"
	status_client "open.chat/app/service/status/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
	"open.chat/pkg/util"

	"open.chat/app/json/consts"
	"open.chat/app/json/helper"
	"open.chat/app/json/service/handler"
)

func (s *Service) onSyncUpdates(ctx context.Context, r *syncpb.TLSyncSyncUpdates) error {
	log.Debugf("onSyncUpdates - request: {%s}", logger.JsonDebugData(r))

	notification, err := s.processUpdatesRequest(ctx, r.GetUserId(), false, r.GetUpdates())
	if err == nil {
		userId := r.GetUserId()
		authKeyId := r.GetAuthKeyId()
		serverId := r.GetServerId()
		sessionId := r.GetSessionId().GetValue()
		if serverId == nil {
			s.pushUpdatesToSession(ctx, syncTypeUserNotMe, userId, authKeyId, 0, r.Updates, "", notification)
		} else {
			if authKeyId != 0 {
				s.pushUpdatesToSession(ctx, syncTypeUserMe, userId, authKeyId, sessionId, r.Updates, serverId.Value, notification)
			} else {
				s.pushUpdatesToSession(ctx, syncTypeUser, userId, 0, 0, r.Updates, serverId.Value, notification)
			}
		}
	} else {
		log.Error(err.Error())
		return nil
	}

	log.Debugf("sync.syncUpdates#3a077679 - reply: {true}")
	return nil
}

func (s *Service) onPushUpdates(ctx context.Context, r *syncpb.TLSyncPushUpdates) error {
	log.Debugf("sync.pushUpdates#5c612649 - request: {%s}", logger.JsonDebugData(r))

	notification, err := s.processUpdatesRequest(ctx, r.GetUserId(), r.GetIsBot(), r.GetUpdates())
	if err == nil {
		s.pushUpdatesToSession(ctx, syncTypeUser, r.GetUserId(), 0, 0, r.Updates, "", notification)
	} else {
		log.Errorf(err.Error())
		return nil
	}

	log.Debugf("sync.pushUpdates#5c612649 - reply: {true}")
	return nil
}

func (s *Service) onBroadcastUpdates(ctx context.Context, r *syncpb.TLSyncBroadcastUpdates) error {
	log.Debugf("sync.pushUpdates#5c612649 - request: {%s}", logger.JsonDebugData(r))

	switch r.BroadcastType {
	case model.BroadcastTypeChat:
		idList, _ := s.ChatFacade.GetChatParticipantIdList(ctx, r.ChatId)
		for _, id := range idList {
			if ok, _ := util.Contains(id, r.ExcludeIds); !ok {
				notification, err := s.processUpdatesRequest(ctx, id, false, r.GetUpdates())
				if err == nil {
					s.pushUpdatesToSession(ctx, syncTypeUser, id, 0, 0, r.Updates, "", notification)
				} else {
					log.Errorf(err.Error())
					return nil
				}
			}
		}
	case model.BroadcastTypeChannel:
		idList := s.ChannelFacade.GetChannelParticipantIdList(ctx, r.ChatId)
		for _, id := range idList {
			if ok, _ := util.Contains(id, r.ExcludeIds); !ok {
				notification, err := s.processUpdatesRequest(ctx, id, false, r.GetUpdates())
				if err == nil {
					s.pushUpdatesToSession(ctx, syncTypeUser, id, 0, 0, r.Updates, "", notification)
				} else {
					log.Errorf(err.Error())
					return nil
				}
			}
		}
	case model.BroadcastTypeChannelAdmins:
		idList := s.GetChannelAdminParticipantIdList(ctx, r.ChatId)
		for _, id := range idList {
			if ok, _ := util.Contains(id, r.ExcludeIds); !ok {
				notification, err := s.processUpdatesRequest(ctx, id, false, r.GetUpdates())
				if err == nil {
					s.pushUpdatesToSession(ctx, syncTypeUser, id, 0, 0, r.Updates, "", notification)
				} else {
					log.Errorf(err.Error())
					return nil
				}
			}
		}
	default:

	}

	log.Debugf("sync.pushUpdates#5c612649 - reply: {true}")
	return nil
}

func (s *Service) processUpdatesRequest(ctx context.Context, userId int32, isBot bool, ups *mtproto.Updates) (notification bool, err error) {
	switch ups.PredicateName {
	case mtproto.Predicate_updateAccountResetAuthorization:
	case mtproto.Predicate_updateShortMessage:
		shortMessage := ups.To_UpdateShortMessage()
		if isBot {
			shortMessage.Data2.Pts = s.Dao.AddToBotUpdateQueue(ctx, userId, updateShortToUpdateNewMessage(userId, shortMessage))
		} else {
			s.Dao.AddToPtsQueue(ctx, userId, shortMessage.GetPts(), shortMessage.GetPtsCount(), updateShortToUpdateNewMessage(userId, shortMessage))
			notification = true
		}
	case mtproto.Predicate_updateShortChatMessage:
		shortMessage := ups.To_UpdateShortChatMessage()
		if isBot {
			shortMessage.Data2.Pts = s.Dao.AddToBotUpdateQueue(ctx, userId, updateShortChatToUpdateNewMessage(userId, shortMessage))
		} else {
			s.Dao.AddToPtsQueue(ctx, userId, shortMessage.GetPts(), shortMessage.GetPtsCount(), updateShortChatToUpdateNewMessage(userId, shortMessage))
			notification = true
		}
	case mtproto.Predicate_updateShort:
	case mtproto.Predicate_updates:
		updates2 := ups.To_Updates()
		for _, update := range updates2.GetUpdates() {
			switch update.PredicateName {
			case mtproto.Predicate_updateNewMessage,
				mtproto.Predicate_updateEditMessage:

				if isBot {
					update.Pts_INT32 = s.Dao.AddToBotUpdateQueue(ctx, userId, update)
				} else {
					s.Dao.AddToPtsQueue(ctx, userId, update.Pts_INT32, update.PtsCount, update)
					notification = true
					msg := update.Message_MESSAGE
					if msg.PredicateName == mtproto.Predicate_message {
						peer := model.FromPeer(msg.GetPeerId())
						if peer.PeerType == model.PEER_USER {
							peer.PeerId = msg.GetFromId_FLAGPEER().GetUserId()
						}
						notifySettings, err := s.UserFacade.GetNotifySettings(ctx, userId, peer)
						if err == nil && notifySettings != nil {
							notification = notifySettings.GetSilent_FLAGBOOL() == nil
						}
					}
				}
			case mtproto.Predicate_updateReadHistoryOutbox,
				mtproto.Predicate_updateReadHistoryInbox,
				mtproto.Predicate_updateWebPage,
				mtproto.Predicate_updateReadMessagesContents:

				s.Dao.AddToPtsQueue(ctx, userId, update.Pts_INT32, update.PtsCount, update)
				if !isBot {
					notification = true
				}
			case mtproto.Predicate_updateDeleteMessages:
				s.Dao.AddToPtsQueue(ctx, userId, update.Pts_INT32, update.PtsCount, update)
				if !isBot {
					notification = true
				}
			case mtproto.Predicate_updateBotCallbackQuery:
				if isBot {
					update.Pts_INT32 = s.Dao.AddToBotUpdateQueue(ctx, userId, update)
				}
			case mtproto.Predicate_updatePhoneCall:
				if update.GetPhoneCall().GetPredicateName() == mtproto.Predicate_phoneCallRequested {
					notification = true
				}
			case mtproto.Predicate_updateEncryption,
				mtproto.Predicate_updateNewEncryptedMessage:
				if !isBot {
					notification = true
				}
			case mtproto.Predicate_updateNewBlog,
				mtproto.Predicate_updateDeleteBlog,
				mtproto.Predicate_updateBlogFollow,
				mtproto.Predicate_updateBlogComment,
				mtproto.Predicate_updateBlogLike:
				if !isBot {
					s.Dao.AddToBlogPtsQueue(ctx, userId, update.Pts_INT32, update.PtsCount, update)
				}
			}
		}
	default:
		err = fmt.Errorf("invalid updates data: {%d}", ups.GetConstructor())
		return
	}

	return
}

func (s *Service) pushUpdatesToSession(ctx context.Context, syncType SyncType, userId int32, authKeyId, clientMsgId int64, pushData *mtproto.Updates, hasServerId string, notification bool) {
	if pushData.Update != nil {
		if pushData.Update.Message_MESSAGE != nil {
			pushData.Update.Message_MESSAGE = model.MessageUpdate(pushData.Update.Message_MESSAGE)
		}
	}
	for _, u := range pushData.Updates {
		if u.Message_MESSAGE != nil {
			u.Message_MESSAGE = model.MessageUpdate(u.Message_MESSAGE)
		}
	}
	if syncType == syncTypeUserMe && hasServerId != "" {
		log.Infof("pushUpdatesToSession - pushData: {server_id: %d, auth_key_id: %d}", hasServerId, authKeyId)
		if clientMsgId != 0 {
			s.PushSessionUpdatesToSession(
				ctx,
				hasServerId,
				&sessionpb.PushSessionUpdatesData{
					AuthKeyId: authKeyId,
					SessionId: clientMsgId,
					Updates:   pushData,
				})
		} else {
			s.PushUpdatesToSession(
				ctx,
				hasServerId,
				&sessionpb.PushUpdatesData{
					AuthKeyId:    authKeyId,
					Notification: notification,
					Updates:      pushData,
				})
		}
	} else {
		var (
			pushExcludeList   = make([]int64, 0)
			serverIdKeyIdList = make(map[string][]int64)
		)

		statusList, _ := status_client.GetOnlineListByUser(ctx, userId)
		log.Debugf("statusList - #%v", statusList)
		for keyId, serverId := range statusList {
			if syncType == syncTypeUserNotMe && authKeyId == keyId {
				continue
			}
			pushExcludeList = append(pushExcludeList, keyId)
			if keyIdList, ok := serverIdKeyIdList[serverId]; ok {
				keyIdList = append(keyIdList, keyId)
				serverIdKeyIdList[serverId] = keyIdList
			} else {
				serverIdKeyIdList[serverId] = []int64{keyId}
			}
		}

		var ispush bool = false
		log.Debugf("serverIdKeyIdList - #%v", serverIdKeyIdList)
		for serverId, keyIdList := range serverIdKeyIdList {
			for _, keyId := range keyIdList {
				ispush = true
				s.PushUpdatesToSession(
					ctx,
					serverId,
					&sessionpb.PushUpdatesData{
						AuthKeyId:    keyId,
						Notification: notification,
						Updates:      pushData,
					})
			}
		}

		if !ispush && !notification {
			notification, _ = s.needActionNotify(ctx, userId, pushData)
		}

		if syncType == syncTypeUser {
			if model.IsBotFather(userId) {
				s.botsClient.PushBotUpdates(ctx, &botpb.BotUpdates{
					BotId:   userId,
					Updates: pushData,
				})
			} else if model.IsBotBing(userId) {
				s.bingClient.PushBotUpdates(ctx, &botpb.BotUpdates{
					BotId:   userId,
					Updates: pushData,
				})
			} else if model.IsBotGif(userId) {
				s.gifClient.PushBotUpdates(ctx, &botpb.BotUpdates{
					BotId:   userId,
					Updates: pushData,
				})
			} else if model.IsBotPic(userId) {
				s.picClient.PushBotUpdates(ctx, &botpb.BotUpdates{
					BotId:   userId,
					Updates: pushData,
				})
			} else if model.IsBotFoursquare(userId) {
				s.foursquareClient.PushBotUpdates(ctx, &botpb.BotUpdates{
					BotId:   userId,
					Updates: pushData,
				})
			} else if notification {
				s.PushClient.PushUpdatesIfNot(ctx, &pushpb.PushUpdatesIfNot{
					UserId:   userId,
					Excludes: pushExcludeList,
					Updates:  pushData,
				})
			}
		}
	}
}

func (s *Service) needActionNotify(ctx context.Context, userID int32, pushData *mtproto.Updates) (bool, error) {
	log.Debugf("needActionNotify - #%v", pushData)
	switch pushData.PredicateName {
	case mtproto.Predicate_updates:
		ups := pushData.To_Updates()
		for _, up := range ups.GetUpdates() {
			switch up.PredicateName {
			case mtproto.Predicate_updateBotWebhookJSON:
				data := &helper.DataJSON{DataJSON: up.GetData_DATAJSON()}
				var relaydata helper.TrelayData
				if data.GetJSONData(&relaydata) != nil {
					log.Errorf("needActionNotify %v", data)
					continue
				}
				if relaydata.Action == consts.ActionCallOnInvite {
					onInvite := handler.TAvOnInvite{}
					err := relaydata.GetData(&onInvite)
					if err != nil {
						log.Errorf("get TAvOnInvite %v,%s", relaydata, err.Error())
					} else {
						user, err := s.UserFacade.GetUserById(ctx, userID, int32(onInvite.From))
						if err != nil {
							log.Errorf("GetUserById %v,%s", relaydata, err.Error())
							continue
						}
						if user != nil {
							pushData.Users = append(pushData.Users, user)
						}
					}
					return true, nil
				}
			default:
			}
		}
	default:
	}
	return false, nil
}
