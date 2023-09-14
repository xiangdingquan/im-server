package model

import (
	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
)

func MakePeerNotifySettings(settings *mtproto.InputPeerNotifySettings) (*mtproto.PeerNotifySettings, error) {
	notifySettings := mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()

	switch settings.GetConstructor() {
	case mtproto.CRC32_inputPeerNotifySettings_38935eb2:
		if settings.ShowPreviews_FLAGBOOLEAN {
			notifySettings.ShowPreviews_FLAGBOOL = mtproto.ToBool(settings.ShowPreviews_FLAGBOOLEAN)
		}
		if settings.Silent_FLAGBOOLEAN {
			notifySettings.Silent_FLAGBOOL = mtproto.ToBool(settings.Silent_FLAGBOOLEAN)
		}
		if settings.MuteUntil_INT32 > 0 {
			notifySettings.MuteUntil_FLAGINT32 = &types.Int32Value{Value: settings.MuteUntil_INT32}
		}
		if settings.Sound_STRING != "" {
			notifySettings.Sound_FLAGSTRING = &types.StringValue{Value: settings.Sound_STRING}
		}
	case mtproto.CRC32_inputPeerNotifySettings_9c3d198e:
		notifySettings.ShowPreviews_FLAGBOOL = settings.ShowPreviews_FLAGBOOL
		notifySettings.Silent_FLAGBOOL = settings.Silent_FLAGBOOL
		notifySettings.MuteUntil_FLAGINT32 = settings.MuteUntil_FLAGINT32
		notifySettings.Sound_FLAGSTRING = settings.Sound_FLAGSTRING
	default:
		err := mtproto.ErrTypeConstructorInvalid
		return nil, err
	}

	return notifySettings, nil
}

func MakeDefaultPeerNotifySettings(peerType int32) *mtproto.PeerNotifySettings {
	settings := mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()

	if peerType == PEER_USERS || peerType == PEER_CHATS || peerType == PEER_BROADCASTS {
		settings.ShowPreviews_FLAGBOOL = mtproto.ToBool(true)
		settings.Silent_FLAGBOOL = mtproto.ToBool(false)
		settings.MuteUntil_FLAGINT32 = &types.Int32Value{Value: 0}
		settings.Sound_FLAGSTRING = &types.StringValue{Value: "default"}
	}

	return settings
}
