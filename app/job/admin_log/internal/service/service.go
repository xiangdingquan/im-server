package service

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"encoding/json"
	"fmt"

	"github.com/gogo/protobuf/proto"

	"open.chat/app/infra/databus/pkg/queue/databus"
	"open.chat/app/job/admin_log/adminlogpb"
	"open.chat/app/job/admin_log/internal/dao"
	"open.chat/app/pkg/databus_util"
	_ "open.chat/app/service/biz_service/message/facade"
	"open.chat/pkg/log"
	"runtime/debug"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Service struct {
	conf *Config

	*dao.Dao
	databus        *databus.Databus
	databusHandler *databus_util.DatabusHandler
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
		Dao:            dao.New(),
		databus:        databus.New(ac.Databus),
		databusHandler: databus_util.NewDatabusHandler(),
	}

	s.databusHandler.GoWatch(s.databus, func(msg *databus.Message) error {
		log.Debugf("recv {key: %s, value: %s", msg.Key, string(msg.Value))
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("handle panic: %s", debug.Stack())
			}
		}()

		switch msg.Key {
		case proto.MessageName((*adminlogpb.ChannelAdminLogEventData)(nil)):
			r := new(adminlogpb.ChannelAdminLogEventData)
			if err := json.Unmarshal(msg.Value, r); err != nil {
				log.Error("unmarshal adminlogpb.ChannelAdminLogEventData error: %v", err.Error())
				return err
			}
			return s.onChannelAdminEventAction(context.Background(), r.LogUserId, r.ChannelId, r.Event)
		default:
			err := fmt.Errorf("invalid key: %s", msg.Key)
			log.Error(err.Error())
			return err
		}
	})

	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return nil
}

// Close close the resources.
func (s *Service) Close() error {
	s.databusHandler.Close()
	return nil
}

// Close close the resources.
func (s *Service) RunLoop() {
}
