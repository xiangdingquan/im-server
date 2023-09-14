package model

import (
	"encoding/json"

	"open.chat/mtproto"
	"open.chat/pkg/hack"
)

type MediaPoll struct {
	UserId  int32                `json:"user_id"`
	Poll    *mtproto.Poll        `json:"poll"`
	Results *mtproto.PollResults `json:"results"`
}

func (m *MediaPoll) DebugString() string {
	return hack.String(GetFirstValue(json.Marshal(m)).([]byte))
}
func (m *MediaPoll) ToMessageMedia() *mtproto.MessageMedia {
	return mtproto.MakeTLMessageMediaPoll(&mtproto.MessageMedia{
		Poll:    m.Poll,
		Results: m.Results,
	}).To_MessageMedia()
}

func (m *MediaPoll) ToUpdateMessagePoll() *mtproto.Update {
	return mtproto.MakeTLUpdateMessagePoll(&mtproto.Update{
		PollId:  m.Poll.Id,
		Poll:    m.Poll,
		Results: m.Results,
	}).To_Update()
}

func GetPollIdByMessage(mediaPoll *mtproto.MessageMedia) (int64, error) {
	if mediaPoll == nil {
		return 0, mtproto.ErrMediaEmpty
	}
	if mediaPoll.PredicateName != mtproto.Predicate_messageMediaPoll {
		return 0, mtproto.ErrMediaInvalid
	}

	return mediaPoll.GetPoll().GetId(), nil
}

func GetPollByMessage(mediaPoll *mtproto.MessageMedia) (*mtproto.Poll, error) {
	if mediaPoll == nil {
		return nil, mtproto.ErrMediaEmpty
	}
	if mediaPoll.PredicateName != mtproto.Predicate_messageMediaPoll {
		return nil, mtproto.ErrMediaInvalid
	}

	return mediaPoll.GetPoll(), nil
}
