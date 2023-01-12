package cmd

import (
	"context"
	"fmt"
	"github.com/getground/tech-tasks/backend/boot"
	"github.com/getground/tech-tasks/backend/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func API() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "start get ground party service in api mode",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Starting get ground service api")
			runAPI()
		},
	}
}

func runAPI() {
	cfg, err := config.NewAPI()
	if err != nil {
		log.Fatalln(err)
	}

	engine := boot.API(cfg)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler: engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down the server")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("server is forced to shutdown")
	}
}
