package dataobject

type PhoneCallRatingsDO struct {
	Id                   int64  `db:"id"`
	CallId               int64  `db:"call_id"`
	ParticipantId        int32  `db:"participant_id"`
	ParticipantAuthKeyId int64  `db:"participant_auth_key_id"`
	UserInitiative       int8   `db:"user_initiative"`
	Rating               int32  `db:"rating"`
	Comment              string `db:"comment"`
	DebugData            string `db:"debug_data"`
}
