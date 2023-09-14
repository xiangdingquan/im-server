package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"open.chat/pkg/log"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/job/scheduled/internal/core"
	redpacket_core "open.chat/app/json/services/handler/redpacket/core"
	remittance_core "open.chat/app/json/services/handler/remittance/core"
	msg_facade "open.chat/app/messenger/msg/facade"
	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	_ "open.chat/app/service/biz_service/message/facade"
	message_facade "open.chat/app/service/biz_service/message/facade"
	"open.chat/model"
	"open.chat/mtproto"
)

const (
	system_uid int32 = 777000
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Service struct {
	running bool
	conf    *Config
	msg_facade.MsgFacade
	message_facade.MessageFacade
	*core.ScheduledCore
	*redpacket_core.RedPacketCore
	*remittance_core.RemittanceCore
}

func New() *Service {
	var (
		ac  = &Config{}
		err error
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s := &Service{
		conf:           ac,
		ScheduledCore:  core.New(nil),
		RedPacketCore:  redpacket_core.New(nil),
		RemittanceCore: remittance_core.New(nil),
	}

	s.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)
	s.MessageFacade, err = message_facade.NewMessageFacade("local")
	checkErr(err)
	sync_client.New()

	s.running = true
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return nil
}

// Close close the resources.
func (s *Service) Close() error {
	s.running = false
	return nil
}

// Close close the resources.
func (s *Service) RunLoop() {
	for s.running {
		<-time.After(time.Second)
		s.onTimer()
	}
}

func (s *Service) onTimer() {
	s.onScheduledMessage()
	s.onRedPacket()
	s.onHandleCountDownMsg()
	s.onTransfer()
}

func (s *Service) onScheduledMessage() {
	boxList := s.MessageFacade.GetScheduledTimeoutMessageList(context.Background(), int32(time.Now().Unix()))
	for _, box := range boxList {
		var (
			peer          = model.FromPeer(box.Message.ToId)
			fromAuthKeyId = int64(0)
		)
		box.Message.FromScheduled = true
		meUpdates, err := s.MsgFacade.SendMessage(context.Background(), box.SendUserId, fromAuthKeyId, peer, &msgpb.OutboxMessage{
			NoWebpage:    false,
			Background:   false,
			RandomId:     rand.Int63(),
			Message:      box.Message,
			ScheduleDate: nil,
		})
		if err != nil {
			log.Errorf("sendMessage error: %v", err)
			continue
		} else {
			log.Debugf("sendMessage - result: %s", meUpdates.DebugString())
		}

		err = s.MessageFacade.DeleteScheduledMessageList(context.Background(), box.SelfUserId, peer, []int32{box.MessageId})
		if err != nil {
			log.Errorf("sendMessage error: %v", err)
			continue
		}

		var updateMessageID *mtproto.Update
		for _, update := range meUpdates.Updates {
			if update.PredicateName == mtproto.Predicate_updateMessageID {
				updateMessageID = update
				break
			}
		}

		updates := model.MakeUpdatesHelper(mtproto.MakeTLUpdateDeleteScheduledMessages(&mtproto.Update{
			Peer_PEER: box.Message.ToId,
			Messages:  []int32{box.MessageId},
		}).To_Update())

		if updateMessageID != nil {
			updates.PushFrontUpdate(mtproto.MakeTLUpdateMessageID(&mtproto.Update{
				Id_INT32: box.MessageId,
				RandomId: box.RandomId,
			}).To_Update())
		}

		sync_client.PushUpdates(context.Background(), box.SelfUserId, updates.GetUpdates())
	}
}

const giveBackMessageTplCN = `【红包退款通知】
您有[%.2f]红包超时未领取完毕，已退回至您的钱包余额。`
const giveBackMessageTplEN = `[Refund red packet notice]
You have [%.2f] red packet that are overdue and have not been collected. They have been returned to your wallet balance.`

func (s *Service) onRedPacket() {
	var (
		ctx = context.Background()
		err error
	)
	redPackeds := s.RedPacketCore.GetRedPacketTimeoutList(ctx, int32(time.Now().Unix())-24*3600)
	for _, r := range redPackeds {
		err = s.RedPacketCore.GiveBackRedPacket(ctx, r.Rid, uint32(r.Uid), r.Price)
		if err != nil {
			log.Errorf("giveBackRedPacket error: %v", err)
			continue
		}

		messageTpl := giveBackMessageTplEN
		langCode := s.ScheduledCore.GetUserLangCode(ctx, r.Uid)
		if langCode != "en" {
			messageTpl = giveBackMessageTplCN
		}

		message := mtproto.MakeTLMessage(&mtproto.Message{
			Out:             true,
			Date:            int32(time.Now().Unix()),
			FromId_FLAGPEER: model.MakePeerUser(system_uid),
			ToId:            model.MakePeerUser(r.Uid),
			Message:         fmt.Sprintf(messageTpl, r.Price),
			Entities: []*mtproto.MessageEntity{
				mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
					Offset: 1,
					Length: 6,
				}).To_MessageEntity(),
			},
		}).To_Message()
		err = s.MsgFacade.PushUserMessage(ctx, 1, system_uid, r.Uid, rand.Int63(), message)
		if err != nil {
			log.Errorf("PushUserMessage%s", err.Error())
		}
	}
}

func (s *Service) onHandleCountDownMsg() {
	ctx := context.Background()
	messages := s.MessageFacade.GetEphemeralExpireList(ctx)
	if len(messages) > 0 {
		for _, message := range messages {
			peer := &model.PeerUtil{PeerType: model.PEER_EMPTY}
			revoke := message.PeerType != model.PEER_CHAT
			if message.PeerType == model.PEER_CHANNEL {
				peer.PeerType = model.PEER_CHANNEL
				peer.PeerId = message.PeerId
				revoke = false
			}
			s.MsgFacade.DeleteMessages(ctx, message.UId, 0, peer, revoke, []int32{message.MsgId})
		}
		s.MessageFacade.DelEphemeralList(ctx, messages)
	}
}

const trefundMessageTplCN = `【转账退款通知】
您有[%.2f]转账超时未领取完毕，已退回至您的钱包余额。`
const trefundMessageTplEN = `[Refund transfer notice]
You have [%.2f] transfer timeout and have not received it. It has been returned to your wallet balance.`

func (s *Service) onTransfer() {
	var (
		ctx = context.Background()
		err error
	)
	tfs := s.RemittanceCore.GetRemittanceTimeoutList(ctx, int32(time.Now().Unix())-24*3600)
	for _, r := range tfs {
		err = s.RemittanceCore.RefundByTimeout(ctx, r.Rid, uint32(r.Uid), r.Price)
		if err != nil {
			log.Errorf("giveBackRedPacket error: %v", err)
			continue
		}

		messageTpl := trefundMessageTplEN
		langCode := s.ScheduledCore.GetUserLangCode(ctx, r.Uid)
		if langCode != "en" {
			messageTpl = trefundMessageTplCN
		}

		message := mtproto.MakeTLMessage(&mtproto.Message{
			Out:             true,
			Date:            int32(time.Now().Unix()),
			FromId_FLAGPEER: model.MakePeerUser(system_uid),
			ToId:            model.MakePeerUser(r.Uid),
			Message:         fmt.Sprintf(messageTpl, r.Price),
			Entities: []*mtproto.MessageEntity{
				mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
					Offset: 1,
					Length: 6,
				}).To_MessageEntity(),
			},
		}).To_Message()
		err = s.MsgFacade.PushUserMessage(ctx, 1, system_uid, r.Uid, rand.Int63(), message)
		if err != nil {
			log.Errorf("PushUserMessage%s", err.Error())
		}
	}
}
