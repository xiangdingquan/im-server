package dataobject

type ThemesDO struct {
	Id            int32  `db:"id"`
	ThemeId       int64  `db:"theme_id"`
	AccessHash    int64  `db:"access_hash"`
	Creator       int32  `db:"creator"`
	Default2      int8   `db:"default2"`
	Slug          string `db:"slug"`
	Title         string `db:"title"`
	Settings      string `db:"settings"`
	InstallsCount int32  `db:"installs_count"`
}
