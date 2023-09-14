package main

import (
	"open.chat/app/service/media/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
