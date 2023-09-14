package dataobject

type LangPackLanguagesDO struct {
	Id              int32  `db:"id"`
	LangPack        string `db:"lang_pack"`
	LangCode        string `db:"lang_code"`
	Version         int32  `db:"version"`
	BaseLangCode    string `db:"base_lang_code"`
	Official        int8   `db:"official"`
	Rtl             int8   `db:"rtl"`
	Beta            int8   `db:"beta"`
	Name            string `db:"name"`
	NativeName      string `db:"native_name"`
	PluralCode      string `db:"plural_code"`
	StringsCount    int32  `db:"strings_count"`
	TranslatedCount int32  `db:"translated_count"`
	TranslationsUrl string `db:"translations_url"`
	State           int8   `db:"state"`
}
