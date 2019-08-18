#!/bin/sh
set -e

: ${HTTP_IP:="0.0.0.0"}
: ${HTTP_PORT:=8080}
: ${PG_HOST:="127.0.0.1"}
: ${PG_PORT:=5432}
: ${PG_USER:=""}
: ${PG_PASS:=""}
: ${PG_DBNAME:="pronuntio"}

if [ "$1" = 'pronuntio-server' ]; then
    exec pronuntio-server \
        -addr=${HTTP_IP} \
	-port=${HTTP_PORT} \
	-pg.host=${PG_HOST} \
	-pg.port=${PG_PORT} \
	-pg.user=${PG_USER} \
	-pg.pass=${PG_PASS} \
	-pg.dbname=${PG_DBNAME}
fi

exec "$@"
