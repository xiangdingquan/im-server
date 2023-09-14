package core

import (
	"context"

	"open.chat/model"
)

func (m *UserCore) CheckLastSeenOnline(ctx context.Context, selfId, userId int32, isContact bool) bool {
	return m.CheckPrivacy(ctx, model.STATUS_TIMESTAMP, selfId, userId, isContact)
}

func (m *UserCore) CheckAllowCalls(ctx context.Context, selfId, userId int32, isContact bool) bool {
	return m.CheckPrivacy(ctx, model.PHONE_CALL, selfId, userId, isContact)
}

func (m *UserCore) CheckAllowProfilePhoto(ctx context.Context, selfId, userId int32, isContact bool) bool {
	return m.CheckPrivacy(ctx, model.PROFILE_PHOTO, selfId, userId, isContact)
}
