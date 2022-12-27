#!/bin/bash -eu

NAMESPACE="liam-sh"
SERVICE="liam-sh-postgres-ro"
SECRET="liam-sh-postgres-app"
LOCAL_PORT="5433"
REMOTE_PORT="5432"

# get credentials used for the local app
source .env
: "$POSTGRES_DB"
: "$POSTGRES_USER"
: "$POSTGRES_PASSWORD"

function port_forward {
	kubectl -n "$NAMESPACE" port-forward "service/${SERVICE}" --address 0.0.0.0 "${LOCAL_PORT}:${REMOTE_PORT}"
}

function get_remote_credentials {
	export REMOTE_USERNAME=$(kubectl -n "$NAMESPACE" get secret "$SECRET" --template '{{ .data.username }}' | base64 --decode)
	export REMOTE_PASSWORD=$(kubectl -n "$NAMESPACE" get secret "$SECRET" --template '{{ .data.password }}' | base64 --decode)
}

trap "sleep 1;kill 0" EXIT

get_remote_credentials
port_forward &

until nc -z localhost "$LOCAL_PORT"; do
	sleep 1
done

docker run -it --rm --network host --add-host=host.docker.internal:host-gateway postgres:14 /bin/bash -c "
    PGPASSWORD='$REMOTE_PASSWORD' pg_dump \
        --no-owner \
        --format='p' \
        --clean \
        --host='host.docker.internal' \
        --port='$LOCAL_PORT' \
        --username='$REMOTE_USERNAME' \
        --dbname='app' \
    | PGPASSWORD='$POSTGRES_PASSWORD' psql \
        --host='host.docker.internal' \
        --port='5432' \
        --username='$POSTGRES_USER' \
        --dbname='$POSTGRES_DB'
"

echo "successfully mirrored database"
