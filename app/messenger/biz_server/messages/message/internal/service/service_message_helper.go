package service

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/mvdan/xurls"

	"open.chat/app/json/services/handler/chats/core"
	sync_client "open.chat/app/messenger/sync/client"
	media_client "open.chat/app/service/media/client"
	"open.chat/app/service/media/mediapb"
	"open.chat/app/sysconfig"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
	"open.chat/pkg/mention"
	"open.chat/pkg/phonenumber"
	"open.chat/pkg/util"
)

func (s *Service) doClearDraft(ctx context.Context, userId int32, authKeyId int64, peer *model.PeerUtil) {
	var hasClearDraft bool
	switch peer.PeerType {
	case model.PEER_USER:
		s.PrivateFacade.ClearDraftMessage(ctx, userId, peer.PeerId)
	case model.PEER_CHAT:
		s.ChatFacade.ClearDraftMessage(ctx, userId, peer.PeerId)
	case model.PEER_CHANNEL:
		s.ChannelFacade.ClearDraftMessage(ctx, userId, peer.PeerId)
	default:
	}

	if hasClearDraft {
		updateDraftMessage := mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
			Peer_PEER: peer.ToPeer(),
			Draft:     mtproto.MakeTLDraftMessageEmpty(nil).To_DraftMessage(),
		})

		updates := mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{updateDraftMessage.To_Update()},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		})

		sync_client.SyncUpdatesNotMe(ctx, userId, authKeyId, updates.To_Updates())
	}
}

func (s *Service) makeMediaByInputMedia(ctx context.Context, userId int32, authKeyId int64, peer *model.PeerUtil, media *mtproto.InputMedia) (messageMedia *mtproto.MessageMedia, err error) {
	var (
		now = int32(time.Now().Unix())
	)

	switch media.PredicateName {
	case mtproto.Predicate_inputMediaUploadedPhoto:
		if (peer != nil) && (peer.PeerType == model.PEER_CHAT || peer.PeerType == model.PEER_CHANNEL) {
			if sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysBanImages, false, 0) {
				return nil, errors.New(string(sysconfig.ConfigKeysBanImages))
			}
		}

		var result *mediapb.PhotoDataRsp
		result, err = media_client.UploadPhotoFile(authKeyId, media.File)
		if err != nil {
			log.Errorf("UploadPhoto error: %v, by %s", err, media.DebugString())
			return
		}

		photo := mtproto.MakeTLPhoto(&mtproto.Photo{
			Id:          result.PhotoId,
			HasStickers: len(media.Stickers) > 0,
			AccessHash:  result.AccessHash,
			Date:        now,
			Sizes:       result.SizeList,
			DcId:        2,
		})

		messageMedia = mtproto.MakeTLMessageMediaPhoto(&mtproto.MessageMedia{
			Photo_FLAGPHOTO: photo.To_Photo(),
			TtlSeconds:      media.TtlSeconds,
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaPhoto:
		if (peer != nil) && (peer.PeerType == model.PEER_CHAT || peer.PeerType == model.PEER_CHANNEL) {
			if sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysBanImages, false, 0) {
				return nil, errors.New(string(sysconfig.ConfigKeysBanImages))
			}
		}

		mediaPhoto := media.To_InputMediaPhoto()
		sizeList, _ := media_client.GetPhotoSizeList(mediaPhoto.GetId_INPUTPHOTO().GetId())

		photo := mtproto.MakeTLPhoto(&mtproto.Photo{
			Id:          mediaPhoto.GetId_INPUTPHOTO().GetId(),
			HasStickers: false,
			AccessHash:  mediaPhoto.GetId_INPUTPHOTO().GetAccessHash(),
			Date:        now,
			Sizes:       sizeList,
			DcId:        2,
		})

		messageMedia = mtproto.MakeTLMessageMediaPhoto(&mtproto.MessageMedia{
			Photo_FLAGPHOTO: photo.To_Photo(),
			TtlSeconds:      mediaPhoto.GetTtlSeconds(),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaGeoPoint:
		messageMedia = mtproto.MakeTLMessageMediaGeo(&mtproto.MessageMedia{
			Geo: model.MakeGeoPointByInput(media.To_InputMediaGeoPoint().GetGeoPoint()),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaContact:
		contact := media.To_InputMediaContact()
		var phoneNumber string = contact.GetPhoneNumber()
		reqJson := sendContactByPhone{}
		err = json.Unmarshal([]byte(phoneNumber), &reqJson)
		useJson := err == nil
		messageMedia = mtproto.MakeTLMessageMediaContact(&mtproto.MessageMedia{
			PhoneNumber: phoneNumber,
			FirstName:   contact.GetFirstName(),
			LastName:    contact.GetLastName(),
			Vcard:       contact.GetVcard(),
			UserId:      0,
		}).To_MessageMedia()
		if useJson {
			messageMedia.PhoneNumber = ""
			contactUser, _ := s.UserFacade.GetUserById(ctx, userId, int32(reqJson.Uid))
			if contactUser != nil {
				messageMedia.UserId = contactUser.Id
			}
		} else {
			phoneNumber, err := phonenumber.CheckAndGetPhoneNumber(contact.GetPhoneNumber())
			if err == nil {
				contactUser, _ := s.UserFacade.GetUserSelfByPhoneNumber(ctx, phoneNumber)
				if contactUser != nil {
					messageMedia.UserId = contactUser.Id
				}
			}
		}
		log.Errorf("reqJson:%v,%v,%s", useJson, reqJson, logger.JsonDebugData(messageMedia))
	case mtproto.Predicate_inputMediaUploadedDocument:
		uploadedDocument := media.To_InputMediaUploadedDocument()
		documentMedia, _ := media_client.UploadedDocumentMedia(authKeyId, uploadedDocument)
		if documentMedia == nil {
			err = mtproto.ErrMediaInvalid
			return
		}
		messageMedia = documentMedia.To_MessageMedia()
	case mtproto.Predicate_inputMediaDocument:
		id := media.To_InputMediaDocument().GetId_INPUTDOCUMENT()
		document3, _ := media_client.GetDocumentById(id.GetId(), id.GetAccessHash())
		messageMedia = mtproto.MakeTLMessageMediaDocument(&mtproto.MessageMedia{
			Document:   document3,
			TtlSeconds: media.To_InputMediaDocument().GetTtlSeconds(),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaVenue:
		venue := media.To_InputMediaVenue()
		messageMedia = mtproto.MakeTLMessageMediaVenue(&mtproto.MessageMedia{
			Geo:       model.MakeGeoPointByInput(venue.GetGeoPoint()),
			Title:     venue.GetTitle(),
			Address:   venue.GetAddress(),
			Provider:  venue.GetProvider(),
			VenueId:   venue.GetVenueId(),
			VenueType: venue.GetVenueType(),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaGifExternal:
		messageMedia = mtproto.MakeTLMessageMediaUnsupported(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaDocumentExternal:
		messageMedia = mtproto.MakeTLMessageMediaUnsupported(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaPhotoExternal:
		messageMedia = mtproto.MakeTLMessageMediaUnsupported(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaGame:
		messageMedia = mtproto.MakeTLMessageMediaUnsupported(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaInvoice:
		messageMedia = mtproto.MakeTLMessageMediaUnsupported(nil).To_MessageMedia()
	case mtproto.Predicate_inputMediaGeoLive:
		messageMedia = mtproto.MakeTLMessageMediaGeoLive(&mtproto.MessageMedia{
			Geo:    model.MakeGeoPointByInput(media.To_InputMediaGeoLive().GetGeoPoint()),
			Period: media.To_InputMediaGeoLive().GetPeriod_INT32(),
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaPoll:
		messageMedia = mtproto.MakeTLMessageMediaPoll(&mtproto.MessageMedia{
			Poll:    media.Poll,
			Results: nil,
		}).To_MessageMedia()
	case mtproto.Predicate_inputMediaDice:
		if media.Emoticon == "🎲" {
			messageMedia = mtproto.MakeTLMessageMediaDice(&mtproto.MessageMedia{
				Value:    rand.Int31()%6 + 1,
				Emoticon: media.Emoticon,
			}).To_MessageMedia()
		} else if media.Emoticon == "🎯" {
			messageMedia = mtproto.MakeTLMessageMediaDice(&mtproto.MessageMedia{
				Value:    rand.Int31()%6 + 1,
				Emoticon: media.Emoticon,
			}).To_MessageMedia()
		} else if media.Emoticon == "🏀" {
			messageMedia = mtproto.MakeTLMessageMediaDice(&mtproto.MessageMedia{
				Value:    rand.Int31()%5 + 1,
				Emoticon: media.Emoticon,
			}).To_MessageMedia()
		} else {
			messageMedia = mtproto.MakeTLMessageMediaDice(&mtproto.MessageMedia{
				Value:    rand.Int31()%6 + 1,
				Emoticon: media.Emoticon,
			}).To_MessageMedia()
		}
	case mtproto.Predicate_inputMediaEmpty:
		err = mtproto.ErrMediaEmpty
	default:
		err = mtproto.ErrMediaInvalid
	}
	return
}

func (s *Service) fixMessageEntities(ctx context.Context, fromId int32, peer *model.PeerUtil, noWebpage bool, message *mtproto.Message, hasBot bool) (*mtproto.Message, error) {
	var (
		entities mtproto.MessageEntitySlice
		idxList  []int
	)
	getIdxList := func() []int {
		if len(idxList) == 0 {
			idxList = mention.EncodeStringToUTF16Index(message.Message)
		}
		return idxList
	}

	//替换关键字
	if peer.PeerType == model.PEER_CHAT || peer.PeerType == model.PEER_CHANNEL {
		words := sysconfig.GetConfig2StringArray(ctx, sysconfig.ConfigKeysBanWords, nil, 0)
		replaces := make([]string, 0)
		for _, w := range words {
			if len(w) > 0 && strings.Contains(message.Message, w) {
				replaces = append(replaces, w)
			}
		}

		for _, r := range replaces {
			s := ""
			for i := 0; i < len(r); i++ {
				s += "*"
			}
			message.Message = strings.Replace(message.Message, r, s, -1)
		}
	}

	var firstUrl string
	rIndexes := xurls.Relaxed().FindAllStringIndex(message.Message, -1)
	if len(rIndexes) > 0 {
		if len(idxList) == 0 {
			getIdxList()
		}
		if (peer.PeerType == model.PEER_CHAT || peer.PeerType == model.PEER_CHANNEL) && sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysBanLinks, false, 0) {
			urls := make([]string, 0)
			for i := len(rIndexes) - 1; i >= 0; i-- {
				v := rIndexes[i]
				urls = append(urls, message.Message[v[0]:v[1]])
			}
			for _, r := range urls {
				message.Message = strings.Replace(message.Message, r, s.getIllegalMessage(ctx), 1)
			}
		} else {
			for idx, v := range rIndexes {
				if idx == 0 {
					firstUrl = message.Message[v[0]:v[1]]
				}
				entityUrl := mtproto.MakeTLMessageEntityUrl(&mtproto.MessageEntity{
					Offset: int32(idxList[v[0]]),
					Length: int32(idxList[v[1]] - idxList[v[0]]),
				})
				entities = append(entities, entityUrl.To_MessageEntity())
			}
		}
	}

	if !noWebpage && firstUrl != "" {
		canEmbedLink := true
		if canEmbedLink {
			if webpage, err := s.GetWebPagePreview(ctx, &mtproto.TLMessagesGetWebPagePreview{
				Message: firstUrl,
			}); err != nil {
			} else {
				message.Media = mtproto.MakeTLMessageMediaWebPage(&mtproto.MessageMedia{
					Webpage: webpage,
				}).To_MessageMedia()
			}
		}
	}

	for _, entity := range message.Entities {
		switch entity.PredicateName {
		case mtproto.Predicate_inputMessageEntityMentionName:
			entityMentionName := mtproto.MakeTLMessageEntityMentionName(&mtproto.MessageEntity{
				Offset:       entity.Offset,
				Length:       entity.Length,
				UserId_INT32: entity.UserId_INPUTUSER.UserId,
			})
			entities = append(entities, entityMentionName.To_MessageEntity())
		default:
			entities = append(entities, entity)
		}
	}

	tags := mention.GetTags('@', message.Message, '(', ')', 'm', 'M')
	if len(tags) > 0 {
		var nameList = make([]string, 0, len(tags))
		for _, tag := range tags {
			nameList = append(nameList, tag.Tag)
		}
		names, _ := s.UsernameFacade.GetListByUsernameList(ctx, nameList)
		log.Debugf("nameList: %v", names)

		for _, tag := range tags {
			if len(idxList) == 0 {
				getIdxList()
			}
			mention2 := mtproto.MakeTLMessageEntityMention(&mtproto.MessageEntity{
				Offset: int32(idxList[tag.Index]),
				Length: int32(idxList[tag.Index+len(tag.Tag)+1] - idxList[tag.Index]),
			}).To_MessageEntity()

			if uname, ok := names[tag.Tag]; ok {
				if uname.PeerType == model.PEER_USER {
					mention2.UserId_INT32 = uname.PeerId
				}
			}
			entities = append(entities, mention2)
			log.Infof("mention2: %v", mention2)
		}
	}

	tags = mention.GetTags('#', message.Message)
	for _, tag := range tags {
		if len(idxList) == 0 {
			getIdxList()
		}
		hashtag := mtproto.MakeTLMessageEntityHashtag(&mtproto.MessageEntity{
			Offset: int32(idxList[tag.Index]),
			Length: int32(idxList[tag.Index+len(tag.Tag)+1] - idxList[tag.Index]),
		}).To_MessageEntity()
		entities = append(entities, hashtag)
	}

	if hasBot {
		tags = mention.GetTags('/', message.Message)
		for _, tag := range tags {
			if len(idxList) == 0 {
				getIdxList()
			}
			hashtag := mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
				Offset: int32(idxList[tag.Index]),
				Length: int32(idxList[tag.Index+len(tag.Tag)+1] - idxList[tag.Index]),
			}).To_MessageEntity()
			entities = append(entities, hashtag)
		}
	}

	sort.Sort(entities)
	message.Entities = entities
	return message, nil
}

func (s *Service) checkUserSendRight(ctx context.Context, userId int32, selfId int32, authKeyId int64) error {
	//log.Debugf("checkUserSendRight, userId:%d, selfId:%d", userId, selfId)
	if userId == selfId {
		return nil
	}

	checkWhisper := func() error {
		contact, mutual := s.UserFacade.GetContactAndMutual(ctx, selfId, userId)
		if !contact || !mutual {
			platform, err := s.AuthFacade.GetPlatform(ctx, authKeyId)
			if err != nil {
				return err
			}
			//log.Debugf("checkUserSendRight, platform:%d", platform)
			// define at gochat/app/services/biz_server/auth/core
			if platform == 1 || platform == 2 {
				log.Debugf("checkUserSendRight, is ios or android, skip whisper checking")
				return nil
			}
			var (
				banWhisper bool = false
			)
			chatsCore := core.New(nil)
			chatIdList := s.ChatFacade.GetUsersChatIdList(ctx, []int32{selfId, userId})
			if len(chatIdList) == 2 {
				commonChats := util.Int32Intersect(chatIdList[selfId], chatIdList[userId])
				//log.Debugf("banned whisper - chat, selfId:%d, userId:%d, commonChats:%v", selfId, userId, commonChats)
				banWhisper = len(commonChats) > 0
				for _, cid := range commonChats {
					bannedRights := chatsCore.GetChatBannedRights(ctx, uint32(cid))
					if !bannedRights.BanWhisper {
						banWhisper = false
						break
					}
				}
			}
			if banWhisper {
				log.Debugf("banned whisper - chat, selfId:%d, userId:%d", selfId, userId)
				return mtproto.ErrUserIdInvalid
			}

			chatIdList = s.ChannelFacade.GetUsersChannelIdList(ctx, []int32{selfId, userId})
			if len(chatIdList) == 2 {
				commonChats := util.Int32Intersect(chatIdList[selfId], chatIdList[userId])
				//log.Debugf("banned whisper - channel, selfId:%d, userId:%d, commonChats:%v", selfId, userId, commonChats)
				banWhisper = len(commonChats) > 0
				for _, cid := range commonChats {
					bannedRights := chatsCore.GetChannelBannedRights(ctx, uint32(cid))
					if !bannedRights.BanWhisper {
						banWhisper = false
						break
					}
				}
			}
			if banWhisper {
				log.Debugf("banned whisper - channel, selfId:%d, userId:%d", selfId, userId)
				return mtproto.ErrUserIdInvalid
			}
		}
		return nil
	}
	err := checkWhisper()
	if err != nil {
		return err
	}

	isHisFriend, _ := s.UserFacade.GetContactAndMutual(ctx, userId, selfId)
	if canSend := s.UserFacade.CheckPrivacy(ctx, model.SEND_MESSAGES, userId, selfId, isHisFriend); !canSend {
		log.Debugf("[UserPrivacy_SendMessage], can not send message, selfId:%d, userId:%d", selfId, userId)
		return mtproto.ErrUserPrivacyRestricted
	}
	return nil
}

func (s *Service) checkChatSendRight(ctx context.Context, chatId int32, selfId int32, media *mtproto.InputMedia) error {
	chat, err := s.ChatFacade.GetMutableChat(ctx, chatId, selfId)
	if err != nil {
		log.Errorf("checkChatSendRight error - %s: (%d - %d)", err, chatId, selfId)
		return mtproto.ErrChatIdInvalid
	}

	me := chat.GetImmutableChatParticipant(selfId)
	if me == nil {
		log.Errorf("GetImmutableChatParticipant error - %s: (%d - %d)", err, chatId, selfId)
		return mtproto.ErrChannelPrivate
	}

	if media == nil {
		if !me.CanSendMessages() {
			return mtproto.ErrChatRestricted
		}
	} else {
		if !me.CanSendMedia() {
			return mtproto.ErrChatSendMediaForbidden
		}
		//if !me.CanSendGifs() {
		//	return mtproto.ErrChatSendGifsForbidden
		//}
		//if !me.CanSendStickers() {
		//	return mtproto.ErrChatSendStickersForbidden
		//}
	}

	return nil
}

func (s *Service) checkChannelSendRight(ctx context.Context, channelId int32, selfId int32, media *mtproto.InputMedia) error {
	channel, err := s.ChannelFacade.GetMutableChannel(ctx, channelId, selfId)
	if err != nil {
		log.Errorf("checkChannelSendRight error - %s: (%d - %d)", err, channelId, selfId)
		return mtproto.ErrChannelInvalid
	}

	me := channel.GetImmutableChannelParticipant(selfId)
	if me == nil {
		log.Errorf("GetImmutableChannelParticipant error - %s: (%d - %d)", err, channelId, selfId)
		return mtproto.ErrChannelPrivate
	}

	if media == nil {
		if !me.CanSendMessages(math.MaxInt32) {
			return mtproto.ErrChatRestricted
		}
	} else {
		if !me.CanSendMedia(math.MaxInt32) {
			return mtproto.ErrChatSendMediaForbidden
		}
		//if !me.CanSendGifs(math.MaxInt32) {
		//	return mtproto.ErrChatSendGifsForbidden
		//}
		//if !me.CanSendStickers(math.MaxInt32) {
		//	return mtproto.ErrChatSendStickersForbidden
		//}
	}

	return nil
}

func (s *Service) isContainBanWord(ctx context.Context, peer *model.PeerUtil, message string) (bool, error) {
	log.Debugf("isContainBanWord, message:%s", message)
	var words []string
	var err error

	if peer.PeerType == model.PEER_CHAT {
		words, err = s.ChatFacade.GetFilterKeywords(ctx, uint32(peer.PeerId))
		if err != nil {
			return false, err
		}
	} else if peer.PeerType == model.PEER_CHANNEL {
		words, err = s.ChannelFacade.GetFilterKeywords(ctx, uint32(peer.PeerId))
		if err != nil {
			return false, err
		}
	}

	log.Debugf("isContainBanWord, words:%v", words)
	for _, w := range words {
		if len(w) > 0 && strings.Contains(message, w) {
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) getIllegalMessage(ctx context.Context) string {
	return model.Localize(ctx, s.AuthSessionRpcClient, model.LocalizationWords{
		model.LocalizationEN:      "[Illegal image, blocked]",
		model.LocalizationCN:      "[违规发图片,被屏蔽]",
		model.LocalizationDefault: "[Illegal image, blocked]",
	})
}
