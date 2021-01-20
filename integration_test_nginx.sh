#!/bin/bash
# integration_test_nginx.sh

await_postgres () {
    until docker-compose exec db psql -U postgres -c '\q'; do
        >&2 echo "Postgres is unavailable - sleeping"
        sleep 1
    done

    >&2 echo "Postgres is up"
}

# stop everything
docker-compose down -v

# start
docker-compose up -d


# curl nginx until success

success_cnt=0

for iter in {1..60}
do
    sleep 1

    # request to nginx
    code=$(curl --write-out '%{http_code}' --silent --output /dev/stderr http://localhost:9000/healthcheck)

    if [[ "$code" == "200" ]]; then
        echo 'Success!'

        ((success_cnt++))

        if [[ $success_cnt -ge 2 ]]; then
            exit 0
        fi
    fi
done

echo 'Retries limit exceeded'
exit 1
