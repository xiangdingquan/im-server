package model

import (
	"encoding/json"
	"fmt"

	"open.chat/mtproto"
)

// 可见类型 0公开 1自己 2联系人 3粉丝 4可见 5不可见
const (
	VisibleType_Public   = 0 // 公开
	VisibleType_Private  = 1 // 自己
	VisibleType_Friend   = 2 // 联系人
	VisibleType_Fans     = 3 // 粉丝
	VisibleType_Allow    = 4 // 可见
	VisibleType_NotAllow = 5 // 不可见
	VisibleType_Follow   = 6 // 关注
	VisibleType_User     = 7 // 用户
	VisibleType_Topic    = 8 //话题
)

type VisibleTypeUtil struct {
	VisibleType int8
	UserId      int32
	GroupTags   IDList
	UserIds     IDList
	Topic       string
	TopicId     int32
	Tab         int32
}

func (p VisibleTypeUtil) String() (s string) {
	switch p.VisibleType {
	case VisibleType_Public:
		return fmt.Sprintf("VisibleType_Public: {type: %d}", p.VisibleType)
	case VisibleType_Private:
		return fmt.Sprintf("VisibleType_Private: {type: %d}", p.VisibleType)
	case VisibleType_Friend:
		return fmt.Sprintf("VisibleType_Friend: {type: %d}", p.VisibleType)
	case VisibleType_Fans:
		return fmt.Sprintf("VisibleType_Fans: {type: %d}", p.VisibleType)
	case VisibleType_Allow:
		return fmt.Sprintf("VisibleType_Allow: {type: %d}", p.VisibleType)
	case VisibleType_NotAllow:
		return fmt.Sprintf("VisibleType_NotAllow: {type: %d}", p.VisibleType)
	case VisibleType_Follow:
		return fmt.Sprintf("VisibleType_Follow: {type: %d}", p.VisibleType)
	case VisibleType_User:
		return fmt.Sprintf("VisibleType_User: {type: %d}", p.VisibleType)
	case VisibleType_Topic:
		return fmt.Sprintf("VisibleType_Topic: {type: %d}", p.VisibleType)
	default:
		return fmt.Sprintf("VisibleType_UNKNOWN: {type: %d}", p.VisibleType)
	}
	// return
}

func FromVisibleType(selfId int32, visible *mtproto.VisibleType) (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{}
	switch visible.PredicateName {
	case mtproto.Predicate_visibleTypePublic:
		v.VisibleType = VisibleType_Public
	case mtproto.Predicate_visibleTypePrivate:
		v.VisibleType = VisibleType_Private
		v.UserId = selfId
	case mtproto.Predicate_visibleTypeFriend:
		v.VisibleType = VisibleType_Friend
	case mtproto.Predicate_visibleTypeFans:
		v.VisibleType = VisibleType_Fans
	case mtproto.Predicate_visibleTypeAllow:
		v.VisibleType = VisibleType_Allow
		v.GroupTags.AddIfNot(visible.GetTags()...)
		v.UserIds.AddIfNot(visible.GetUsers()...)
	case mtproto.Predicate_visibleTypeNotAllow:
		v.VisibleType = VisibleType_NotAllow
		v.GroupTags.AddIfNot(visible.GetTags()...)
		v.UserIds.AddIfNot(visible.GetUsers()...)
	case mtproto.Predicate_visibleTypeFollow:
		v.VisibleType = VisibleType_Follow
	case mtproto.Predicate_visibleTypeUser:
		v.VisibleType = VisibleType_User
		v.UserId = visible.GetUserId()
	case mtproto.Predicate_visibleTypeTopic:
		v.VisibleType = VisibleType_Topic
		s := struct {
			Topic string `json:"topic"`
			Tab   int32  `json:"tab"`
		}{}
		err := json.Unmarshal([]byte(visible.GetTopic()), &s)
		if err != nil {
			v.Topic = err.Error()
			v.Tab = 0
		}
		v.Topic = s.Topic
		v.Tab = s.Tab
	default:
		panic(fmt.Sprintf("FromVisibleType(%v) error!", visible))
	}
	return
}

func (v *VisibleTypeUtil) ToVisibleType() (visible *mtproto.VisibleType) {
	switch v.VisibleType {
	case VisibleType_Public:
		visible = mtproto.MakeTLVisibleTypePublic(nil).To_VisibleType()
	case VisibleType_Private:
		visible = mtproto.MakeTLVisibleTypePrivate(nil).To_VisibleType()
	case VisibleType_Friend:
		visible = mtproto.MakeTLVisibleTypeFriend(nil).To_VisibleType()
	case VisibleType_Fans:
		visible = mtproto.MakeTLVisibleTypeFans(nil).To_VisibleType()
	case VisibleType_Allow:
		visible = mtproto.MakeTLVisibleTypeAllow(&mtproto.VisibleType{
			Tags:  v.GroupTags,
			Users: v.UserIds,
		}).To_VisibleType()
	case VisibleType_NotAllow:
		visible = mtproto.MakeTLVisibleTypeNotAllow(&mtproto.VisibleType{
			Tags:  v.GroupTags,
			Users: v.UserIds,
		}).To_VisibleType()
	case VisibleType_Follow:
		visible = mtproto.MakeTLVisibleTypeFollow(nil).To_VisibleType()
	case VisibleType_User:
		visible = mtproto.MakeTLVisibleTypeUser(&mtproto.VisibleType{
			UserId: v.UserId,
		}).To_VisibleType()
	case VisibleType_Topic:
		visible = mtproto.MakeTLVisibleTypeTopic(&mtproto.VisibleType{
			Topic: v.Topic,
		}).To_VisibleType()
	default:
		panic(fmt.Sprintf("ToVisibleType(%v) error!", visible))
	}
	return
}

func MakeVisibleTypePublic() (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{
		VisibleType: VisibleType_Public,
	}
	return
}

func MakeVisibleTypePrivate(selfId int32) (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{
		VisibleType: VisibleType_Private,
		UserId:      selfId,
	}
	return
}

func MakeVisibleTypeFriend() (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{
		VisibleType: VisibleType_Friend,
	}
	return
}

func MakeVisibleTypeFans() (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{
		VisibleType: VisibleType_Fans,
	}
	return
}

func MakeVisibleTypeAllow(GroupTags, UserIds []int32) (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{
		VisibleType: VisibleType_Allow,
		GroupTags:   GroupTags,
		UserIds:     UserIds,
	}
	return
}

func MakeVisibleTypeNotAllow(GroupTags, UserIds []int32) (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{
		VisibleType: VisibleType_NotAllow,
		GroupTags:   GroupTags,
		UserIds:     UserIds,
	}
	return
}

func MakeTLVisibleTypeFollow(GroupTags, UserIds []int32) (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{
		VisibleType: VisibleType_Follow,
	}
	return
}

func MakeTLVisibleTypeUser(UserId int32) (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{
		VisibleType: VisibleType_User,
		UserId:      UserId,
	}
	return
}
