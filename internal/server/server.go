package server

import (
	"api-buddy/config"
	"api-buddy/presentation/settings"
	"api-buddy/server/route"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
)

func Run(ctx context.Context, conf *config.Config) {
	// validatorの初期化
	settings.InitValidationSettings()

	api := settings.NewGinEngine()
	route.InitRoute(api)

	// CORSの設定
	api.Use(cors.New(cors.Config{
		AllowOrigins:     conf.Server.AllowOrigins,
		AllowMethods:     conf.Server.AllowMethods,
		AllowHeaders:     conf.Server.AllowHeaders,
		AllowCredentials: conf.Server.AllowCredentials,
		MaxAge:           conf.Server.MaxAge,
	}))

	address := conf.Server.Address + ":" + conf.Server.Port
	log.Printf("Starting server on %s...\n", address)

	srv := &http.Server{
		Addr:              address,
		Handler:           api,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		os.Exit(1)
	}
}
