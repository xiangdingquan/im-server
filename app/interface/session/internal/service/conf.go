package service

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
)

type Config struct {
	WardenClient *warden.ClientConfig
	Routine      *Routine
	HTTP         *bm.ServerConfig
}

// Routine routine.
type Routine struct {
	Size uint64
	Chan uint64
}
