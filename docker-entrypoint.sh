#!/bin/sh
set -e

: ${ADDR:="0.0.0.0:8080"}

if [ "$1" = 'pronuntio-server' ]; then
    exec pronuntio-server \
        -addr=${ADDR}
fi

exec "$@"