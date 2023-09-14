package service

import (
	"encoding/json"
	"io/ioutil"

	"open.chat/app/messenger/biz_server/help/internal/dao"

	"open.chat/mtproto"
)

const (
	configFile     = "./config.json"
	expiresTimeout = 3600 // 超时时间设置为3600秒
)

var config mtproto.TLConfig

func init() {
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(configData), &config)
	if err != nil {
		panic(err)
	}
}

type Service struct {
	*dao.Dao
}

func New() *Service {
	return &Service{
		Dao: dao.New(),
	}
}

// Close close the resource.
func (s *Service) Close() {
	// s.Dao.Close()
}
