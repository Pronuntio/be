package configuration

import "flag"

type ApplicationConfiguration struct {
	HttpAddr string
}

func ParseArgs() *ApplicationConfiguration {
	c := &ApplicationConfiguration{}

	flag.StringVar(&c.HttpAddr, "addr", "127.0.0.1:8080", "server addr <ip>:<port>")

	flag.Parse()

	return c
}
