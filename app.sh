#!/bin/sh

DB_STRING="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB_NAME?sslmode=$POSTGRES_SSL"

/app/migrate -path /app/migrations -database "$DB_STRING" up
/app/mnemosyne