package model

import (
	"time"

	"open.chat/mtproto"
)

var (
	userStatusEmpty = mtproto.MakeTLUserStatusEmpty(nil).To_UserStatus()
)

func MakeUserStatusOnline() *mtproto.UserStatus {
	now := time.Now().Unix()
	status := mtproto.MakeTLUserStatusOnline(&mtproto.UserStatus{
		Expires: int32(now + 60),
	}).To_UserStatus()
	return status
}

func MakeUserStatusOffline(lastSeenAt int32) *mtproto.UserStatus {
	return mtproto.MakeTLUserStatusOffline(&mtproto.UserStatus{
		WasOnline: int32(lastSeenAt),
	}).To_UserStatus()
}

func MakeUserStatus(lastSeenAt int32, allowTimestamp bool) *mtproto.UserStatus {
	now := int32(time.Now().Unix())

	if allowTimestamp {
		if now <= lastSeenAt+60 {
			status := mtproto.MakeTLUserStatusOnline(&mtproto.UserStatus{
				Expires: int32(lastSeenAt + 60),
			}).To_UserStatus()
			return status
		} else {
			status := mtproto.MakeTLUserStatusOffline(&mtproto.UserStatus{
				WasOnline: int32(lastSeenAt),
			}).To_UserStatus()
			return status
		}
	} else {
		if now-lastSeenAt >= 60*60*24*30 {
			return nil
		} else if now-lastSeenAt >= 60*60*24*7 {
			return mtproto.MakeTLUserStatusLastMonth(nil).To_UserStatus()
		} else if now-lastSeenAt >= 60*60*24*3 {
			return mtproto.MakeTLUserStatusLastWeek(nil).To_UserStatus()
		} else {
			return mtproto.MakeTLUserStatusRecently(nil).To_UserStatus()
		}
	}
}
