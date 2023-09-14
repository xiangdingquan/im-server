package dataobject

type GiphyDatasDO struct {
	Id         int32  `db:"id"`
	GiphyId    string `db:"giphy_id"`
	DocumentId int64  `db:"document_id"`
	PhotoId    int64  `db:"photo_id"`
}
