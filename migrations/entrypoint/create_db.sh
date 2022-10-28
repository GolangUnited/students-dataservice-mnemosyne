#!/bin/sh

set -xue

wait_for_db()
{
    for i in $(seq 1 30); do
        echo "SELECT 1" | psql -h $POSTGRES_HOST -U $POSTGRES_USER && return
        sleep 1
    done

    exit 1
}

wait_for_db

echo "SELECT 'CREATE DATABASE $POSTGRES_DB_NAME' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '$POSTGRES_DB_NAME')\gexec" | psql -h $POSTGRES_HOST -U $POSTGRES_USER -v ON_ERROR_STOP=1
