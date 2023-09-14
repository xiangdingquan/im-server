package dataobject

type AuthsDO struct {
	Id             int32  `db:"id"`
	AuthKeyId      int64  `db:"auth_key_id"`
	Layer          int32  `db:"layer"`
	ApiId          int32  `db:"api_id"`
	DeviceModel    string `db:"device_model"`
	SystemVersion  string `db:"system_version"`
	AppVersion     string `db:"app_version"`
	SystemLangCode string `db:"system_lang_code"`
	LangPack       string `db:"lang_pack"`
	LangCode       string `db:"lang_code"`
	SystemCode     string `db:"system_code"`
	Proxy          string `db:"proxy"`
	Params         string `db:"params"`
	ClientIp       string `db:"client_ip"`
	Deleted        int8   `db:"deleted"`
}
