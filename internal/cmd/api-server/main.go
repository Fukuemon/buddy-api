package main

// @title       Buddy-API
// @version     1.0
// @description BuddyのAPIサーバー
// @host http://localhost:8080
// @BasePath  /v1

import (
	"api-buddy/config"
	_ "api-buddy/docs"
	"api-buddy/infrastructure/mysql/db"
	"api-buddy/server"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfig()

	db.DBOpen(conf.DBConfig)
	defer db.DBClose()

	server.Run(ctx, conf)
}
