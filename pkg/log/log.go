package log

import (
	"github.com/go-kratos/kratos/pkg/log"
)

type Config = log.Config

// Init create logger with context.
func Init(conf *log.Config) {
	log.Init(conf)
}

// Close close resource.
func Close() (err error) {
	return log.Close()
}

var (
	Debug  = log.Info
	Debugf = log.Info

	Info  = log.Info
	Infof = log.Info
	Infov = log.Infov

	Warn  = log.Warn
	Warnf = log.Warn

	Error  = log.Error
	Errorf = log.Error

	KV = log.KV

	SetFormat = log.SetFormat
)
