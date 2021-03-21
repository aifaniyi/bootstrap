package gotemplate

// main package contents
const (
	Main = `
	package main

import (
	"github.com/go-chi/chi"
	"gitlab.com/aifaniyi/go-libs/logger"
	"gitlab.com/aifaniyi/go-libs/postgres"
)

func main() {
	// load app config
	config := settings.LoadSettings()

	// create db service
	conn, err := postgres.GetPostgresORM()
	if err != nil {
		logger.Error.Fatal("Postgres connection fail:" + err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		logger.Error.Fatal("Postgres connection fail:" + err.Error())
	}
	defer sqlDB.Close()

	// enables auto creation of tables
	db.Migrate(conn, config.IsDev)

	// create dependencies
	dependencies := &settings.Dependencies{
		Db:             newDBService(conn),
	}

	// create router
	router := chi.NewRouter()
	server := newServer(config.APIPort, router)

	// setup route controllers
	web.SetupRouteControllers(router, config, dependencies)

	// start server
	go startServer(config.APIPort, server)

	awaitTermination(server)
}
`
	Server = `
	package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"gitlab.com/aifaniyi/go-libs/logger"
	"gorm.io/gorm"
)

const (
	shutdownWait = 10 // seconds to wait for pending requests to be completed before shut down
)

func newServer(port string, router *chi.Mux) *http.Server {
	return &http.Server{
		Addr: port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
}

func startServer(port string, server *http.Server) {
	logger.Info.Printf("starting server on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		logger.Info.Printf("http server error: %s", err.Error())
	}
}

// wait until shutdown signal is received
// and run graceful shutdown to allow existing
// connections terminate before shutdown
func awaitTermination(server *http.Server) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM,
		syscall.SIGINT)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(),
		shutdownWait*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	logger.Info.Println("shutting down")
}

var newDBService = func(conn *gorm.DB) db.Service {
	return db.NewService(conn)
}
`
)
