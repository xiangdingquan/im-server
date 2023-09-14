package service

import (
	"context"
	"encoding/json"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

type searchPublicChannelQuery struct {
	Ver    int   `json:"ver"`
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

func (s *Service) ContactsSearch(ctx context.Context, request *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.search - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		limit = request.Limit
	)
	if limit > 50 {
		limit = 50
	}

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("contacts.search - error: %v", err)
		return nil, err
	}

	q := request.Q

	if q == "" {
		err := mtproto.ErrSearchQueryEmpty
		log.Errorf("contacts.search - error: %v", err)
		return nil, err
	}

	if q[0] == '@' {
		q = q[1:]
	}

	if len(q) < 5 {
		err := mtproto.ErrQueryTooShort
		log.Errorf("contacts.search - error: %v", err)
		return nil, err
	}

	var (
		userIdList    []int32
		channelIdList []int32
	)

	found := mtproto.MakeTLContactsFound(&mtproto.Contacts_Found{
		MyResults: []*mtproto.Peer{},
		Results:   []*mtproto.Peer{},
		Users:     []*mtproto.User{},
		Chats:     []*mtproto.Chat{},
	}).To_Contacts_Found()
	if len(q) >= 5 && limit > 0 {
		userIdList, channelIdList = s.UserFacade.SearchContacts(ctx, md.UserId, q, limit)
	}

	j := searchPublicChannelQuery{}
	err := json.Unmarshal([]byte(q), &j)
	if err != nil {
		channelIdList = append(channelIdList, s.ChannelFacade.SearchChannelByTitle(ctx, q)...)
	} else {
		channelIdList = append(channelIdList, s.ChannelFacade.SearchPublicChannels(ctx, j.Offset, j.Limit)...)
	}

	if len(userIdList) > 0 {
		userList := s.UserFacade.GetUserListByIdList(ctx, md.UserId, userIdList)
		found.Users = userList
		for _, u := range userList {
			peer := mtproto.MakeTLPeerUser(&mtproto.Peer{
				UserId: u.GetId(),
			})
			if u.GetContact() {
				found.MyResults = append(found.MyResults, peer.To_Peer())
			} else {
				found.Results = append(found.Results, peer.To_Peer())
			}
		}
	}

	if len(channelIdList) > 0 {
		channelList := s.ChannelFacade.GetChannelListByIdList(ctx, md.UserId, channelIdList...)
		for _, c := range channelList {
			if c.PredicateName == mtproto.Predicate_chatEmpty {
				continue
			}
			found.Chats = append(found.Chats, c)
			peer := mtproto.MakeTLPeerChannel(&mtproto.Peer{
				ChannelId: c.GetId(),
			})
			found.Results = append(found.Results, peer.To_Peer())
		}
	}

	log.Debugf("contacts.search#11f812d8 - reply: %s", logger.JsonDebugData(found))
	return found, nil
}
