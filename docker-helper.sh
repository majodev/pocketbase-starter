#!/bin/bash
set -Eeo pipefail

# Get first service name from docker-compose.yml
get_first_service() {
    local first_service=$(docker compose config --services | head -n1)
    if [ -z "$first_service" ]; then
        echo "Error: No services found in docker-compose.yml" >&2
        exit 1
    fi
    echo "$first_service"
}

SERVICE_NAME=$(get_first_service)

if [ "$1" = "--up" ]; then
    docker compose up --no-start
    docker compose start # ensure we are started, handle also allowed to be consumed by vscode
    docker compose exec $SERVICE_NAME bash
fi

if [ "$1" = "--halt" ]; then
    docker compose stop
fi

if [ "$1" = "--rebuild" ]; then
    docker compose up -d --force-recreate --no-deps --build $SERVICE_NAME
fi

if [ "$1" = "--destroy" ]; then
    docker compose down --rmi local -v --remove-orphans
fi

[ -n "$1" -a \( "$1" = "--up" -o "$1" = "--halt" -o "$1" = "--rebuild" -o "$1" = "--destroy" \) ] \
    || { echo "usage: $0 --up | --halt | --rebuild | --destroy" >&2; exit 1; }