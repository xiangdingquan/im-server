@echo off
echo �����������л���Ϊlinux...
go env -w CGO_ENABLED=0
go env -w GOOS=windows
go env -w GOARCH=amd64
echo �������
echo ���뿪ʼ........
echo ���ڱ���databus,����رմ˴���!!!
go build -o .\bin_windows\databus.exe  .\app\infra\databus\cmd
echo ����ɹ�
echo ���б���õ��ļ������binĿ¼
echo ���ڱ���ʣ�µ��ļ�,������ɺ��Զ��˳�,����ǿ�йرմ���!
start /b go build -o .\bin_windows\media.exe  .\app\service\media\cmd && go build -o .\bin_windows\auth_session.exe  .\app\service\auth_session\cmd && go build -o .\bin_windows\inbox.exe  .\app\messenger\msg\inbox\cmd
start /b go build -o .\bin_windows\sync.exe  .\app\messenger\sync\cmd && go build -o .\bin_windows\push.exe  .\app\messenger\push\cmd && go build -o .\bin_windows\webpage.exe  .\app\messenger\webpage\cmd
start /b go build -o .\bin_windows\gif.exe  .\app\bots\gif\cmd && go build -o .\bin_windows\admin_log.exe  .\app\job\admin_log\cmd && go build -o .\bin_windows\scheduled.exe  .\app\job\scheduled\cmd
start /b go build -o .\bin_windows\session.exe  .\app\interface\session\cmd && go build -o .\bin_windows\botway.exe  .\app\interface\botway\cmd && go build -o .\bin_windows\api_server.exe  .\app\admin\api_server\cmd
start /b go build -o .\bin_windows\gateway.exe  .\app\interface\gateway\cmd && go build -o .\bin_windows\botfather.exe  .\app\bots\botfather\cmd && go build -o .\bin_windows\msg.exe  .\app\messenger\msg\msg\cmd
go build -o .\bin_windows\relay.exe  .\app\interface\relay\cmd
go build -o .\bin_windows\wsserver.exe  .\app\interface\wsserver\cmd
go build -o .\bin_windows\biz_server.exe  .\app\messenger\biz_server\cmd

echo ���ڻ�ԭ���л���
go env -w CGO_ENABLED=1
go env -w GOOS=windows
go env -w GOARCH=amd64
echo �������