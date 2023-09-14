package dataobject

type UsernameDO struct {
	Id       int64  `db:"id"`
	Username string `db:"username"`
	PeerType int8   `db:"peer_type"`
	PeerId   int32  `db:"peer_id"`
	Deleted  int8   `db:"deleted"`
	UpdateAt string `db:"update_at"`
}
