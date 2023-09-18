package main

import (
	"notes-api-server/internal/app"
	"notes-api-server/server"
)

func main() {
	app := app.NewApp()
	server.Init(app)
}
