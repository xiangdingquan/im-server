package service

import (
	"github.com/BurntSushi/toml"
	"open.chat/app/infra/databus/pkg/queue/databus"
)

type Config struct {
	Databus *databus.Config
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
