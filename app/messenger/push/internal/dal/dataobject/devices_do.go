package dataobject

type DevicesDO struct {
	Id           int64  `db:"id"`
	AuthKeyId    int64  `db:"auth_key_id"`
	UserId       int32  `db:"user_id"`
	TokenType    int8   `db:"token_type"`
	Token        string `db:"token"`
	NoMuted      int8   `db:"no_muted"`
	LockedPeriod int32  `db:"locked_period"`
	AppSandbox   int8   `db:"app_sandbox"`
	Secret       string `db:"secret"`
	OtherUids    string `db:"other_uids"`
	State        int8   `db:"state"`
}
