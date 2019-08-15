package main

import (
	"fmt"
	"os"

	"github.com/pronuntio/core/configuration"
	"github.com/pronuntio/core/domain/user"
	"github.com/pronuntio/core/domain/word"
	"github.com/pronuntio/core/infrastructure/databases"
	"go.uber.org/zap"
)

func main() {
	appConf := configuration.ParseArgs()

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
	pgWorkDao := word.NewPostgresWordDao(pgConn, l.Named("pg_word_dao"))
}
