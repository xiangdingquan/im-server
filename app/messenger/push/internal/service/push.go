package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"open.chat/app/messenger/push/pushpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
	"open.chat/pkg/util"

	"open.chat/app/json/consts"
	"open.chat/app/json/helper"
	"open.chat/app/json/service/handler"
)

type PushMessage struct {
	UserId  string          `json:"user_id,omitempty"`
	Title   string          `json:"title,omitempty"`
	LocKey  string          `json:"loc_key,omitempty"`
	LocArgs []string        `json:"loc_args,omitempty"`
	Custom  *Custom         `json:"custom,omitempty"`
	Message string          `json:"message,omitempty"`
	Badge   int32           `json:"badge,omitempty"`
	Sound   string          `json:"sound,omitempty"`
	Update  *mtproto.Update `json:"-"`
}

type Custom struct {
	Dc           int32  `json:"dc,omitempty"`
	Addr         string `json:"addr,omitempty"`
	ChannelId    int32  `json:"channel_id,omitempty"`
	FromId       int32  `json:"from_id,omitempty"`
	ChatId       int32  `json:"chat_id,omitempty"`
	EncryptionId int32  `json:"encryption_id,omitempty"`
	ChatFromId   int32  `json:"chat_from_id,omitempty"`
	Mention      int32  `json:"mention,omitempty"`
	RandomId     int32  `json:"random_id,omitempty"`
	MsgId        int32  `json:"msg_id,omitempty"`
	MaxId        int32  `json:"max_id,omitempty"`
	Silent       int32  `json:"silent,omitempty"`
	EditDate     int32  `json:"edit_date,omitempty"`
	Schedule     int32  `json:"schedule,omitempty"`
	Messages     string `json:"messages,omitempty"`
	CallId       int64  `json:"call_id,omitempty"`
	CallAh       int64  `json:"call_ah,omitempty"`
	Attachb64    string `json:"attachb64,omitempty"`
}

func getUserNameById(id int32, userList []*mtproto.User) string {
	var me *mtproto.User
	for _, u := range userList {
		if u.Id == id {
			me = u
			break
		}
	}
	return model.GetUserName(me)
}

func makeNotificationMessage(userId int32, msg *mtproto.Message, locKey string, locArgs []string) *PushMessage {
	pushMsg := &PushMessage{
		UserId:  util.Int32ToString(userId),
		Title:   "notice",
		LocKey:  locKey,
		LocArgs: locArgs,
		Custom: &Custom{
			Dc:           0,
			Addr:         "",
			ChannelId:    0,
			FromId:       0,
			ChatId:       0,
			EncryptionId: 0,
			ChatFromId:   0,
			Mention:      0,
			RandomId:     0,
			MsgId:        0,
			MaxId:        0,
			Silent:       0,
			EditDate:     0,
			Schedule:     0,
			Messages:     "",
			CallId:       0,
			CallAh:       0,
		},
		Message: "",
		Badge:   1,
		Sound:   "0.m4a",
	}
	if msg != nil {
		custom := pushMsg.Custom
		custom.ChatFromId = msg.GetFromId_FLAGPEER().GetUserId()
		custom.Mention = int32(util.BoolToInt8(msg.Mentioned))
		custom.MsgId = msg.Id
		custom.Silent = int32(util.BoolToInt8(msg.Silent))
		custom.EditDate = msg.GetEditDate().GetValue()
		custom.Schedule = int32(util.BoolToInt8(msg.FromScheduled))
		pushMsg.Message = msg.Message
		switch msg.GetToId().GetPredicateName() {
		case mtproto.Predicate_peerUser:
			pushMsg.Custom.FromId = msg.GetFromId_FLAGPEER().GetUserId()
		case mtproto.Predicate_peerChat:
			pushMsg.Custom.ChatId = msg.ToId.ChatId
		case mtproto.Predicate_peerChannel:
			pushMsg.Custom.ChannelId = msg.ToId.ChannelId
		}
	}
	return pushMsg
}

func (s *Service) makePushMessageList2(ctx context.Context, userId int32, updates *mtproto.Updates) (notificationList []*PushMessage) {
	langCode := s.dao.GetUserLangCode(ctx, userId)
	title := "notice"
	if langCode != "en" {
		title = "通知"
	}
	var (
		// READ_HISTORY
		readHistoryF = func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			readHistory := &PushMessage{
				UserId: util.Int32ToString(userId),
				Title:  title,
				LocKey: "READ_HISTORY",
				Custom: &Custom{
					MaxId:     update.MaxId,
					ChannelId: update.ChannelId,
					FromId:    0,
					ChatId:    0,
				},
			}

			if update.Peer_PEER != nil {
				peer := update.Peer_PEER
				switch peer.PredicateName {
				case mtproto.Predicate_peerUser:
					readHistory.Custom.FromId = peer.UserId
				case mtproto.Predicate_peerChat:
					readHistory.Custom.ChatId = peer.ChatId
				}
			}
			notificationList = append(notificationList, readHistory)
		}

		messageDeletedF = func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			messageDeleted := &PushMessage{
				UserId: util.Int32ToString(userId),
				Title:  title,
				LocKey: "MESSAGE_DELETED",
				Custom: &Custom{
					ChannelId: update.ChannelId,
					Messages:  util.JoinInt32List(update.Messages, ","),
				},
			}

			if update.Peer_PEER != nil {
				peer := update.Peer_PEER
				switch peer.PredicateName {
				case mtproto.Predicate_peerUser:
					messageDeleted.Custom.FromId = peer.UserId
				case mtproto.Predicate_peerChat:
					messageDeleted.Custom.ChatId = peer.ChatId
				}
			}
			notificationList = append(notificationList, messageDeleted)
		}

		NewMessageF = func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			var newMessage *PushMessage
			inboxMsg := update.Message_MESSAGE
			switch inboxMsg.PredicateName {
			case mtproto.Predicate_messageService:
				action := inboxMsg.GetAction()
				switch action.GetPredicateName() {
				case mtproto.Predicate_messageActionContactSignUp:
				}
			case mtproto.Predicate_message:
				if inboxMsg.GetMedia() == nil {
					if inboxMsg.Message == "" {
						newMessage = makeNotificationMessage(
							userId,
							inboxMsg,
							"MESSAGE_NOTEXT",
							[]string{
								getUserNameById(inboxMsg.GetFromId_FLAGPEER().GetUserId(), users),
							})
					} else {
						cm := s.customNotification(inboxMsg.Message)
						if cm != "" {
							newMessage = makeNotificationMessage(
								userId,
								inboxMsg,
								"MESSAGE_TEXT",
								[]string{
									getUserNameById(inboxMsg.GetFromId_FLAGPEER().GetUserId(), users),
									cm,
								})
							newMessage.Message = cm
						}
					}
				} else {
					media := inboxMsg.GetMedia()
					newMessage = makeNotificationMessage(userId, inboxMsg, "", []string{
						getUserNameById(inboxMsg.GetFromId_FLAGPEER().GetUserId(), users),
					})
					if media.GetPredicateName() == mtproto.Predicate_messageMediaPhoto {
						newMessage.LocKey = "MESSAGE_PHOTO"
						newMessage.Message = "[photo]"
						if langCode != "en" {
							newMessage.Message = "[图片]"
						}
					} else if model.IsVideoMessage(inboxMsg) {
						newMessage.LocKey = "MESSAGE_VIDEO"
						newMessage.Message = "[video]"
						if langCode != "en" {
							newMessage.Message = "[视频]"
						}
					} else if model.IsGameMessage(inboxMsg) {

					} else if model.IsVoiceMessage(inboxMsg) {
						newMessage.LocKey = "MESSAGE_AUDIO"
						newMessage.Message = "[audio]"
						if langCode != "en" {
							newMessage.Message = "[语音]"
						}
					} else if model.IsRoundVideoMessage(inboxMsg) {

					} else if model.IsMusicMessage(inboxMsg) {
						newMessage.LocKey = "MESSAGE_MUSIC"
						newMessage.Message = "[music]"
						if langCode != "en" {
							newMessage.Message = "[音乐]"
						}
					} else if media.GetPredicateName() == mtproto.Predicate_messageMediaContact {

					} else if media.GetPredicateName() == mtproto.Predicate_messageMediaPoll {

					} else if media.GetPredicateName() == mtproto.Predicate_messageMediaGeo ||
						media.GetPredicateName() == mtproto.Predicate_messageMediaVenue {

					} else if media.GetPredicateName() == mtproto.Predicate_messageMediaGeoLive {

					} else if media.GetPredicateName() == mtproto.Predicate_messageMediaDocument {

					}
				}
			case mtproto.Predicate_messageEmpty:
			}
			if newMessage != nil {
				if chats != nil {
					newMessage.Message = newMessage.LocArgs[0] + ":" + newMessage.Message
					newMessage.LocArgs[0] = chats[0].Title
				}
				newMessage.Title = title
				notificationList = append(notificationList, newMessage)
			}
		}
	)

	model.VisitUpdates(userId, updates, map[string]model.UpdateVisitedFunc{
		mtproto.Predicate_updateDcOptions: func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			for _, dc := range update.DcOptions {
				notificationList = append(notificationList, &PushMessage{
					UserId: util.Int32ToString(userId),
					Title:  title,
					LocKey: "DC_UPDATE",
					Custom: &Custom{
						Dc:        dc.Id,
						Addr:      fmt.Sprintf("%s:%d", dc.IpAddress, dc.Port),
						ChannelId: dc.Port,
					},
				})
			}
		},

		mtproto.Predicate_updateServiceNotification: func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			notificationList = append(notificationList, &PushMessage{
				UserId:  util.Int32ToString(userId),
				Title:   title,
				LocKey:  "MESSAGE_ANNOUNCEMENT",
				Message: update.Message_STRING,
			})

		},

		"SESSION_REVOKE": func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			notificationList = append(notificationList, &PushMessage{
				UserId: util.Int32ToString(userId),
				Title:  title,
				LocKey: "SESSION_REVOKE",
			})
		},
		mtproto.Predicate_updateGeoLiveViewed: func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			notificationList = append(notificationList, &PushMessage{
				UserId: util.Int32ToString(userId),
				Title:  title,
				LocKey: "GEO_LIVE_PENDING",
			})
		},

		mtproto.Predicate_updateReadHistoryInbox: readHistoryF,
		mtproto.Predicate_updateReadChannelInbox: readHistoryF,

		mtproto.Predicate_updateDeleteMessages:        messageDeletedF,
		mtproto.Predicate_updateDeleteChannelMessages: messageDeletedF,
		mtproto.Predicate_updateNewMessage:            NewMessageF,
		mtproto.Predicate_updateNewChannelMessage:     NewMessageF,

		mtproto.Predicate_updateNewEncryptedMessage: func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			newMessage := makeNotificationMessage(userId, nil, "MESSAGE_TEXT", []string{})
			newMessage.Message = "[Received a new message]"
			if langCode != "en" {
				newMessage.Title = "通知"
				newMessage.Message = "[收到一条新消息]"
			}
			notificationList = append(notificationList, newMessage)
		},

		mtproto.Predicate_updateEncryption: func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			newMessage := makeNotificationMessage(userId, nil, "MESSAGE_TEXT", []string{})
			if update.Chat.GetParticipantId() == userId {
				if langCode != "en" {
					newMessage.Message = "【" + getUserNameById(update.Chat.GetAdminId(), users) + " 邀请你加入私密聊天】"
				} else {
					newMessage.Message = "[" + getUserNameById(update.Chat.GetAdminId(), users) + " Invite you to join private chat]"
				}
			} else {
				if langCode != "en" {
					newMessage.Message = "【" + getUserNameById(update.Chat.GetParticipantId(), users) + " 已加入私密聊天】"
				} else {
					newMessage.Message = "[" + getUserNameById(update.Chat.GetParticipantId(), users) + " Joined private chat]"
				}
			}
			newMessage.Title = title
			notificationList = append(notificationList, newMessage)
		},

		mtproto.Predicate_updatePhoneCall: func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {

			log.Debugf("updatePhoneCall: %s", update.DebugString())
			phoneCall := update.PhoneCall
			switch phoneCall.PredicateName {
			case mtproto.Predicate_phoneCallRequested:
				notificationList = append(notificationList, &PushMessage{
					UserId:  util.Int32ToString(userId),
					Title:   title,
					LocKey:  "PHONE_CALL_REQUEST",
					LocArgs: []string{getUserNameById(phoneCall.AdminId, users)},
					Custom: &Custom{
						FromId: phoneCall.AdminId,
						CallId: phoneCall.Id,
						CallAh: phoneCall.AccessHash,
					},
					Update: update,
				})
			}
		},

		mtproto.Predicate_updatePeerSettings: func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			notificationList = append(notificationList, &PushMessage{
				Title:  title,
				LocKey: "CONTACT_ADDED",
				Update: update,
			})
		},

		mtproto.Predicate_updateBotWebhookJSON: func(
			userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {
			data := &helper.DataJSON{DataJSON: update.GetData_DATAJSON()}
			var relaydata helper.TrelayData
			var err error
			if err = data.GetJSONData(&relaydata); err != nil {
				log.Errorf("Updates updateBotWebhookJSON %v,%v", err, data)
			} else {
				if relaydata.Action == consts.ActionCallOnInvite {
					var v handler.TAvOnInvite
					if err = relaydata.GetData(&v); err != nil {
						log.Errorf("onTAvOnInvite %v,%v,", err, data)
					} else {
						userName := ""
						for _, user := range users {
							if user.GetId() == int32(v.From) {
								userName = user.GetFirstName().GetValue()
							}
						}
						msg := userName + " Invite you to voice call"
						if langCode != "en" {
							msg = userName + "邀请你语音通话"
						}
						if v.IsVideo {
							msg = userName + "Invite you to video call"
							if langCode != "en" {
								msg = userName + "邀请你视频通话"

							}
						} else if v.IsMeetingAV {
							msg = userName + "Voice call initiated"
							if langCode != "en" {
								msg = userName + "发起了语音通话"
							}
						}

						newMessage := &PushMessage{
							UserId: util.Int32ToString(userId),
							Title:  title,
							LocKey: "MESSAGE_TEXT",
							LocArgs: []string{
								"Single chat",
								msg,
							},
							Custom: &Custom{
								FromId:     int32(v.From),
								ChatId:     int32(v.ChatID),
								ChatFromId: int32(v.From),
							},
							Message: msg,
							Badge:   1,
							Sound:   "incoming.caf",
						}
						if langCode != "en" {
							newMessage.LocArgs[0] = "单聊"
						}
						if v.IsMeetingAV {
							newMessage.LocArgs[0] = "Group chat"
							if langCode != "en" {
								newMessage.LocArgs[0] = "群聊"
							}
						}
						notificationList = append(notificationList, newMessage)
					}
				}
			}
		},
	})

	return
}

func Base64URLDecode(data string) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	res, err := base64.URLEncoding.DecodeString(data)
	fmt.Println("Base64URLDecode is :", string(res), err)
	return base64.URLEncoding.DecodeString(data)
}

func Base64UrlSafeEncode(source []byte) string {
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}

func encodeNotificationMessageData(key *crypto.AuthKey, pushMsg *PushMessage) (pData string) {
	jBuf, err := json.Marshal(pushMsg)
	if err != nil {
		log.Errorf("json marshal error: %v", err)
		return
	}
	x2 := mtproto.NewEncodeBuf(4 + len(jBuf))
	x2.Int(int32(len(jBuf)))
	x2.Bytes(jBuf)
	msgKey, data, err := key.AesIgeEncrypt(x2.GetBuf())
	if err != nil {
		log.Errorf("aesIgeEncrypt data error: %v", err)
		return
	}

	x := mtproto.NewEncodeBuf(len(msgKey) + len(data) + 8)
	x.Long(key.AuthKeyId())
	x.Bytes(msgKey)
	x.Bytes(data)
	pData = Base64UrlSafeEncode(x.GetBuf())
	return
}

func (s *Service) onPushUpdatesIfNot(ctx context.Context, r *pushpb.PushUpdatesIfNot) error {
	log.Debugf("onPushUpdatesIfNot - request: {%s}", logger.JsonDebugData(r))

	pushMsgList := s.makePushMessageList2(ctx, r.UserId, r.Updates)
	for _, pushMsg := range pushMsgList {
		jBuf, _ := json.Marshal(pushMsg)
		log.Debugf("onPushUpdatesIfNot - pushMsg: {%s}", string(jBuf))
		x2 := mtproto.NewEncodeBuf(4 + len(jBuf))
		x2.Int(int32(len(jBuf)))
		x2.Bytes(jBuf)

		s.dao.WalkDevices(ctx,
			r.UserId,
			r.Excludes,
			pushMsg,
			func(ctx context.Context, pushAuthKeyId int64, tokenType int8, token string, secret []byte, pushMsg interface{}) {
				m := pushMsg.(*PushMessage)
				if m.LocKey == "CONTACT_ADDED" || m.LocKey == "PHONE_CALL_REQUEST" {
					s.dao.AddSeqToUpdatesQueue(ctx, pushAuthKeyId, r.UserId, 0, model.TLObjectToJson(m.Update))
				}

				switch tokenType {
				case model.PushTypeAPNS:
					s.onPushAPNS(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeFCM:
					s.onPushFCM(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeMPNS:
					s.onPushMPNS(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeSimplePush:
					s.onPushDeprecated(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeUbuntuPhone:
					s.onPushUbuntu(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeBlackberry:
					s.onPushBlackberry(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeInternealPush:
					s.onPushMTProto(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeWindowsPush:
					s.onPushWNS(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeAPNSVoip:
					s.onPushAPNSVoIP(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeWebPush:
					s.onPushWeb(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeMPNSVoip:
					s.onPushMPNSVoIP(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeTizen:
					s.onPushTizen(ctx, pushAuthKeyId, token, secret, m)
				case model.PushTypeJPush:
					s.onPushJPush(ctx, pushAuthKeyId, token, secret, m)
				default:
					log.Errorf("invalid tokenType %d", tokenType)
					return
				}
			})
	}
	return nil
}

func (s *Service) onPushUpdates(ctx context.Context, r *pushpb.PushUpdates) error {
	log.Debugf("onPushUpdates - request: {%s}", logger.JsonDebugData(r))
	return nil
}
