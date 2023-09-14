package dataobject

type AuthSeqUpdatesDO struct {
	Id         int64  `db:"id"`
	AuthId     int64  `db:"auth_id"`
	UserId     int32  `db:"user_id"`
	Seq        int32  `db:"seq"`
	UpdateType int32  `db:"update_type"`
	UpdateData []byte `db:"update_data"`
	Date2      int32  `db:"date2"`
}
