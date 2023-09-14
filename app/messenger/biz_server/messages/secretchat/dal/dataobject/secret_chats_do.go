package dataobject

type SecretChatsDO struct {
	Id                   int32  `db:"id"`
	AccessHash           int64  `db:"access_hash"`
	AdminId              int32  `db:"admin_id"`
	ParticipantId        int32  `db:"participant_id"`
	AdminAuthKeyId       int64  `db:"admin_auth_key_id"`
	ParticipantAuthKeyId int64  `db:"participant_auth_key_id"`
	RandomId             int32  `db:"random_id"`
	GA                   string `db:"g_a"`
	GB                   string `db:"g_b"`
	KeyFingerprint       int64  `db:"key_fingerprint"`
	State                int8   `db:"state"`
	Date                 int32  `db:"date"`
}
