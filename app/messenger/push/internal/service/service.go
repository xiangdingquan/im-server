package service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/gogo/protobuf/proto"
	"open.chat/app/messenger/push/internal/dao/apns2"
	"open.chat/app/messenger/push/internal/dao/tpns"

	"open.chat/app/infra/databus/pkg/queue/databus"
	"open.chat/app/messenger/push/internal/dao"
	"open.chat/app/messenger/push/internal/dao/fcm"
	"open.chat/app/messenger/push/internal/dao/jpush"
	"open.chat/app/messenger/push/pushpb"
	"open.chat/pkg/log"

	"open.chat/app/messenger/push/internal/service/iospushcert"
)

const (
	apiKey = "AAAA"
)

type Service struct {
	conf        *Config
	consumer    *databus.Databus
	dao         *dao.Dao
	fcmClient   *fcm.Client
	apnsClient  *apns2.Client
	jpushClient *jpush.Client
	tpnsClient  *tpns.Client
}

func New() *Service {
	var (
		err error
		ac  = &Config{}
		s   = new(Service)
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s.conf = ac
	s.consumer = databus.New(ac.Databus)
	s.dao = dao.New()
	if s.conf.PushFCM != nil {
		if s.conf.PushFCM.ServerURL == "" {
			s.conf.PushFCM.ServerURL = fcm.ServerURL
		}
		s.fcmClient = fcm.NewClient(s.conf.PushFCM.Key,
			time.Duration(s.conf.PushFCM.Timeout)*time.Second,
			s.conf.PushFCM.ServerURL)
	}

	s.jpushClient = jpush.NewClient(ac.JPush.AppKey, ac.JPush.Secret, time.Duration(ac.JPush.Timeout)*time.Second)
	iosBundID := ac.PushIOS.BundID
	var certPEMBlock []byte
	var keyPEMBlock []byte

	switch iosBundID {
	default:
		iosBundID = "bh.tg.ios"
		certPEMBlock = iospushcert.CertPEMBlock_bh_tg_ios
		keyPEMBlock = iospushcert.KeyPEMBlock_bh_tg_ios
	}
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		log.Errorf("onPushAPNS error: %v", err)
		panic(err)
	}

	s.apnsClient = apns2.NewClient(iosBundID, cert, 10*time.Second).Production() //Development()

	go s.consume()
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resources.
func (s *Service) Close() error {
	if err := s.consumer.Close(); err != nil {
		return err
	}
	s.dao.Close()
	return nil
}

func (s *Service) consume() {
	msgs := s.consumer.Messages()
	for {
		msg, ok := <-msgs
		if !ok {
			log.Warn("[job] consumer has been closed")
			return
		}
		if msg.Topic != s.conf.Databus.Topic {
			log.Error("unknown message:%v", msg)
			continue
		}

		s.onPushData(context.Background(), msg.Key, msg.Value)

		msg.Commit()
	}
}

func (s *Service) onPushData(ctx context.Context, key string, value []byte) error {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("handle panic: %s", debug.Stack())
		}
	}()

	switch key {
	case proto.MessageName((*pushpb.PushUpdatesIfNot)(nil)):
		r := new(pushpb.PushUpdatesIfNot)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.onPushUpdatesIfNot(ctx, r)
	case proto.MessageName((*pushpb.PushUpdates)(nil)):
		r := new(pushpb.PushUpdates)
		if err := json.Unmarshal(value, r); err != nil {
			log.Error(err.Error())
			return err
		}
		return s.onPushUpdates(ctx, r)
	default:
		err := fmt.Errorf("invalid recv {key: %s, value: %s}", key, string(value))
		log.Error(err.Error())
		return err
	}
}
