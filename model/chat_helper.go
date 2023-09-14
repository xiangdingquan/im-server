package model

import (
	"context"
)

type ChatHelper interface {
	GetMutableChat(ctx context.Context, chatId int32, id ...int32) (chat *MutableChat, err error)
}
