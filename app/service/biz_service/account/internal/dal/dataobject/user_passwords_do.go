package dataobject

type UserPasswordsDO struct {
	Id           int64  `db:"id"`
	UserId       int32  `db:"user_id"`
	NewAlgoSalt1 string `db:"new_algo_salt1"`
	V            string `db:"v"`
	SrpId        int64  `db:"srp_id"`
	SrpB         string `db:"srp_b"`
	B            string `db:"B"`
	Hint         string `db:"hint"`
	Email        string `db:"email"`
	HasRecovery  int8   `db:"has_recovery"`
	Code         string `db:"code"`
	CodeExpired  int32  `db:"code_expired"`
	Attempts     int32  `db:"attempts"`
	State        int8   `db:"state"`
}
