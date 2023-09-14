#!/bin/bash
#Array Declaration
arr=(
databus
media
sync
gif
session
gateway
relay
wsserver
biz_server
auth_session
push
admin_log
botway
botfather
inbox
webpage
scheduled
api_server
msg
)

for i in "${arr[@]}"
do
mkdir -p logs/$i
done
