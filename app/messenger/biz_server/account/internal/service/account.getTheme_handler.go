package service

import (
	"context"

	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetTheme(ctx context.Context, request *mtproto.TLAccountGetTheme) (*mtproto.Theme, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getTheme - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		theme *model.Theme
		err   error
	)

	switch request.GetTheme().GetPredicateName() {
	case mtproto.Predicate_inputTheme:
		theme, err = s.AccountCore.GetThemeByIdAndFormat(ctx, request.GetTheme().GetId(), request.Format)
		if err != nil {
			log.Errorf("account.getTheme - error: %v", err)
			return nil, err
		}
	case mtproto.Predicate_inputThemeSlug:
		theme, err = s.AccountCore.GetThemeBySlugAndFormat(ctx, request.GetTheme().GetSlug(), request.Format)
		if err != nil {
			log.Errorf("account.getTheme - error: %v", err)
			return nil, err
		}
	default:
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("account.getTheme - error: %v", err)
		return nil, err
	}

	dList, _ := media_client.GetDocumentByIdList([]int64{theme.DocumentId})
	if len(dList) != 1 {
		log.Errorf("account.getTheme - error: invalid documentId(%d)", theme.DocumentId)
		// err = mtproto.ErrThemeInvalid
		return nil, mtproto.ErrThemeInvalid
	}

	theme.Document = dList[0]

	log.Debugf("account.getTheme - reply: %s", theme.GetTheme().DebugString())
	return theme.GetTheme(), nil
}
