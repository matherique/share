package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/matherique/share/internal/infra/web"
	"github.com/matherique/share/internal/store"
	"github.com/matherique/share/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	slog.Info("database connected")

	makeHandlers(r, client.Database("share"))

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

	ctx, cancel = context.WithCancel(context.Background())

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("fail on shutdown", "err", err)
		panic(err)
	}
}

func makeHandlers(r *chi.Mux, db *mongo.Database) {
	hashesStore := store.NewHashesStore(db)
	snipetStore := store.NewSnipetStore(db)
	importUc := usecase.NewCreateUseCase(hashesStore, snipetStore)
	web.RegisterCreateHandler(r, importUc)
}
