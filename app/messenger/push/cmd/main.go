package main

import (
	"open.chat/app/messenger/push/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
