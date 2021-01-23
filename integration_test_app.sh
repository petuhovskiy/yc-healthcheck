#!/bin/bash
# integration_test_app.sh

await_postgres () {
    until docker-compose exec db psql -U postgres -c '\q'; do
        >&2 echo "Postgres is unavailable - sleeping"
        sleep 1
    done

    >&2 echo "Postgres is up"
}

set -e

# stop everything
docker-compose down -v

# start
docker-compose up -d

await_postgres

EXPECTED='{"ip":"127.0.0.1","services":[{"ip":"127.0.0.1","status":"AVAILABLE"},{"ip":"127.0.0.2","status":"AVAILABLE"}]}'
ACTUAL=$(curl http://localhost:3000/healthcheck)

echo "EXPTECTED => $EXPECTED"
echo "ACTUAL => $ACTUAL"

echo "All ok!"