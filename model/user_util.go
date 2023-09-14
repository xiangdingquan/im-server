package model

import (
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func ToUserIdByInput(userSelfId int32, inputUser *mtproto.InputUser) int32 {
	switch inputUser.PredicateName {
	case mtproto.Predicate_inputUserEmpty:
		return 0
	case mtproto.Predicate_inputUserSelf:
		return userSelfId
	case mtproto.Predicate_inputUser:
		return inputUser.UserId
	default:
		log.Errorf("invalid inputUser classID - %v", inputUser)
		return 0
	}
}

func ToUserIdListByInput(userSelfId int32, inputUsers []*mtproto.InputUser) []int32 {
	idList := make([]int32, 0, len(inputUsers))
	for _, user := range inputUsers {
		id := ToUserIdByInput(userSelfId, user)
		if id > 0 {
			idList = append(idList, id)
		} else {
			// ignore in
		}
	}
	return idList
}
