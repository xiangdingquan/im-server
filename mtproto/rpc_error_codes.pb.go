package mtproto

import (
	"fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	"math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TLRpcErrorCodes int32

const (
	// 	Error handling
	//
	// 	There will be errors when working with the API, and they must be correctly handled on the client.
	//
	// 	An error is characterized by several parameters:
	//
	// 	Error Code
	// 	Similar to HTTP status. Contains information on the type of error that occurred: for example,
	// 	a data input error, privacy error, or server error. This is a required parameter.
	//
	// 	Error Type
	// 	A string literal in the form of /[A-Z_0-9]+/, which summarizes the problem. For example, AUTH_KEY_UNREGISTERED.
	// 	This is an optional parameter.
	//
	// 	Error Description
	// 	May contain more detailed information on the error and how to resolve it,
	// 	for example: authorization required, use auth.* methods. Please note that the description text is subject to change,
	// 	one should avoid tying application logic to these messages. This is an optional parameter.
	//
	// 	Error Constructors
	// 	There should be a way to handle errors that are returned in rpc_error constructors.
	//
	// 	If an error constructor does not differentiate between type and description
	// 	but instead contains a single field called error_message (as in the example above),
	// 	it must be split into 2 components, for example, using the following regular expression: /^([A-Z_0-9]+)(: (.+))?/.
	//
	// Protobuf3 enum first is 0
	TLRpcErrorCodes_ERROR_CODE_OK TLRpcErrorCodes = 0
	// 303 ERROR_SEE_OTHER
	//
	// The request must be repeated, but directed to a different data center.
	//
	// FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
	// PHONE_MIGRATE_X: the phone number a user is trying to use for authorization is associated with a different data center.
	// NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
	// USER_MIGRATE_X: the user whose identity is being used to execute
	// 				   queries is associated with a different data center (for registration)
	//
	// In all these cases, the error description’s string literal
	// 		contains the number of the data center (instead of the X) to which the repeated query must be sent.
	// More information about redirects between data centers »
	//
	TLRpcErrorCodes_FILE_MIGRATE_X    TLRpcErrorCodes = 303000
	TLRpcErrorCodes_PHONE_MIGRATE_X   TLRpcErrorCodes = 303001
	TLRpcErrorCodes_NETWORK_MIGRATE_X TLRpcErrorCodes = 303002
	TLRpcErrorCodes_USER_MIGRATE_X    TLRpcErrorCodes = 303003
	TLRpcErrorCodes_ERROR_SEE_OTHER   TLRpcErrorCodes = 303
	// 400 BAD_REQUEST
	//
	// The query contains errors. In the event that a request was created using a form
	// and contains user generated data,
	// the user should be notified that the data must be corrected before the query is repeated.
	//
	//
	// Examples of Errors:
	// 	FIRSTNAME_INVALID: The first name is invalid
	// 	LASTNAME_INVALID: The last name is invalid
	// 	PHONE_NUMBER_INVALID: The phone number is invalid
	// 	PHONE_CODE_HASH_EMPTY: phone_code_hash is missing
	// 	PHONE_CODE_EMPTY: phone_code is missing
	// 	PHONE_CODE_EXPIRED: The confirmation code has expired
	// 	API_ID_INVALID: The api_id/api_hash combination is invalid
	// 	PHONE_NUMBER_OCCUPIED: The phone number is already in use
	// 	PHONE_NUMBER_UNOCCUPIED: The phone number is not yet being used
	// 	USERS_TOO_FEW: Not enough users (to create a chat, for example)
	// 	USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
	// 	TYPE_CONSTRUCTOR_INVALID: The type constructor is invalid
	// 	FILE_PART_INVALID: The file part number is invalid
	// 	FILE_PARTS_INVALID: The number of file parts is invalid
	// 	FILE_PART_Х_MISSING: Part X (where X is a number) of the file is missing from storage
	// 	MD5_CHECKSUM_INVALID: The MD5 checksums do not match
	// 	PHOTO_INVALID_DIMENSIONS: The photo dimensions are invalid
	// 	FIELD_NAME_INVALID: The field with the name FIELD_NAME is invalid
	// 	FIELD_NAME_EMPTY: The field with the name FIELD_NAME is missing
	// 	MSG_WAIT_FAILED: A waiting call returned an error
	//
	TLRpcErrorCodes_FIRSTNAME_INVALID        TLRpcErrorCodes = 400000
	TLRpcErrorCodes_LASTNAME_INVALID         TLRpcErrorCodes = 400001
	TLRpcErrorCodes_PHONE_NUMBER_INVALID     TLRpcErrorCodes = 400002
	TLRpcErrorCodes_PHONE_CODE_HASH_EMPTY    TLRpcErrorCodes = 400003
	TLRpcErrorCodes_PHONE_CODE_EMPTY         TLRpcErrorCodes = 400004
	TLRpcErrorCodes_PHONE_CODE_EXPIRED       TLRpcErrorCodes = 400005
	TLRpcErrorCodes_API_ID_INVALID           TLRpcErrorCodes = 400006
	TLRpcErrorCodes_PHONE_NUMBER_OCCUPIED    TLRpcErrorCodes = 400007
	TLRpcErrorCodes_PHONE_NUMBER_UNOCCUPIED  TLRpcErrorCodes = 400008
	TLRpcErrorCodes_USERS_TOO_FEW            TLRpcErrorCodes = 400009
	TLRpcErrorCodes_USERS_TOO_MUCH           TLRpcErrorCodes = 400010
	TLRpcErrorCodes_TYPE_CONSTRUCTOR_INVALID TLRpcErrorCodes = 400011
	TLRpcErrorCodes_FILE_PART_INVALID        TLRpcErrorCodes = 400012
	TLRpcErrorCodes_FILE_PART_X_MISSING      TLRpcErrorCodes = 400013
	TLRpcErrorCodes_MD5_CHECKSUM_INVALID     TLRpcErrorCodes = 400014
	TLRpcErrorCodes_PHOTO_INVALID_DIMENSIONS TLRpcErrorCodes = 400015
	TLRpcErrorCodes_FIELD_NAME_INVALID       TLRpcErrorCodes = 400016
	TLRpcErrorCodes_FIELD_NAME_EMPTY         TLRpcErrorCodes = 400017
	TLRpcErrorCodes_MSG_WAIT_FAILED          TLRpcErrorCodes = 400018
	// PHONE_NUMBER_BANNED = 4000;
	// android client code:
	//    if (error.code == 400 && "PARTICIPANT_VERSION_OUTDATED".equals(error.text)) {
	//        callFailed(VoIPController.ERROR_PEER_OUTDATED);
	TLRpcErrorCodes_PARTICIPANT_VERSION_OUTDATED TLRpcErrorCodes = 400019
	TLRpcErrorCodes_USER_RESTRICTED              TLRpcErrorCodes = 400020
	TLRpcErrorCodes_NAME_NOT_MODIFIED            TLRpcErrorCodes = 400021
	TLRpcErrorCodes_USER_NOT_MUTUAL_CONTACT      TLRpcErrorCodes = 400022
	TLRpcErrorCodes_BOT_GROUPS_BLOCKED           TLRpcErrorCodes = 400023
	TLRpcErrorCodes_FILE_REFERENCE_X             TLRpcErrorCodes = 400500
	TLRpcErrorCodes_FILE_TOKEN_INVALID           TLRpcErrorCodes = 400501
	TLRpcErrorCodes_REQUEST_TOKEN_INVALID        TLRpcErrorCodes = 400502
	//
	TLRpcErrorCodes_PHONE_CODE_INVALID      TLRpcErrorCodes = 400025
	TLRpcErrorCodes_PHONE_NUMBER_BANNED     TLRpcErrorCodes = 400030
	TLRpcErrorCodes_SESSION_PASSWORD_NEEDED TLRpcErrorCodes = 400040
	// password
	TLRpcErrorCodes_CODE_INVALID          TLRpcErrorCodes = 400050
	TLRpcErrorCodes_PASSWORD_HASH_INVALID TLRpcErrorCodes = 400051
	TLRpcErrorCodes_NEW_PASSWORD_BAD      TLRpcErrorCodes = 400052
	TLRpcErrorCodes_NEW_SALT_INVALID      TLRpcErrorCodes = 400053
	TLRpcErrorCodes_EMAIL_INVALID         TLRpcErrorCodes = 400054
	TLRpcErrorCodes_EMAIL_UNCONFIRMED     TLRpcErrorCodes = 400055
	TLRpcErrorCodes_SRP_PASSWORD_CHANGED  TLRpcErrorCodes = 400056
	TLRpcErrorCodes_SRP_ID_INVALID        TLRpcErrorCodes = 400057
	// username
	TLRpcErrorCodes_USERNAME_INVALID      TLRpcErrorCodes = 400060
	TLRpcErrorCodes_USERNAME_OCCUPIED     TLRpcErrorCodes = 400061
	TLRpcErrorCodes_USERNAMES_UNAVAILABLE TLRpcErrorCodes = 400062
	TLRpcErrorCodes_USERNAME_NOT_MODIFIED TLRpcErrorCodes = 400063
	TLRpcErrorCodes_USERNAME_NOT_OCCUPIED TLRpcErrorCodes = 400064
	// chat
	TLRpcErrorCodes_CHAT_ID_INVALID         TLRpcErrorCodes = 400070
	TLRpcErrorCodes_CHAT_NOT_MODIFIED       TLRpcErrorCodes = 400071
	TLRpcErrorCodes_PARTICIPANT_NOT_EXISTS  TLRpcErrorCodes = 400072
	TLRpcErrorCodes_NO_EDIT_CHAT_PERMISSION TLRpcErrorCodes = 400073
	TLRpcErrorCodes_CHAT_TITLE_NOT_MODIFIED TLRpcErrorCodes = 400074
	TLRpcErrorCodes_NO_CHAT_TITLE           TLRpcErrorCodes = 400075
	TLRpcErrorCodes_CHAT_ABOUT_NOT_MODIFIED TLRpcErrorCodes = 400076
	TLRpcErrorCodes_CHAT_ADMIN_REQUIRED     TLRpcErrorCodes = 400077
	TLRpcErrorCodes_PARTICIPANT_EXISTED     TLRpcErrorCodes = 400078
	// channel
	TLRpcErrorCodes_CHANNEL_PRIVATE                TLRpcErrorCodes = 400080
	TLRpcErrorCodes_CHANNEL_PUBLIC_GROUP_NA        TLRpcErrorCodes = 400081
	TLRpcErrorCodes_USER_BANNED_IN_CHANNEL         TLRpcErrorCodes = 400082
	TLRpcErrorCodes_CHANNELS_ADMIN_PUBLIC_TOO_MUCH TLRpcErrorCodes = 40083
	TLRpcErrorCodes_CHANNELS_TOO_MUCH              TLRpcErrorCodes = 400084
	TLRpcErrorCodes_NO_INVITE_CHANNEL_PERMISSION   TLRpcErrorCodes = 400085
	// invite, user banned.
	TLRpcErrorCodes_INVITE_HASH_EXPIRED TLRpcErrorCodes = 400090
	TLRpcErrorCodes_INVITE_HASH_INVALID TLRpcErrorCodes = 400091
	// access
	TLRpcErrorCodes_ACCESS_HASH_INVALID  TLRpcErrorCodes = 400200
	TLRpcErrorCodes_INPUT_CHANNEL_EMPTY  TLRpcErrorCodes = 400201
	TLRpcErrorCodes_USER_NOT_PARTICIPANT TLRpcErrorCodes = 400202
	TLRpcErrorCodes_PEER_ID_INVALID      TLRpcErrorCodes = 400203
	TLRpcErrorCodes_CHANNEL_ID_INVALID   TLRpcErrorCodes = 400204
	// message
	TLRpcErrorCodes_MESSAGE_ID_INVALID        TLRpcErrorCodes = 400210
	TLRpcErrorCodes_MESSAGE_EDIT_TIME_EXPIRED TLRpcErrorCodes = 400211
	TLRpcErrorCodes_MESSAGE_NOT_MODIFIED      TLRpcErrorCodes = 400212
	TLRpcErrorCodes_MESSAGE_EMPTY             TLRpcErrorCodes = 400213
	TLRpcErrorCodes_USER_LEFT_CHAT            TLRpcErrorCodes = 400300
	TLRpcErrorCodes_USER_KICKED               TLRpcErrorCodes = 400301
	TLRpcErrorCodes_USER_ALREADY_PARTICIPANT  TLRpcErrorCodes = 400302
	// LANG_CODE_NOT_SUPPORTED
	TLRpcErrorCodes_LANG_CODE_NOT_SUPPORTED TLRpcErrorCodes = 400400
	// STICKERSET_INVALID: The provided sticker set is invalid
	TLRpcErrorCodes_STICKERSET_INVALID TLRpcErrorCodes = 400401
	TLRpcErrorCodes_BAD_REQUEST        TLRpcErrorCodes = 400
	// There was an unauthorized attempt to use functionality available only to authorized users.
	//
	// Examples of Errors:
	// 	AUTH_KEY_UNREGISTERED: The key is not registered in the system
	// 	AUTH_KEY_INVALID: The key is invalid
	// 	USER_DEACTIVATED: The user has been deleted/deactivated
	// 	SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
	// 	SESSION_EXPIRED: The authorization has expired
	// 	ACTIVE_USER_REQUIRED: The method is only available to already activated users
	// 	AUTH_KEY_PERM_EMPTY: The method is unavailble for temporary authorization key, not bound to permanent
	//
	TLRpcErrorCodes_AUTH_KEY_UNREGISTERED TLRpcErrorCodes = 401000
	TLRpcErrorCodes_AUTH_KEY_INVALID      TLRpcErrorCodes = 401001
	TLRpcErrorCodes_USER_DEACTIVATED      TLRpcErrorCodes = 401002
	TLRpcErrorCodes_SESSION_REVOKED       TLRpcErrorCodes = 401003
	TLRpcErrorCodes_SESSION_EXPIRED       TLRpcErrorCodes = 401004
	TLRpcErrorCodes_ACTIVE_USER_REQUIRED  TLRpcErrorCodes = 401005
	TLRpcErrorCodes_AUTH_KEY_PERM_EMPTY   TLRpcErrorCodes = 401006
	// Only a small portion of the API methods are available to unauthorized users:
	//
	// auth.sendCode
	// auth.sendCall
	// auth.checkPhone
	// auth.signUp
	// auth.signIn
	// auth.importAuthorization
	// help.getConfig
	// help.getNearestDc
	//
	// Other methods will result in an error: 401 UNAUTHORIZED.
	TLRpcErrorCodes_UNAUTHORIZED TLRpcErrorCodes = 401
	// Privacy violation. For example, an attempt to write a message to someone who has blacklisted the current user.
	//
	//
	// android client code:
	//    } else if(error.code==403 && "USER_PRIVACY_RESTRICTED".equals(error.text)){
	//        callFailed(VoIPController.ERROR_PRIVACY);
	TLRpcErrorCodes_USER_PRIVACY_RESTRICTED      TLRpcErrorCodes = 403001
	TLRpcErrorCodes_CALL_PROTOCOL_LAYER_INVALID  TLRpcErrorCodes = 403002
	TLRpcErrorCodes_CHAT_SEND_STICKERS_FORBIDDEN TLRpcErrorCodes = 403003
	TLRpcErrorCodes_FORBIDDEN                    TLRpcErrorCodes = 403
	// 406
	// android client code:
	// }else if(error.code==406){
	//     callFailed(VoIPController.ERROR_LOCALIZED);
	TLRpcErrorCodes_ERROR_LOCALIZED TLRpcErrorCodes = 406000
	TLRpcErrorCodes_LOCALIZED       TLRpcErrorCodes = 406
	// The maximum allowed number of attempts to invoke the given method with the given input parameters has been exceeded.
	// For example, in an attempt to request a large number of text messages (SMS) for the same phone number.
	//
	// Error Example:
	// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
	//
	TLRpcErrorCodes_FLOOD_WAIT_X TLRpcErrorCodes = 420000
	// PEER_FLOOD
	// FLOOD_WAIT
	TLRpcErrorCodes_FLOOD TLRpcErrorCodes = 420
	// An internal server error occurred while a request was being processed;
	// for example, there was a disruption while accessing a database or file storage.
	//
	// If a client receives a 500 error, or you believe this error should not have occurred,
	// please collect as much information as possible about the query and error and send it to the developers.
	TLRpcErrorCodes_INTERNAL              TLRpcErrorCodes = 500
	TLRpcErrorCodes_INTERNAL_SERVER_ERROR TLRpcErrorCodes = 500000
	// If a server returns an error with a code other than the ones listed above,
	// it may be considered the same as a 500 error and treated as an internal server error.
	//
	TLRpcErrorCodes_OTHER TLRpcErrorCodes = 501
	//    // OFFSET_INVALID
	//    // RETRY_LIMIT
	//    // FILE_TOKEN_INVALID
	//    // REQUEST_TOKEN_INVALID
	//
	//    // CHANNEL_PRIVATE
	//    // CHANNEL_PUBLIC_GROUP_NA
	//    // USER_BANNED_IN_CHANNEL
	//
	//
	//    // MESSAGE_NOT_MODIFIED
	//
	//    // USERS_TOO_MUCH
	//
	//    // -1000
	//
	//    /////////////////////////////////////////////////////////////
	//     // android client code:
	//       } else if (request instanceof TLRPC.TL_auth_resendCode) {
	//        if (error.text.contains("PHONE_NUMBER_INVALID")) {
	//            showSimpleAlert(fragment, LocaleController.getString("InvalidPhoneNumber", R.string.InvalidPhoneNumber));
	//        } else if (error.text.contains("PHONE_CODE_EMPTY") || error.text.contains("PHONE_CODE_INVALID")) {
	//            showSimpleAlert(fragment, LocaleController.getString("InvalidCode", R.string.InvalidCode));
	//        } else if (error.text.contains("PHONE_CODE_EXPIRED")) {
	//            showSimpleAlert(fragment, LocaleController.getString("CodeExpired", R.string.CodeExpired));
	//        } else if (error.text.startsWith("FLOOD_WAIT")) {
	//            showSimpleAlert(fragment, LocaleController.getString("FloodWait", R.string.FloodWait));
	//        } else if (error.code != -1000) {
	//            showSimpleAlert(fragment, LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred) + "\n" + error.text);
	//        }
	//
	//     /////////////////////////////////////////////////////////////
	//        } else if (request instanceof TLRPC.TL_updateUserName) {
	//            switch (error.text) {
	//                case "USERNAME_INVALID":
	//                    showSimpleAlert(fragment, LocaleController.getString("UsernameInvalid", R.string.UsernameInvalid));
	//                    break;
	//                case "USERNAME_OCCUPIED":
	//                    showSimpleAlert(fragment, LocaleController.getString("UsernameInUse", R.string.UsernameInUse));
	//                    break;
	//                case "USERNAMES_UNAVAILABLE":
	//                    showSimpleAlert(fragment, LocaleController.getString("FeatureUnavailable", R.string.FeatureUnavailable));
	//                    break;
	//                default:
	//                    showSimpleAlert(fragment, LocaleController.getString("ErrorOccurred", R.string.ErrorOccurred));
	//                    break;
	//            }
	//
	//     /////////////////////////////////////////////////////////////
	//            } else if (request instanceof TLRPC.TL_payments_sendPaymentForm) {
	//            switch (error.text) {
	//                case "BOT_PRECHECKOUT_FAILED":
	//                    showSimpleToast(fragment, LocaleController.getString("PaymentPrecheckoutFailed", R.string.PaymentPrecheckoutFailed));
	//                    break;
	//                case "PAYMENT_FAILED":
	//                    showSimpleToast(fragment, LocaleController.getString("PaymentFailed", R.string.PaymentFailed));
	//                    break;
	//                default:
	//                    showSimpleToast(fragment, error.text);
	//                    break;
	//            }
	//        } else if (request instanceof TLRPC.TL_payments_validateRequestedInfo) {
	//            switch (error.text) {
	//                case "SHIPPING_NOT_AVAILABLE":
	//                    showSimpleToast(fragment, LocaleController.getString("PaymentNoShippingMethod", R.string.PaymentNoShippingMethod));
	//                    break;
	//                default:
	//                    showSimpleToast(fragment, error.text);
	//                    break;
	//            }
	//        }
	//
	//     /////////////////////////////////////////////////////////////
	//
	//        } else {
	//            if (error.text.equals("2FA_RECENT_CONFIRM")) {
	//                needShowAlert(LocaleController.getString("AppName", R.string.AppName), LocaleController.getString("ResetAccountCancelledAlert", R.string.ResetAccountCancelledAlert));
	//            } else if (error.text.startsWith("2FA_CONFIRM_WAIT_")) {
	//                Bundle params = new Bundle();
	//                params.putString("phoneFormated", requestPhone);
	//                params.putString("phoneHash", phoneHash);
	//                params.putString("code", phoneCode);
	//                params.putInt("startTime", ConnectionsManager.getInstance().getCurrentTime());
	//                params.putInt("waitTime", Utilities.parseInt(error.text.replace("2FA_CONFIRM_WAIT_", "")));
	//                setPage(8, true, params, false);
	//            } else {
	TLRpcErrorCodes_OTHER2 TLRpcErrorCodes = 502
	// db error
	TLRpcErrorCodes_DBERR      TLRpcErrorCodes = 600
	TLRpcErrorCodes_DBERR_SQL  TLRpcErrorCodes = 600000
	TLRpcErrorCodes_DBERR_CONN TLRpcErrorCodes = 600001
	// db error
	TLRpcErrorCodes_NOTRETURN_CLIENT TLRpcErrorCodes = 700
)

var TLRpcErrorCodes_name = map[int32]string{
	0:      "ERROR_CODE_OK",
	303000: "FILE_MIGRATE_X",
	303001: "PHONE_MIGRATE_X",
	303002: "NETWORK_MIGRATE_X",
	303003: "USER_MIGRATE_X",
	303:    "ERROR_SEE_OTHER",
	400000: "FIRSTNAME_INVALID",
	400001: "LASTNAME_INVALID",
	400002: "PHONE_NUMBER_INVALID",
	400003: "PHONE_CODE_HASH_EMPTY",
	400004: "PHONE_CODE_EMPTY",
	400005: "PHONE_CODE_EXPIRED",
	400006: "API_ID_INVALID",
	400007: "PHONE_NUMBER_OCCUPIED",
	400008: "PHONE_NUMBER_UNOCCUPIED",
	400009: "USERS_TOO_FEW",
	400010: "USERS_TOO_MUCH",
	400011: "TYPE_CONSTRUCTOR_INVALID",
	400012: "FILE_PART_INVALID",
	400013: "FILE_PART_X_MISSING",
	400014: "MD5_CHECKSUM_INVALID",
	400015: "PHOTO_INVALID_DIMENSIONS",
	400016: "FIELD_NAME_INVALID",
	400017: "FIELD_NAME_EMPTY",
	400018: "MSG_WAIT_FAILED",
	400019: "PARTICIPANT_VERSION_OUTDATED",
	400020: "USER_RESTRICTED",
	400021: "NAME_NOT_MODIFIED",
	400022: "USER_NOT_MUTUAL_CONTACT",
	400023: "BOT_GROUPS_BLOCKED",
	400500: "FILE_REFERENCE_X",
	400501: "FILE_TOKEN_INVALID",
	400502: "REQUEST_TOKEN_INVALID",
	400025: "PHONE_CODE_INVALID",
	400030: "PHONE_NUMBER_BANNED",
	400040: "SESSION_PASSWORD_NEEDED",
	400050: "CODE_INVALID",
	400051: "PASSWORD_HASH_INVALID",
	400052: "NEW_PASSWORD_BAD",
	400053: "NEW_SALT_INVALID",
	400054: "EMAIL_INVALID",
	400055: "EMAIL_UNCONFIRMED",
	400056: "SRP_PASSWORD_CHANGED",
	400057: "SRP_ID_INVALID",
	400060: "USERNAME_INVALID",
	400061: "USERNAME_OCCUPIED",
	400062: "USERNAMES_UNAVAILABLE",
	400063: "USERNAME_NOT_MODIFIED",
	400064: "USERNAME_NOT_OCCUPIED",
	400070: "CHAT_ID_INVALID",
	400071: "CHAT_NOT_MODIFIED",
	400072: "PARTICIPANT_NOT_EXISTS",
	400073: "NO_EDIT_CHAT_PERMISSION",
	400074: "CHAT_TITLE_NOT_MODIFIED",
	400075: "NO_CHAT_TITLE",
	400076: "CHAT_ABOUT_NOT_MODIFIED",
	400077: "CHAT_ADMIN_REQUIRED",
	400078: "PARTICIPANT_EXISTED",
	400080: "CHANNEL_PRIVATE",
	400081: "CHANNEL_PUBLIC_GROUP_NA",
	400082: "USER_BANNED_IN_CHANNEL",
	40083:  "CHANNELS_ADMIN_PUBLIC_TOO_MUCH",
	400084: "CHANNELS_TOO_MUCH",
	400085: "NO_INVITE_CHANNEL_PERMISSION",
	400090: "INVITE_HASH_EXPIRED",
	400091: "INVITE_HASH_INVALID",
	400200: "ACCESS_HASH_INVALID",
	400201: "INPUT_CHANNEL_EMPTY",
	400202: "USER_NOT_PARTICIPANT",
	400203: "PEER_ID_INVALID",
	400204: "CHANNEL_ID_INVALID",
	400210: "MESSAGE_ID_INVALID",
	400211: "MESSAGE_EDIT_TIME_EXPIRED",
	400212: "MESSAGE_NOT_MODIFIED",
	400213: "MESSAGE_EMPTY",
	400300: "USER_LEFT_CHAT",
	400301: "USER_KICKED",
	400302: "USER_ALREADY_PARTICIPANT",
	400400: "LANG_CODE_NOT_SUPPORTED",
	400401: "STICKERSET_INVALID",
	400:    "BAD_REQUEST",
	401000: "AUTH_KEY_UNREGISTERED",
	401001: "AUTH_KEY_INVALID",
	401002: "USER_DEACTIVATED",
	401003: "SESSION_REVOKED",
	401004: "SESSION_EXPIRED",
	401005: "ACTIVE_USER_REQUIRED",
	401006: "AUTH_KEY_PERM_EMPTY",
	401:    "UNAUTHORIZED",
	403001: "USER_PRIVACY_RESTRICTED",
	403002: "CALL_PROTOCOL_LAYER_INVALID",
	403003: "CHAT_SEND_STICKERS_FORBIDDEN",
	403:    "FORBIDDEN",
	406000: "ERROR_LOCALIZED",
	406:    "LOCALIZED",
	420000: "FLOOD_WAIT_X",
	420:    "FLOOD",
	500:    "INTERNAL",
	500000: "INTERNAL_SERVER_ERROR",
	501:    "OTHER",
	502:    "OTHER2",
	600:    "DBERR",
	600000: "DBERR_SQL",
	600001: "DBERR_CONN",
	700:    "NOTRETURN_CLIENT",
}
var TLRpcErrorCodes_value = map[string]int32{
	"ERROR_CODE_OK":                  0,
	"FILE_MIGRATE_X":                 303000,
	"PHONE_MIGRATE_X":                303001,
	"NETWORK_MIGRATE_X":              303002,
	"USER_MIGRATE_X":                 303003,
	"ERROR_SEE_OTHER":                303,
	"FIRSTNAME_INVALID":              400000,
	"LASTNAME_INVALID":               400001,
	"PHONE_NUMBER_INVALID":           400002,
	"PHONE_CODE_HASH_EMPTY":          400003,
	"PHONE_CODE_EMPTY":               400004,
	"PHONE_CODE_EXPIRED":             400005,
	"API_ID_INVALID":                 400006,
	"PHONE_NUMBER_OCCUPIED":          400007,
	"PHONE_NUMBER_UNOCCUPIED":        400008,
	"USERS_TOO_FEW":                  400009,
	"USERS_TOO_MUCH":                 400010,
	"TYPE_CONSTRUCTOR_INVALID":       400011,
	"FILE_PART_INVALID":              400012,
	"FILE_PART_X_MISSING":            400013,
	"MD5_CHECKSUM_INVALID":           400014,
	"PHOTO_INVALID_DIMENSIONS":       400015,
	"FIELD_NAME_INVALID":             400016,
	"FIELD_NAME_EMPTY":               400017,
	"MSG_WAIT_FAILED":                400018,
	"PARTICIPANT_VERSION_OUTDATED":   400019,
	"USER_RESTRICTED":                400020,
	"NAME_NOT_MODIFIED":              400021,
	"USER_NOT_MUTUAL_CONTACT":        400022,
	"BOT_GROUPS_BLOCKED":             400023,
	"FILE_REFERENCE_X":               400500,
	"FILE_TOKEN_INVALID":             400501,
	"REQUEST_TOKEN_INVALID":          400502,
	"PHONE_CODE_INVALID":             400025,
	"PHONE_NUMBER_BANNED":            400030,
	"SESSION_PASSWORD_NEEDED":        400040,
	"CODE_INVALID":                   400050,
	"PASSWORD_HASH_INVALID":          400051,
	"NEW_PASSWORD_BAD":               400052,
	"NEW_SALT_INVALID":               400053,
	"EMAIL_INVALID":                  400054,
	"EMAIL_UNCONFIRMED":              400055,
	"SRP_PASSWORD_CHANGED":           400056,
	"SRP_ID_INVALID":                 400057,
	"USERNAME_INVALID":               400060,
	"USERNAME_OCCUPIED":              400061,
	"USERNAMES_UNAVAILABLE":          400062,
	"USERNAME_NOT_MODIFIED":          400063,
	"USERNAME_NOT_OCCUPIED":          400064,
	"CHAT_ID_INVALID":                400070,
	"CHAT_NOT_MODIFIED":              400071,
	"PARTICIPANT_NOT_EXISTS":         400072,
	"NO_EDIT_CHAT_PERMISSION":        400073,
	"CHAT_TITLE_NOT_MODIFIED":        400074,
	"NO_CHAT_TITLE":                  400075,
	"CHAT_ABOUT_NOT_MODIFIED":        400076,
	"CHAT_ADMIN_REQUIRED":            400077,
	"PARTICIPANT_EXISTED":            400078,
	"CHANNEL_PRIVATE":                400080,
	"CHANNEL_PUBLIC_GROUP_NA":        400081,
	"USER_BANNED_IN_CHANNEL":         400082,
	"CHANNELS_ADMIN_PUBLIC_TOO_MUCH": 40083,
	"CHANNELS_TOO_MUCH":              400084,
	"NO_INVITE_CHANNEL_PERMISSION":   400085,
	"INVITE_HASH_EXPIRED":            400090,
	"INVITE_HASH_INVALID":            400091,
	"ACCESS_HASH_INVALID":            400200,
	"INPUT_CHANNEL_EMPTY":            400201,
	"USER_NOT_PARTICIPANT":           400202,
	"PEER_ID_INVALID":                400203,
	"CHANNEL_ID_INVALID":             400204,
	"MESSAGE_ID_INVALID":             400210,
	"MESSAGE_EDIT_TIME_EXPIRED":      400211,
	"MESSAGE_NOT_MODIFIED":           400212,
	"MESSAGE_EMPTY":                  400213,
	"USER_LEFT_CHAT":                 400300,
	"USER_KICKED":                    400301,
	"USER_ALREADY_PARTICIPANT":       400302,
	"LANG_CODE_NOT_SUPPORTED":        400400,
	"STICKERSET_INVALID":             400401,
	"BAD_REQUEST":                    400,
	"AUTH_KEY_UNREGISTERED":          401000,
	"AUTH_KEY_INVALID":               401001,
	"USER_DEACTIVATED":               401002,
	"SESSION_REVOKED":                401003,
	"SESSION_EXPIRED":                401004,
	"ACTIVE_USER_REQUIRED":           401005,
	"AUTH_KEY_PERM_EMPTY":            401006,
	"UNAUTHORIZED":                   401,
	"USER_PRIVACY_RESTRICTED":        403001,
	"CALL_PROTOCOL_LAYER_INVALID":    403002,
	"CHAT_SEND_STICKERS_FORBIDDEN":   403003,
	"FORBIDDEN":                      403,
	"ERROR_LOCALIZED":                406000,
	"LOCALIZED":                      406,
	"FLOOD_WAIT_X":                   420000,
	"FLOOD":                          420,
	"INTERNAL":                       500,
	"INTERNAL_SERVER_ERROR":          500000,
	"OTHER":                          501,
	"OTHER2":                         502,
	"DBERR":                          600,
	"DBERR_SQL":                      600000,
	"DBERR_CONN":                     600001,
	"NOTRETURN_CLIENT":               700,
}

func (x TLRpcErrorCodes) String() string {
	return proto.EnumName(TLRpcErrorCodes_name, int32(x))
}
func (TLRpcErrorCodes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rpc_error_codes_825e2504028f1a87, []int{0}
}

func init() {
	proto.RegisterEnum("mtproto.TLRpcErrorCodes", TLRpcErrorCodes_name, TLRpcErrorCodes_value)
}

func init() {
	proto.RegisterFile("rpc_error_codes.proto", fileDescriptor_rpc_error_codes_825e2504028f1a87)
}

var fileDescriptor_rpc_error_codes_825e2504028f1a87 = []byte{
	// 1437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x56, 0xc9, 0x8f, 0x15, 0xc7,
	0x19, 0x67, 0x98, 0x40, 0x42, 0xb3, 0x4c, 0x51, 0x30, 0xd0, 0x2c, 0x79, 0x51, 0x22, 0x4e, 0x91,
	0xc2, 0x28, 0x89, 0xf2, 0x07, 0xd4, 0xeb, 0xfa, 0xde, 0x7b, 0xa5, 0xe9, 0xae, 0x6a, 0xaa, 0xaa,
	0x67, 0xe1, 0x52, 0x62, 0x26, 0x13, 0x88, 0x04, 0x19, 0x34, 0xcc, 0xdc, 0xb3, 0xef, 0x0b, 0x84,
	0x84, 0x84, 0x44, 0x8a, 0x38, 0xe4, 0x90, 0x43, 0x16, 0x29, 0x76, 0xbf, 0x19, 0xec, 0x67, 0x0f,
	0xf6, 0xc1, 0x1e, 0x60, 0x6c, 0x16, 0x83, 0x64, 0xd9, 0x17, 0x83, 0x2f, 0xde, 0xc5, 0x01, 0x38,
	0x5b, 0x55, 0xd5, 0xdd, 0xaf, 0xdf, 0xd8, 0xa7, 0xee, 0xfe, 0xfd, 0xaa, 0xbe, 0xad, 0x7e, 0xdf,
	0x57, 0x1d, 0x8c, 0x2e, 0x9c, 0x9d, 0x35, 0x73, 0x0b, 0x0b, 0xf3, 0x0b, 0x66, 0x76, 0xfe, 0xbb,
	0x73, 0xe7, 0x8e, 0x9e, 0x5d, 0x98, 0x5f, 0x9c, 0xc7, 0x5f, 0x3c, 0xb3, 0xe8, 0x5e, 0x0e, 0x7e,
	0xe3, 0xe4, 0xf7, 0x17, 0x4f, 0x2d, 0xcd, 0x1c, 0x9d, 0x9d, 0x3f, 0x33, 0x76, 0x72, 0xfe, 0xe4,
	0xfc, 0x98, 0x83, 0x67, 0x96, 0xbe, 0xe7, 0xbe, 0xdc, 0x87, 0x7b, 0xf3, 0xfb, 0xbe, 0xfe, 0x78,
	0x6f, 0x30, 0xa2, 0x63, 0x79, 0x76, 0x16, 0xac, 0xc9, 0xc8, 0x5a, 0xc4, 0xbb, 0x83, 0x9d, 0x20,
	0xa5, 0x90, 0x26, 0x12, 0x14, 0x8c, 0x18, 0x47, 0x9b, 0xf0, 0xde, 0x60, 0x57, 0x8b, 0xc5, 0x60,
	0x12, 0xd6, 0x96, 0x44, 0x83, 0x99, 0x42, 0x7f, 0x5d, 0xc5, 0x78, 0x34, 0x18, 0x49, 0x3b, 0x82,
	0xd7, 0xe1, 0xcb, 0xab, 0x18, 0xef, 0x0f, 0x76, 0x73, 0xd0, 0x93, 0x42, 0x8e, 0xd7, 0x88, 0xbf,
	0xad, 0x62, 0x6b, 0x25, 0x53, 0x20, 0x6b, 0xe8, 0xdf, 0x1d, 0x3a, 0xe2, 0xdd, 0x29, 0x00, 0x23,
	0x74, 0x07, 0x24, 0xfa, 0xdf, 0x66, 0x6b, 0xa4, 0xc5, 0xa4, 0xd2, 0x9c, 0x24, 0x60, 0x18, 0x9f,
	0x20, 0x31, 0xa3, 0xe8, 0x87, 0x79, 0x88, 0xf7, 0x05, 0x28, 0x26, 0x1b, 0xf0, 0x1f, 0xe5, 0x21,
	0x3e, 0x18, 0xec, 0xf5, 0xc1, 0xf0, 0x2c, 0x69, 0x82, 0xac, 0xb8, 0x1f, 0xe7, 0x21, 0x3e, 0x14,
	0x8c, 0x7a, 0xce, 0x65, 0xd4, 0x21, 0xaa, 0x63, 0x20, 0x49, 0xf5, 0x34, 0xfa, 0x89, 0x37, 0x58,
	0x23, 0x3d, 0xfe, 0xd3, 0x3c, 0xc4, 0x61, 0x80, 0xeb, 0xf8, 0x54, 0xca, 0x24, 0x50, 0xf4, 0xb3,
	0x3c, 0xb4, 0x79, 0x90, 0x94, 0x19, 0x46, 0x2b, 0x27, 0x3f, 0xaf, 0x3b, 0x29, 0x02, 0x10, 0x51,
	0x94, 0xa5, 0x0c, 0x28, 0xfa, 0x45, 0x1e, 0xe2, 0x2f, 0x07, 0xfb, 0x07, 0xc8, 0x8c, 0x57, 0xf4,
	0x2f, 0xf3, 0x10, 0xef, 0x09, 0x76, 0xda, 0xca, 0x28, 0xa3, 0x85, 0x30, 0x2d, 0x98, 0x44, 0xbf,
	0xf2, 0x6e, 0xfa, 0x60, 0x92, 0x45, 0x1d, 0xf4, 0xeb, 0x3c, 0xc4, 0x8d, 0x20, 0xd4, 0xd3, 0xa9,
	0x8d, 0x8a, 0x2b, 0x2d, 0xb3, 0x48, 0x8b, 0x7e, 0xae, 0xbf, 0xc9, 0x43, 0x5f, 0xb8, 0x18, 0x4c,
	0x4a, 0xa4, 0xae, 0x88, 0xdf, 0xe6, 0x21, 0x3e, 0x10, 0xec, 0xe9, 0x13, 0x53, 0x26, 0x61, 0x4a,
	0x31, 0xde, 0x46, 0xbf, 0xf3, 0xb5, 0x4b, 0xe8, 0x77, 0x4c, 0xd4, 0x81, 0x68, 0x5c, 0x65, 0x49,
	0xb5, 0xed, 0xf7, 0xde, 0x5f, 0xda, 0x11, 0x5a, 0x94, 0xa0, 0xa1, 0x2c, 0x01, 0xae, 0x98, 0xe0,
	0x0a, 0xfd, 0xc1, 0x97, 0xa9, 0xc5, 0x20, 0xa6, 0x66, 0xe0, 0x44, 0xce, 0xfb, 0xc2, 0xd6, 0x18,
	0x5f, 0xd8, 0x0b, 0x79, 0x68, 0x65, 0x93, 0xa8, 0xb6, 0x99, 0x24, 0x4c, 0x9b, 0x16, 0x61, 0x31,
	0x50, 0xf4, 0xc7, 0x3c, 0xc4, 0x5f, 0x0b, 0x0e, 0xdb, 0xd0, 0x58, 0xc4, 0x52, 0xc2, 0xb5, 0x99,
	0x00, 0x69, 0x9d, 0x18, 0x91, 0x69, 0x4a, 0x34, 0x50, 0x74, 0xd1, 0x6f, 0x75, 0x0a, 0x92, 0xa0,
	0xb4, 0x64, 0x91, 0x85, 0xff, 0xe4, 0x73, 0x76, 0x3e, 0xb8, 0xd0, 0x26, 0x11, 0x94, 0xb5, 0x6c,
	0x5d, 0xff, 0xec, 0xcb, 0xee, 0xd6, 0x3b, 0x22, 0xd3, 0x19, 0x89, 0x6d, 0xdd, 0x34, 0x89, 0x34,
	0xba, 0xe4, 0x63, 0x6f, 0x0a, 0x6d, 0xda, 0x52, 0x64, 0xa9, 0x32, 0xcd, 0x58, 0x44, 0xe3, 0x40,
	0xd1, 0x5f, 0xca, 0xd8, 0x63, 0x30, 0x12, 0x5a, 0x20, 0x81, 0x47, 0x56, 0xac, 0x8f, 0x57, 0x8a,
	0x6c, 0x63, 0x30, 0x5a, 0x8c, 0x03, 0xaf, 0xb2, 0x7d, 0xb2, 0xe2, 0x8e, 0x5f, 0xc2, 0xb1, 0x0c,
	0x94, 0xde, 0x40, 0x3e, 0x5d, 0xd9, 0xa8, 0xa5, 0x92, 0xb9, 0xec, 0x4f, 0x65, 0x40, 0x18, 0x4d,
	0xc2, 0x39, 0x50, 0xf4, 0x0f, 0x1f, 0xbc, 0x02, 0xe5, 0x8a, 0x90, 0x12, 0xa5, 0x26, 0x85, 0xa4,
	0x86, 0x03, 0x50, 0xa0, 0xe8, 0x5f, 0x79, 0x88, 0x71, 0xb0, 0x63, 0xc0, 0xda, 0xff, 0x0b, 0x0d,
	0x96, 0x4b, 0x9d, 0xcc, 0x4b, 0xf2, 0x19, 0x9f, 0x13, 0x87, 0xc9, 0xbe, 0xad, 0x26, 0xa1, 0xe8,
	0xd9, 0x3e, 0xae, 0x48, 0xdc, 0x17, 0x4c, 0xee, 0x45, 0x09, 0x09, 0x61, 0x71, 0x05, 0x76, 0x7d,
	0xa9, 0x3d, 0x98, 0xf1, 0x48, 0xf0, 0x16, 0x93, 0x09, 0x50, 0xb4, 0xec, 0x35, 0xa4, 0x64, 0xda,
	0xb7, 0x1e, 0x75, 0x08, 0x6f, 0x03, 0x45, 0x2b, 0x5e, 0xc9, 0x96, 0xab, 0x35, 0xcc, 0x55, 0xef,
	0xd7, 0x1e, 0xce, 0x80, 0x6e, 0x7a, 0xde, 0x45, 0x85, 0x57, 0x5d, 0xf2, 0x82, 0xcf, 0xae, 0x24,
	0x94, 0xc9, 0x38, 0x99, 0x20, 0x2c, 0x26, 0xcd, 0x18, 0xd0, 0x8b, 0x83, 0xe4, 0xa0, 0x0e, 0x56,
	0x3f, 0x87, 0xac, 0xcc, 0x5e, 0xf3, 0xa2, 0x8a, 0x3a, 0x44, 0xd7, 0xc3, 0x7b, 0xc5, 0x87, 0xe1,
	0xe0, 0x01, 0x63, 0xaf, 0xe6, 0x21, 0x3e, 0x1c, 0xec, 0xab, 0x0b, 0xd5, 0xf2, 0x30, 0xc5, 0x94,
	0x56, 0x68, 0xcd, 0x9f, 0x1a, 0x17, 0x06, 0x28, 0xd3, 0xc6, 0x6d, 0x4f, 0x41, 0xba, 0x56, 0x13,
	0x1c, 0x5d, 0xf7, 0xb4, 0x83, 0x35, 0xd3, 0xf1, 0x86, 0x40, 0x6f, 0xf8, 0x9a, 0x73, 0x61, 0xfa,
	0x2b, 0xd0, 0xcd, 0xda, 0x1e, 0xd2, 0x14, 0xd9, 0x86, 0x78, 0xd6, 0xbd, 0x84, 0x3c, 0x4d, 0x13,
	0xc6, 0x8d, 0x15, 0xa1, 0x9b, 0x54, 0xaf, 0x15, 0xea, 0xaa, 0x85, 0xea, 0xc2, 0x04, 0x8a, 0x5e,
	0xaf, 0xb2, 0xe6, 0x1c, 0x62, 0x93, 0x4a, 0x36, 0x41, 0x34, 0xa0, 0xdb, 0x95, 0x2f, 0x0f, 0x67,
	0xcd, 0x98, 0x45, 0xbe, 0x3b, 0x0c, 0x27, 0xe8, 0x8e, 0xcf, 0xdd, 0x35, 0x94, 0x97, 0xa9, 0x61,
	0xdc, 0x14, 0xab, 0xd1, 0xdd, 0x3c, 0xc4, 0x47, 0x82, 0x46, 0xf1, 0xa9, 0x8a, 0x68, 0x0a, 0x1b,
	0xd5, 0x04, 0xbb, 0x78, 0x75, 0x73, 0x51, 0x58, 0xbf, 0xaa, 0x22, 0xee, 0xf9, 0x09, 0xc0, 0xdd,
	0x9c, 0x61, 0x1a, 0x4c, 0x15, 0x45, 0xbf, 0x7e, 0xf7, 0x7d, 0x46, 0xc5, 0x02, 0x3f, 0xc6, 0x8b,
	0xb1, 0xfc, 0xd6, 0x67, 0xa9, 0xf2, 0x2c, 0xdf, 0xf6, 0x14, 0x89, 0x22, 0x50, 0x6a, 0x90, 0x5a,
	0xeb, 0x16, 0xbb, 0xd2, 0x4c, 0x57, 0x0e, 0xfd, 0xa0, 0xba, 0xde, 0x75, 0x92, 0xae, 0xa6, 0x47,
	0xad, 0x8c, 0xe8, 0x46, 0xd7, 0x95, 0x2f, 0x05, 0x7b, 0xcd, 0xf4, 0x45, 0x73, 0xb3, 0xeb, 0x1a,
	0xbd, 0xb4, 0x53, 0x63, 0xd6, 0x3d, 0x93, 0x80, 0x52, 0xa4, 0x0d, 0x75, 0xe6, 0x6e, 0x37, 0xc4,
	0x5f, 0x09, 0x0e, 0x94, 0x8c, 0x93, 0x8d, 0x66, 0x49, 0xff, 0xbe, 0x79, 0xc3, 0xc7, 0x51, 0x2e,
	0x18, 0x38, 0xfc, 0x7b, 0x5d, 0x27, 0x98, 0x6a, 0xb3, 0x0b, 0xfc, 0x7e, 0xb7, 0xba, 0x39, 0x4c,
	0x0c, 0x2d, 0xaf, 0x42, 0xf4, 0xef, 0xe5, 0x10, 0xef, 0x0e, 0xb6, 0x3b, 0x74, 0x9c, 0xb9, 0x31,
	0xf7, 0x9f, 0x65, 0x37, 0xdc, 0x1d, 0x44, 0x62, 0x09, 0x84, 0x4e, 0x0f, 0x64, 0xf9, 0xdf, 0x65,
	0xa7, 0x86, 0x98, 0xf0, 0xb6, 0x1f, 0x5b, 0xd6, 0xb7, 0xca, 0xd2, 0x54, 0x48, 0xab, 0xa1, 0xf3,
	0x7e, 0xac, 0x29, 0x6d, 0xad, 0x49, 0x05, 0xfd, 0xd9, 0x71, 0x61, 0x25, 0xc4, 0x28, 0xd8, 0xde,
	0x24, 0xd4, 0x14, 0x13, 0x11, 0x9d, 0x1f, 0xb6, 0x2d, 0x48, 0x32, 0xdd, 0x31, 0xe3, 0x30, 0x6d,
	0x32, 0x2e, 0xa1, 0x6d, 0xa5, 0x68, 0x33, 0x7c, 0xaf, 0xe7, 0x46, 0x41, 0x45, 0x96, 0x66, 0xde,
	0xef, 0x55, 0x23, 0xc2, 0x50, 0x20, 0x91, 0x76, 0x22, 0xa5, 0xe8, 0x83, 0x9e, 0xab, 0x7e, 0x39,
	0x1a, 0x25, 0x4c, 0x08, 0x9b, 0xce, 0x87, 0x83, 0x70, 0x59, 0xbf, 0x8f, 0x7a, 0xae, 0x7e, 0x6e,
	0x3b, 0x98, 0xe2, 0xf2, 0x28, 0x3a, 0xe4, 0xe3, 0x9e, 0x57, 0x46, 0xe9, 0xd9, 0x4a, 0xad, 0xa8,
	0xe2, 0x27, 0x3d, 0x5b, 0xaf, 0x1d, 0x19, 0xb7, 0xa4, 0x90, 0xec, 0x38, 0x50, 0x74, 0x61, 0xb8,
	0xba, 0x4f, 0x5c, 0xc7, 0x44, 0xd3, 0xf5, 0x7b, 0xe8, 0xea, 0x7a, 0x88, 0xbf, 0x1a, 0x1c, 0x8a,
	0x48, 0x6c, 0x1b, 0x4a, 0x68, 0x11, 0x89, 0xd8, 0xc4, 0x64, 0xba, 0xf6, 0x2b, 0xf2, 0xdc, 0xba,
	0xd3, 0xb8, 0x6b, 0x56, 0x05, 0x9c, 0x9a, 0xb2, 0x78, 0xa6, 0x25, 0x64, 0x93, 0x51, 0x0a, 0x1c,
	0x3d, 0xbf, 0x1e, 0xe2, 0x5d, 0xc1, 0xb6, 0x3e, 0x70, 0x71, 0xd8, 0xa6, 0xe5, 0xff, 0x90, 0x62,
	0x11, 0x91, 0xd8, 0xc5, 0xf2, 0xe8, 0x5d, 0xb7, 0xac, 0x0f, 0x5c, 0x1a, 0xb6, 0x17, 0x42, 0x2b,
	0x16, 0x82, 0xfa, 0x9b, 0x75, 0x0a, 0x5d, 0xb9, 0x73, 0x00, 0x07, 0xc1, 0x16, 0x87, 0xa1, 0x7f,
	0x0e, 0xe3, 0x9d, 0xc1, 0x97, 0x18, 0xd7, 0x76, 0x0a, 0xc6, 0xe8, 0xb1, 0x3b, 0x90, 0xf2, 0xd3,
	0x28, 0x90, 0x13, 0x20, 0x8d, 0xf3, 0x82, 0xae, 0xbc, 0xdc, 0xb0, 0xfb, 0xfc, 0xaf, 0xd8, 0x93,
	0x61, 0xbc, 0x3d, 0xd8, 0xea, 0xde, 0xbf, 0x85, 0x9e, 0x0e, 0x5b, 0x82, 0x36, 0x41, 0x4a, 0xf4,
	0xe6, 0x17, 0xf0, 0x48, 0xb0, 0xcd, 0xbd, 0x1b, 0x75, 0x2c, 0x46, 0xd7, 0x6e, 0x1d, 0xc1, 0x28,
	0x08, 0x3c, 0x10, 0x09, 0xce, 0xd1, 0x4b, 0xb7, 0x8e, 0xe0, 0xd1, 0x00, 0x71, 0xa1, 0x25, 0xe8,
	0x4c, 0x72, 0x13, 0xc5, 0x0c, 0xb8, 0x46, 0xbd, 0x2d, 0xcd, 0x6f, 0xae, 0x3d, 0x68, 0x6c, 0x7a,
	0xf4, 0xa0, 0x31, 0xb4, 0xf6, 0xb0, 0x31, 0x74, 0xfb, 0x61, 0x63, 0xe8, 0x9d, 0x87, 0x8d, 0xa1,
	0xe3, 0x87, 0x7e, 0x30, 0x37, 0xb3, 0x74, 0xfa, 0xc4, 0xd1, 0xd9, 0x53, 0x27, 0x16, 0xc7, 0x66,
	0x4f, 0x2f, 0x9d, 0x5b, 0x9c, 0x5b, 0x18, 0x2b, 0x7e, 0x6c, 0x67, 0xb6, 0xba, 0xc7, 0xb7, 0x3f,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x42, 0x2b, 0x91, 0x84, 0x01, 0x0b, 0x00, 0x00,
}
