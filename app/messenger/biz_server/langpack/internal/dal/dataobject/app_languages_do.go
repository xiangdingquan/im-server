package dataobject

type AppLanguagesDO struct {
	Id              int32  `db:"id"`
	App             string `db:"app"`
	LangCode        string `db:"lang_code"`
	Version         int32  `db:"version"`
	StringsCount    int32  `db:"strings_count"`
	TranslatedCount int32  `db:"translated_count"`
	State           int8   `db:"state"`
}
