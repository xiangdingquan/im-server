package dataobject

type LangPackStringsDO struct {
	Id         int32  `db:"id"`
	LangPack   string `db:"lang_pack"`
	LangCode   string `db:"lang_code"`
	Version    int32  `db:"version"`
	Key2       string `db:"key2"`
	Pluralized int8   `db:"pluralized"`
	Value      string `db:"value"`
	ZeroValue  string `db:"zero_value"`
	OneValue   string `db:"one_value"`
	TwoValue   string `db:"two_value"`
	FewValue   string `db:"few_value"`
	ManyValue  string `db:"many_value"`
	OtherValue string `db:"other_value"`
	Deleted    int8   `db:"deleted"`
}
