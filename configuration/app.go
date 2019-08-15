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
	c := &ApplicationConfiguration{}

	flag.StringVar(&c.HttpHost, "addr", "127.0.0.1", "server ip")
	flag.UintVar(&c.HttpPort, "port", 8080, "server port")

	flag.Parse()

	return c
}
