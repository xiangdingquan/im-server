package helper

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/messenger/msg/msgpb"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/util"

	msg_facade "open.chat/app/messenger/msg/facade"
)

type (
	// TJsonMessage2 .
	TJsonMessage struct {
		Mid uint32
		Sid uint32
		Msg string
	}

	// TSendjson .
	TSendjson struct {
		From      uint32
		AuthKeyId int64
		Mid       uint32
		Sid       uint32
	}
)

func MakeSender(from uint32, authKeyId int64, mid uint32, sid uint32) *TSendjson {
	return &TSendjson{
		From:      from,
		AuthKeyId: authKeyId,
		Mid:       mid,
		Sid:       sid,
	}
}

// ParsingMessage .
func (m *TJsonMessage) Parse(mid, sid uint32, v interface{}) bool {
	str := `^ct!([a-f\dA-F]{32})!(\d+)!(\d+)!([^$]+)`
	r := regexp.MustCompile(str)
	matchs := r.FindStringSubmatch(m.Msg)
	if len(matchs) != 5 {
		return false
	}
	text := matchs[4]
	if text == "" {
		return false
	}
	vv := md5.Sum([]byte(text))
	sMd5 := hex.EncodeToString(vv[:])
	if sMd5 != strings.ToLower(matchs[1]) {
		return false
	}
	m.Mid, _ = util.StringToUint32(matchs[2])
	m.Sid, _ = util.StringToUint32(matchs[3])
	checked := (mid == 0 || m.Mid == mid) && (sid == 0 || m.Sid == sid)
	s, ok := v.(*string)
	if ok {
		*s = text
	} else if checked {
		err := json.Unmarshal([]byte(text), v)
		if err != nil {
			return false
		}
	}
	return checked
}

// ToString .
func (m *TJsonMessage) ToMessage(v interface{}) bool {
	//Data map[string]interface{} `json:"data"`
	m.Msg = ""
	if v == nil {
		v = struct{}{}
	}
	cbJson, err := json.Marshal(v)
	if err != nil {
		return false
	}
	vv := md5.Sum(cbJson)
	sMd5 := hex.EncodeToString(vv[:])
	m.Msg = "ct!" + sMd5
	m.Msg += "!" + util.Int64ToString(int64(m.Mid))
	m.Msg += "!" + util.Int64ToString(int64(m.Sid))
	m.Msg += "!" + string(cbJson)
	return true
}

var msgFacade msg_facade.MsgFacade

func (m *TSendjson) pushUserMessage(ctx context.Context, toUserId int32, v interface{}) error {
	jsonMsg := TJsonMessage{
		Mid: m.Mid,
		Sid: m.Sid,
	}
	if !jsonMsg.ToMessage(v) {
		return errors.New("make message fail")
	}
	message := mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		Date:            int32(time.Now().Unix()),
		FromId_FLAGPEER: model.MakePeerUser((int32)(m.From)),
		ToId:            model.MakePeerUser(toUserId),
		Message:         jsonMsg.Msg,
	}).To_Message()

	if msgFacade == nil {
		var err error
		msgFacade, err = msg_facade.NewMsgFacade("emsg")
		if err != nil {
			panic(err)
		}
	}
	return msgFacade.PushUserMessage(ctx, 1, (int32)(m.From), toUserId, rand.Int63(), message)
}

func (m *TSendjson) send(ctx context.Context, peer *model.PeerUtil, v interface{}) error {
	jsonMsg := TJsonMessage{
		Mid: m.Mid,
		Sid: m.Sid,
	}
	if !jsonMsg.ToMessage(v) {
		return errors.New("make message fail")
	}
	outboxMsg := msgpb.OutboxMessage{
		NoWebpage:  true,
		Background: false,
		RandomId:   rand.Int63(),
		Message: mtproto.MakeTLMessage(&mtproto.Message{
			Out:             true,
			FromId_FLAGPEER: model.MakePeerUser(int32(m.From)),
			ToId:            peer.ToPeer(),
			Date:            int32(time.Now().Unix()),
			Message:         jsonMsg.Msg,
		}).To_Message(),
		ScheduleDate: &types.Int32Value{Value: 0},
	}

	if msgFacade == nil {
		var err error
		msgFacade, err = msg_facade.NewMsgFacade("emsg")
		if err != nil {
			panic(err)
		}
	}

	reply, err := msgFacade.SendMessage(ctx, int32(m.From), m.AuthKeyId, peer, &outboxMsg)
	go func() {
		sync_client.SyncUpdatesMe(ctx, (int32)(m.From), m.AuthKeyId, 0, "", reply)
		//fmt.Printf("%v,%d,%d,%v", ctx, m.From, m.AuthKeyId, reply)
	}()
	return err
}

// SendToUser .
func (m *TSendjson) SendToUser(ctx context.Context, userID uint32, v interface{}) error {
	return m.send(ctx, model.MakeUserPeerUtil(int32(userID)), v)
}

// SendToChat .
func (m *TSendjson) SendToChat(ctx context.Context, chatID uint32, v interface{}) error {
	return m.send(ctx, model.MakeChatPeerUtil(int32(chatID)), v)
}

// SendToChannel .
func (m *TSendjson) SendToChannel(ctx context.Context, channelID uint32, v interface{}) error {
	return m.send(ctx, model.MakeChannelPeerUtil(int32(channelID)), v)
}

// PushMessage .
func (m *TSendjson) PushMessage(ctx context.Context, toUserId uint32, v interface{}) error {
	return m.pushUserMessage(ctx, int32(toUserId), v)
}
