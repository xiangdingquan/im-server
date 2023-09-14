package accounts

import (
	"context"
	"math"
	"open.chat/app/json/helper"
	"open.chat/app/json/service/handler"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"strconv"
)

const (
	multiOnlineKey  = "multiOnline"
	notificationKey = "notification"
	usingSoundKey   = "usingSound"
)

func (s *cls) ToggleMultiOnline(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TToggleMultiOnline) *helper.ResultJSON {
	var value string
	if r.IsOn {
		value = "true"
	} else {
		value = "false"
	}

	err := s.AccountFacade.SetSettingValue(ctx, md.UserId, multiOnlineKey, value)
	if err != nil {
		log.Errorf("accounts.toggleMultiOnline, update db failed, error: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "save setting failed"}
	}

	if !r.IsOn {
		tKeyIdList, _ := s.RPCSessionClient.SessionResetAuthorization(ctx, &authsessionpb.TLSessionResetAuthorization{
			UserId:    md.UserId,
			AuthKeyId: md.AuthId,
			Hash:      0,
		})

		for _, id := range tKeyIdList.Datas {
			upds := mtproto.MakeTLUpdateAccountResetAuthorization(&mtproto.Updates{
				UserId:    md.UserId,
				AuthKeyId: id,
			}).To_Updates()
			sync_client.SyncUpdatesMe(ctx, md.UserId, id, 0, "", upds)
		}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) GetMultiOnline(ctx context.Context, md *grpc_util.RpcMetadata) *helper.ResultJSON {
	isOn := s.AccountFacade.GetSettingValueBool(ctx, md.UserId, multiOnlineKey, true)
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: struct {
		IsOn bool `json:"isOn"`
	}{
		IsOn: isOn,
	}}
}

func (s *cls) ModifyNotificationSettings(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TNotificationSettings) *helper.ResultJSON {
	notificationNum, usingSound := toDb(r)

	var err error
	err = s.AccountFacade.SetSettingValue(ctx, md.UserId, notificationKey, strconv.Itoa(notificationNum))
	if err != nil {
		log.Errorf("accounts.modifyNotificationSettings, update notification settings to db failed, error: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "save notification failed"}
	}

	err = s.AccountFacade.SetSettingValue(ctx, md.UserId, usingSoundKey, strconv.Itoa(usingSound))
	if err != nil {
		log.Errorf("accounts.modifyNotificationSettings, update using sound to db failed, error: %v", err)
		return &helper.ResultJSON{Code: -2, Msg: "save sound failed"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) GetNotificationSettings(ctx context.Context, md *grpc_util.RpcMetadata) *helper.ResultJSON {
	notificationNum := s.AccountFacade.GetSettingValueInt32(ctx, md.UserId, notificationKey, math.MaxInt32)
	usingSound := s.AccountFacade.GetSettingValueInt32(ctx, md.UserId, usingSoundKey, 0)

	data := fromDb(notificationNum, usingSound)

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func toDb(r *handler.TNotificationSettings) (notificationNumber, usingSound int) {
	if r.ShowNotification {
		notificationNumber |= 1 << 0
	}
	if r.ShowPreview {
		notificationNumber |= 1 << 1
	}
	if r.InAppSound {
		notificationNumber |= 1 << 2
	}
	if r.InAppVibrate {
		notificationNumber |= 1 << 3
	}
	if r.InAppPreview {
		notificationNumber |= 1 << 4
	}
	if r.CountClosed {
		notificationNumber |= 1 << 5
	}

	usingSound = int(r.Sound)

	return
}

func fromDb(notificationNumber, usingSound int32) *handler.TNotificationSettings {
	return &handler.TNotificationSettings{
		ShowNotification: (notificationNumber & (1 << 0)) != 0,
		ShowPreview:      (notificationNumber & (1 << 1)) != 0,
		InAppSound:       (notificationNumber & (1 << 2)) != 0,
		InAppVibrate:     (notificationNumber & (1 << 3)) != 0,
		InAppPreview:     (notificationNumber & (1 << 4)) != 0,
		CountClosed:      (notificationNumber & (1 << 5)) != 0,

		Sound: usingSound,
	}
}
