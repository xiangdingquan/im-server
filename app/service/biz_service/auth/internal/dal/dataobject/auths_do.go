package dataobject

type AuthsDo struct {
	ApiId         int64  `db:"api_id"`
	DeviceModel   string `db:"device_model"`
	SystemVersion string `db:"system_version"`
}
