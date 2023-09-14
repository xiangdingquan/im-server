package main

import (
	"open.chat/app/bots/gif/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
