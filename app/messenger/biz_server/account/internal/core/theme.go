package core

import (
	"context"

	"open.chat/app/messenger/biz_server/account/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/pkg/database/sqlx"

	"open.chat/mtproto"
)

func (m *AccountCore) GetThemeBySlugAndFormat(ctx context.Context, slug, format string) (*model.Theme, error) {
	themeDO, err := m.ThemesDAO.SelectBySlug(ctx, slug)
	if err != nil {
		return nil, err
	} else if themeDO == nil {
		return nil, mtproto.ErrThemeInvalid
	}

	themeFormatDO, err := m.ThemeFormatsDAO.SelectByThemeIdAndFormat(ctx, themeDO.ThemeId, format)
	if err != nil {
		return nil, err
	} else if themeFormatDO == nil {
		return nil, mtproto.ErrThemeFormatInvalid
	}

	theme := mtproto.MakeTLTheme(&mtproto.Theme{
		Creator:       false,
		Default:       false,
		Id:            themeDO.ThemeId,
		AccessHash:    themeDO.AccessHash,
		Slug:          themeDO.Slug,
		Title:         themeDO.Title,
		Document:      nil,
		Settings:      nil,
		InstallsCount: themeDO.InstallsCount,
	}).To_Theme()
	return model.MakeTheme(theme, themeFormatDO.DocumentId), nil
}

func (m *AccountCore) GetThemeByIdAndFormat(ctx context.Context, id int64, format string) (*model.Theme, error) {
	themeDO, err := m.ThemesDAO.SelectByThemeId(ctx, id)
	if err != nil {
		return nil, err
	} else if themeDO == nil {
		return nil, mtproto.ErrThemeInvalid
	}

	themeFormatDO, err := m.ThemeFormatsDAO.SelectByThemeIdAndFormat(ctx, themeDO.ThemeId, format)
	if err != nil {
		return nil, err
	} else if themeFormatDO == nil {
		return nil, mtproto.ErrThemeFormatInvalid
	}

	theme := mtproto.MakeTLTheme(&mtproto.Theme{
		Creator:       false,
		Default:       false,
		Id:            themeDO.ThemeId,
		AccessHash:    themeDO.AccessHash,
		Slug:          themeDO.Slug,
		Title:         themeDO.Title,
		Document:      nil,
		Settings:      nil,
		InstallsCount: themeDO.InstallsCount,
	}).To_Theme()
	return model.MakeTheme(theme, themeFormatDO.DocumentId), nil
}

func (m *AccountCore) InstallTheme(ctx context.Context, userId int32, themeId int64, format string) error {
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, _, result.Err = m.UserThemesDAO.InsertIgnoreTx(tx, &dataobject.UserThemesDO{
			UserId:  userId,
			ThemeId: themeId,
			Format:  format,
		})
		if result.Err != nil {
			return
		}

		_, result.Err = m.ThemesDAO.AddInstallsCountTx(tx, themeId)
	})
	return tR.Err
}

func (m *AccountCore) GetInstalledThemes(ctx context.Context, userId int32, format string) []*model.Theme {
	doList, _ := m.UserThemesDAO.SelectByUserIdAndFormat(ctx, userId, format)

	themes := make([]*model.Theme, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		theme, _ := m.GetThemeByIdAndFormat(ctx, doList[i].ThemeId, format)
		if theme != nil {
			themes = append(themes, theme)
		}
	}

	return themes
}
