package service

import (
	"context"
	"open.chat/model"
	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// blogs.getHistory#8386e43c visible_type:VisibleType offset_id:int add_offset:int limit:int hash:int = MicroBlogs;
func (s *Service) BlogsGetHistory(ctx context.Context, request *mtproto.TLBlogsGetHistory) (*mtproto.MicroBlogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.getHistory - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		reply *mtproto.MicroBlogs
		err   error
	)

	me := md.GetUserId()
	visible := model.FromVisibleType(me, request.VisibleType)
	switch visible.VisibleType {
	case model.VisibleType_Public:
	case model.VisibleType_Private:
	case model.VisibleType_Friend:
	case model.VisibleType_Follow:
	case model.VisibleType_User:
		if me != visible.UserId {
			user, err := s.UserFacade.GetUserById(ctx, me, visible.UserId)
			if err != nil || user == nil || user.Deleted {
				if err == nil {
					err = mtproto.ErrUserIdInvalid
				}
				log.Errorf("invalid user: %v", visible.UserId)
				return nil, err
			}
		}
	case model.VisibleType_Topic:
		if visible.Tab == 0 {
			log.Errorf("blogs.getHistory, parse visibleTypeTopic failed, error: %s", visible.Topic)
			return nil, mtproto.ErrButtonTypeInvalid
		}
		var topicId int32
		topicId, err = s.BlogFacade.GetTopicIdByName(ctx, visible.Topic)
		if err != nil || topicId == 0 {
			return nil, mtproto.ErrButtonTypeInvalid
		}
		visible.TopicId = topicId
	default:
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.getHistory: %v", err)
		return nil, err
	}

	if visible.VisibleType == model.VisibleType_Private || visible.VisibleType == model.VisibleType_User {
		isHisFriend, mutual := true, true
		if visible.VisibleType == model.VisibleType_User {
			_, mutual = s.UserFacade.GetContactAndMutual(ctx, me, visible.UserId)
			isHisFriend, _ = s.UserFacade.GetContactAndMutual(ctx, visible.UserId, me)
		}
		reply, err = s.BlogFacade.GetUserHistory(ctx, me, visible.UserId, isHisFriend, mutual, request.GetOffsetId(), request.GetAddOffset(), request.GetLimit())
	} else {
		reply, err = s.BlogFacade.GetCircleHistory(ctx, me, visible, request.GetOffsetId(), request.GetAddOffset(), request.GetLimit(), func() []int32 {
			return s.UserFacade.GetMutualContactUserIdList(ctx, true, me)
		}, func(uid int32) bool {
			isHisFriend, _ := s.UserFacade.GetContactAndMutual(ctx, uid, me)
			return isHisFriend
		})
	}

	if err != nil {
		log.Errorf("blogs.getHistory#8386e43c - error: %v", err)
		return nil, err
	}

	log.Debugf("blogs.getHistory#8386e43c - reply: %s", reply.DebugString())
	return reply, err
}
