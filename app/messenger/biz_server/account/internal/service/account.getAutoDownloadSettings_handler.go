package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetAutoDownloadSettings(ctx context.Context, request *mtproto.TLAccountGetAutoDownloadSettings) (*mtproto.Account_AutoDownloadSettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getAutoDownloadSettings#56da0b3f - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	accountAutoDownloadSettings := mtproto.MakeTLAccountAutoDownloadSettings(&mtproto.Account_AutoDownloadSettings{
		Low:    makeAutoDownloadSettings(false, true, true, true, 1048576, 512000, 512000),
		Medium: makeAutoDownloadSettings(false, true, true, false, 1048576, 10485760, 1048576),
		High:   makeAutoDownloadSettings(false, true, true, false, 1048576, 15728640, 3145728),
	})

	log.Debugf("account.getAutoDownloadSettings#56da0b3f - reply: %s\n", accountAutoDownloadSettings.DebugString())
	return accountAutoDownloadSettings.To_Account_AutoDownloadSettings(), nil
}

func makeAutoDownloadSettings(disabled, videoPreloadLarge, audioPreloadNext, phonecallsLessData bool, photoSizeMax, videoSizeMax, fileSizeMax int32) *mtproto.AutoDownloadSettings {
	return mtproto.MakeTLAutoDownloadSettings(&mtproto.AutoDownloadSettings{
		Disabled:           disabled,
		VideoPreloadLarge:  videoPreloadLarge,
		AudioPreloadNext:   audioPreloadNext,
		PhonecallsLessData: phonecallsLessData,
		PhotoSizeMax:       photoSizeMax,
		VideoSizeMax:       videoSizeMax,
		FileSizeMax:        fileSizeMax,
	}).To_AutoDownloadSettings()
}
