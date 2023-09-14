package model

import (
	"context"
)

type ChannelHelper interface {
	GetMutableChannel(ctx context.Context, channelId int32, id ...int32) (chat *MutableChannel, err error)
}
