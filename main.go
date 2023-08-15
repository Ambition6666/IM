package main

import (
	"im/config"
	"im/internal/http/routes"
	websokcet "im/internal/service/websocket"
	"im/sql"
)

func init() {
	config.ConfigInit()
	websokcet.Manager = *websokcet.NewClientManager()
	sql.InitSql()
}
func main() {
	go websokcet.Manager.Start()
	sql.RForm()
	r := routes.InitRoute()
	r.Run(":" + config.ServerPort)
}
