package service

import (
	"context"

	"open.chat/app/messenger/biz_server/auth/internal/logic"
	"open.chat/app/messenger/biz_server/auth/internal/model"
	gmodel "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthResendCode(ctx context.Context, request *mtproto.TLAuthResendCode) (reply *mtproto.Auth_SentCode, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.resendCode - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("auth.resendCode - error: %v", err)
		return
	}

	if request.PhoneCodeHash == "" {
		log.Errorf("check phone_code_hash error - empty")
		err = mtproto.ErrPhoneCodeHashEmpty
		return
	}

	phoneNumber, err := checkPhoneNumberInvalid(request.PhoneNumber)
	if err != nil {
		log.Errorf("check phone_number(%s) error - %v", request.PhoneNumber, err)
		err = mtproto.ErrPhoneNumberInvalid
		return
	}

	banned := s.BannedFacade.CheckPhoneNumberBanned(ctx, phoneNumber)
	if banned {
		log.Warnf("{phone_number: %s} banned: %v", phoneNumber, err)
		err = mtproto.ErrPhoneNumberBanned
		return
	}

	if s.AuthCore.IsBannedIp(ctx, md.ClientAddr) {
		err = mtproto.ErrPhoneNumberBanned
		return
	}

	actionType := logic.GetActionType(request)
	if err = s.AuthCore.CheckCanDoAction(ctx, md.AuthId, phoneNumber, actionType); err != nil {
		log.Warnf("check can do action - %s: %v", phoneNumber, err)
		return
	}

	codeData, err2 := s.DoAuthReSendCode(ctx,
		md.AuthId,
		phoneNumber,
		request.PhoneCodeHash,
		func(codeData2 *model.PhoneCodeTransaction) error {
			extraData, err := s.VerifyCodeInterface.SendSmsVerifyCode(context.Background(), phoneNumber, codeData2.PhoneCode, codeData2.PhoneCodeHash, gmodel.GetLangType(ctx, s.AuthSessionRpcClient))
			if err != nil {
				log.Errorf("sendSmsVerifyCode error: %v", err)
				return err
			}

			codeData2.SentCodeType = model.CodeTypeSms
			codeData2.NextCodeType = model.CodeTypeSms
			codeData2.State = model.CodeStateSent
			codeData2.PhoneCodeExtraData = extraData

			return nil
		})
	if err2 != nil {
		log.Error(err2.Error())
		err = err2
		return
	}

	// log
	logic.DoLogAuthAction(s.AuthCore, md, request.PhoneNumber, logic.GetActionType(request), "auth.resendCode#3ef1a9bf")

	reply = codeData.ToAuthSentCode()
	log.Debugf("auth.resendCode - reply: %s", reply.DebugString())
	return
}
