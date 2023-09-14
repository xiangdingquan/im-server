package consts

type OnAction string

// call
const (
	// ActionCallOnInvite 邀请用户通话
	ActionCallOnInvite OnAction = "call.onInvite"
	// ActionCallOnCancel 当前通话已取消
	ActionCallOnCancel OnAction = "call.onCancel"
	// ActionCallOnLeave 当前通话已取消
	ActionCallOnLeave OnAction = "call.onLeave"
	// ActionChatRight
	ActionChatsRightOnUpdate OnAction = "chats.rights.onUpdate"
	// ActionChatKeywordsOnUpdate
	ActionChatsKeywordsOnUpdate OnAction = "chats.keywords.onUpdate"
	ActionChatsNicknameOnUpdate OnAction = "chats.nickname,onUpdate"

	ActionMessageBatchSendOnSend OnAction = "messages.batchSend.onSend"
	ActionMessageBatchSendDelete OnAction = "messages.batchSend.OnDelete"

	ActionMessageReactionOnUpdate OnAction = "messages.reaction.onUpdate"
)
