package dataobject

type UserPeerSettingsDO struct {
	Id                    int64 `db:"id"`
	UserId                int32 `db:"user_id"`
	PeerType              int8  `db:"peer_type"`
	PeerId                int32 `db:"peer_id"`
	Hide                  int8  `db:"hide"`
	ReportSpam            int8  `db:"report_spam"`
	AddContact            int8  `db:"add_contact"`
	BlockContact          int8  `db:"block_contact"`
	ShareContact          int8  `db:"share_contact"`
	NeedContactsException int8  `db:"need_contacts_exception"`
	ReportGeo             int8  `db:"report_geo"`
	Autoarchived          int8  `db:"autoarchived"`
	GeoDistance           int32 `db:"geo_distance"`
}
