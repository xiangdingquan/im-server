package model

import (
	"fmt"
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/pkg/env2"
	"open.chat/mtproto"
)

func MakeSignInServiceNotification(user *mtproto.User, authId int64, client, region, clientIp string) *mtproto.Update {
	now := time.Now()
	notification := mtproto.MakeTLUpdateServiceNotification(&mtproto.Update{
		Popup:          false,
		InboxDate:      &types.Int32Value{Value: int32(now.Unix())},
		Type:           fmt.Sprintf("auth%d_%d", authId, now.Unix()),
		Message_STRING: "",
		Media:          mtproto.MakeTLMessageMediaEmpty(nil).To_MessageMedia(),
		Entities:       nil,
	}).To_Update()

	notification.Message_STRING, notification.Entities = MakeTextAndMessageEntities([]MessageBuildEntry{
		{
			Text:       "",
			Param:      "新登录.",
			EntityType: mtproto.Predicate_messageEntityBold,
		},
		{
			Text: fmt.Sprintf("%s,\n\n  系统于%s时间检测到有新设备登录您的帐户。\n\n设备：%s\n位置：%s (IP = ",
				GetUserName(user),
				now.UTC(),
				client,
				region),
			Param:      clientIp,
			EntityType: mtproto.Predicate_messageEntityTextUrl,
		},
		{
			Text: ")\n\n感谢您的支持！\n\n    " + env2.MY_APP_NAME,
		},
	})

	return notification
}
