package main

import (
	"api-buddy/config"
	"api-buddy/server"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfig()

	server.Run(ctx, conf)
}
