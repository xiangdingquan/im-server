package service

import (
	"github.com/BurntSushi/toml"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/pkg/log"
)

type Config struct {
	Log  *log.Config
	HTTP *bm.ServerConfig
	RPC  *warden.ClientConfig
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
