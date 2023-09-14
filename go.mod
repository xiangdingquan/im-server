module open.chat

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Shopify/sarama v1.23.1
	github.com/beorn7/perks v1.0.1
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/bwmarrin/snowflake v0.3.0
	github.com/coreos/go-systemd v0.0.0-20180511133405-39ca1b05acc7 // indirect
	github.com/coreos/pkg v0.0.0-20160727233714-3ac0863d7acf // indirect
	github.com/disintegration/imaging v1.6.2
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-ini/ini v1.62.0 // indirect
	github.com/go-kratos/kratos v0.5.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.0
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/gorm v1.9.13
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/mvdan/xurls v0.0.0-00010101000000-000000000000
	github.com/nyaruka/phonenumbers v1.0.55
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/oschwald/geoip2-golang v1.4.0
	github.com/panjf2000/gnet v1.2.7
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.9.1
	github.com/prometheus/procfs v0.0.11
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a
	github.com/smartystreets/goconvey v0.0.0-20181108003508-044398e4856c // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1
	github.com/tidwall/gjson v1.6.0
	github.com/vjeantet/grok v1.0.0
	go.etcd.io/etcd v0.0.0-20200402134248-51bdeb39e698
	go.uber.org/atomic v1.6.0
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200420201142-3c4aac89819a
	golang.org/x/net v0.0.0-20200421231249-e086a090c8fd
	google.golang.org/grpc v1.28.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15
)

replace (
	github.com/go-kratos/kratos => ./third_party/github.com/go-kratos/kratos
	github.com/mvdan/xurls => ./third_party/github.com/mvdan/xurls
	github.com/panjf2000/gnet => ./third_party/github.com/panjf2000/gnet
)
