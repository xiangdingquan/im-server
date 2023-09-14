package dataobject

type ChannelParticipantsExtDO struct {
	ChannelParticipantsDO
	TopMessage      int32 `db:"top_message"`
	ReadOutboxMaxId int32 `db:"read_outbox_max_id"`
	Pts             int32 `db:"pts"`
	Date            int32 `db:"date"`
}
