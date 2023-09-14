package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/gogo/protobuf/types"

	relay_client "open.chat/app/interface/relay/client"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) PhoneDiscardCall(ctx context.Context, request *mtproto.TLPhoneDiscardCall) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("phone.discardCall - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	CALL_ALREADY_ACCEPTED	The call was already accepted
	// 400	CALL_PEER_INVALID	The provided call peer object is invalid
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("phone.discardCall - error: %v", err)
		return nil, err
	}

	peer := request.GetPeer().To_InputPhoneCall()

	callData, err := s.Dao.GetPhoneCallSession(ctx, peer.GetId())
	if err != nil || callData == nil {
		log.Errorf("invalid peer: {%v}, err: %v", peer, err)
		err = mtproto.ErrCallPeerInvalid
		return nil, err
	}

	var duration *types.Int32Value = nil
	switch request.Reason.PredicateName {
	case mtproto.Predicate_phoneCallDiscardReasonBusy:
	case mtproto.Predicate_phoneCallDiscardReasonDisconnect:
	case mtproto.Predicate_phoneCallDiscardReasonMissed:
	case mtproto.Predicate_phoneCallDiscardReasonHangup:
		duration = &types.Int32Value{Value: request.GetDuration()}
	}

	phoneCallDiscarded := mtproto.MakeTLPhoneCallDiscarded(&mtproto.PhoneCall{
		Id:        callData.Id,
		NeedDebug: true,
		Reason:    request.GetReason(),
		Duration:  duration,
	}).To_PhoneCall()

	var toId int32
	if md.UserId == callData.AdminId {
		toId = callData.ParticipantId
	} else {
		toId = callData.AdminId
	}

	// check
	discardOk, err := relay_client.DiscardCallSession(ctx, callData.Id)
	_ = discardOk

	updatePhoneCall := mtproto.MakeTLUpdatePhoneCall(&mtproto.Update{
		PhoneCall: phoneCallDiscarded,
	}).To_Update()

	go func() {
		ctx2 := context.Background()

		toUpdates := model.NewUpdatesLogic(toId)
		// 1. add phoneCallRequested
		toUpdates.AddUpdate(updatePhoneCall)
		// 2. add Users
		toUsers := s.UserFacade.GetUserListByIdList(ctx2, toId, []int32{callData.AdminId, callData.ParticipantId})
		toUpdates.AddUsers(toUsers)
		sync_client.PushUpdates(ctx2, toId, toUpdates.ToUpdates())
	}()

	// 消息去重
	go func() {
		ctx2 := context.Background()

		action := mtproto.MakeTLMessageActionPhoneCall(&mtproto.MessageAction{
			CallId:   callData.Id,
			Reason:   request.GetReason(),
			Duration: duration,
		}).To_MessageAction()
		if request.GetDuration() > 0 {
			action.Duration = &types.Int32Value{Value: request.GetDuration()}
		}

		message := mtproto.MakeTLMessageService(&mtproto.Message{
			Out:             true,
			Date:            int32(time.Now().Unix()),
			FromId_FLAGPEER: model.MakePeerUser(callData.AdminId),
			ToId:            model.MakePeerUser(callData.ParticipantId),
			Action:          action,
		}).To_Message()

		randomId := rand.Int63()
		s.MsgFacade.PushUserMessage(ctx2, 0, callData.AdminId, callData.ParticipantId, randomId, message)
	}()
	// }

	/////////////////////////////////////////////////////////////////////////////////
	// reply
	replyUpdates := model.NewUpdatesLogic(md.UserId)
	replyUpdates.AddUpdate(updatePhoneCall)
	replyUsers := s.UserFacade.GetUserListByIdList(ctx, md.UserId, []int32{callData.AdminId, callData.ParticipantId})
	replyUpdates.AddUsers(replyUsers)
	reply := replyUpdates.ToUpdates()

	log.Debugf("phone.discardCall#78d413a6 - reply {%s}", logger.JsonDebugData(reply))
	return reply, nil
}
