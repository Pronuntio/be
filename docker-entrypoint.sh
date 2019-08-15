#!/bin/sh
set -e

: ${HTTP_IP:="0.0.0.0"}
: ${HTTP_PORT:=8080}

if [ "$1" = 'pronuntio-server' ]; then
    exec pronuntio-server \
        -addr=${HTTP_IP} \
	-port=${HTTP_PORT}
fi

exec "$@"
