package dataobject

type BotsDO struct {
	Id                   int32  `db:"id"`
	BotId                int32  `db:"bot_id"`
	BotType              int8   `db:"bot_type"`
	CreatorUserId        int32  `db:"creator_user_id"`
	Token                string `db:"token"`
	Description          string `db:"description"`
	BotChatHistory       int8   `db:"bot_chat_history"`
	BotNochats           int8   `db:"bot_nochats"`
	Verified             int8   `db:"verified"`
	BotInlineGeo         int8   `db:"bot_inline_geo"`
	BotInfoVersion       int32  `db:"bot_info_version"`
	BotInlinePlaceholder string `db:"bot_inline_placeholder"`
}
