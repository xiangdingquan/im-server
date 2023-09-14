package dataobject

type UserThemesDO struct {
	Id      int32  `db:"id"`
	UserId  int32  `db:"user_id"`
	ThemeId int64  `db:"theme_id"`
	Format  string `db:"format"`
}
