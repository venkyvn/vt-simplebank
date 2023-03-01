#!/bin/sh

# scrip exits immediately if error occurred
set -e

# migrate db first
#migrate -path db/migration -database "postgresql://root:123@localhost:5432/simple_bank?sslmode=disable" -verbose up
echo "run db migration"
source /app/app.env
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up


# take all parameter and run it
echo "start the app"
exec "$@"
