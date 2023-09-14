package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/phonenumber"
)

func (s *Service) AccountSendChangePhoneCode(ctx context.Context, request *mtproto.TLAccountSendChangePhoneCode) (*mtproto.Auth_SentCode, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.sendChangePhoneCode#8e57deb - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		sentCode *mtproto.Auth_SentCode
		err      error
	)

	switch request.GetConstructor() {
	case mtproto.CRC32_account_sendChangePhoneCode_8e57deb:
		request.Settings = &mtproto.CodeSettings{
			AllowFlashcall: request.AllowFlashcall,
			CurrentNumber:  mtproto.FromBool(request.CurrentNumber),
		}
	case mtproto.CRC32_account_sendChangePhoneCode_82574ae5:
	default:
		err = mtproto.ErrTypeConstructorInvalid
		log.Errorf("invalid constructor request: %s", request.DebugString())
		return nil, err
	}

	// 3. check number
	// 3.1. empty
	if request.PhoneNumber == "" {
		log.Errorf("check phone_number error - empty")
		return nil, mtproto.ErrPhoneNumberInvalid
	}

	// 3.2. check phone_number
	// We need getRegionCode from phone_number
	pNumber, err := phonenumber.MakePhoneNumberHelper(request.PhoneNumber, "")
	if err != nil {
		log.Errorf("check phone_number error - %v", err)
		return nil, mtproto.ErrPhoneNumberInvalid
	}

	phoneNumber := pNumber.GetNormalizeDigits()

	// 5. banned phone number
	banned := s.BannedFacade.CheckPhoneNumberBanned(ctx, phoneNumber)
	if banned {
		err = mtproto.ErrPhoneNumberBanned
		log.Warnf("{phone_number: %s} banned: %v", phoneNumber, err)
		return nil, err
	}

	if s.AccountCore.IsBannedIp(ctx, md.ClientAddr) {
		err = mtproto.ErrPhoneNumberBanned
		return nil, err
	}

	phoneRegistered, _ := s.UserFacade.CheckPhoneNumberExist(ctx, phoneNumber)
	if phoneRegistered {
		err = mtproto.ErrPhoneNumberOccupied
		log.Warnf("{phone_number: %s} registered: %v", phoneNumber, err)
		return nil, err
	}

	_ = sentCode
	err = mtproto.ErrMethodNotImpl
	return nil, err
}
