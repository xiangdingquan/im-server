package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetThemes(ctx context.Context, request *mtproto.TLAccountGetThemes) (*mtproto.Account_Themes, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getThemes - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	themes := s.AccountCore.GetInstalledThemes(ctx, md.UserId, request.Format)

	accountThemes := mtproto.MakeTLAccountThemes(&mtproto.Account_Themes{
		Hash:   0,
		Themes: make([]*mtproto.Theme, 0, len(themes)),
	}).To_Account_Themes()

	for _, theme := range themes {
		accountThemes.Themes = append(accountThemes.Themes, theme.GetTheme())
	}

	log.Debugf("account.getThemes - reply: %s", accountThemes.DebugString())
	return accountThemes, nil
}
