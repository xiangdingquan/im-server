package core

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/util"
)

func (m *UserCore) AddPeerSettings(ctx context.Context, selfId int32, peer *model.PeerUtil, settings *mtproto.PeerSettings) error {
	_, _, err := m.UserPeerSettingsDAO.InsertIgnore(ctx, &dataobject.UserPeerSettingsDO{
		UserId:                selfId,
		PeerType:              int8(peer.PeerType),
		PeerId:                peer.PeerId,
		Hide:                  0,
		ReportSpam:            util.BoolToInt8(settings.ReportSpam),
		AddContact:            util.BoolToInt8(settings.AddContact),
		BlockContact:          util.BoolToInt8(settings.BlockContact),
		ShareContact:          util.BoolToInt8(settings.ShareContact),
		NeedContactsException: util.BoolToInt8(settings.NeedContactsException),
		ReportGeo:             util.BoolToInt8(settings.ReportGeo),
		Autoarchived:          util.BoolToInt8(settings.Autoarchived),
		GeoDistance:           settings.GetGeoDistance().GetValue(),
	})

	return err
}

func (m *UserCore) GetPeerSettings(ctx context.Context, selfId int32, peer *model.PeerUtil) (*mtproto.PeerSettings, error) {
	peerSettingsDO, err := m.UserPeerSettingsDAO.Select(ctx, selfId, int8(peer.PeerType), peer.PeerId)
	if err != nil {
		return nil, err
	}

	var (
		peerSettings *mtproto.PeerSettings
	)

	if peerSettingsDO != nil {
		peerSettings = &mtproto.PeerSettings{
			ReportSpam:            util.Int8ToBool(peerSettingsDO.ReportSpam),
			AddContact:            util.Int8ToBool(peerSettingsDO.AddContact),
			BlockContact:          util.Int8ToBool(peerSettingsDO.BlockContact),
			ShareContact:          util.Int8ToBool(peerSettingsDO.ShareContact),
			NeedContactsException: util.Int8ToBool(peerSettingsDO.NeedContactsException),
			ReportGeo:             util.Int8ToBool(peerSettingsDO.ReportGeo),
			Autoarchived:          util.Int8ToBool(peerSettingsDO.Autoarchived),
			GeoDistance:           nil,
		}

		if peerSettingsDO.GeoDistance != 0 {
			peerSettings.GeoDistance = &types.Int32Value{Value: peerSettingsDO.GeoDistance}
		}
	}

	return mtproto.MakeTLPeerSettings(peerSettings).To_PeerSettings(), nil
}

func (m *UserCore) DeletePeerSettings(ctx context.Context, selfId int32, peer *model.PeerUtil) error {
	_, err := m.UserPeerSettingsDAO.Delete(ctx, selfId, int8(peer.PeerType), peer.PeerId)
	return err
}
