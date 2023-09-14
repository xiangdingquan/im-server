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

// TrelayData .
type TrelayData struct {
	Action consts.OnAction        `json:"action"`
	From   uint32                 `json:"from"`
	To     []uint32               `json:"to"`
	Data   map[string]interface{} `json:"data"`
}

// GetData .
func (m *TrelayData) GetData(v interface{}) error {
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

// PushUpdates .
func (m *TrelayData) PushUpdates(ctx context.Context, v interface{}) error {
	if v != nil {
		sjson, err := json.Marshal(v)
		if err != nil {
			return err
		}

		err = json.Unmarshal(sjson, &m.Data)
		if err != nil {
			return err
		}
	}

	strJSON, err := json.Marshal(m)
	if err != nil {
		log.Errorf(err.Error())
		return err
	}
	//sync_client.SyncUpdatesMe(ctx, md.UserId, 0, 0, md.ServerId, model.MakeUpdatesByUpdates(upd))
	//sync_client.SyncUpdatesNotMe(ctx, uid, 0, model.MakeUpdatesByUpdates(upd))
	upd := mtproto.MakeTLUpdateBotWebhookJSON(&mtproto.Update{
		Data_DATAJSON: mtproto.MakeTLDataJSON(&mtproto.DataJSON{
			Data: string(strJSON),
		}).To_DataJSON(),
	}).To_Update()

	go func() {
		for _, uid := range m.To {
			if uid != m.From {
				sync_client.PushUpdates(ctx, (int32)(uid), model.MakeUpdatesByUpdates(upd))
				//fmt.Printf("%v,%d,%v", ctx, uid, model.MakeUpdatesByUpdates(upd))
			}
		}
	}()
	return nil
}
