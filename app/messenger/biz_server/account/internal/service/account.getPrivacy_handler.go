package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetPrivacy(ctx context.Context, request *mtproto.TLAccountGetPrivacy) (reply *mtproto.Account_PrivacyRules, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getPrivacy - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		ruleList []*mtproto.PrivacyRule
		// rules    *mtproto.TLAccountPrivacyRules
	)

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("account.getPrivacy - error: %v", err)
		return
	}

	key := model.FromInputPrivacyKeyType(request.Key)
	// Check request valid.
	if key == model.KEY_TYPE_INVALID {
		err = mtproto.ErrPrivacyKeyInvalid
		log.Errorf("account.getPrivacy - error: %v", err)
		return
	}

	ruleList, _ = s.UserFacade.GetPrivacy(ctx, md.UserId, key)
	if len(ruleList) == 0 {
		if key == model.PHONE_NUMBER {
			reply = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
				Rules: []*mtproto.PrivacyRule{mtproto.MakeTLPrivacyValueDisallowAll(nil).To_PrivacyRule()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
			}).To_Account_PrivacyRules()
		} else {
			reply = mtproto.MakeTLAccountPrivacyRules(&mtproto.Account_PrivacyRules{
				Rules: []*mtproto.PrivacyRule{mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule()},
				Users: []*mtproto.User{},
				Chats: []*mtproto.Chat{},
			}).To_Account_PrivacyRules()
		}
	} else {
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
	}

	log.Debugf("account.getPrivacy#dadbc950 - reply: %s", reply.DebugString())
	return
}
