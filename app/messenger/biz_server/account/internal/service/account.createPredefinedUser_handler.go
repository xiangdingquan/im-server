package service

import (
	"context"

	"github.com/gogo/protobuf/types"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/phonenumber"
)

func (s *Service) AccountCreatePredefinedUser(ctx context.Context, request *mtproto.TLAccountCreatePredefinedUser) (*mtproto.PredefinedUser, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.createPredefinedUser - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// check
	if !md.IsAdmin {
		err := mtproto.ErrApiServerNeeded
		log.Errorf("account.createPredefinedUser - error: %v", err)
		return nil, err
	}

	if request.Phone == "" {
		err := mtproto.ErrPhoneNumberInvalid
		log.Errorf("account.createPredefinedUser - error: %v", err)
		return nil, err
	}

	// register name ruler
	// check first name invalid
	if request.GetFirstName().GetValue() == "" {
		err := mtproto.ErrFirstNameInvalid
		log.Errorf("check first_name error - error: %v", err)
		return nil, err
	}

	// 3.2. check phone_number
	// We need getRegionCode from phone_number
	pNumber, err := phonenumber.MakePhoneNumberHelper(request.Phone, "")
	if err != nil {
		log.Errorf("check phone_number error - %v", err)
		err = mtproto.ErrPhoneNumberInvalid
		return nil, err
	}
	phoneNumber := pNumber.GetNormalizeDigits()

	if ok, _ := s.UserFacade.CheckPhoneNumberExist(ctx, phoneNumber); ok {
		err = mtproto.ErrPhoneNumberOccupied
		return nil, err
	}

	// user, err := s.UserFacade.CreateNewUser(ctx, )
	key := crypto.CreateAuthKey()
	_, err = s.RPCSessionClient.SessionSetAuthKey(ctx, &authsessionpb.TLSessionSetAuthKey{
		AuthKey: &authsessionpb.AuthKeyInfo{
			AuthKeyId:          key.AuthKeyId(),
			AuthKey:            key.AuthKey(),
			AuthKeyType:        model.AuthKeyTypePerm,
			PermAuthKeyId:      key.AuthKeyId(),
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		},
		FutureSalt: nil,
	})
	if err != nil {
		log.Errorf("create user secret key error")
		return nil, err
	}
	user, err := s.UserFacade.CreateNewUser(
		ctx,
		key.AuthKeyId(),
		phoneNumber,
		pNumber.GetRegionCode(),
		request.GetFirstName().GetValue(),
		request.GetLastName().GetValue())
	if err != nil {
		log.Errorf("create user error: %v", err)
		return nil, err
	}

	//  check
	r, err := s.UserFacade.CreateNewPredefinedUser(ctx,
		phoneNumber,
		request.GetFirstName().GetValue(),
		request.GetLastName().GetValue(),
		request.GetUsername().GetValue(),
		request.GetCode(),
		request.GetVerified())
	if err != nil {
		log.Errorf("account.createPredefinedUser - error: %v", err)
		return nil, err
	}

	s.UserFacade.PredefinedBindRegisteredUserId(ctx, phoneNumber, user.GetId())
	r.RegisteredUserId = &types.Int32Value{Value: user.GetId()}
	_, err = s.updateUsername(ctx, md, request.GetUsername().GetValue())

	log.Debugf("account.createPredefinedUser - reply: %s", r.DebugString())
	return r, nil
}
