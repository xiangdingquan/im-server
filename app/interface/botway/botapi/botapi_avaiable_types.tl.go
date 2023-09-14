package botapi

import (
	"encoding/json"
)

type ChatId struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func (m *ChatId) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ChatId) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (m *Location) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Location) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type CallbackQuery struct {
	Id              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageId string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

func (m *CallbackQuery) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *CallbackQuery) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type PollOption struct {
	Text       string `json:"text,omitempty"`
	VoterCount int32  `json:"voter_count"`
}

func (m *PollOption) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *PollOption) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int32 `json:"retry_after,omitempty"`
}

func (m *ResponseParameters) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ResponseParameters) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type InputMedia struct {
	Type              string             `json:"type"`
	Media             string             `json:"media"`
	Caption           string             `json:"caption,omitempty"`
	ParseMode         string             `json:"parse_mode,omitempty"`
	Thumb             *InputFileOrFileId `json:"thumb,omitempty"`
	Width             int32              `json:"width,omitempty"`
	Height            int32              `json:"height,omitempty"`
	Duration          int32              `json:"duration,omitempty"`
	SupportsStreaming bool               `json:"supports_streaming,omitempty"`
	Performer         string             `json:"performer,omitempty"`
	Title             string             `json:"title,omitempty"`
}

func (m *InputMedia) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *InputMedia) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

func (m *LoginUrl) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *LoginUrl) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

func (m *ChatPhoto) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ChatPhoto) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type InputFileOrFileId struct {
	FileId string `json:"file_id"`
	Url    string `json:"url"`
}

func (m *InputFileOrFileId) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *InputFileOrFileId) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type User struct {
	Id                      int32  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

func (m *User) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *User) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Animation struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int32      `json:"width"`
	Height       int32      `json:"height"`
	Duration     int32      `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int32      `json:"file_size,omitempty"`
}

func (m *Animation) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Animation) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type VideoNote struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       int32      `json:"length"`
	Duration     int32      `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileSize     int32      `json:"file_size,omitempty"`
}

func (m *VideoNote) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *VideoNote) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Poll struct {
	Id                    string        `json:"id"`
	Question              string        `json:"question"`
	Options               []*PollOption `json:"options"`
	TotalVoterCount       int32         `json:"total_voter_count"`
	IsClosed              bool          `json:"is_closed"`
	IsAnonymous           bool          `json:"is_anonymous"`
	Type                  string        `json:"type"`
	AllowsMultipleAnswers bool          `json:"allows_multiple_answers"`
	CorrectOptionId       int32         `json:"correct_option_id,omitempty"`
}

func (m *Poll) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Poll) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

func (m *InlineKeyboardMarkup) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *InlineKeyboardMarkup) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Chat struct {
	Id               int64            `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title,omitempty"`
	Username         string           `json:"username,omitempty"`
	FirstName        string           `json:"first_name,omitempty"`
	LastName         string           `json:"last_name,omitempty"`
	Photo            *ChatPhoto       `json:"photo,omitempty"`
	Description      string           `json:"description,omitempty"`
	InviteLink       string           `json:"invite_link,omitempty"`
	PinnedMessage    *Message         `json:"pinned_message,omitempty"`
	Permissions      *ChatPermissions `json:"permissions,omitempty"`
	StickerSetName   string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet bool             `json:"can_set_sticker_set,omitempty"`
}

func (m *Chat) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Chat) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int32  `json:"offset"`
	Length   int32  `json:"length"`
	Url      string `json:"url,omitempty"`
	User     *User  `json:"user,omitempty"`
	Language string `json:"language,omitempty"`
}

func (m *MessageEntity) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *MessageEntity) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Audio struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Duration     int32      `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int32      `json:"file_size,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
}

func (m *Audio) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Audio) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Venue struct {
	Location       *Location `json:"location"`
	Title          string    `json:"title"`
	Address        string    `json:"address"`
	FoursquareId   string    `json:"foursquare_id,omitempty"`
	FoursquareType string    `json:"foursquare_type,omitempty"`
}

func (m *Venue) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Venue) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type PollAnswer struct {
	PollId    string  `json:"poll_id"`
	User      *User   `json:"user"`
	OptionIds []int32 `json:"option_ids"`
}

func (m *PollAnswer) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *PollAnswer) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func (m *BotCommand) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *BotCommand) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int32  `json:"width"`
	Height       int32  `json:"height"`
	FileSize     int32  `json:"file_size,omitempty"`
}

func (m *PhotoSize) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *PhotoSize) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Video struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int32      `json:"width"`
	Height       int32      `json:"height"`
	Duration     int32      `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int32      `json:"file_size,omitempty"`
}

func (m *Video) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Video) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type File struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int32  `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

func (m *File) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *File) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

func (m *ForceReply) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ForceReply) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int32      `json:"file_size,omitempty"`
}

func (m *Document) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Document) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int32  `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int32  `json:"file_size,omitempty"`
}

func (m *Voice) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Voice) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserId      int32  `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

func (m *Contact) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Contact) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type UserProfilePhotos struct {
	TotalCount int32          `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}

func (m *UserProfilePhotos) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *UserProfilePhotos) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	Url                          string        `json:"url,omitempty"`
	LoginUrl                     *LoginUrl     `json:"login_url,omitempty"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

func (m *InlineKeyboardButton) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *InlineKeyboardButton) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title,omitempty"`
	UntilDate             int64  `json:"until_date,omitempty"`
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
}

func (m *ChatMember) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ChatMember) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
}

func (m *ChatPermissions) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ChatPermissions) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Message struct {
	MessageId             int32                 `json:"message_id"`
	From                  *User                 `json:"from,omitempty"`
	Date                  int32                 `json:"date"`
	Chat                  *Chat                 `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat                 `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId  int32                 `json:"forward_from_message_id,omitempty"`
	ForwardSignature      string                `json:"forward_signature,omitempty"`
	ForwardSenderName     string                `json:"forward_sender_name,omitempty"`
	ForwardDate           int32                 `json:"forward_date,omitempty"`
	ReplyToMessage        *Message              `json:"reply_to_message,omitempty"`
	EditDate              int32                 `json:"edit_date,omitempty"`
	MediaGroupId          string                `json:"media_group_id,omitempty"`
	AuthorSignature       string                `json:"author_signature,omitempty"`
	Text                  string                `json:"text,omitempty"`
	Entities              []*MessageEntity      `json:"entities,omitempty"`
	CaptionEntities       []*MessageEntity      `json:"caption_entities,omitempty"`
	Audio                 *Audio                `json:"audio,omitempty"`
	Document              *Document             `json:"document,omitempty"`
	Animation             *Animation            `json:"animation,omitempty"`
	Game                  *Game                 `json:"game,omitempty"`
	Photo                 []*PhotoSize          `json:"photo,omitempty"`
	Sticker               *Sticker              `json:"sticker,omitempty"`
	Video                 *Video                `json:"video,omitempty"`
	Voice                 *Voice                `json:"voice,omitempty"`
	VideoNote             *VideoNote            `json:"video_note,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	Contact               *Contact              `json:"contact,omitempty"`
	Location              *Location             `json:"location,omitempty"`
	Venue                 *Venue                `json:"venue,omitempty"`
	Poll                  *Poll                 `json:"poll,omitempty"`
	Dice                  *Dice                 `json:"dice,omitempty"`
	NewChatMembers        []*User               `json:"new_chat_members,omitempty"`
	LeftChatMember        *User                 `json:"left_chat_member,omitempty"`
	NewChatTitle          string                `json:"new_chat_title,omitempty"`
	NewChatPhoto          []*PhotoSize          `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool                  `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool                  `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool                  `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool                  `json:"channel_chat_created,omitempty"`
	MigrateToChatId       int64                 `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId     int64                 `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message              `json:"pinned_message,omitempty"`
	Invoice               *Invoice              `json:"invoice,omitempty"`
	SuccessfulPayment     *SuccessfulPayment    `json:"successful_payment,omitempty"`
	ConnectedWebsite      string                `json:"connected_website,omitempty"`
	PassportData          *PassportData         `json:"passport_data,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *Message) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Message) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Dice struct {
	Value int32 `json:"value"`
}

func (m *Dice) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Dice) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ReplyKeyboardMarkup struct {
	ResizeKeyboard  bool `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	Selective       bool `json:"selective,omitempty"`
}

func (m *ReplyKeyboardMarkup) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ReplyKeyboardMarkup) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact,omitempty"`
	RequestLocation bool                    `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

func (m *KeyboardButton) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *KeyboardButton) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type KeyboardButtonPollType struct {
	Type string `json:"type,omitempty"`
}

func (m *KeyboardButtonPollType) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *KeyboardButtonPollType) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

func (m *ReplyKeyboardRemove) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *ReplyKeyboardRemove) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}
