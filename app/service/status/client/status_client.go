package status_client

import (
	"context"

	status_facade "open.chat/app/service/status/facade"
	_ "open.chat/app/service/status/facade/redis"
)

var cli status_facade.StatusFacade

func New() (err error) {
	if cli == nil {
		cli, err = status_facade.NewStatusClient("redis")
	}
	return
}

func AddOnline(ctx context.Context, userId int32, authKeyId int64, serverId string) error {
	return cli.AddOnline(ctx, userId, authKeyId, serverId)
}

func ExpireOnline(ctx context.Context, userId int32, authKeyId int64) (bool, error) {
	return cli.ExpireOnline(ctx, userId, authKeyId)
}

func DelOnline(ctx context.Context, userId int32, authKeyId int64) error {
	return cli.DelOnline(ctx, userId, authKeyId)
}

func GetOnlineListByKeyIdList(ctx context.Context, authKeyIds []int64) (res []string, err error) {
	return cli.GetOnlineListByKeyIdList(ctx, authKeyIds)
}

func GetOnlineByKeyId(ctx context.Context, authKeyId int64) (res string, err error) {
	return cli.GetOnlineByKeyId(ctx, authKeyId)
}

func GetOnlineListExcludeKeyId(ctx context.Context, userId int32, authKeyId int64) (res map[int64]string, err error) {
	return cli.GetOnlineListExcludeKeyId(ctx, userId, authKeyId)
}

func GetOnlineListByUser(ctx context.Context, userId int32) (res map[int64]string, err error) {
	return cli.GetOnlineListByUser(ctx, userId)
}

func GetOnlineMapByUserList(ctx context.Context, userIdList []int32) (ress map[int64]string, onUserList []int32, err error) {
	return cli.GetOnlineMapByUserList(ctx, userIdList)
}

func CheckUserOnline(ctx context.Context, userId int32) bool {
	r, _ := GetOnlineListByUser(ctx, userId)
	return len(r) > 0
}
