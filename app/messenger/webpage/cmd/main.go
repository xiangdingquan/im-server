package main

import (
	"open.chat/app/messenger/webpage/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
