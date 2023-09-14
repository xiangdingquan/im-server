package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountSetPrivacy(ctx context.Context, request *mtproto.TLAccountSetPrivacy) (reply *mtproto.Account_PrivacyRules, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.setPrivacy - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("account.setPrivacy - error: %v", err)
		return
	}

	key := model.FromInputPrivacyKeyType(request.Key)
	// Check request valid.
	if key == model.KEY_TYPE_INVALID {
		err = mtproto.ErrPrivacyKeyInvalid
		log.Errorf("account.setPrivacy - error: %v", err)
		return
	}

	ruleList := model.ToPrivacyRuleListByInput(md.UserId, request.Rules)

	if key != model.PHONE_NUMBER && len(ruleList) == 0 {
		ruleList = append(ruleList, mtproto.MakeTLPrivacyValueDisallowAll(nil).To_PrivacyRule())
	}

	if err = s.UserFacade.SetPrivacy(ctx, md.UserId, int(key), ruleList); err != nil {
		log.Errorf("account.setPrivacy - error: %v", err)
		return
	}

	reply = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
		Rules: ruleList,
	}).To_Account_PrivacyRules()

	userIdList, chatIdList, channelIdList := model.PickAllIdListByRules(ruleList)
	if len(userIdList) > 0 {
		// check error
		reply.Users = s.UserFacade.GetUserListByIdList(ctx, md.UserId, userIdList)
	} else {
		reply.Users = []*mtproto.User{}
	}
	reply.Chats = s.ChatFacade.GetChatListByIdList(ctx, md.UserId, chatIdList)
	reply.Chats = append(reply.Chats, s.ChannelFacade.GetChannelListByIdList(ctx, md.UserId, channelIdList...)...)

	go func() {
		updatePrivacy := &mtproto.TLUpdatePrivacy{Data2: &mtproto.Update{
			Key:   model.ToPrivacyKey(key),
			Rules: reply.Rules,
		}}

		syncUpdates := model.NewUpdatesLogic(md.UserId)
		syncUpdates.AddUpdate(updatePrivacy.To_Update())
		syncUpdates.AddUsers(reply.Users)
		syncUpdates.AddChats(reply.Chats)

		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, syncUpdates.ToUpdates())
	}()

	log.Debugf("account.setPrivacy#c9f81ce8 - reply: %s", reply.DebugString())
	return
}
