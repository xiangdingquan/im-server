package commands

import (
	"errors"
	"flag"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/trace/jaeger"
	"github.com/go-kratos/kratos/pkg/net/trace/zipkin"

	"open.chat/pkg/log"
)

// //////////////////////////////////////////////////////////////
var (
	GMainInst MainInstance
	GSignal   chan os.Signal

	enableZipkin bool
	enableJaeger bool
)

func init() {
	flag.BoolVar(&enableZipkin, "zipkin", false, "zipkin")
	flag.BoolVar(&enableJaeger, "jaeger", false, "jaeger")
}

type MainInstance interface {
	Initialize() error
	RunLoop()
	Destroy()
}

func Run(inst MainInstance) {
	flag.Parse()
	if err := paladin.Init(); err != nil {
		panic(err)
	}
	log.Init(nil)
	defer log.Close()

	if inst == nil {
		panic(errors.New("inst is nil, exit"))
	}

	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	if enableZipkin {
		zipkin.Init(&zipkin.Config{
			Endpoint: "http://localhost:9411/api/v2/spans",
		})
	} else if enableJaeger {
		jaeger.Init()
	}

	log.SetFormat("[%D %T] [%L] [%S] %M")
	log.Info("instance initialize...")
	err := inst.Initialize()
	log.Info("inited")
	if err != nil {
		panic(err)
	}

	// global
	GMainInst = inst

	log.Info("instance run_loop...")
	go inst.RunLoop()

	GSignal = make(chan os.Signal, 1)
	signal.Notify(GSignal, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-GSignal
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Infof("instance exit...")
			inst.Destroy()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
