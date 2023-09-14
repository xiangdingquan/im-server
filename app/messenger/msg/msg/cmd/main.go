package main

import (
	"open.chat/app/messenger/msg/internal/msg/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
