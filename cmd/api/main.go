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
	"github.com/matherique/share/internal/infra/repository"
	"github.com/matherique/share/internal/infra/web"
	"github.com/matherique/share/internal/usecase"
	"github.com/matherique/share/pkg/secure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectNoSqlDatabase(database, url string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	return client.Database(database)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	slog.Info("database connected")

	db := connectNoSqlDatabase("share", "mongodb://root:123@mongo:27017")

	makeHandlers(r, db)

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

	if err := db.Client().Disconnect(ctx); err != nil {
		slog.Error("fail on shutdown database", "err", err)
		panic(err)
	}

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("fail on shutdown server", "err", err)
		panic(err)
	}
}

func makeHandlers(r *chi.Mux, db *mongo.Database) {
	// secure pkg
	sec := secure.NewAESSecure()

	// repository
	hashesStore := repository.NewHashesRepositoryMongo(db)
	snipetStore := repository.NewSnipetRepositoryMongo(db)

	// use case
	generateHash := usecase.NewGenerateHashUseCase(hashesStore)
	getSecureSnipetUsecase := usecase.NewGetSecureSnipetUseCase(snipetStore, sec)
	createSecureUserCase := usecase.NewCreateSecureUseCase(snipetStore, generateHash, sec)
	createUseCase := usecase.NewCreateUseCase(snipetStore, generateHash)
	getSnipetUseCase := usecase.NewGetSnipetUseCase(snipetStore)

	// handlers
	getSecureSnipetHandler := web.NewGetSecureSnipetHandler(getSecureSnipetUsecase)
	createSecureHandler := web.RegisterCreateSecureHandler(createSecureUserCase)
	createHandler := web.RegisterCreateHandler(createUseCase)
	getSnipetHandler := web.NewGetSnipetHandler(getSnipetUseCase)

	// routers
	r.Get("/secure/{hash}", getSecureSnipetHandler.Do)
	r.Get("/{hash}", getSnipetHandler.Do)
	r.Post("/{hash}", createHandler.Do)
	r.Post("/secure/{hash}", createSecureHandler.Do)
}
