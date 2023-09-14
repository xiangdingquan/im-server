package dataobject

type UsersDO struct {
	Id                int32  `db:"id"`
	UserType          int8   `db:"user_type"`
	AccessHash        int64  `db:"access_hash"`
	SecretKeyId       int64  `db:"secret_key_id"`
	ChannelId         int32  `db:"channel_id"`
	InviterUId        int32  `db:"inviter_uid"`
	FirstName         string `db:"first_name"`
	LastName          string `db:"last_name"`
	Username          string `db:"username"`
	Password          string `db:"password"`
	Phone             string `db:"phone"`
	CountryCode       string `db:"country_code"`
	Country           string `db:"country"`
	Province          string `db:"province"`
	City              string `db:"city"`
	CityCode          string `db:"city_code"`
	Gender            int8   `db:"gender"`
	Birth             string `db:"birth"`
	Verified          int8   `db:"verified"`
	About             string `db:"about"`
	State             int32  `db:"state"`
	IsBot             int8   `db:"is_bot"`
	IsInternal        int8   `db:"is_internal"`
	IsVirtualUser     int8   `db:"is_virtual"`
	AccountDaysTtl    int32  `db:"account_days_ttl"`
	Photo             string `db:"photo"`
	ProfilePhoto      string `db:"profile_photo"`
	Photos            string `db:"photos"`
	Min               int8   `db:"min"`
	Restricted        int8   `db:"restricted"`
	RestrictionReason string `db:"restriction_reason"`
	Deleted           int8   `db:"deleted"`
	DeleteReason      string `db:"delete_reason"`
}
