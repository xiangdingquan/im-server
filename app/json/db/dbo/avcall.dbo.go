package dbo

import "time"

type (
	// AvcallRecordDO .
	AvcallRecordDO struct {
		ID       uint32    `db:"id"`
		CallID   uint32    `db:"call_id"`
		UserID   uint32    `db:"user_id"`
		IsRead   bool      `db:"is_read"`
		CallTime uint32    `db:"call_time"`
		EnterAt  uint32    `db:"enter_at"`
		LeaveAt  uint32    `db:"leave_at"`
		Deleted  bool      `db:"deleted"`
		CreateAt time.Time `db:"created_at"`
	}

	// AvcallDO from .
	AvcallDO struct {
		ID          uint32    `db:"id"`
		ChannelName string    `db:"channel_name"`
		ChatID      uint32    `db:"chat_id"`
		OwnerUID    uint32    `db:"owner_uid"`
		MemberInfo  string    `db:"member_uids"`
		Members     []uint32  `db:"-"`
		CreateAt    uint32    `db:"-"`
		StartAt     uint32    `db:"start_at"`
		IsVideo     bool      `db:"is_video"`
		IsMeet      bool      `db:"is_meet"`
		CloseAt     uint32    `db:"close_at"`
		Deleted     bool      `db:"deleted"`
		CreateTime  time.Time `db:"created_at"`
	}
)
