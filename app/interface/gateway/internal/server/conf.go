package server

import (
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/pkg/net2"
)

var (
	Conf = &Config{}
)

type Config struct {
	ServerId       int32
	MaxProc        int
	KeyFile        string
	KeyFingerprint string
	Server         *net2.TcpServerConfig
	WardenClient   *warden.ClientConfig
}
