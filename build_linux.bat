@echo off

echo set build environment linux

go env -w CGO_ENABLED=0

go env -w GOOS=linux

go env -w GOARCH=amd64

echo build databus

go build -o .\bin\databus  .\app\infra\databus\cmd

echo build media

go build -o .\bin\media  .\app\service\media\cmd

echo build auth_session

go build -o .\bin\auth_session .\app\service\auth_session\cmd

echo build inbox

go build -o .\bin\inbox  .\app\messenger\msg\inbox\cmd

echo build sync

go build -o .\bin\sync  .\app\messenger\sync\cmd

echo build push

go build -o .\bin\push  .\app\messenger\push\cmd

echo build webpage

go build -o .\bin\webpage  .\app\messenger\webpage\cmd

echo build gif

go build -o .\bin\gif  .\app\bots\gif\cmd

echo build admin_log

go build -o .\bin\admin_log  .\app\job\admin_log\cmd

echo build scheduled

go build -o .\bin\scheduled  .\app\job\scheduled\cmd 

echo build session

go build -o .\bin\session  .\app\interface\session\cmd

echo build botway

go build -o .\bin\botway  .\app\interface\botway\cmd

echo build api_server

go build -o .\bin\api_server  .\app\admin\api_server\cmd 

echo build gateway

go build -o .\bin\gateway  .\app\interface\gateway\cmd

echo build botfather

go build -o .\bin\botfather  .\app\bots\botfather\cmd

echo build msg

go build -o .\bin\msg  .\app\messenger\msg\msg\cmd

echo build relay

go build -o .\bin\relay  .\app\interface\relay\cmd

echo build wsserver

go build -o .\bin\wsserver  .\app\interface\wsserver\cmd

echo build biz_server

go build -o .\bin\biz_server  .\app\messenger\biz_server\cmd

echo set build environment windows

go env -w CGO_ENABLED=0

go env -w GOOS=windows

go env -w GOARCH=amd64

echo finish

pause