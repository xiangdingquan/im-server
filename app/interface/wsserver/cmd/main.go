package main

import (
	"open.chat/app/interface/wsserver/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
