package core

import (
	"context"

	"github.com/gogo/protobuf/types"

	"open.chat/app/messenger/biz_server/langpack/internal/dal/dataobject"
	"open.chat/mtproto"
	"open.chat/pkg/util"
)

func makeLangPackString(do *dataobject.StringsDO) (langPackString *mtproto.LangPackString) {
	if do.Pluralized == 1 {
		langPackString = mtproto.MakeTLLangPackString(&mtproto.LangPackString{
			Key:        do.Key2,
			ZeroValue:  nil,
			OneValue:   nil,
			TwoValue:   nil,
			FewValue:   nil,
			ManyValue:  nil,
			OtherValue: do.OtherValue,
		}).To_LangPackString()
		if do.ZeroValue != "" {
			langPackString.ZeroValue = &types.StringValue{Value: do.ZeroValue}
		}
		if do.OneValue != "" {
			langPackString.OneValue = &types.StringValue{Value: do.OneValue}
		}
		if do.TwoValue != "" {
			langPackString.TwoValue = &types.StringValue{Value: do.TwoValue}
		}
		if do.FewValue != "" {
			langPackString.FewValue = &types.StringValue{Value: do.FewValue}
		}
		if do.ManyValue != "" {
			langPackString.ManyValue = &types.StringValue{Value: do.ManyValue}
		}
	} else {
		langPackString = mtproto.MakeTLLangPackString(&mtproto.LangPackString{
			Key:   do.Key2,
			Value: do.Value,
		}).To_LangPackString()
	}

	return
}

func makeLangPackLanguage(do1 *dataobject.LanguagesDO, do2 *dataobject.AppLanguagesDO) *mtproto.LangPackLanguage {
	language := mtproto.MakeTLLangPackLanguage(&mtproto.LangPackLanguage{
		Official:        util.Int8ToBool(do1.Official),
		Rtl:             util.Int8ToBool(do1.Rtl),
		Beta:            util.Int8ToBool(do1.Beta),
		Name:            do1.Name,
		NativeName:      do1.NativeName,
		LangCode:        do1.LangCode,
		BaseLangCode:    nil,
		PluralCode:      do1.PluralCode,
		StringsCount:    do2.StringsCount,
		TranslatedCount: do2.TranslatedCount,
		TranslationsUrl: do1.TranslationsUrl,
	}).To_LangPackLanguage()

	if do1.BaseLangCode != "" {
		language.BaseLangCode = &types.StringValue{Value: do1.BaseLangCode}
	}
	return language
}

func (c *LangPackCore) GetLanguages(ctx context.Context, pack string) []*mtproto.LangPackLanguage {
	appLanguageList, _ := c.AppLanguagesDAO.SelectLanguageList(ctx, pack)

	var (
		codeList  = make([]string, len(appLanguageList))
		languages = make([]*mtproto.LangPackLanguage, 0, len(appLanguageList))
	)

	if len(appLanguageList) == 0 {
		return languages
	} else {
		for i := 0; i < len(appLanguageList); i++ {
			codeList[i] = appLanguageList[i].LangCode
		}
	}

	var findAppLanguage = func(code string) *dataobject.AppLanguagesDO {
		for i := 0; i < len(appLanguageList); i++ {
			if appLanguageList[i].LangCode == code {
				return &appLanguageList[i]
			}
		}
		return nil
	}

	languageList, _ := c.LanguagesDAO.SelectLanguageList(ctx, codeList)
	for i := 0; i < len(languageList); i++ {
		if appLanguage := findAppLanguage(languageList[i].LangCode); appLanguage != nil {
			languages = append(languages, makeLangPackLanguage(&languageList[i], appLanguage))
		}
	}

	return languages
}

func (c *LangPackCore) GetLanguage(ctx context.Context, pack, code string) (int32, *mtproto.LangPackLanguage, error) {
	appLanguage, err := c.AppLanguagesDAO.SelectLanguage(ctx, pack, code)
	if err != nil {
		return 0, nil, err
	} else if appLanguage == nil {
		err = mtproto.ErrLangCodeNotSupported
		return 0, nil, err
	}

	language, err := c.LanguagesDAO.SelectLanguage(ctx, code)
	if err != nil {
		return 0, nil, err
	} else if language == nil {
		err = mtproto.ErrLangCodeNotSupported
		return 0, nil, err
	}

	return appLanguage.Version, makeLangPackLanguage(language, appLanguage), nil
}

func (c *LangPackCore) GetDifference(ctx context.Context, pack, code string, version int32) []*mtproto.LangPackString {
	doList, _ := c.Dao.StringsDAO.SelectGTVersionList(ctx, pack, code, version)
	sList := make([]*mtproto.LangPackString, len(doList))
	for i := 0; i < len(doList); i++ {
		sList[i] = makeLangPackString(&doList[i])
	}
	return sList
}

func (c *LangPackCore) GetStringListByIdList(ctx context.Context, pack, code string, keys []string) []*mtproto.LangPackString {
	doList, _ := c.Dao.StringsDAO.SelectListByKeyList(ctx, pack, code, keys)
	sList := make([]*mtproto.LangPackString, len(doList))
	for i := 0; i < len(doList); i++ {
		sList[i] = makeLangPackString(&doList[i])
	}
	return sList
}

func (c *LangPackCore) GetStrings(ctx context.Context, pack, code string) []*mtproto.LangPackString {
	doList, _ := c.Dao.StringsDAO.SelectList(ctx, pack, code)
	sList := make([]*mtproto.LangPackString, len(doList))
	for i := 0; i < len(doList); i++ {
		sList[i] = makeLangPackString(&doList[i])
	}
	return sList
}
