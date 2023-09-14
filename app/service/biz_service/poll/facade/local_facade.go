package poll_facade

import (
	"context"

	"open.chat/app/service/biz_service/poll/internal/core"
	"open.chat/app/service/biz_service/poll/internal/dao"
	"open.chat/model"
	"open.chat/mtproto"
)

func init() {
	Register("local", localPollFacadeInstance)
}

type localPollFacade struct {
	*core.PollCore
}

func localPollFacadeInstance() PollFacade {
	return &localPollFacade{
		PollCore: core.New(dao.New()),
	}
}

func (c *localPollFacade) CreateMediaPoll(ctx context.Context, userId int32, correctAnswers []int, poll *mtproto.Poll) (*model.MediaPoll, error) {
	return c.PollCore.CreateMediaPoll(ctx, userId, correctAnswers, poll)
}

func (c *localPollFacade) CloseMediaPoll(ctx context.Context, userId int32, pollId int64) (*model.MediaPoll, error) {
	return c.PollCore.CloseMediaPoll(ctx, userId, pollId)
}

func (c *localPollFacade) GetMediaPoll(ctx context.Context, userId int32, pollId int64) (*model.MediaPoll, error) {
	return c.PollCore.GetMediaPoll(ctx, userId, pollId)
}

func (c *localPollFacade) SendVote(ctx context.Context, userId int32, pollId int64, options []string) (*model.MediaPoll, error) {
	return c.PollCore.SendVote(ctx, userId, pollId, options)
}

func (c *localPollFacade) GetPollVoters(ctx context.Context, userId int32, pollId int64, option string, offset string, limit int32) (*mtproto.Messages_VotesList, error) {
	return c.PollCore.GetPollVoters(ctx, userId, pollId, option, offset, limit)
}
