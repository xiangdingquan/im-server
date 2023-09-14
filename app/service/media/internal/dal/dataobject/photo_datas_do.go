package dataobject

type PhotoDatasDO struct {
	Id         int32  `db:"id"`
	PhotoId    int64  `db:"photo_id"`
	PhotoType  int8   `db:"photo_type"`
	DcId       int32  `db:"dc_id"`
	VolumeId   int64  `db:"volume_id"`
	LocalId    int32  `db:"local_id"`
	AccessHash int64  `db:"access_hash"`
	Width      int32  `db:"width"`
	Height     int32  `db:"height"`
	FileSize   int32  `db:"file_size"`
	FilePath   string `db:"file_path"`
	Ext        string `db:"ext"`
}
