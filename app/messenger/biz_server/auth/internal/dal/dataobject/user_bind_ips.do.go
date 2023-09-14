package dataobject

type UserBindIps struct {
	ID      uint32   `db:"id"`
	UserID  int32    `db:"user_id"`
	IpAddrs string   `db:"ip_addrs"`
	IpList  []string `db:"-"`
}
