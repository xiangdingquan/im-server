package core

import (
	"context"

	"github.com/gogo/protobuf/types"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/util"
)

func makePeerNotifySettingsByDO(do *dataobject.UserNotifySettingsDO) (settings *mtproto.PeerNotifySettings) {
	settings = mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()
	if do.ShowPreviews != -1 {
		settings.ShowPreviews_FLAGBOOL = mtproto.ToBool(do.ShowPreviews == 1)
	}
	if do.Silent != -1 {
		settings.Silent_FLAGBOOL = mtproto.ToBool(do.Silent == 1)
	}
	if do.MuteUntil != -1 {
		settings.MuteUntil_FLAGINT32 = &types.Int32Value{Value: do.MuteUntil}
	}
	if do.Sound != "-1" {
		settings.Sound_FLAGSTRING = &types.StringValue{Value: do.Sound}
	}
	return
}

func makeDOByPeerNotifySettings(settings *mtproto.PeerNotifySettings) (doMap map[string]interface{}) {
	doMap = map[string]interface{}{}

	if settings.ShowPreviews_FLAGBOOL != nil {
		doMap["show_previews"] = util.BoolToInt8(mtproto.FromBool(settings.ShowPreviews_FLAGBOOL))
	} else {
		doMap["show_previews"] = -1
	}

	if settings.Silent_FLAGBOOL != nil {
		doMap["silent"] = util.BoolToInt8(mtproto.FromBool(settings.Silent_FLAGBOOL))
	} else {
		doMap["silent"] = -1
	}

	if settings.MuteUntil_FLAGINT32 != nil {
		doMap["mute_until"] = settings.MuteUntil_FLAGINT32.Value
	} else {
		doMap["mute_until"] = -1
	}

	if settings.Sound_FLAGSTRING != nil {
		doMap["sound"] = settings.Sound_FLAGSTRING.Value
	} else {
		doMap["sound"] = "-1"
	}

	return
}

func (m *UserCore) GetNotifySettingsByPeerList(
	ctx context.Context,
	userId int32,
	userIdList, chatIdList, channelIdList []int32) (settingsList []map[int32]*mtproto.PeerNotifySettings, err error) {

	var (
		cacheError        = false
		peerType          int32
		peerId            int32
		cacheSettingsList map[int64]*mtproto.PeerNotifySettings

		userSettingsList    = make(map[int32]*mtproto.PeerNotifySettings)
		chatSettingsList    = make(map[int32]*mtproto.PeerNotifySettings)
		channelSettingsList = make(map[int32]*mtproto.PeerNotifySettings)

		missUserIdList    []int32
		missChatIdList    []int32
		missChannelIdList []int32
	)

	if cacheSettingsList, err = m.Dao.Redis.GetPeerNotifySettingsByPeerList(ctx, userId, map[int32][]int32{
		model.PEER_USER:    userIdList,
		model.PEER_CHAT:    chatIdList,
		model.PEER_CHANNEL: channelIdList,
	}); err != nil {
		cacheError = true

		missUserIdList = userIdList
		missChatIdList = chatIdList
		missChannelIdList = channelIdList
	} else {
		missUserIdList = make([]int32, 0, len(userIdList))
		missChatIdList = make([]int32, 0, len(chatIdList))
		missChannelIdList = make([]int32, 0, len(channelIdList))

		for k, v := range cacheSettingsList {
			peerType = int32(k >> 32)
			peerId = int32(k & 0xffffffff)
			switch peerType {
			case model.PEER_USER:
				userSettingsList[peerId] = v
				if v == nil {
					missUserIdList = append(missUserIdList, peerId)
				}
			case model.PEER_CHAT:
				chatSettingsList[peerId] = v
				if v == nil {
					missChatIdList = append(missChatIdList, peerId)
				}
			case model.PEER_CHANNEL:
				channelSettingsList[peerId] = v
				if v == nil {
					missChannelIdList = append(missChannelIdList, peerId)
				}
			}
		}

		// all cache
		if len(missUserIdList) == 0 && len(missChatIdList) == 0 && len(missChannelIdList) == 0 {
			return
		}
	}

	// miss
	var doList []dataobject.UserNotifySettingsDO
	if doList, err = m.UserNotifySettingsDAO.SelectList(ctx, userId, missUserIdList, missChatIdList, missChannelIdList); err != nil {
		return
	}

	for i := 0; i < len(doList); i++ {
		do := &doList[i]
		switch do.PeerType {
		case model.PEER_USER:
			userSettingsList[do.PeerId] = makePeerNotifySettingsByDO(do)
		case model.PEER_CHAT:
			chatSettingsList[do.PeerId] = makePeerNotifySettingsByDO(do)
		case model.PEER_CHANNEL:
			channelSettingsList[do.PeerId] = makePeerNotifySettingsByDO(do)
		}
	}

	for _, id := range userIdList {
		if _, ok := userSettingsList[id]; !ok {
			userSettingsList[id] = model.MakeDefaultPeerNotifySettings(model.PEER_USER)
		}
	}
	for _, id := range chatIdList {
		if _, ok := chatSettingsList[id]; !ok {
			chatSettingsList[id] = model.MakeDefaultPeerNotifySettings(model.PEER_CHAT)
		}
	}
	for _, id := range channelIdList {
		if _, ok := channelSettingsList[id]; !ok {
			channelSettingsList[id] = model.MakeDefaultPeerNotifySettings(model.PEER_CHANNEL)
		}
	}

	if !cacheError {
		missSettings := make(map[int64]*mtproto.PeerNotifySettings)
		for _, id := range missUserIdList {
			missSettings[int64(model.PEER_USER)<<32|int64(id)] = userSettingsList[id]
		}
		for _, id := range missUserIdList {
			missSettings[int64(model.PEER_CHAT)<<32|int64(id)] = chatSettingsList[id]
		}
		for _, id := range missUserIdList {
			missSettings[int64(model.PEER_CHANNEL)<<32|int64(id)] = channelSettingsList[id]
		}
		m.Dao.Redis.SetPeerNotifySettingsList(ctx, userId, missSettings)
	}

	settingsList = []map[int32]*mtproto.PeerNotifySettings{
		userSettingsList,
		chatSettingsList,
		channelSettingsList,
	}
	return
}

func (m *UserCore) GetNotifySettings(ctx context.Context, userId int32, peer *model.PeerUtil) (settings *mtproto.PeerNotifySettings, err error) {
	var (
		cacheError = false
		do         *dataobject.UserNotifySettingsDO
	)

	if settings, err = m.Dao.Redis.GetPeerNotifySettings(ctx, userId, peer); err != nil {
		cacheError = true
	} else if settings != nil {
		// hit
		return
	}

	// miss or redis error
	if do, err = m.UserNotifySettingsDAO.Select(ctx, userId, int8(peer.PeerType), peer.PeerId); err != nil {
		// log.Errorf("")
		return
	}

	if do == nil {
		settings = model.MakeDefaultPeerNotifySettings(peer.PeerType)
	} else {
		settings = makePeerNotifySettingsByDO(do)
	}

	if !cacheError {
		// put cache
		m.Dao.Redis.SetPeerNotifySettings(ctx, userId, peer, settings)
	}
	return
}

func (m *UserCore) SetNotifySettings(ctx context.Context, userId int32, peer *model.PeerUtil, settings *mtproto.PeerNotifySettings) (err error) {
	cMap := makeDOByPeerNotifySettings(settings)
	if _, _, err = m.UserNotifySettingsDAO.InsertOrUpdateExt(ctx, userId, peer.PeerType, peer.PeerId, cMap); err != nil {
		return
	}

	// putCache
	m.Dao.Redis.SetPeerNotifySettings(ctx, userId, peer, settings)
	return
}

func (m *UserCore) ResetNotifySettings(ctx context.Context, userId int32) (err error) {
	if _, err = m.UserNotifySettingsDAO.DeleteAll(ctx, userId); err != nil {
		return
	}
	// del cache
	m.Dao.Redis.DelAllPeerNotifySettings(ctx, userId)

	return err
}

func (m *UserCore) GetAllNotifySettings(ctx context.Context, userId int32) (settings map[int64]*mtproto.PeerNotifySettings, err error) {
	var doList []dataobject.UserNotifySettingsDO

	if doList, err = m.UserNotifySettingsDAO.SelectAll(ctx, userId); err != nil {
		return
	}

	settings = make(map[int64]*mtproto.PeerNotifySettings)
	for i := 0; i < len(doList); i++ {
		settings[int64(doList[i].PeerType)<<32|int64(doList[i].PeerId)] = makePeerNotifySettingsByDO(&doList[i])
	}

	return
}
