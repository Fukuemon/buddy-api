package main

// @title       buddy-api
// @version     1.0
// @description BuddyのAPIサーバー

// @host      localhost:8080
// @BasePath  /v1

import (
	"api-buddy/config"
	_ "api-buddy/docs"
	"api-buddy/server"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfig()

	server.Run(ctx, conf)
}
