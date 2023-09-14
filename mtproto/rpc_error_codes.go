package mtproto

import (
	"github.com/go-kratos/kratos/pkg/ecode"
)

// https://core.telegram.org/api/errors
/*
Error handling
There will be errors when working with the API, and they must be correctly handled on the client.

An error is characterized by several parameters:

Error Code
Numerical value similar to HTTP status. Contains information on the type of error that occurred: for example, a data input error, privacy error, or server error. This is a required parameter.

Error Type
A string literal in the form of /[A-Z_0-9]+/, which summarizes the problem. For example, AUTH_KEY_UNREGISTERED. This is an optional parameter.

Error Constructors
There should be a way to handle errors that are returned in rpc_error constructors.

Below is a list of error codes and their meanings:
*/

var (
	/*
		ERROR_CODE_OK = 0;
	*/
	// ErrCodeOk = ecode.New(0)

	/*
		303 SEE_OTHER
		The request must be repeated, but directed to a different data center.

		Examples of Errors:
			FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
			PHONE_MIGRATE_X: the phone number a user is trying to use for authorization is associated with a different data center.
			NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
			USER_MIGRATE_X: the user whose identity is being used to execute queries is associated with a different data center (for registration)

		In all these cases, the error description’s string literal contains the number of the data center (instead of the X) to which the repeated query must be sent.
		More information about redirects between data centers »
	*/
	ErrSeeOther = ecode.New(303)

	/*
		400 BAD_REQUEST
		The query contains errors. In the event that a request was created using a form and contains user generated data, the user should be notified that the data must be corrected before the query is repeated.
	*/
	ErrBadRequest = ecode.New(400)

	/*
		401 UNAUTHORIZED
		There was an unauthorized attempt to use functionality available only to authorized users.

		Examples of Errors:
			AUTH_KEY_UNREGISTERED: The key is not registered in the system
			AUTH_KEY_INVALID: The key is invalid
			USER_DEACTIVATED: The user has been deleted/deactivated
			SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
			SESSION_EXPIRED: The authorization has expired
			AUTH_KEY_PERM_EMPTY: The method is unavailble for temporary authorization key, not bound to permanent
	*/
	ErrUnauthorized = ecode.New(401)

	/*
		403 FORBIDDEN
		Privacy violation. For example, an attempt to write a message to someone who has blacklisted the current user.
	*/
	ErrForbidden = ecode.New(403)

	/*
		404 NOT_FOUND
		An attempt to invoke a non-existent object, such as a method.
	*/
	ErrNotFound = ecode.New(404)

	/*
		406 NOT_ACCEPTABLE
		Similar to 400 BAD_REQUEST, but the app should not display any error messages to user in UI as a result of this response. The error message will be delivered via updateServiceNotification instead.
	*/
	ErrNotAcceptable = ecode.New(406)

	/*
		420 FLOOD
		The maximum allowed number of attempts to invoke the given method with the given input parameters has been exceeded. For example, in an attempt to request a large number of text messages (SMS) for the same phone number.

		Error Example:
			FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
	*/
	ErrFlood = ecode.New(420)

	/*
		500 INTERNAL
		An internal server error occurred while a request was being processed; for example, there was a disruption while accessing a database or file storage.

		If a client receives a 500 error, or you believe this error should not have occurred, please collect as much information as possible about the query and error and send it to the developers.
	*/
	ErrInternal = ecode.New(500)

	/*
		Other Error Codes
		If a server returns an error with a code other than the ones listed above, it may be considered the same as a 500 error and treated as an internal server error.
	*/

	// | -503 | Timeout | Timeout while fetching data |
	ErrTimeOut503 = ecode.New(5030000)

	// db error
	ErrNotReturnClient = ecode.New(700)
)

// 303 SEE_OTHER
//
// | 303 | NETWORK_MIGRATE_X | Repeat the query to data-center X |
// | 303 | PHONE_MIGRATE_X | Repeat the query to data-center X |
//
// FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
func NewErrFileMigrateX(dc int32) *ecode.Status {
	return ecode.Errorf(ErrSeeOther, "FILE_MIGRATE_%d", dc)
}

// PHONE_MIGRATE_X: the phone number a user is trying to use for authorization is associated with a different data center.
func NewErrPhoneMigrateX(dc int32) *ecode.Status {
	return ecode.Errorf(ErrSeeOther, "PHONE_MIGRATE_%d", dc)
}

// NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
func NewErrNetworkMigrateX(dc int32) *ecode.Status {
	return ecode.Errorf(ErrSeeOther, "NETWORK_MIGRATE_%d", dc)
}

// USER_MIGRATE_X: the user whose identity is being used to execute queries is associated with a different data center (for registration)
func NewErrUserMigrateX(dc int32) *ecode.Status {
	return ecode.Errorf(ErrSeeOther, "USER_MIGRATE_%d", dc)
}

// 420 FLOOD
//
// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
func NewErrFloodWaitX(second int32) *ecode.Status {
	return ecode.Errorf(ErrFlood, "FLOOD_WAIT_%d", second)
}

// 420	SLOWMODE_WAIT_X	Slowmode is enabled in this chat: you must wait for the specified number of seconds before sending another message to the chat.
func NewSlowModeWaitX(second int32) *ecode.Status {
	return ecode.Errorf(ErrFlood, "SLOWMODE_WAIT_%d", second)
}

// 400 BAD_REQUEST
var (
	// ABOUT_TOO_LONG	400	The provided bio is too long
	ErrAboutTooLong = ecode.Error(ErrBadRequest, "ABOUT_TOO_LONG")

	// ACCESS_TOKEN_EXPIRED	400	Bot token expired
	ErrAccessTokenExpired = ecode.Error(ErrBadRequest, "ACCESS_TOKEN_EXPIRED")

	// ACCESS_TOKEN_INVALID	400	The provided token is not valid
	ErrAccessTokenInvalid = ecode.Error(ErrBadRequest, "ACCESS_TOKEN_INVALID")

	// METHOD_NOT_IMPL: The method not impl
	ErrInputRequestInvalid = ecode.Error(ErrBadRequest, "INPUT_REQUEST_INVALID")

	// METHOD_NOT_IMPL: The method not impl
	ErrMethodNotImpl = ecode.Error(ErrBadRequest, "METHOD_NOT_IMPL")

	// TYPE_CONSTRUCTOR_INVALID: The type constructor is invalid
	ErrTypeConstructorInvalid = ecode.Error(ErrBadRequest, "TYPE_CONSTRUCTOR_INVALID")

	// | 400 | API_ID_INVALID | API ID invalid |
	ErrApiIdInvalid = ecode.Error(ErrBadRequest, "API_ID_INVALID")

	// | 400 | API_ID_PUBLISHED_FLOOD | This API id was published somewhere, you can't use it now |
	ErrApiIdPublishedFlood = ecode.Error(ErrBadRequest, "API_ID_PUBLISHED_FLOOD")

	// 400	BOT_INLINE_DISABLED	This bot can't be used in inline mode
	ErrBotInlineDisabled = ecode.Error(ErrBadRequest, "BOT_INLINE_DISABLED")

	// 400	BOT_INVALID	This is not a valid bot
	ErrBotInvalid = ecode.Error(ErrBadRequest, "BOT_INVALID")

	// | 400 | BOT_METHOD_INVALID | This method can't be used by a bot |
	ErrBotMethodInvalid = ecode.Error(ErrBadRequest, "BOT_METHOD_INVALID")

	// | 400 | API_SERVER_NEEDED | This method be used by api server |
	ErrApiServerNeeded = ecode.Error(ErrBadRequest, "API_SERVER_NEEDED")

	// | 400 | INPUT_REQUEST_TOO_LONG | The request is too big |
	ErrInputRequestTooLong = ecode.Error(ErrBadRequest, "INPUT_REQUEST_TOO_LONG")

	// | 400 | PHONE_NUMBER_APP_SIGNUP_FORBIDDEN | You can't sign up using this app |
	ErrPhoneNumberAppSignupForbidden = ecode.Error(ErrBadRequest, "PHONE_NUMBER_APP_SIGNUP_FORBIDDEN")

	// | 400 | PHONE_NUMBER_BANNED | The provided phone number is banned from telegram |
	ErrPhoneNumberBanned = ecode.Error(ErrBadRequest, "PHONE_NUMBER_BANNED")

	// | 400 | PHONE_NUMBER_FLOOD | You asked for the code too many times. |
	ErrPhoneNumberFlood = ecode.Error(ErrBadRequest, "PHONE_NUMBER_FLOOD")

	// | 400 | PHONE_NUMBER_INVALID | Invalid phone number |
	ErrPhoneNumberInvalid = ecode.Error(ErrBadRequest, "PHONE_NUMBER_INVALID")

	// | 400 | PHONE_PASSWORD_PROTECTED | This phone is password protected |
	ErrPhonePasswordProtected = ecode.Error(ErrBadRequest, "PHONE_PASSWORD_PROTECTED")

	// | 400 | SMS_CODE_CREATE_FAILED | An error occurred while creating the SMS code |
	ErrSmsCreateFailed = ecode.Error(ErrBadRequest, "SMS_CODE_CREATE_FAILED")

	// 400	FIRSTNAME_INVALID	Invalid first name
	ErrFirstNameInvalid = ecode.Error(ErrBadRequest, "FIRSTNAME_INVALID")

	// 400	LASTNAME_INVALID	Invalid last name
	ErrLastNameInvalid = ecode.Error(ErrBadRequest, "LASTNAME_INVALID")

	// 400	PHONE_CODE_EMPTY	phone_code from a SMS is empty
	ErrPhoneCodeEmpty = ecode.Error(ErrBadRequest, "PHONE_CODE_EMPTY")

	// 400	PHONE_CODE_EXPIRED	SMS expired
	ErrPhoneCodeExpired = ecode.Error(ErrBadRequest, "PHONE_CODE_EXPIRED")

	// 400	PHONE_CODE_INVALID	Invalid SMS code was sent
	ErrPhoneCodeInvalid = ecode.Error(ErrBadRequest, "PHONE_CODE_INVALID")

	// 400	PHONE_NUMBER_OCCUPIED	The phone number is already in use)
	ErrPhoneNumberOccupied = ecode.Error(ErrBadRequest, "PHONE_NUMBER_OCCUPIED")

	// PHONE_NUMBER_UNOCCUPIED: The phone number is not yet being used
	ErrPhoneNumberUnoccupied = ecode.Error(ErrBadRequest, "PHONE_NUMBER_UNOCCUPIED")

	// | 400 | CHANNEL_INVALID | The provided channel is invalid |
	ErrChannelInvalid = ecode.Error(ErrBadRequest, "CHANNEL_INVALID")

	// | 400 | CHANNEL_PRIVATE | You haven't joined this channel/supergroup |
	ErrChannelPrivate = ecode.Error(ErrBadRequest, "CHANNEL_PRIVATE")

	// | 400 | CHAT_ADMIN_REQUIRED | You must be an admin in this chat to do this |
	ErrChatAdminRequired = ecode.Error(ErrBadRequest, "CHAT_ADMIN_REQUIRED")

	// | 400 | EXTERNAL_URL_INVALID | External URL invalid |
	ErrExternalUrlInvalid = ecode.Error(ErrBadRequest, "EXTERNAL_URL_INVALID")

	// | 400 | FILE_PARTS_INVALID | The number of file parts is invalid |
	ErrFilePartsInvalid = ecode.Error(ErrBadRequest, "FILE_PARTS_INVALID")

	// FILE_PART_INVALID: The file part number is invalid. The value is not between 0 and 2,999.
	ErrFilePartInvalid = ecode.Error(ErrBadRequest, "FILE_PART_INVALID")

	// | 400 | FILE_PART_LENGTH_INVALID | The length of a file part is invalid |
	ErrFilePartLengthInvalid = ecode.Error(ErrBadRequest, "FILE_PART_LENGTH_INVALID")

	// FILE_PART_SIZE_INVALID - 512KB cannot be evenly divided by part_size
	ErrFilePartSizeInvalid = ecode.Error(ErrBadRequest, "FILE_PART_SIZE_INVALID")

	// FILE_PART_TOO_BIG: The size limit (512 KB) for the content of the file part has been exceeded
	ErrFilePartTooBig = ecode.Error(ErrBadRequest, "FILE_PART_TOO_BIG")

	// FILE_PART_SIZE_CHANGED - The part size is different from the size of one of the previous parts in the same file
	ErrFilePartSizeChanged = ecode.Error(ErrBadRequest, "FILE_PART_SIZE_CHANGED")

	// MD5_CHECKSUM_INVALID: The file’s checksum did not match the md5_checksum parameter
	ErrCheckSumInvalid = ecode.Error(ErrBadRequest, "MD5_CHECKSUM_INVALID")

	// | 400 | IMAGE_PROCESS_FAILED | Failure while processing image |
	ErrImageProcessFailed = ecode.Error(ErrBadRequest, "IMAGE_PROCESS_FAILED")

	// | 400 | INPUT_USER_DEACTIVATED | The specified user was deleted |
	ErrInputUserDeactivated = ecode.Error(ErrBadRequest, "INPUT_USER_DEACTIVATED")

	// | 400 | MEDIA_CAPTION_TOO_LONG | The caption is too long |
	ErrMediaCaptionTooLong = ecode.Error(ErrBadRequest, "MEDIA_CAPTION_TOO_LONG")

	// | 400 | MEDIA_EMPTY | The provided media object is invalid |
	ErrMediaEmpty = ecode.Error(ErrBadRequest, "MEDIA_EMPTY")

	// | 400 | MEDIA_INVALID | Media invalid |
	ErrMediaInvalid = ecode.Error(ErrBadRequest, "MEDIA_INVALID")

	// | 400 | PHOTO_EXT_INVALID | The extension of the photo is invalid |
	ErrPhotoExtInvalid = ecode.Error(ErrBadRequest, "PHOTO_EXT_INVALID")

	// | 400 | PHOTO_INVALID_DIMENSIONS | The photo dimensions are invalid |
	ErrPhotoInvalidDimensions = ecode.Error(ErrBadRequest, "PHOTO_INVALID_DIMENSIONS")

	// | 400 | PHOTO_SAVE_FILE_INVALID |
	ErrPhotoSaveFileInvalid = ecode.Error(ErrBadRequest, "PHOTO_SAVE_FILE_INVALID")

	// | 400 | USER_BANNED_IN_CHANNEL | You're banned from sending messages in supergroups/channels |
	ErrUserBannedInChannel = ecode.Error(ErrBadRequest, "USER_BANNED_IN_CHANNEL")

	// | 400 | USER_IS_BLOCKED | You were blocked by this user |
	ErrUserIsBlocked = ecode.Error(ErrBadRequest, "USER_IS_BLOCKED")

	// | 400 | USER_IS_BOT | Bots can't send messages to other bots |
	ErrUserIsBot = ecode.Error(ErrBadRequest, "USER_IS_BOT")

	// | 400 | WEBPAGE_CURL_FAILED | Failure while fetching the webpage with cURL |
	ErrWebpageCurlFailed = ecode.Error(ErrBadRequest, "WEBPAGE_CURL_FAILED")

	// | 400 | WEBPAGE_MEDIA_EMPTY | Webpage media empty |
	ErrWebpageMediaEmpty = ecode.Error(ErrBadRequest, "WEBPAGE_MEDIA_EMPTY")

	// | `INPUT_METHOD_INVALID` | The method called is invalid
	ErrInputMethodInvalid = ecode.Error(ErrBadRequest, "INPUT_METHOD_INVALID")

	// | 400 | ENCRYPTED_MESSAGE_INVALID | Encrypted message is incorrect |
	ErrEncryptedMessageInvalid = ecode.Error(ErrBadRequest, "ENCRYPTED_MESSAGE_INVALID")

	// | 400 | TEMP_AUTH_KEY_ALREADY_BOUND | The passed temporary key is already bound to another perm_auth_key_id |
	ErrTempAuthKeyAlreadyBound = ecode.Error(ErrBadRequest, "TEMP_AUTH_KEY_ALREADY_BOUND")

	// | 400 | TEMP_AUTH_KEY_EMPTY | The request was not performed with a temporary authorization key |
	ErrTempAuthKeyEmpty = ecode.Error(ErrBadRequest, "TEMP_AUTH_KEY_EMPTY")

	// 400	TOKEN_INVALID	The provided token is invalid
	ErrTokenInvalid = ecode.Error(ErrBadRequest, "TOKEN_INVALID")

	// 400	PEER_ID_INVALID	The provided peer id is invalid
	ErrPeerIdInvalid = ecode.Error(ErrBadRequest, "PEER_ID_INVALID")

	// 400	TTL_DAYS_INVALID	The provided TTL is invalid
	ErrTtlDaysInvalid = ecode.Error(ErrBadRequest, "TTL_DAYS_INVALID")

	// 400	PRIVACY_KEY_INVALID	The privacy key is invalid
	ErrPrivacyKeyInvalid = ecode.Error(ErrBadRequest, "PRIVACY_KEY_INVALID")

	// 400	CONTACT_ID_INVALID	The provided contact ID is invalid
	ErrContactIdInvalid = ecode.Error(ErrBadRequest, "CONTACT_ID_INVALID")

	// PHONE_CODE_HASH_EMPTY: phone_code_hash is missing
	ErrPhoneCodeHashEmpty = ecode.Error(ErrBadRequest, "PHONE_CODE_HASH_EMPTY")

	// USERS_TOO_FEW: Not enough users (to create a chat, for example)
	ErrUsersTooFew = ecode.Error(ErrBadRequest, "USERS_TOO_FEW")

	// USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
	ErrUsersTooMuch = ecode.Error(ErrBadRequest, "USERS_TOO_MUCH")

	// 400	SESSION_PASSWORD_NEEDED	The user has enabled 2FA, more steps are needed
	ErrSessionPasswordNeeded = ecode.Error(ErrBadRequest, "SESSION_PASSWORD_NEEDED")

	// 400	BUTTON_DATA_INVALID	The data of one or more of the buttons you provided is invalid
	ErrButtonDataInvalid = ecode.Error(ErrBadRequest, "BUTTON_DATA_INVALID")

	// 400	BUTTON_TYPE_INVALID	The type of one or more of the buttons you provided is invalid
	ErrButtonTypeInvalid = ecode.Error(ErrBadRequest, "BUTTON_TYPE_INVALID")

	// 400	BUTTON_URL_INVALID	Button URL invalid
	ErrButtonUrlInvalid = ecode.Error(ErrBadRequest, "BUTTON_URL_INVALID")

	// 400	CHAT_ID_INVALID	The provided chat id is invalid
	ErrChatIdInvalid = ecode.Error(ErrBadRequest, "CHAT_ID_INVALID")

	// 400	CHAT_RESTRICTED	You can't send messages in this chat, you were restricted
	ErrChatRestricted = ecode.Error(ErrBadRequest, "CHAT_RESTRICTED")

	// 400	ENTITY_MENTION_USER_INVALID	You mentioned an invalid user
	ErrEntityMentionUserInvalid = ecode.Error(ErrBadRequest, "ENTITY_MENTION_USER_INVALID")

	// 400	MESSAGE_EMPTY	The provided message is empty
	ErrMessageEmpty = ecode.Error(ErrBadRequest, "MESSAGE_EMPTY")

	// 400	MESSAGE_TOO_LONG	The provided message is too long
	ErrMessageTooLong = ecode.Error(ErrBadRequest, "MESSAGE_TOO_LONG")

	// 400	MSG_ID_INVALID	Provided reply_to_msg_id is invalid
	ErrMsgIdInvalid = ecode.Error(ErrBadRequest, "MSG_ID_INVALID")

	// 400	MESSAGE_ID_INVALID	400	The specified message ID is invalid or you can't do that operation on such message
	ErrMessageIdInvalid = ecode.Error(ErrBadRequest, "MESSAGE_ID_INVALID")

	// 400	REPLY_MARKUP_INVALID	The provided reply markup is invalid
	ErrReplyMarkupInvalid = ecode.Error(ErrBadRequest, "REPLY_MARKUP_INVALID")

	// 400	YOU_BLOCKED_USER	You blocked this user
	ErrYouBlockedUser = ecode.Error(ErrBadRequest, "YOU_BLOCKED_USER")

	// 400	CHAT_TITLE_EMPTY	No chat title provided
	ErrChatTitleEmpty = ecode.Error(ErrBadRequest, "CHAT_TITLE_EMPTY")

	// 400	FOLDER_ID_INVALID	Invalid folder ID
	ErrFolderIdInvalid = ecode.Error(ErrBadRequest, "FOLDER_ID_INVALID")

	// 400	INPUT_CONSTRUCTOR_INVALID	The provided constructor is invalid
	ErrInputConstructorInvalid = ecode.Error(ErrBadRequest, "INPUT_CONSTRUCTOR_INVALID")

	// 400	USER_ALREADY_PARTICIPANT	The user is already in the group
	ErrUserAlreadyParticipant = ecode.Error(ErrBadRequest, "USER_ALREADY_PARTICIPANT")

	// CHAT_NOT_MODIFIED
	// | `CHAT_NOT_MODIFIED` | The chat settings were not modified
	ErrChatNotModified = ecode.Error(ErrBadRequest, "CHAT_NOT_MODIFIED")

	// 400	USER_NOT_PARTICIPANT	You're not a member of this supergroup/channel
	ErrUserNotParticipant = ecode.Error(ErrBadRequest, "USER_NOT_PARTICIPANT")

	// 400	INVITE_HASH_EMPTY	The invite hash is empty
	ErrInviteHashEmpty = ecode.Error(ErrBadRequest, "INVITE_HASH_EMPTY")

	// 400	INVITE_HASH_EXPIRED	The invite link has expired
	ErrInviteHashExpired = ecode.Error(ErrBadRequest, "INVITE_HASH_EXPIRED")

	// 400	INVITE_HASH_INVALID	The invite hash is invalid
	ErrInviteHashInvalid = ecode.Error(ErrBadRequest, "INVITE_HASH_INVALID")

	// 400	USER_KICKED	This user was kicked from this supergroup/channel
	ErrUserKicked = ecode.Error(ErrBadRequest, "USER_KICKED")

	// 400	USER_ID_INVALID	The provided user ID is invalid
	ErrUserIdInvalid = ecode.Error(ErrBadRequest, "USER_ID_INVALID")

	// 400	ADMINS_TOO_MUCH	There are too many admins
	ErrAdminsTooMuch = ecode.Error(ErrBadRequest, "ADMINS_TOO_MUCH")

	// 400	BOT_CHANNELS_NA	Bots can't edit admin privileges
	ErrBotChannelsNA = ecode.Error(ErrBadRequest, "BOT_CHANNELS_NA")

	// 400	USER_CREATOR	You can't leave this channel, because you're its creator
	ErrUserCreator = ecode.Error(ErrBadRequest, "USER_CREATOR")

	// 400	USER_LEFT_CHAT
	// } else if (error == qstr("USER_LEFT_CHAT")) {
	ErrUserLeftChat = ecode.Error(ErrBadRequest, "USER_LEFT_CHAT")

	// 400	USER_ADMIN_INVALID	You're not an admin
	ErrUserAdminInvalid = ecode.Error(ErrBadRequest, "USER_ADMIN_INVALID")

	// 400	USERNAME_INVALID	The provided username is not valid
	ErrUsernameInvalid = ecode.Error(ErrBadRequest, "USERNAME_INVALID")

	// 400	MESSAGE_IDS_EMPTY	No message ids were provided
	ErrMessageIdsEmpty = ecode.Error(ErrBadRequest, "MESSAGE_IDS_EMPTY")

	// PIN_RESTRICTED	400	You can't pin messages in private chats with other people
	ErrPinRestricted = ecode.Error(ErrBadRequest, "PIN_RESTRICTED")

	// poll
	// BOT_POLLS_DISABLED	400	You cannot create polls under a bot account
	ErrBotPollsDisabled = ecode.Error(ErrBadRequest, "BOT_POLLS_DISABLED")

	// OPTIONS_TOO_MUCH	400	You defined too many options for the poll
	ErrOptionsTooMuch = ecode.Error(ErrBadRequest, "OPTIONS_TOO_MUCH")

	// POLL_OPTION_DUPLICATE	400	A duplicate option was sent in the same poll
	ErrOptionDuplicate = ecode.Error(ErrBadRequest, "POLL_OPTION_DUPLICATE")

	// POLL_UNSUPPORTED	400	This layer does not support polls in the issued method
	ErrPollUnsupported = ecode.Error(ErrBadRequest, "POLL_UNSUPPORTED")

	// 400	SEARCH_QUERY_EMPTY	The search query is empty
	ErrSearchQueryEmpty = ecode.Error(ErrBadRequest, "SEARCH_QUERY_EMPTY")

	// 400	QUERY_TOO_SHORT	The query string is too short
	ErrQueryTooShort = ecode.Error(ErrBadRequest, "QUERY_TOO_SHORT")

	// 400	LOCATION_INVALID	The provided location is invalid
	ErrLocationInvalid = ecode.Error(ErrBadRequest, "LOCATION_INVALID")

	// 400	RANDOM_LENGTH_INVALID	Random length invalid
	ErrRandomLengthInvalid = ecode.Error(ErrBadRequest, "RANDOM_LENGTH_INVALID")

	// 400	CALL_ALREADY_ACCEPTED	The call was already accepted
	ErrCallAlreadyAccepted = ecode.Error(ErrBadRequest, "CALL_ALREADY_ACCEPTED")

	// 400	CALL_ALREADY_DECLINED	The call was already declined
	ErrCallAlreadyDeclined = ecode.Error(ErrBadRequest, "CALL_ALREADY_DECLINED")

	// 400	CALL_PEER_INVALID	The provided call peer object is invalid
	ErrCallPeerInvalid = ecode.Error(ErrBadRequest, "CALL_PEER_INVALID")

	// 400	CALL_PROTOCOL_FLAGS_INVALID	Call protocol flags invalid
	ErrCallProtocolFlagsInvalid = ecode.Error(ErrBadRequest, "CALL_PROTOCOL_FLAGS_INVALID")

	// 400	PARTICIPANT_VERSION_OUTDATED	The other participant does not use an up to date telegram client with support for calls
	ErrParticipantVersionOutdated = ecode.Error(ErrBadRequest, "PARTICIPANT_VERSION_OUTDATED")

	// STICKERSET_INVALID	400	The provided sticker set is invalid
	ErrStickersetInvalid = ecode.Error(ErrBadRequest, "STICKERSET_INVALID")

	// STICKERS_EMPTY	400	No sticker provided
	ErrStickersEmpty = ecode.Error(ErrBadRequest, "STICKERS_EMPTY")

	// STICKER_EMOJI_INVALID	400	Sticker emoji invalid
	ErrStickerEmojiInvalid = ecode.Error(ErrBadRequest, "STICKER_EMOJI_INVALID")

	// STICKER_FILE_INVALID	400	Sticker file invalid
	ErrStickerFileInvalid = ecode.Error(ErrBadRequest, "STICKER_FILE_INVALID")

	//STICKER_ID_INVALID	400	The provided sticker ID is invalid
	ErrStickerIdInvalid = ecode.Error(ErrBadRequest, "STICKER_ID_INVALID")

	// STICKER_INVALID	400	The provided sticker is invalid
	ErrStickerInvalid = ecode.Error(ErrBadRequest, "STICKER_INVALID")

	// STICKER_PNG_DIMENSIONS	400	Sticker png dimensions invalid
	ErrStickerPngDimensions = ecode.Error(ErrBadRequest, "STICKER_PNG_DIMENSIONS")

	// 400	LANG_PACK_INVALID	The provided language pack is invalid
	ErrLangPackInvalid = ecode.Error(ErrBadRequest, "LANG_PACK_INVALID")

	// if ("LANG_CODE_NOT_SUPPORTED".equals(error.text)) {
	ErrLangCodeNotSupported = ecode.Error(ErrBadRequest, "LANG_CODE_NOT_SUPPORTED")

	// 400	OFFSET_INVALID	The provided offset is invalid
	ErrOffsetInvalid = ecode.Error(ErrBadRequest, "OFFSET_INVALID")

	// 400	QUERY_ID_EMPTY	The query ID is empty
	ErrQueryIdEmpty = ecode.Error(ErrBadRequest, "QUERY_ID_EMPTY")

	// 400 - AUTH_TOKEN_INVALID, an invalid authorization token was provided
	ErrAuthTokenInvalid = ecode.Error(ErrBadRequest, "AUTH_TOKEN_INVALID")

	// 400 - AUTH_TOKEN_EXPIRED, the provided authorization token has expired and the updated QR-code must be re-scanned
	ErrAuthTokenExpired = ecode.Error(ErrBadRequest, "AUTH_TOKEN_EXPIRED")

	// 400 - AUTH_TOKEN_ALREADY_ACCEPTED, the authorization token was already used
	ErrAuthTokenAccepted = ecode.Error(ErrBadRequest, "AUTH_TOKEN_ALREADY_ACCEPTED")

	// 400 - THEME_INVALID
	ErrThemeInvalid = ecode.Error(ErrBadRequest, "THEME_INVALID")

	// 400 - THEME_FORMAT_INVALID
	ErrThemeFormatInvalid = ecode.Error(ErrBadRequest, "THEME_FORMAT_INVALID")

	// 400
	ErrUserPasswordNeeded    = ecode.Error(ErrBadRequest, "USER_PASSWORD_NEEDED")
	ErrUserNameNotExist      = ecode.Error(ErrBadRequest, "USERNAME_NOT_EXIST")
	ErrUserPasswordNotSet    = ecode.Error(ErrBadRequest, "USER_PASSWORD_NOT_SET")
	ErrPasswordVerifyInvalid = ecode.Error(ErrBadRequest, "PASSWORD_VERIFY_INVALID")
	ErrIpAddressBanned       = ecode.Error(ErrBadRequest, "IP_ADDRESS_BANNED")
	ErrUserBindedIpAddress   = ecode.Error(ErrBadRequest, "USER_BINDED_IP_ADDRESS")
	ErrInviteCodeInvalid     = ecode.Error(ErrBadRequest, "INVITE_CODE_INVALID")
	ErrNotModifyUserName     = ecode.Error(ErrBadRequest, "NOT_MODIFY_USERNAME")
)

// FILE_PART_Х_MISSING: Part X (where X is a number) of the file is missing from storage. Try repeating the method call to resave the part.
func NewFilePartXMissing(x int32) *ecode.Status {
	return ecode.Errorf(ErrBadRequest, "FILE_PART_%d_MISSING", x)
}

// 401 UNAUTHORIZED
var (
	// AUTH_KEY_UNREGISTERED: The key is not registered in the system
	ErrAuthKeyUnregistered = ecode.Error(ErrUnauthorized, "AUTH_KEY_UNREGISTERED")

	// AUTH_KEY_INVALID: The key is invalid
	ErrAuthKeyInvalid = ecode.Error(ErrUnauthorized, "AUTH_KEY_INVALID")

	// USER_DEACTIVATED: The user has been deleted/deactivated
	ErrUserDeactivated = ecode.Error(ErrUnauthorized, "USER_DEACTIVATED")

	// SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
	ErrSessionRevoked = ecode.Error(ErrUnauthorized, "SESSION_REVOKED")

	// SESSION_EXPIRED: The authorization has expired
	ErrSessionExpired = ecode.Error(ErrUnauthorized, "SESSION_EXPIRED")

	// AUTH_KEY_PERM_EMPTY: The method is unavailble for temporary authorization key, not bound to permanent
	ErrAuthKeyPermEmpty = ecode.Error(ErrUnauthorized, "AUTH_KEY_PERM_EMPTY")
)

// 403 FORBIDDEN
var (
	// | 403 | CHAT_SEND_MEDIA_FORBIDDEN | You can't send media in this chat |
	ErrChatSendMediaForbidden = ecode.Error(ErrForbidden, "CHAT_SEND_MEDIA_FORBIDDEN")

	// 403	You can't send gifs in this chat
	ErrChatSendGifsForbidden = ecode.Error(ErrForbidden, "CHAT_SEND_GIFS_FORBIDDEN")

	// 403	You can't send stickers in this chat
	ErrChatSendStickersForbidden = ecode.Error(ErrForbidden, "CHAT_SEND_STICKERS_FORBIDDEN")

	// | 403 | CHAT_WRITE_FORBIDDEN | You can't write in this chat |
	ErrChatWriteForbidden = ecode.Error(ErrForbidden, "CHAT_WRITE_FORBIDDEN")

	// 403	USER_RESTRICTED	You're spamreported, you can't create channels or chats.
	ErrUserRestricted = ecode.Error(ErrForbidden, "USER_RESTRICTED")

	// 403	MESSAGE_AUTHOR_REQUIRED	Message author required
	ErrMessageAuthorRequired = ecode.Error(ErrForbidden, "MESSAGE_AUTHOR_REQUIRED")

	// 403	USER_NOT_MUTUAL_CONTACT	The provided user is not a mutual contact
	ErrUserNotMutualContact = ecode.Error(ErrForbidden, "USER_NOT_MUTUAL_CONTACT")

	// 403	USER_PRIVACY_RESTRICTED	The user's privacy settings do not allow you to do this
	ErrUserPrivacyRestricted = ecode.Error(ErrForbidden, "USER_PRIVACY_RESTRICTED")

	// 403	CHAT_ADMIN_INVITE_REQUIRED	You do not have the rights to do this
	ErrChatAdminInviteRequired = ecode.Error(ErrForbidden, "CHAT_ADMIN_INVITE_REQUIRED")

	// 403	RIGHT_FORBIDDEN	Your admin rights do not allow you to do this
	ErrRightForbidden = ecode.Error(ErrForbidden, "RIGHT_FORBIDDEN")

	// 403	MESSAGE_DELETE_FORBIDDEN	You can't delete one of the messages you tried to delete, most likely because it is a service message.
	ErrMessageDeleteForbidden = ecode.Error(ErrForbidden, "MESSAGE_DELETE_FORBIDDEN")
)

// 406 NOT_ACCEPTABLE
var (
	// | 406 | PHONE_PASSWORD_FLOOD | You have tried logging in too many times |
	ErrPhonePasswordFlood = ecode.Error(ErrNotAcceptable, "PHONE_PASSWORD_FLOOD")
)

// 500 InternalServerError
var (
	// | 500 | INTERNAL_SERVER_ERROR |  |
	ErrInternelServerError = ecode.Error(ErrInternal, "INTERNAL_SERVER_ERROR")
)

// 503 Timeout
var (
	// | -503 | Timeout | Timeout while fetching data |
	ErrTimeout = ecode.Error(ErrTimeOut503, "Timeout")
)

var (
	// db error
	// TLRpcErrorCodes_NOTRETURN_CLIENT TLRpcErrorCodes = 700
	ErrPushRpcClient = ecode.Error(ErrNotReturnClient, "NOTRETURN_CLIENT")
)

//)

/*
ACTIVE_USER_REQUIRED	401	The method is only available to already activated users
ADMINS_TOO_MUCH	400	Too many admins
ADMIN_RANK_EMOJI_NOT_ALLOWED	400	The given admin title or rank was invalid (possibly larger than 16 characters)
API_ID_INVALID	400	The api_id/api_hash combination is invalid
API_ID_PUBLISHED_FLOOD	400	This API id was published somewhere, you can't use it now
ARTICLE_TITLE_EMPTY	400	The title of the article is empty
AUTH_BYTES_INVALID	400	The provided authorization is invalid
AUTH_KEY_DUPLICATED	406	The authorization key (session file) was used under two different IP addresses simultaneously, and can no longer be used. Use the same session exclusively, or use different sessions
AUTH_KEY_INVALID	401	The key is invalid
AUTH_KEY_PERM_EMPTY	401	The method is unavailable for temporary authorization key, not bound to permanent
AUTH_KEY_UNREGISTERED	401	The key is not registered in the system
AUTH_RESTART	500	Restart the authorization process
BANNED_RIGHTS_INVALID	400	You cannot use that set of permissions in this request, i.e. restricting view_messages as a default
BOTS_TOO_MUCH	400	There are too many bots in this chat/channel
BOT_CHANNELS_NA	400	Bots can't edit admin privileges
BOT_GROUPS_BLOCKED	400	This bot can't be added to groups
BOT_INLINE_DISABLED	400	This bot can't be used in inline mode
BOT_INVALID	400	BOT_PAYMENTS_DISABLED
BOT_METHOD_INVALID	400	The API access for bot users is restricted. The method you tried to invoke cannot be executed as a bot
BOT_MISSING	400	This method can only be run by a bot
BOT_PAYMENTS_DISABLED	400	This method can only be run by a bot
BOT_POLLS_DISABLED	400	You cannot create polls under a bot account
BROADCAST_ID_INVALID	400	The channel is invalid
BUTTON_DATA_INVALID	400	The provided button data is invalid
BUTTON_TYPE_INVALID	400	The type of one of the buttons you provided is invalid
BUTTON_URL_INVALID	400	Button URL invalid
CALL_ALREADY_ACCEPTED	400	The call was already accepted
CALL_ALREADY_DECLINED	400	The call was already declined
CALL_OCCUPY_FAILED	500	The call failed because the user is already making another call
CALL_PEER_INVALID	400	The provided call peer object is invalid
CALL_PROTOCOL_FLAGS_INVALID	400	Call protocol flags invalid
CDN_METHOD_INVALID	400	This method cannot be invoked on a CDN server. Refer to https://core.telegram.org/cdn#schema for available methods
CHANNELS_ADMIN_PUBLIC_TOO_MUCH	400	You're admin of too many public channels, make some channels private to change the username of this channel
CHANNELS_TOO_MUCH	400	You have joined too many channels/supergroups
CHANNEL_INVALID	400	Invalid channel object. Make sure to pass the right types, for instance making sure that the request is designed for channels or otherwise look for a different one more suited
CHANNEL_PRIVATE	400	The channel specified is private and you lack permission to access it. Another reason may be that you were banned from it
CHANNEL_PUBLIC_GROUP_NA	403	Channel/supergroup not available
CHAT_ABOUT_NOT_MODIFIED	400	About text has not changed
CHAT_ABOUT_TOO_LONG	400	Chat about too long
CHAT_ADMIN_INVITE_REQUIRED	403	You do not have the rights to do this
CHAT_ADMIN_REQUIRED	400 403	Chat admin privileges are required to do that in the specified chat (for example, to send a message in a channel which is not yours), or invalid permissions used for the channel or group
CHAT_FORBIDDEN	N/A	You cannot write in this chat
CHAT_ID_EMPTY	400	The provided chat ID is empty
CHAT_ID_INVALID	400	Invalid object ID for a chat. Make sure to pass the right types, for instance making sure that the request is designed for chats (not channels/megagroups) or otherwise look for a different one more suited\nAn example working with a megagroup and AddChatUserRequest, it will fail because megagroups are channels. Use InviteToChannelRequest instead
CHAT_INVALID	400	The chat is invalid for this request
CHAT_LINK_EXISTS	400	The chat is linked to a channel and cannot be used in that request
CHAT_NOT_MODIFIED	400	The chat or channel wasn't modified (title, invites, username, admins, etc. are the same)
CHAT_RESTRICTED	400	The chat is restricted and cannot be used in that request
CHAT_SEND_GIFS_FORBIDDEN	403	You can't send gifs in this chat
CHAT_SEND_INLINE_FORBIDDEN	400	You cannot send inline results in this chat
CHAT_SEND_MEDIA_FORBIDDEN	403	You can't send media in this chat
CHAT_SEND_STICKERS_FORBIDDEN	403	You can't send stickers in this chat
CHAT_TITLE_EMPTY	400	No chat title provided
CHAT_WRITE_FORBIDDEN	403	You can't write in this chat
CODE_EMPTY	400	The provided code is empty
CODE_HASH_INVALID	400	Code hash invalid
CODE_INVALID	400	Code invalid (i.e. from email)
CONNECTION_API_ID_INVALID	400	The provided API id is invalid
CONNECTION_DEVICE_MODEL_EMPTY	400	Device model empty
CONNECTION_LANG_PACK_INVALID	400	The specified language pack is not valid. This is meant to be used by official applications only so far, leave it empty
CONNECTION_LAYER_INVALID	400	The very first request must always be InvokeWithLayerRequest
CONNECTION_NOT_INITED	400	Connection not initialized
CONNECTION_SYSTEM_EMPTY	400	Connection system empty
CONTACT_ID_INVALID	400	The provided contact ID is invalid
DATA_INVALID	400	Encrypted data invalid
DATA_JSON_INVALID	400	The provided JSON data is invalid
DATE_EMPTY	400	Date empty
DC_ID_INVALID	400	This occurs when an authorization is tried to be exported for the same data center one is currently connected to
DH_G_A_INVALID	400	g_a invalid
EMAIL_HASH_EXPIRED	400	The email hash expired and cannot be used to verify it
EMAIL_INVALID	400	The given email is invalid
EMAIL_UNCONFIRMED_X	400	Email unconfirmed, the length of the code must be {code_length}
EMOTICON_EMPTY	400	The emoticon field cannot be empty
ENCRYPTED_MESSAGE_INVALID	400	Encrypted message invalid
ENCRYPTION_ALREADY_ACCEPTED	400	Secret chat already accepted
ENCRYPTION_ALREADY_DECLINED	400	The secret chat was already declined
ENCRYPTION_DECLINED	400	The secret chat was declined
ENCRYPTION_ID_INVALID	400	The provided secret chat ID is invalid
ENCRYPTION_OCCUPY_FAILED	500	TDLib developer claimed it is not an error while accepting secret chats and 500 is used instead of 420
ENTITIES_TOO_LONG	400	It is no longer possible to send such long data inside entity tags (for example inline text URLs)
ENTITY_MENTION_USER_INVALID	400	You can't use this entity
ERROR_TEXT_EMPTY	400	The provided error message is empty
EXPORT_CARD_INVALID	400	Provided card is invalid
EXTERNAL_URL_INVALID	400	External URL invalid
FIELD_NAME_EMPTY	N/A	The field with the name FIELD_NAME is missing
FIELD_NAME_INVALID	N/A	The field with the name FIELD_NAME is invalid
FILE_ID_INVALID	400	The provided file id is invalid
FILE_MIGRATE_X	303	The file to be accessed is currently stored in DC {new_dc}
FILE_PARTS_INVALID	400	The number of file parts is invalid
FILE_PART_0_MISSING	N/A	File part 0 missing
FILE_PART_EMPTY	400	The provided file part is empty
FILE_PART_LENGTH_INVALID	400	The length of a file part is invalid
FILE_PART_SIZE_INVALID	400	The provided file part size is invalid
FILE_PART_X_MISSING	400	Part {which} of the file is missing from storage
FILEREF_UPGRADE_NEEDED	406	The file reference needs to be refreshed before being used again
FIRSTNAME_INVALID	400	The first name is invalid
FLOOD_TEST_PHONE_WAIT_X	420	A wait of {seconds} seconds is required in the test servers
FLOOD_WAIT_X	420	A wait of {seconds} seconds is required
FOLDER_ID_EMPTY	400	The folder you tried to use was not valid
FRESH_RESET_AUTHORISATION_FORBIDDEN	406	The current session is too new and cannot be used to reset other authorisations yet
GIF_ID_INVALID	400	The provided GIF ID is invalid
GROUPED_MEDIA_INVALID	400	Invalid grouped media
HASH_INVALID	400	The provided hash is invalid
HISTORY_GET_FAILED	500	Fetching of history failed
IMAGE_PROCESS_FAILED	400	Failure while processing image
INLINE_RESULT_EXPIRED	400	The inline query expired
INPUT_CONSTRUCTOR_INVALID	400	The provided constructor is invalid
INPUT_FETCH_ERROR	N/A	An error occurred while deserializing TL parameters
INPUT_FETCH_FAIL	400	Failed deserializing TL payload
INPUT_LAYER_INVALID	400	The provided layer is invalid
INPUT_METHOD_INVALID	N/A	The invoked method does not exist anymore or has never existed
INPUT_REQUEST_TOO_LONG	400	The input request was too long. This may be a bug in the library as it can occur when serializing more bytes than it should (like appending the vector constructor code at the end of a message)
INPUT_USER_DEACTIVATED	400	The specified user was deleted
INTERDC_X_CALL_ERROR	N/A	An error occurred while communicating with DC {dc}
INTERDC_X_CALL_RICH_ERROR	N/A	A rich error occurred while communicating with DC {dc}
INVITE_HASH_EMPTY	400	The invite hash is empty
INVITE_HASH_EXPIRED	400	The chat the user tried to join has expired and is not valid anymore
INVITE_HASH_INVALID	400	The invite hash is invalid
LANG_PACK_INVALID	400	The provided language pack is invalid
LASTNAME_INVALID	N/A	The last name is invalid
LIMIT_INVALID	400	An invalid limit was provided. See https://core.telegram.org/api/files#downloading-files
LINK_NOT_MODIFIED	400	The channel is already linked to this group
LOCATION_INVALID	400	The location given for a file was invalid. See https://core.telegram.org/api/files#downloading-files
MAX_ID_INVALID	400	The provided max ID is invalid
MAX_QTS_INVALID	400	The provided QTS were invalid
MD5_CHECKSUM_INVALID	N/A	The MD5 check-sums do not match
MEDIA_CAPTION_TOO_LONG	400	The caption is too long
MEDIA_EMPTY	400	The provided media object is invalid
MEDIA_INVALID	400	Media invalid
MEDIA_NEW_INVALID	400	The new media to edit the message with is invalid (such as stickers or voice notes)
MEDIA_PREV_INVALID	400	The old media cannot be edited with anything else (such as stickers or voice notes)
MEGAGROUP_ID_INVALID	400	The group is invalid
MEGAGROUP_PREHISTORY_HIDDEN	400	You can't set this discussion group because it's history is hidden
MEMBER_NO_LOCATION	500	An internal failure occurred while fetching user info (couldn't find location)
MEMBER_OCCUPY_PRIMARY_LOC_FAILED	500	Occupation of primary member location failed
MESSAGE_AUTHOR_REQUIRED	403	Message author required
MESSAGE_DELETE_FORBIDDEN	403	You can't delete one of the messages you tried to delete, most likely because it is a service message.
MESSAGE_EDIT_TIME_EXPIRED	400	You can't edit this message anymore, too much time has passed since its creation.
MESSAGE_EMPTY	400	Empty or invalid UTF-8 message was sent
MESSAGE_IDS_EMPTY	400	No message ids were provided
MESSAGE_ID_INVALID	400	The specified message ID is invalid or you can't do that operation on such message
MESSAGE_NOT_MODIFIED	400	Content of the message was not modified
MESSAGE_TOO_LONG	400	Message was too long. Current maximum length is 4096 UTF-8 characters
MSG_WAIT_FAILED	400	A waiting call returned an error
MT_SEND_QUEUE_TOO_LONG	500	(no description given)
NEED_CHAT_INVALID	500	The provided chat is invalid
NEED_MEMBER_INVALID	500	The provided member is invalid or does not exist (for example a thumb size)
NETWORK_MIGRATE_X	303	The source IP address is associated with DC {new_dc}
NEW_SALT_INVALID	400	The new salt is invalid
NEW_SETTINGS_INVALID	400	The new settings are invalid
CONTACT_ID_INVALID	400	The provided contact ID is invalid
OFFSET_INVALID	400	The given offset was invalid, it must be divisible by 1KB. See https://core.telegram.org/api/files#downloading-files
OFFSET_PEER_ID_INVALID	400	The provided offset peer is invalid
OPTIONS_TOO_MUCH	400	You defined too many options for the poll
PACK_SHORT_NAME_INVALID	400	Invalid sticker pack name. It must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot username>".
PACK_SHORT_NAME_OCCUPIED	400	A stickerpack with this name already exists
PARTICIPANTS_TOO_FEW	400	Not enough participants
PARTICIPANT_CALL_FAILED	500	Failure while making call
PARTICIPANT_VERSION_OUTDATED	400	The other participant does not use an up to date telegram client with support for calls
PASSWORD_EMPTY	400	The provided password is empty
PASSWORD_HASH_INVALID	400	The password (and thus its hash value) you entered is invalid
PASSWORD_REQUIRED	400	The account must have 2-factor authentication enabled (a password) before this method can be used
PAYMENT_PROVIDER_INVALID	400	The payment provider was not recognised or its token was invalid
PEER_FLOOD	N/A	Too many requests
PEER_ID_INVALID	400	An invalid Peer was used. Make sure to pass the right peer type
PEER_ID_NOT_SUPPORTED	400	The provided peer ID is not supported
PERSISTENT_TIMESTAMP_EMPTY	400	Persistent timestamp empty
PERSISTENT_TIMESTAMP_INVALID	400	Persistent timestamp invalid
PERSISTENT_TIMESTAMP_OUTDATED	500	Persistent timestamp outdated
PHONE_CODE_EMPTY	400	The phone code is missing
PHONE_CODE_EXPIRED	400	The confirmation code has expired
PHONE_CODE_HASH_EMPTY	N/A	The phone code hash is missing
PHONE_CODE_INVALID	400	The phone code entered was invalid
PHONE_MIGRATE_X	303	The phone number a user is trying to use for authorization is associated with DC {new_dc}
PHONE_NUMBER_APP_SIGNUP_FORBIDDEN	400	(no description is given)
PHONE_NUMBER_BANNED	400	The used phone number has been banned from Telegram and cannot be used anymore. Maybe check https://www.telegram.org/faq_spam
PHONE_NUMBER_FLOOD	400	You asked for the code too many times.
PHONE_NUMBER_INVALID	400 406	The phone number is invalid
PHONE_NUMBER_OCCUPIED	400	The phone number is already in use
PHONE_NUMBER_UNOCCUPIED	400	The phone number is not yet being used
PHONE_PASSWORD_FLOOD	406	You have tried logging in too many times
PHONE_PASSWORD_PROTECTED	400	This phone is password protected
PHOTO_CONTENT_URL_EMPTY	400	The content from the URL used as a photo appears to be empty or has caused another HTTP error
PHOTO_CROP_SIZE_SMALL	400	Photo is too small
PHOTO_EXT_INVALID	400	The extension of the photo is invalid
PHOTO_INVALID	400	Photo invalid
PHOTO_INVALID_DIMENSIONS	400	The photo dimensions are invalid (hint: `pip install pillow` for `send_file` to resize images)
PHOTO_SAVE_FILE_INVALID	400	The photo you tried to send cannot be saved by Telegram. A reason may be that it exceeds 10MB. Try resizing it locally
PHOTO_THUMB_URL_EMPTY	400	The URL used as a thumbnail appears to be empty or has caused another HTTP error
PIN_RESTRICTED	400	You can't pin messages in private chats with other people
POLL_OPTION_DUPLICATE	400	A duplicate option was sent in the same poll
POLL_UNSUPPORTED	400	This layer does not support polls in the issued method
PRIVACY_KEY_INVALID	400	The privacy key is invalid
PTS_CHANGE_EMPTY	500	No PTS change
QUERY_ID_EMPTY	400	The query ID is empty
QUERY_ID_INVALID	400	The query ID is invalid
QUERY_TOO_SHORT	400	The query string is too short
RANDOM_ID_DUPLICATE	500	You provided a random ID that was already used
RANDOM_ID_INVALID	400	A provided random ID is invalid
RANDOM_LENGTH_INVALID	400	Random length invalid
RANGES_INVALID	400	Invalid range provided
REG_ID_GENERATE_FAILED	500	Failure while generating registration ID
REPLY_MARKUP_INVALID	400	The provided reply markup is invalid
REPLY_MARKUP_TOO_LONG	400	The data embedded in the reply markup buttons was too much
RESULT_ID_DUPLICATE	400	Duplicated IDs on the sent results. Make sure to use unique IDs.
RESULT_TYPE_INVALID	400	Result type invalid
RESULTS_TOO_MUCH	400	You sent too many results. See https://core.telegram.org/bots/api#answerinlinequery for the current limit.
RIGHT_FORBIDDEN	403	Either your admin rights do not allow you to do this or you passed the wrong rights combination (some rights only apply to channels and vice versa)
RPC_CALL_FAIL	N/A	Telegram is having internal issues, please try again later.
RPC_MCGET_FAIL	N/A	Telegram is having internal issues, please try again later.
RSA_DECRYPT_FAILED	400	Internal RSA decryption failed
SEARCH_QUERY_EMPTY	400	The search query is empty
SEND_MESSAGE_MEDIA_INVALID	400	The message media was invalid or not specified
SEND_MESSAGE_TYPE_INVALID	400	The message type is invalid
SESSION_EXPIRED	401	The authorization has expired
SESSION_PASSWORD_NEEDED	401	Two-steps verification is enabled and a password is required
SESSION_REVOKED	401	The authorization has been invalidated, because of the user terminating all sessions
SHA256_HASH_INVALID	400	The provided SHA256 hash is invalid
SHORTNAME_OCCUPY_FAILED	400	An error occurred when trying to register the short-name used for the sticker pack. Try a different name
START_PARAM_EMPTY	400	The start parameter is empty
START_PARAM_INVALID	400	Start parameter invalid
STICKERSET_INVALID	400	The provided sticker set is invalid
STICKERS_EMPTY	400	No sticker provided
STICKER_EMOJI_INVALID	400	Sticker emoji invalid
STICKER_FILE_INVALID	400	Sticker file invalid
STICKER_ID_INVALID	400	The provided sticker ID is invalid
STICKER_INVALID	400	The provided sticker is invalid
STICKER_PNG_DIMENSIONS	400	Sticker png dimensions invalid
STORAGE_CHECK_FAILED	500	Server storage check failed
STORE_INVALID_SCALAR_TYPE	500	(no description is provided)
TAKEOUT_INIT_DELAY_X	420	A wait of {seconds} seconds is required before being able to initiate the takeout
TAKEOUT_INVALID	400	The takeout session has been invalidated by another data export session
TAKEOUT_REQUIRED	400	You must initialize a takeout request first
TIMEOUT	503	A timeout occurred while fetching data from the bot
TMP_PASSWORD_DISABLED	400	The temporary password is disabled
TOKEN_INVALID	400	The provided token is invalid
TTL_DAYS_INVALID	400	The provided TTL is invalid
TYPES_EMPTY	400	The types field is empty
TYPE_CONSTRUCTOR_INVALID	N/A	The type constructor is invalid
UNKNOWN_METHOD	500	The method you tried to call cannot be called on non-CDN DCs
UNTIL_DATE_INVALID	400	That date cannot be specified in this request (try using None)
URL_INVALID	400	The URL used was invalid (e.g. when answering a callback with an URL that's not t.me/yourbot or your game's URL)
USERNAME_INVALID	400	Nobody is using this username, or the username is unacceptable. If the latter, it must match r"[a-zA-Z][\w\d]{3,30}[a-zA-Z\d]"
USERNAME_NOT_MODIFIED	400	The username is not different from the current username
USERNAME_NOT_OCCUPIED	400	The username is not in use by anyone else yet
USERNAME_OCCUPIED	400	The username is already taken
USERS_TOO_FEW	400	Not enough users (to create a chat, for example)
USERS_TOO_MUCH	400	The maximum number of users has been exceeded (to create a chat, for example)
USER_ADMIN_INVALID	400	Either you're not an admin or you tried to ban an admin that you didn't promote
USER_ALREADY_PARTICIPANT	400	The authenticated user is already a participant of the chat
USER_BANNED_IN_CHANNEL	400	You're banned from sending messages in supergroups/channels
USER_BLOCKED	400	User blocked
USER_BOT	400	Bots can only be admins in channels.
USER_BOT_INVALID	400 403	This method can only be called by a bot
USER_BOT_REQUIRED	400	This method can only be called by a bot
USER_CHANNELS_TOO_MUCH	403	One of the users you tried to add is already in too many channels/supergroups
USER_CREATOR	400	You can't leave this channel, because you're its creator
USER_DEACTIVATED	401	The user has been deleted/deactivated
USER_DEACTIVATED_BAN	401	The user has been deleted/deactivated
USER_ID_INVALID	400	Invalid object ID for a user. Make sure to pass the right types, for instance making sure that the request is designed for users or otherwise look for a different one more suited
USER_INVALID	400	The given user was invalid
USER_IS_BLOCKED	400 403	User is blocked
USER_IS_BOT	400	Bots can't send messages to other bots
USER_KICKED	400	This user was kicked from this supergroup/channel
USER_MIGRATE_X	303	The user whose identity is being used to execute queries is associated with DC {new_dc}
USER_NOT_MUTUAL_CONTACT	400 403	The provided user is not a mutual contact
USER_NOT_PARTICIPANT	400	The target user is not a member of the specified megagroup or channel
USER_PRIVACY_RESTRICTED	403	The user's privacy settings do not allow you to do this
USER_RESTRICTED	403	You're spamreported, you can't create channels or chats
VIDEO_CONTENT_TYPE_INVALID	400	The video content type is not supported with the given parameters (i.e. supports_streaming)
WALLPAPER_FILE_INVALID	400	The given file cannot be used as a wallpaper
WALLPAPER_INVALID	400	The input wallpaper was not valid
WC_CONVERT_URL_INVALID	400	WC convert URL invalid
WEBPAGE_CURL_FAILED	400	Failure while fetching the webpage with cURL
WEBPAGE_MEDIA_EMPTY	400	Webpage media empty
WORKER_BUSY_TOO_LONG_RETRY	500	Telegram workers are too busy to respond immediately
YOU_BLOCKED_USER	400	You blocked this user
*/
