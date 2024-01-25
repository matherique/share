package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/matherique/share/internal/usecase"
	"github.com/matherique/share/internal/web"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	makeHandlers(r)

	server := &http.Server{Addr: ":8080", Handler: r}

	go func() {
		slog.Info("server start", "port", ":8080")
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

func makeHandlers(r *chi.Mux) {
	importUc := usecase.NewImportUseCase()
	web.RegisterImportHandler(r, importUc)
}
