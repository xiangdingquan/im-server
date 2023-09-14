package core

import (
	"context"
	"math"
	"math/rand"
	"time"

	"encoding/json"

	"github.com/gogo/protobuf/types"
	"open.chat/app/service/biz_service/channel/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

const (
	createChatFlood = 10 // 10s
)

func makeImmutableChannel(do *dataobject.ChannelsDO) (channel *model.ImmutableChannel) {
	channel = &model.ImmutableChannel{
		Id:                   do.Id,
		AccessHash:           do.AccessHash,
		SecretKeyId:          do.SecretKeyId,
		Title:                do.Title,
		Username:             do.Username,
		Photo:                nil,
		CreatorId:            do.CreatorUserId,
		TopMessage:           do.TopMessage,
		ReadOutboxMaxId:      do.ReadOutboxMaxId,
		Broadcast:            util.Int8ToBool(do.Broadcast),
		Democracy:            util.Int8ToBool(do.Democracy),
		Verified:             util.Int8ToBool(do.Verified),
		Megagroup:            util.Int8ToBool(do.Megagroup),
		Signatures:           util.Int8ToBool(do.Signatures),
		Min:                  false,
		Scam:                 false,
		HasLink:              util.Int8ToBool(do.HasLink),
		HasGeo:               util.Int8ToBool(do.HasGeo),
		SlowmodeEnabled:      util.Int8ToBool(do.SlowmodeEnabled),
		Date:                 do.Date,
		Version:              do.Version,
		DefaultBannedRights:  model.ChatBannedRights{Rights: do.DefaultBannedRights, UntilDate: math.MaxInt32},
		ParticipantsCount:    do.ParticipantsCount,
		About:                do.About,
		Notice:               do.Notice,
		HiddenPrehistory:     util.Int8ToBool(do.PreHistoryHidden),
		AdminsCount:          do.AdminsCount,
		KickedCount:          do.KickedCount,
		BannedCount:          do.BannedCount,
		OnlineCount:          0,
		ChatPhoto:            mtproto.MakeTLPhotoEmpty(nil).To_Photo(),
		Link:                 do.Link,
		BotInfo:              nil,
		MigratedFromChatId:   do.MigratedFromChatId,
		MigratedFromMaxId:    0,
		PinnedMsgId:          do.PinnedMsgId,
		StickerSet:           nil,
		LinkedChatId:         do.LinkedChatId,
		Location:             nil,
		SlowmodeSeconds:      do.SlowmodeSeconds,
		SlowmodeNextSendDate: 0,
		Pts:                  do.Pts,
		Deleted:              util.Int8ToBool(do.Deleted),
	}

	if channel.HasGeo {
		channel.Location = mtproto.MakeTLChannelLocation(&mtproto.ChannelLocation{
			GeoPoint: mtproto.MakeTLGeoPoint(&mtproto.GeoPoint{
				Lat:  do.Lat,
				Long: do.Long,
			}).To_GeoPoint(),
			Address: do.Address,
		}).To_ChannelLocation()
	}

	if do.Photo != "" {
		json.Unmarshal(hack.Bytes(do.Photo), channel.ChatPhoto)
	}
	channel.Photo = model.MakeChatPhotoByPhoto(channel.ChatPhoto)

	return
}

func makeImmutableChannelParticipants(channel *model.ImmutableChannel, doList []dataobject.ChannelParticipantsDO) (participants map[int32]*model.ImmutableChannelParticipant) {
	participants = make(map[int32]*model.ImmutableChannelParticipant, len(doList))
	for i := 0; i < len(doList); i++ {
		participants[doList[i].UserId] = makeImmutableChannelParticipant(channel, &doList[i])
	}
	return
}

func makeImmutableChannelParticipant(channel *model.ImmutableChannel, do *dataobject.ChannelParticipantsDO) (participant *model.ImmutableChannelParticipant) {
	participant = &model.ImmutableChannelParticipant{
		Channel:             channel,
		Creator:             util.Int8ToBool(do.IsCreator),
		State:               int(do.State),
		Id:                  do.Id,
		UserId:              do.UserId,
		Date:                do.Date2,
		InviterId:           do.InviterUserId,
		Rank:                "",
		CanEdit:             false,
		PromotedBy:          do.PromotedBy,
		AdminRights:         model.ChatAdminRights(do.AdminRights),
		KickedBy:            do.KickedBy,
		BannedRights:        model.ChatBannedRights{Rights: do.BannedRights, UntilDate: do.BannedUntilDate},
		Pinned:              util.Int8ToBool(do.IsPinned),
		UnreadMark:          util.Int8ToBool(do.UnreadMark),
		ChannelId:           do.ChannelId,
		TopMessage:          channel.TopMessage,
		ReadInboxMaxId:      do.ReadInboxMaxId,
		UnreadCount:         channel.TopMessage - do.ReadInboxMaxId,
		UnreadMentionsCount: do.UnreadMentionsCount,
		NotifySettings:      nil,
		Draft:               nil,
		FolderId:            do.FolderId,
		AvailableMinId:      do.AvailableMinId,
		AvailableMinPts:     do.AvailableMinPts,
		MigratedFromMaxId:   do.MigratedFromMaxId,
		Nickname:            do.Nickname,
	}
	return
}

func makeMutableChannel(do *dataobject.ChannelsDO, doList []dataobject.ChannelParticipantsDO) *model.MutableChannel {
	channel := makeImmutableChannel(do)
	participants := makeImmutableChannelParticipants(channel, doList)

	return &model.MutableChannel{
		Channel:      channel,
		Participants: participants,
	}
}

func (m *ChannelCore) GetMutableChannel(ctx context.Context, channelId int32, id ...int32) (channel *model.MutableChannel, err error) {
	channel = new(model.MutableChannel)

	channelDO, err := m.ChannelsDAO.Select(ctx, channelId)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	} else if channelDO == nil {
		return nil, mtproto.ErrChannelInvalid
	}

	channel.Channel = makeImmutableChannel(channelDO)
	if len(id) > 0 {
		doList, err := m.ChannelParticipantsDAO.SelectByUserIdList(ctx, channelId, id)
		if err != nil {
			return nil, mtproto.ErrInternelServerError
		}
		channel.Participants = makeImmutableChannelParticipants(channel.Channel, doList)
	}

	return
}

func (m *ChannelCore) CreateChannel(ctx context.Context, creatorId int32, secretKeyId int64, isBroadcast bool, title, about, notice string, geo *mtproto.InputGeoPoint, address *types.StringValue, randomId int64) (*model.MutableChannel, error) {
	var (
		channelDO *dataobject.ChannelsDO
		date      = int32(time.Now().Unix())
	)

	// 2. check title
	if title == "" {
		err := mtproto.ErrChatTitleEmpty
		log.Error("title empty")
		return nil, err
	}

	// check flood
	if channelDO, err := m.ChannelsDAO.SelectLastCreator(ctx, creatorId); err != nil {
		return nil, err
	} else if channelDO != nil {
		if date-channelDO.Date < createChatFlood {
			err = mtproto.NewErrFloodWaitX(date - channelDO.Date)
			log.Errorf("createChannel error: %v. lastCreate = ", err, channelDO.Date)
			return nil, err
		}
	}

	hasGeo := geo != nil
	channelDO = &dataobject.ChannelsDO{
		Id:                0,
		CreatorUserId:     creatorId,
		AccessHash:        rand.Int63(),
		SecretKeyId:       secretKeyId,
		RandomId:          randomId,
		ParticipantsCount: 1,
		AdminsCount:       1,
		Title:             title,
		About:             about,
		Notice:            notice,
		Link:              "",
		Lat:               geo.GetLat(),
		Long:              geo.GetLong(),
		AccuracyRadius:    0,
		Address:           address.GetValue(),
		PhotoId:           0,
		HasLink:           0,
		HasGeo:            util.BoolToInt8(hasGeo),
		SlowmodeEnabled:   0,
		Broadcast:         util.BoolToInt8(isBroadcast),
		Megagroup:         util.BoolToInt8(!isBroadcast),
		Signatures:        0,
		Version:           1,
		Date:              date,
	}

	// democracy
	// broadcast为true, 必须设置democracy为true
	if isBroadcast {
		channelDO.Democracy = 1
	}

	// if channel is existed.
	participantDO := &dataobject.ChannelParticipantsDO{
		UserId:      creatorId,
		IsCreator:   1,
		AdminRights: int32(model.MakeChatAdminRights(model.MakeChatCreatorAdminRights())),
		State:       model.ChatMemberStateNormal,
		Date2:       date,
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		// 1. insert chat
		channelId, _, err := m.ChannelsDAO.InsertTx(tx, channelDO)
		if err != nil {
			result.Err = err
			return
		}
		channelDO.Id = int32(channelId)

		participantDO.ChannelId = channelDO.Id
		_, _, err = m.ChannelParticipantsDAO.InsertTx(tx, participantDO)
		if err != nil {
			result.Err = err
			return
		}
		return
	})

	if tR.Err != nil {
		return nil, tR.Err
	}

	return makeMutableChannel(channelDO, []dataobject.ChannelParticipantsDO{*participantDO}), nil
}

func (m *ChannelCore) GetMutableChannelByLink(ctx context.Context, link string, id ...int32) (*model.MutableChannel, error) {
	channel := new(model.MutableChannel)

	channelDO, err := m.ChannelsDAO.SelectByLink(ctx, link)
	if err != nil {
		return nil, err
	} else if channelDO == nil {
		err = mtproto.ErrInviteHashInvalid
		return nil, err
	}

	channel.Channel = makeImmutableChannel(channelDO)
	if len(id) > 0 {
		doList, err := m.ChannelParticipantsDAO.SelectByUserIdList(ctx, channel.GetId(), id)
		if err != nil {
			return nil, err
		}
		channel.Participants = makeImmutableChannelParticipants(channel.Channel, doList)
	}

	return channel, nil
}

func (m *ChannelCore) GetAdminLogs(ctx context.Context, channelId int32, q string, eventsFilter int32, admins []int32, maxId, minId int64, limit int32) []*mtproto.ChannelAdminLogEvent {
	now := int32(time.Now().Unix())
	logDOList, err := m.ChannelAdminLogsDAO.SelectByChannelId(ctx, channelId, now-24*60*60)
	if err != nil {
		log.Errorf("getAdminLogs - error: %v", err)
	}

	adminLogs := make([]*mtproto.ChannelAdminLogEvent, 0, len(logDOList))
	for i := 0; i < len(logDOList); i++ {
		action := new(mtproto.ChannelAdminLogEventAction)
		if err = json.Unmarshal(hack.Bytes(logDOList[i].EventData), action); err != nil {
			log.Errorf("getAdminLogs - invalid event_data:{%#v}, error: %v", logDOList[i], err)
			continue
		}
		adminLogs = append(adminLogs, mtproto.MakeTLChannelAdminLogEvent(&mtproto.ChannelAdminLogEvent{
			Id:     logDOList[i].Id,
			Date:   logDOList[i].Date2,
			UserId: logDOList[i].UserId,
			Action: action,
		}).To_ChannelAdminLogEvent())
	}

	return adminLogs
}

func (m *ChannelCore) GetLeftChannelList(ctx context.Context, userId, offset int32) (channels []*model.MutableChannel, err error) {
	pDOList, _ := m.ChannelParticipantsDAO.SelectLeftList(ctx, userId)
	return m.makeMutableChannelList(ctx, userId, pDOList), nil
}

func (m *ChannelCore) GetAdminedPublicChannels(ctx context.Context, userId int32, byLocation, checkLimit bool) (channels []*model.MutableChannel, err error) {
	pDOList, _ := m.ChannelParticipantsDAO.SelectMyAdminPublicList(ctx, userId)
	return m.makeMutableChannelList(ctx, userId, pDOList), nil
}

func (m *ChannelCore) GetMyAdminChannelList(ctx context.Context, userId int32) model.MutableChannels {
	pDOList, _ := m.ChannelParticipantsDAO.SelectMyAdminList(ctx, userId)
	return m.makeMutableChannelList(ctx, userId, pDOList)
}

func (m *ChannelCore) makeMutableChannelList(ctx context.Context, userId int32, pDOList []dataobject.ChannelParticipantsDO) (channels []*model.MutableChannel) {
	if len(pDOList) == 0 {
		return []*model.MutableChannel{}
	}

	idList := make([]int32, 0, len(pDOList))
	for i := 0; i < len(pDOList); i++ {
		idList = append(idList, pDOList[i].ChannelId)
	}

	cDOList, _ := m.ChannelsDAO.SelectByIdList(ctx, idList)
	if len(cDOList) == 0 {
		return []*model.MutableChannel{}
	}

	channels = make([]*model.MutableChannel, 0, len(cDOList))
	for i := 0; i < len(cDOList); i++ {
		for j := 0; j < len(pDOList); j++ {
			if pDOList[j].ChannelId == cDOList[i].Id {
				channels = append(channels, makeMutableChannel(&cDOList[i], pDOList[j:j+1]))
				break
			}
		}
	}

	return
}

func (m *ChannelCore) MigrateFromChat(ctx context.Context, secretKeyId int64, chat *model.MutableChat) (*model.MutableChannel, error) {
	var now = int32(time.Now().Unix())

	channel := &model.MutableChannel{
		Channel: &model.ImmutableChannel{
			Id:                   0,
			AccessHash:           rand.Int63(),
			SecretKeyId:          secretKeyId,
			Title:                chat.Chat.Title,
			Username:             "",
			CreatorId:            chat.Chat.Creator,
			Photo:                chat.Chat.Photo,
			TopMessage:           0,
			Broadcast:            false,
			Democracy:            false,
			Verified:             false,
			Megagroup:            true,
			Signatures:           false,
			Min:                  false,
			Scam:                 false,
			HasLink:              false,
			HasGeo:               false,
			SlowmodeEnabled:      false,
			Date:                 now,
			Version:              1,
			DefaultBannedRights:  model.ChatBannedRights{Rights: 0, UntilDate: 0},
			ParticipantsCount:    0,
			About:                "",
			Notice:               "",
			HiddenPrehistory:     false,
			AdminsCount:          0,
			KickedCount:          0,
			BannedCount:          0,
			OnlineCount:          0,
			ChatPhoto:            chat.Chat.ChatPhoto,
			Link:                 "",
			BotInfo:              nil,
			MigratedFromChatId:   chat.Chat.Id,
			MigratedFromMaxId:    0,
			PinnedMsgId:          0,
			StickerSet:           nil,
			LinkedChatId:         0,
			Location:             nil,
			SlowmodeSeconds:      0,
			SlowmodeNextSendDate: 0,
			Pts:                  0,
			Deleted:              false,
		},
		Participants: make(map[int32]*model.ImmutableChannelParticipant, len(chat.Participants)),
	}

	for k, v := range chat.Participants {
		var isCreator bool
		adminRights := model.MakeChatAdminRights(v.AdminRights)
		if k == channel.Channel.CreatorId {
			isCreator = true
			adminRights = model.MakeChatAdminRights(model.MakeChatCreatorAdminRights())
		}

		channel.Participants[k] = &model.ImmutableChannelParticipant{
			Channel:              channel.Channel,
			Creator:              isCreator,
			State:                0,
			ChannelId:            0,
			Id:                   0,
			UserId:               k,
			Date:                 channel.Channel.Date,
			InviterId:            v.ChatParticipant.InviterId,
			CanEdit:              false,
			PromotedBy:           0,
			AdminRights:          adminRights,
			KickedBy:             0,
			BannedRights:         model.ChatBannedRights{Rights: 0, UntilDate: 0},
			Rank:                 "",
			Pinned:               v.Dialog.Pinned,
			UnreadMark:           false,
			TopMessage:           0,
			ReadInboxMaxId:       1,
			UnreadCount:          0,
			UnreadMentionsCount:  0,
			NotifySettings:       nil,
			Draft:                nil,
			FolderId:             0,
			AvailableMinId:       0,
			AvailableUpdatedDate: 0,
			AvailableMinPts:      0,
		}
	}

	sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		channelDO := &dataobject.ChannelsDO{
			Id:                  0,
			CreatorUserId:       chat.Chat.Creator,
			AccessHash:          rand.Int63(),
			RandomId:            rand.Int63(),
			TopMessage:          0,
			PinnedMsgId:         0,
			Date2:               now,
			Pts:                 0,
			ParticipantsCount:   int32(len(channel.Participants)),
			AdminsCount:         1,
			KickedCount:         0,
			BannedCount:         0,
			Title:               chat.Chat.Title,
			About:               chat.Chat.About,
			Notice:              chat.Chat.Notice,
			PhotoId:             0,
			Public:              0,
			Username:            "",
			Link:                chat.Chat.Link,
			HasLink:             1,
			Broadcast:           0,
			Verified:            0,
			Megagroup:           1,
			Democracy:           0,
			Signatures:          0,
			AdminsEnabled:       0,
			DefaultBannedRights: 0,
			MigratedFromChatId:  chat.Chat.Id,
			PreHistoryHidden:    0,
			Deactivated:         0,
			Version:             1,
			Date:                now,
			Deleted:             0,
		}
		id, _, err := m.ChannelsDAO.InsertTx(tx, channelDO)
		if err != nil {
			result.Err = err
			return
		}
		channelDO.Id = int32(id)
		channel.Channel.Id = channelDO.Id

		pDOList := make([]*dataobject.ChannelParticipantsDO, 0, len(chat.Participants))
		for _, p := range chat.Participants {
			isCreator := p.ChatParticipant.UserId == chat.Chat.Creator
			adminRights := model.MakeChatAdminRights(p.AdminRights)
			if isCreator {
				adminRights = model.MakeChatAdminRights(model.MakeChatCreatorAdminRights())
			}
			pDOList = append(pDOList, &dataobject.ChannelParticipantsDO{
				Id:                        0,
				ChannelId:                 channelDO.Id,
				UserId:                    p.ChatParticipant.UserId,
				IsCreator:                 util.BoolToInt8(isCreator),
				IsPinned:                  0,
				OrderPinned:               0,
				ReadInboxMaxId:            1,
				UnreadCount:               0,
				DraftType:                 0,
				DraftMessageData:          "",
				FolderId:                  0,
				FolderPinned:              0,
				FolderOrderPinned:         0,
				InviterUserId:             p.ChatParticipant.InviterId,
				PromotedBy:                0,
				AdminRights:               int32(adminRights),
				HiddenPrehistory:          0,
				HiddenPrehistoryMessageId: 0,
				KickedBy:                  0,
				BannedRights:              0,
				BannedUntilDate:           0,
				MigratedFromMaxId:         p.Dialog.TopMessage,
				AvailableMinId:            0,
				AvailableMinPts:           0,
				Rank:                      "",
				HasScheduled:              0,
				State:                     0,
				Date2:                     now,
			})
		}
		if _, _, err = m.ChannelParticipantsDAO.InsertBulkTx(tx, pDOList); err != nil {
			result.Err = err
			return
		}
	})

	return channel, nil
}

func (m *ChannelCore) GetAllChannels(ctx context.Context, userId int32) model.MutableChannels {
	pDOList, _ := m.ChannelParticipantsDAO.SelectMyAllChannelList(ctx, userId)
	return m.makeMutableChannelList(ctx, userId, pDOList)
}
