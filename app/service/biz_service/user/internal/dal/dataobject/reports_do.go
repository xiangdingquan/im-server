package dataobject

type ReportsDO struct {
	Id                  int64  `db:"id"`
	UserId              int32  `db:"user_id"`
	ReportType          int8   `db:"report_type"`
	PeerType            int8   `db:"peer_type"`
	PeerId              int32  `db:"peer_id"`
	MessageSenderUserId int32  `db:"message_sender_user_id"`
	MessageId           int32  `db:"message_id"`
	Reason              int8   `db:"reason"`
	Text                string `db:"text"`
}
