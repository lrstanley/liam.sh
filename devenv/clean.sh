#!/bin/bash
#shellcheck disable=SC2155

if [ ! -d ".git" ]; then
	echo >&2 "error: run $0 from the root of the project"
	exit 1
fi

set -x
docker-compose \
	--project-name "$(basename "$PWD")" \
	--file ./devenv/docker-compose.yaml \
	down --volumes --remove-orphans --rmi local --timeout 1
