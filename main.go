package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/pronuntio/core/configuration"
	"github.com/pronuntio/core/domain/user"
	"github.com/pronuntio/core/domain/word"
	"github.com/pronuntio/core/infrastructure/databases"
	"github.com/pronuntio/core/server"
	us "github.com/pronuntio/core/services/user"
	ws "github.com/pronuntio/core/services/word"
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

	pgConn, err := databases.NewPostgresConnection(
		appConf.PgConfig.Host,
		appConf.PgConfig.Port,
		appConf.PgConfig.DBName,
		appConf.PgConfig.Username,
		appConf.PgConfig.Password,
	)
	if err != nil {
		fmt.Printf("unable to create pg db connection: %s\r\n", err.Error())
		os.Exit(-1)
	}

	pgUserDao := user.NewPostgresUserDao(pgConn, l.Named("pg_user_dao"))
	pgWordDao := word.NewPostgresWordDao(pgConn, l.Named("pg_word_dao"))

	userService := us.NewUserService(pgUserDao, l.Named("user_service"))
	wordService := ws.NewWordService(pgWordDao, l.Named("word_service"))

	hRouter := mux.NewRouter()
	userService.GetRoutes(hRouter)
	wordService.GetRoutes(hRouter)
	hServer := server.NewHTTPServer(
		appConf.HttpHost,
		appConf.HttpPort,
		hRouter,
		l,
	)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGTERM)

	hServer.Start(appContext)

	for {
		select {
		case <-sigCh:
			cancelF()
		}
	}
}
