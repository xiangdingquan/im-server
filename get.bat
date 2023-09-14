@echo off

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64


go get github.com/go-kratos/kratos/pkg/conf/dsn@v0.5.0
go get github.com/go-kratos/kratos/pkg/net/http/blademaster@v0.5.0

go get github.com/go-kratos/kratos/pkg/ecode@v0.5.0
go get github.com/go-kratos/kratos/pkg/ecode/types@v0.5.0

go get github.com/go-kratos/kratos/pkg/stat/sys/cpu@v0.5.0
go get github.com/go-kratos/kratos/pkg/log@v0.5.0

go get github.com/go-kratos/kratos/pkg/net/trace@v0.5.0
go get github.com/go-kratos/kratos/pkg/naming@v0.5.0
go get github.com/go-kratos/kratos/pkg/naming/etcd@v1.0.1
go get github.com/go-kratos/kratos/pkg/conf/paladin@v0.5.0

go get github.com/go-kratos/kratos/pkg/net/trace/jaeger
go get github.com/go-kratos/kratos/pkg/net/trace/zipkin@v0.5.0

go get golang.org/x/net/idna@v0.0.0-20200421231249-e086a090c8fd
go get google.golang.org/grpc/internal/transport@v1.28.1

go get open.chat/app/infra/databus/internal/server/tcp
go get open.chat/app/infra/databus/pkg/queue/databus
go get open.chat/app/infra/databus/pkg/stat/prom
go get open.chat/pkg/database/sqlx
go get open.chat/app/service/dfs/internal/dao
go get open.chat/app/service/idgen/facade/snowflake
go get open.chat/app/service/dfs/internal/cachefs
go get open.chat/pkg/phonenumber
go get open.chat/pkg/grpc_util/server
go get open.chat/app/service/auth_session/internal/dao
go get open.chat/app/interface/wsserver/internal/service

go env -w CGO_ENABLED=1
go env -w GOOS=windows
go env -w GOARCH=amd64

