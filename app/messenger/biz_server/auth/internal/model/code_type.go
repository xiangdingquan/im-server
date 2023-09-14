package model

import (
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	CodeTypeNone      = 0
	CodeTypeApp       = 1
	CodeTypeSms       = 2
	CodeTypeCall      = 3
	CodeTypeFlashCall = 4
)

const (
	CodeStateOk      = 1
	CodeStateSend    = 2
	CodeStateSent    = 3
	CodeStateReSent  = 6
	CodeStateSignIn  = 4
	CodeStateSignUp  = 5
	CodeStateDeleted = -1
)

func MakeCodeType(phoneRegistered, allowFlashCall, currentNumber bool) (int, int) {
	_ = phoneRegistered
	_ = allowFlashCall
	_ = currentNumber

	sentCodeType := CodeTypeApp
	nextCodeType := CodeTypeNone
	return sentCodeType, nextCodeType
}

func makeAuthCodeType(codeType int) *mtproto.Auth_CodeType {
	switch codeType {
	case CodeTypeSms:
		return mtproto.MakeTLAuthCodeTypeSms(nil).To_Auth_CodeType()
	case CodeTypeCall:
		return mtproto.MakeTLAuthCodeTypeCall(nil).To_Auth_CodeType()
	case CodeTypeFlashCall:
		return mtproto.MakeTLAuthCodeTypeFlashCall(nil).To_Auth_CodeType()
	default:
		return nil
	}
}

func makeAuthSentCodeType(codeType, codeLength int, pattern string) (authSentCodeType *mtproto.Auth_SentCodeType) {
	switch codeType {
	case CodeTypeApp:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeApp(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case CodeTypeSms:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeSms(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case CodeTypeCall:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeCall(&mtproto.Auth_SentCodeType{
			Length: int32(codeLength),
		}).To_Auth_SentCodeType()
	case CodeTypeFlashCall:
		authSentCodeType = mtproto.MakeTLAuthSentCodeTypeFlashCall(&mtproto.Auth_SentCodeType{
			PredicateName: mtproto.Predicate_auth_sentCodeTypeFlashCall,
			Constructor:   mtproto.CRC32_auth_sentCodeTypeFlashCall,
			Length:        int32(codeLength),
			Pattern:       pattern,
		}).To_Auth_SentCodeType()
	default:
		err := fmt.Errorf("invalid sentCodeType: %d", codeType)
		log.Errorf("makeAuthSentCodeType - %v", err)
		panic(err)
	}

	return
}
