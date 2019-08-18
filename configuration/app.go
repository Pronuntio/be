package configuration

import "flag"

type ApplicationConfiguration struct {
	HttpHost string
	HttpPort uint
	PgConfig *PostgresConfiguration
}

func NewApplicationConfiguration() *ApplicationConfiguration {
	return &ApplicationConfiguration{
		HttpHost: "",
		HttpPort: 0,
		PgConfig: &PostgresConfiguration{},
	}
}

func ParseArgs() *ApplicationConfiguration {
	c := NewApplicationConfiguration()

	flag.StringVar(&c.HttpHost, "addr", "127.0.0.1", "server ip")
	flag.UintVar(&c.HttpPort, "port", 8080, "server port")
	flag.StringVar(&c.PgConfig.Host, "pg.host", "127.0.0.1", "postgres host")
	flag.UintVar(&c.PgConfig.Port, "pg.port", 5432, "postgres port")
	flag.StringVar(&c.PgConfig.Username, "pg.user", "", "pg user")
	flag.StringVar(&c.PgConfig.Password, "pg.pass", "", "pg password")
	flag.StringVar(&c.PgConfig.DBName, "pg.dbname", "", "pg db name")
	flag.Parse()

	return c
}
