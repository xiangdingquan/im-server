package poll_facade

import (
	"context"
	"fmt"

	"open.chat/model"
	"open.chat/mtproto"
)

type PollFacade interface {
	CreateMediaPoll(ctx context.Context, userId int32, correctAnswers []int, poll *mtproto.Poll) (*model.MediaPoll, error)
	CloseMediaPoll(ctx context.Context, userId int32, pollId int64) (*model.MediaPoll, error)
	GetMediaPoll(ctx context.Context, userId int32, pollId int64) (*model.MediaPoll, error)
	SendVote(ctx context.Context, userId int32, pollId int64, options []string) (*model.MediaPoll, error)
	GetPollVoters(ctx context.Context, userId int32, pollId int64, option string, offset string, limit int32) (*mtproto.Messages_VotesList, error)
}

type Instance func() PollFacade

var instances = make(map[string]Instance)

func Register(name string, inst Instance) {
	if inst == nil {
		panic("register instance is nil")
	}
	if _, ok := instances[name]; ok {
		panic("register called twice for instance " + name)
	}
	instances[name] = inst
}

func NewPollFacade(name string) (inst PollFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
