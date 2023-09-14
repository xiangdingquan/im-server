package service

import (
	"github.com/BurntSushi/toml"
	"github.com/go-kratos/kratos/pkg/naming/discovery"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/infra/databus/pkg/queue/databus"
	"open.chat/pkg/log"
)

type Config struct {
	PushIOS   *PushIOS
	PushFCM   *PushFCM
	JPush     *JPush
	Tpns      *PushTpns
	Log       *log.Config
	HTTP      *bm.ServerConfig
	RPC       *warden.ClientConfig
	Databus   *databus.Config
	Discovery *discovery.Config
	Routine   *Routine
}

type PushIOS struct {
	BundID string
}

type PushFCM struct {
	Key       string
	Timeout   int
	ServerURL string
}

type JPush struct {
	AppKey  string
	Secret  string
	Timeout int
}

type PushTpns struct {
	AccessID  string
	SecretKey string
	Timeout   int
}

// Routine routine.
type Routine struct {
	Size uint64
	Chan uint64
}

// Set set config and decode.
func (c *Config) Set(text string) error {
	var tmp Config
	if _, err := toml.Decode(text, &tmp); err != nil {
		return err
	}
	*c = tmp
	return nil
}
