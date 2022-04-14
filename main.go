package main

import (
	"github.com/JSainsburyPLC/third-party-token-server/app"
	"github.com/JSainsburyPLC/third-party-token-server/config"
)


func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3001")
}