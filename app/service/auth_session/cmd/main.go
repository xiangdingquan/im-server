package main

import (
	"open.chat/app/service/auth_session/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
