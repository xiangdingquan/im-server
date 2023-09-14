package dataobject

type LanguagesDO struct {
	Id              int32  `db:"id"`
	LangCode        string `db:"lang_code"`
	BaseLangCode    string `db:"base_lang_code"`
	Link            string `db:"link"`
	Official        int8   `db:"official"`
	Rtl             int8   `db:"rtl"`
	Beta            int8   `db:"beta"`
	Name            string `db:"name"`
	NativeName      string `db:"native_name"`
	PluralCode      string `db:"plural_code"`
	TranslationsUrl string `db:"translations_url"`
	State           int8   `db:"state"`
}
