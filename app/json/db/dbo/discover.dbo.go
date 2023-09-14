package dbo

import (
	"time"
)

type (
	// DiscoverGroupDO .
	DiscoverGroupDO struct {
		ID         uint32    `db:"id"`
		ChannelID  uint32    `db:"channel_id"`
		Name       string    `db:"name"`
		CreateAt   uint32    `db:"-"`
		CreateTime time.Time `db:"created_at"`
	}

	// DiscoverMenus .
	DiscoverMenus struct {
		ID         uint32    `db:"id"`
		ChannelID  uint32    `db:"channel_id"`
		GroupID    uint32    `db:"group_id"`
		Category   int8      `db:"category"`
		Title      string    `db:"title"`
		Logo       string    `db:"logo"`
		Url        string    `db:"url"`
		CreateAt   uint32    `db:"-"`
		CreateTime time.Time `db:"created_at"`
	}
)
