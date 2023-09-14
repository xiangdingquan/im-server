package main

import (
	"context"
	"flag"

	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	// 解析flag
	flag.Parse()
	// 初始化日志模块
	log.Init(nil)
	defer log.Close()

	log.SetFormat("[%D %T] [%L] [%S] %M")
	// jsonCodec :=
	// 打印日志
	log.V(5).Info("hi:%s", "kratos")
	log.Infoc(context.TODO(), "hi:%s", "kratos")
	log.Infov(context.TODO(), log.KV2("key1", 100), log.KV2("key2", "test value"))
	log.Warn("hi:%s", "kratos")
	log.Warnc(context.TODO(), "hi:%s", "kratos")
	log.Warnv(context.TODO(), log.KV2("key1", 100), log.KV2("key2", "test value"))
	log.Error("hi:%s", "kratos")
	log.Errorc(context.TODO(), "hi:%s", "kratos")
	log.Errorv(context.TODO(), log.KV2("key1", 100), log.KV2("key2", "test value"))
}
