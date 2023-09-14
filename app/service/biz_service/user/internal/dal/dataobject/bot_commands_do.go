package dataobject

type BotCommandsDO struct {
	Id          int32  `db:"id"`
	BotId       int32  `db:"bot_id"`
	Command     string `db:"command"`
	Description string `db:"description"`
}
