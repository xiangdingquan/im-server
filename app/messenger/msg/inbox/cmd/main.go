package main

import (
	"open.chat/app/messenger/msg/internal/inbox/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
