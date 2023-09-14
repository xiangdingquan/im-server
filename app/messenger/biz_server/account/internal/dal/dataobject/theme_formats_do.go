package dataobject

type ThemeFormatsDO struct {
	Id         int32  `db:"id"`
	ThemeId    int64  `db:"theme_id"`
	Format     string `db:"format"`
	DocumentId int64  `db:"document_id"`
}
