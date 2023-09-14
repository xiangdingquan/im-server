package main

import (
	"open.chat/app/messenger/sync/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
