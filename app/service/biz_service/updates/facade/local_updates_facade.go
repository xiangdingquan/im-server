package updates_facade

import (
	"context"

	"open.chat/app/service/biz_service/updates/internal/core"
	"open.chat/app/service/biz_service/updates/internal/dao"
	"open.chat/mtproto"
)

type localUpdatesFacade struct {
	*core.UpdatesCore
}

func localUpdateFacadeInstance() UpdatesFacade {
	return &localUpdatesFacade{
		UpdatesCore: core.New(dao.New()),
	}
}

func (c *localUpdatesFacade) GetState(ctx context.Context, authKeyId int64, userId int32) (*mtproto.Updates_State, error) {
	return c.UpdatesCore.GetState(ctx, authKeyId, userId)
}

func (c *localUpdatesFacade) GetDifference(ctx context.Context, authKeyId int64, userId, pts, limit int32) (*mtproto.Updates_Difference, error) {
	return c.UpdatesCore.GetDifference(ctx, authKeyId, userId, pts, limit)
}

func init() {
	Register("local", localUpdateFacadeInstance)
}
