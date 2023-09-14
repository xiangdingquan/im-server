package dataobject

type SecretChatCloseRequestsDo struct {
	SecretChatId int32 `db:"secret_chat_id"`
	FromUID      int32 `db:"from_uid"`
	ToUID        int32 `db:"to_uid"`
}
