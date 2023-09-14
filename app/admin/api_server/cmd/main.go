package main

import (
	"open.chat/app/admin/api_server/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
