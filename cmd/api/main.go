package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/matherique/share/internal/app"
)

func main() {
	server := &http.Server{Addr: ":8080"}

	app.Register()

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			slog.Error("fail on listen and serve", "err", err)
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, os.Interrupt)
	<-c

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("fail on shutdown", "err", err)
		panic(err)
	}
}
