package dataobject

type StickerSetsDO struct {
	Id            int32  `db:"id"`
	StickerSetId  int64  `db:"sticker_set_id"`
	AccessHash    int64  `db:"access_hash"`
	Title         string `db:"title"`
	ShortName     string `db:"short_name"`
	Count         int32  `db:"count"`
	Hash          int32  `db:"hash"`
	Official      int8   `db:"official"`
	Mask          int8   `db:"mask"`
	Masks         int8   `db:"masks"`
	Archived      int8   `db:"archived"`
	Animated      int8   `db:"animated"`
	InstalledDate int32  `db:"installed_date"`
	Thumb         string `db:"thumb"`
	ThumbDcId     int32  `db:"thumb_dc_id"`
}
