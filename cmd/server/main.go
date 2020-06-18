package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/pronuntio/core/configuration"
	db "github.com/pronuntio/core/pkg/infra/db"
	server "github.com/pronuntio/core/pkg/infra/httpserver"
	service "github.com/pronuntio/core/pkg/service"
	"github.com/pronuntio/core/pkg/user"
	"github.com/pronuntio/core/pkg/word"
	"github.com/pronuntio/core/version"
	"go.uber.org/zap"
)

func main() {
	appConf := configuration.ParseArgs()
	appContext, cancelF := context.WithCancel(context.Background())

	l, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("unable to instantiate logger: %s\r\n", err.Error())
		os.Exit(-1)
	}

	if appConf.Version {
		fmt.Println("version:", version.Revision)
		fmt.Println("build time:", version.BuildTime)
		os.Exit(0)
	}

	l.Info("stated with configuration", zap.Any("config", appConf))

	pgConn, err := db.NewPostgresConnection(
		appConf.PgConfig.Host,
		appConf.PgConfig.Port,
		appConf.PgConfig.DBName,
		appConf.PgConfig.Username,
		appConf.PgConfig.Password,
	)
	if err != nil {
		l.Error("unable to create pg db connection", zap.Error(err))
		os.Exit(-1)
	}

	pgUserDao := user.NewPostgresUserDao(pgConn, l.Named("pg_user_dao"))
	pgWordDao := word.NewPostgresWordDao(pgConn, l.Named("pg_word_dao"))

	userService := service.NewUserService(pgUserDao, l.Named("user_service"))
	wordService := service.NewWordService(pgWordDao, l.Named("word_service"))

	hRouter := mux.NewRouter()
	userService.GetRoutes(hRouter)
	wordService.GetRoutes(hRouter)
	hServer := server.NewHTTPServer(
		appConf.HttpHost,
		appConf.HttpPort,
		hRouter,
		l.Named("http-server"),
	)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGTERM)

	l.Info("starting server...")
	hServer.Start(appContext)
	l.Info("server started")
	for {
		select {
		case <-sigCh:
			cancelF()
		}
	}
}
