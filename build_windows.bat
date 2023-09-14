@echo off
echo 正在设置运行环境为linux...
go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
echo 设置完成
echo 编译开始........
echo 正在编译databus,请勿关闭此窗口!!!
go build -o .\bin_windows\databus.exe  .\app\infra\databus\cmd
echo 编译成功
echo 所有编译好的文件存放在bin目录
echo 正在编译剩下的文件,编译完成后将自动退出,请勿强行关闭窗口!
start /b go build -o .\bin_windows\media.exe  .\app\service\media\cmd && go build -o .\bin_windows\auth_session.exe  .\app\service\auth_session\cmd && go build -o .\bin_windows\inbox.exe  .\app\messenger\msg\inbox\cmd
start /b go build -o .\bin_windows\sync.exe  .\app\messenger\sync\cmd && go build -o .\bin_windows\push.exe  .\app\messenger\push\cmd && go build -o .\bin_windows\webpage.exe  .\app\messenger\webpage\cmd
start /b go build -o .\bin_windows\gif.exe  .\app\bots\gif\cmd && go build -o .\bin_windows\admin_log.exe  .\app\job\admin_log\cmd && go build -o .\bin_windows\scheduled.exe  .\app\job\scheduled\cmd
start /b go build -o .\bin_windows\session.exe  .\app\interface\session\cmd && go build -o .\bin_windows\botway.exe  .\app\interface\botway\cmd && go build -o .\bin_windows\api_server.exe  .\app\admin\api_server\cmd
start /b go build -o .\bin_windows\gateway.exe  .\app\interface\gateway\cmd && go build -o .\bin_windows\botfather.exe  .\app\bots\botfather\cmd && go build -o .\bin_windows\msg.exe  .\app\messenger\msg\msg\cmd
go build -o .\bin_windows\relay.exe  .\app\interface\relay\cmd
go build -o .\bin_windows\wsserver.exe  .\app\interface\wsserver\cmd
go build -o .\bin_windows\biz_server.exe  .\app\messenger\biz_server\cmd

echo 正在还原运行环境
go env -w CGO_ENABLED=1
go env -w GOOS=windows
go env -w GOARCH=amd64
echo 设置完成