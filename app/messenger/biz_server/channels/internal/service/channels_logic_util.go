package service

import (
	"open.chat/mtproto"
)

func IsInputChannel(channel *mtproto.InputChannel) bool {
	return channel != nil && channel.PredicateName == mtproto.Predicate_inputChannel
}
