package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountInstallTheme(ctx context.Context, request *mtproto.TLAccountInstallTheme) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.installTheme - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		theme *model.Theme
		err   error
	)

	switch request.GetTheme().GetPredicateName() {
	case mtproto.Predicate_inputTheme:
		theme, err = s.AccountCore.GetThemeByIdAndFormat(ctx, request.GetTheme().GetId(), request.Format.GetValue())
		if err != nil {
			log.Errorf("account.installTheme - error: %v", err)
			return nil, err
		}
	case mtproto.Predicate_inputThemeSlug:
		theme, err = s.AccountCore.GetThemeBySlugAndFormat(ctx, request.GetTheme().GetSlug(), request.Format.GetValue())
		if err != nil {
			log.Errorf("account.installTheme - error: %v", err)
			return nil, err
		}
	default:
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("account.installTheme - error: %v", err)
		return nil, err
	}

	err = s.AccountCore.InstallTheme(ctx, md.UserId, theme.GetId(), request.GetFormat().GetValue())
	if err != nil {
		log.Errorf("account.installTheme - error: %v", err)
		return nil, err
	}

	log.Debugf("account.installTheme - reply: {true}")
	return mtproto.BoolTrue, nil
}
