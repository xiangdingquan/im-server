package dataobject

type PhoneCallSessionsDO struct {
	Id                   int64  `db:"id"`
	AccessHash           int64  `db:"access_hash"`
	AdminId              int32  `db:"admin_id"`
	ParticipantId        int32  `db:"participant_id"`
	AdminAuthKeyId       int64  `db:"admin_auth_key_id"`
	ParticipantAuthKeyId int64  `db:"participant_auth_key_id"`
	RandomId             int64  `db:"random_id"`
	AdminProtocol        string `db:"admin_protocol"`
	ParticipantProtocol  string `db:"participant_protocol"`
	GAHash               string `db:"g_a_hash"`
	GA                   string `db:"g_a"`
	GB                   string `db:"g_b"`
	KeyFingerprint       int64  `db:"key_fingerprint"`
	Connections          string `db:"connections"`
	AdminDebugData       string `db:"admin_debug_data"`
	ParticipantDebugData string `db:"participant_debug_data"`
	AdminRating          int32  `db:"admin_rating"`
	AdminComment         string `db:"admin_comment"`
	ParticipantRating    int32  `db:"participant_rating"`
	ParticipantComment   string `db:"participant_comment"`
	Date                 int32  `db:"date"`
	State                int32  `db:"state"`
}
