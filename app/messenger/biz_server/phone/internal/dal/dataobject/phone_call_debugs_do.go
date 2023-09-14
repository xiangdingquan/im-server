package dataobject

type PhoneCallDebugsDO struct {
	Id                   int64  `db:"id"`
	CallId               int64  `db:"call_id"`
	ParticipantId        int32  `db:"participant_id"`
	ParticipantAuthKeyId int64  `db:"participant_auth_key_id"`
	DebugData            string `db:"debug_data"`
}
