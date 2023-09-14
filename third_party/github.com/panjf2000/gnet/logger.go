package gnet

import (
	"github.com/panjf2000/gnet/internal/logging"
)

// Logger logging.Logger
type Logger = logging.Logger

// InitLogger
var (
	InitLogger = logging.InitLogger
	Cleanup    = logging.Cleanup
)
