package dataobject

type BannedIp struct {
	ID     uint32 `db:"id"`
	IpAddr string `db:"ip_addr"`
}
