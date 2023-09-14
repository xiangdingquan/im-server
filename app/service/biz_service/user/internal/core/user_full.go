package core

import (
	"context"

	"github.com/gogo/protobuf/types"
	"open.chat/model"
	"open.chat/mtproto"
)

func (m *UserCore) GetFullUser(ctx context.Context, selfId, userId int32) (*mtproto.UserFull, error) {
	var (
		isMe  = selfId == userId
		users model.MutableUsers
		err   error
	)

	if isMe {
		users = m.GetMutableUsers(ctx, selfId)
	} else {
		users = m.GetMutableUsers(ctx, selfId, userId)
	}

	me, _ := users.GetImmutableUser(selfId)
	_ = me
	user, _ := users.GetImmutableUser(userId)

	if user == nil {
		err = mtproto.ErrUserIdInvalid
		return nil, err
	}

	userFull := mtproto.MakeTLUserFull(&mtproto.UserFull{
		Blocked:             m.CheckBlockUser(ctx, selfId, userId),
		PhoneCallsAvailable: true,
		PhoneCallsPrivate:   false,
		CanPinMessage:       isMe,
		HasScheduled:        false,
		VideoCallsAvailable: true,
		User:                user.ToImmutableUser(me),
		About:               nil,
		Settings:            nil,
		ProfilePhoto:        nil,
		NotifySettings:      nil,
		BotInfo:             nil,
		PinnedMsgId:         nil,
		CommonChatsCount:    0,
		FolderId:            nil,
	}).To_UserFull()
	if user.About() != "" {
		userFull.About = &types.StringValue{Value: user.About()}
	}

	userFull.Settings, err = m.GetPeerSettings(context.Background(), selfId, model.MakeUserPeerUtil(userId))
	if err != nil {
		return nil, err
	}
	userFull.ProfilePhoto = user.Photo()
	notifySettings, err := m.GetNotifySettings(context.Background(), selfId, model.MakeUserPeerUtil(userId))
	if err != nil || notifySettings == nil {
		notifySettings = mtproto.MakeTLPeerNotifySettings(nil).To_PeerNotifySettings()
	}
	userFull.NotifySettings = notifySettings

	if user.IsBot() {
		userFull.PhoneCallsAvailable = false
		userFull.PhoneCallsPrivate = false
		userFull.BotInfo = m.GetBotInfo(ctx, userId)
	}
	return userFull, nil
}
