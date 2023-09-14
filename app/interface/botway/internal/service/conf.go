package service

import (
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
)

var (
	Conf = &Config{}
)

type Config struct {
	ServerId     int32
	MaxProc      int
	WardenClient *warden.ClientConfig
}
