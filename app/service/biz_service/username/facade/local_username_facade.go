package username_facade

import (
	"context"

	"open.chat/app/service/biz_service/username/internal/core"
	"open.chat/app/service/biz_service/username/internal/dal/dataobject"
	"open.chat/app/service/biz_service/username/internal/dao"
)

type localUsernameFacade struct {
	*core.UsernameCore
}

func New() UsernameFacade {
	return &localUsernameFacade{
		UsernameCore: core.New(dao.New()),
	}
}

func (c *localUsernameFacade) GetAccountUsername(ctx context.Context, userId int32) (string, error) {
	return c.UsernameCore.GetAccountUsername(ctx, userId)
}

func (c *localUsernameFacade) CheckAccountUsername(ctx context.Context, userId int32, username string) (int, error) {
	return c.UsernameCore.CheckAccountUsername(ctx, userId, username)
}

func (c *localUsernameFacade) UpdateUsernameByPeer(ctx context.Context, peerType, peerId int32, username string) (bool, error) {
	return c.UsernameCore.UpdateUsernameByPeer(ctx, peerType, peerId, username)
}

func (c *localUsernameFacade) GetChannelUsername(ctx context.Context, channelId int32) (string, error) {
	return c.UsernameCore.GetChannelUsername(ctx, channelId)
}

func (c *localUsernameFacade) CheckChannelUsername(ctx context.Context, channelId int32, username string) (int, error) {
	return c.UsernameCore.CheckChannelUsername(ctx, channelId, username)
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////
func (c *localUsernameFacade) CheckUsername(ctx context.Context, username string) (int, error) {
	return c.UsernameCore.CheckUsername(ctx, username)
}

func (c *localUsernameFacade) UpdateUsername(ctx context.Context, peerType, peerId int32, username string) (bool, error) {
	return c.UsernameCore.UpdateUsernameByPeer(ctx, peerType, peerId, username)
}

func (c *localUsernameFacade) DeleteUsername(ctx context.Context, username string) (bool, error) {
	return c.UsernameCore.DeleteUsername(ctx, username)
}

func (c *localUsernameFacade) ResolveUsername(ctx context.Context, username string) (int32, int32, error) {
	peer, err := c.UsernameCore.ResolveUsername(ctx, username)
	if err != nil {
		return 0, 0, err
	}
	return peer.PeerType, peer.PeerId, err
}

func (c *localUsernameFacade) GetListByUsernameList(ctx context.Context, names []string) (map[string]*dataobject.UsernameDO, error) {
	return c.UsernameCore.GetListByUsernameList(ctx, names)
}

func (c *localUsernameFacade) DeleteUsernameByPeer(ctx context.Context, peerType, peerId int32) error {
	return c.UsernameCore.DeleteUsernameByPeer(ctx, peerType, peerId)
}

func init() {
	Register("local", New)
}
