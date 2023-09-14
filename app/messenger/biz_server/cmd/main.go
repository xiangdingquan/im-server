package main

import (
	"open.chat/app/messenger/biz_server/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
