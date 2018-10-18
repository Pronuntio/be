#!/usr/bin/env bash
set -e

createdb pr_main
psql -d pr_main -f /tmp/init.sql -v ON_ERROR_STOP=1
