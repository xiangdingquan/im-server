package helper

import (
	"context"
	"encoding/json"

	"open.chat/app/json/consts"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

// PushUpdate .
type PushUpdate struct {
	Action consts.OnAction        `json:"action"`
	From   uint32                 `json:"from"`
	Data   map[string]interface{} `json:"data"`
}

// GetData .
func (m *PushUpdate) GetData(v interface{}) error {
	strJSON, err := json.Marshal(m.Data)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}

	err = json.Unmarshal(strJSON, v)
	if err != nil {
		return err
	}

	return nil
}

func (m *PushUpdate) makeUpdate(v interface{}) (*mtproto.Update, error) {
	var (
		bjson []byte
		err   error
	)
	if v != nil {
		bjson, err = json.Marshal(v)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bjson, &m.Data)
		if err != nil {
			return nil, err
		}
	}

	bjson, err = json.Marshal(m)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}

	return mtproto.MakeTLUpdateBotWebhookJSON(&mtproto.Update{
		Data_DATAJSON: mtproto.MakeTLDataJSON(&mtproto.DataJSON{
			Data: string(bjson),
		}).To_DataJSON(),
	}).To_Update(), nil
}

// PushUsers .
func (m *PushUpdate) ToUsers(ctx context.Context, to []uint32, v interface{}) error {
	upd, err := m.makeUpdate(v)
	if err != nil {
		return err
	}
	go func() {
		for _, uid := range to {
			if uid != m.From {
				sync_client.PushUpdates(ctx, (int32)(uid), model.MakeUpdatesByUpdates(upd))
			}
		}
	}()
	return nil
}

func (m *PushUpdate) ToChat(ctx context.Context, to uint32, v interface{}) error {
	upd, err := m.makeUpdate(v)
	if err != nil {
		return err
	}
	return sync_client.BroadcastChatUpdates(ctx, int32(to), model.MakeUpdatesByUpdates(upd))
}

func (m *PushUpdate) ToChannel(ctx context.Context, to uint32, v interface{}) error {
	upd, err := m.makeUpdate(v)
	if err != nil {
		return err
	}
	return sync_client.BroadcastChannelUpdates(ctx, int32(to), model.MakeUpdatesByUpdates(upd))
}
