#!/bin/bash
#shellcheck disable=SC2155

if [ ! -d ".git" ]; then
	echo >&2 "error: run $0 from the root of the project"
	exit 1
fi

export BASE="$(basename "$PWD")"

DANGLING=$(docker images -f "dangling=true" -q)
if [ ! -z "$DANGLING" ]; then
	docker rmi "${DANGLING[*]}" 2> /dev/null
fi

export USER=$(id -u)
export GROUP=$(id -g)
export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

# make sure devimage has been built.
if [ "$(docker image ls devimage | wc -l)" -lt 2 ]; then
	docker compose \
		--project-name "$BASE" \
		--file ./devenv/docker-compose.yaml \
		build devimage
fi

docker compose \
	--project-name "$BASE" \
	--file ./devenv/docker-compose.yaml \
	up \
	--remove-orphans \
	--build \
	--abort-on-container-exit \
	--timeout 0 "$@"
