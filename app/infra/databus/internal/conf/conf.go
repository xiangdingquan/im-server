package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

var (
	// Conf global config variable
	Conf     = &Config{}
	confPath string
	client   *bm.Client
)

// Config databus config struct
type Config struct {
	// base
	Addr     string
	Clusters map[string]*Kafka
	// Log
	Log *log.Config
	// http
	HTTPServer *bm.ServerConfig
	// mysql
	MySQL *sqlx.Config
}

// Kafka contains cluster, brokers, sync.
type Kafka struct {
	Cluster string
	Brokers []string
}

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init int config
func Init() error {
	if confPath == "" {
		panic("")
	}
	return local()
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
