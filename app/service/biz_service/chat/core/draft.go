package core

import (
	"context"
	"encoding/json"
	"time"

	"open.chat/app/service/biz_service/chat/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

func (m *ChatCore) SaveDraftMessage(ctx context.Context, userId, peerId int32, draft *mtproto.DraftMessage) (err error) {
	_, err = m.ChatParticipantsDAO.SaveDraft(ctx, 2, hack.String(model.TLObjectToJson(draft)), userId, peerId)
	return
}

func (m *ChatCore) ClearDraftMessage(ctx context.Context, userId, peerId int32) error {
	_, err := m.ChatParticipantsDAO.SaveDraft(ctx,
		1,
		hack.String(model.TLObjectToJson(model.MakeDraftMessageEmpty(int32(time.Now().Unix())))),
		userId,
		peerId)
	return err
}

func (m *ChatCore) GetAllDrafts(ctx context.Context, userId int32) (peers []int32, drafts []*mtproto.DraftMessage, err error) {
	var doList []dataobject.ChatParticipantsDO
	if doList, err = m.ChatParticipantsDAO.SelectAllDrafts(ctx, userId); err != nil {
		return
	}

	peers = make([]int32, 0, len(doList))
	drafts = make([]*mtproto.DraftMessage, 0, len(doList))

	for i := 0; i < len(doList); i++ {
		if doList[i].DraftMessageData == "" {
			log.Errorf("draft empty: %v", doList[i])
			continue
		}

		draft := new(mtproto.DraftMessage)
		if err2 := json.Unmarshal(hack.Bytes(doList[i].DraftMessageData), draft); err2 != nil {
			log.Errorf("invalid draft: %v", doList[i])
			continue
		}
		peers = append(peers, doList[i].ChatId)
		drafts = append(drafts, draft)
	}

	return
}

func (m *ChatCore) ClearAllDrafts(ctx context.Context, userId int32) error {
	_, err := m.ChatParticipantsDAO.ClearAllDrafts(ctx, userId)
	return err
}
