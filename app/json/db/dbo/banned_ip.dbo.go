package dbo

import (
	"time"
)

type (
	// BannedIp .
	BannedIp struct {
		ID         uint32    `db:"id"`
		IpAddr     string    `db:"ip_addr"`
		Deleted    bool      `db:"deleted"`
		CreateTime time.Time `db:"created_at"`
		CreateAt   uint32    `db:"-"`
	}
)
