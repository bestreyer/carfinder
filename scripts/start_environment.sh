#!/usr/bin/env bash
set -e

DOCKER_COMPOSE_COMMAND=$1

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
pushd ${DIR}/..
${DOCKER_COMPOSE_COMMAND} build
${DOCKER_COMPOSE_COMMAND} up -d

# Wait docker-compose up
NEXT_WAIT_TIME=0
while [[ ! $(${DOCKER_COMPOSE_COMMAND} ps -q postgre) || $NEXT_WAIT_TIME -eq 60 ]]
do
    sleep $(( NEXT_WAIT_TIME++ ))
done

docker cp scripts/init_tables.sh $(${DOCKER_COMPOSE_COMMAND} ps -q postgre):/init_tables.sh

# Wait up database
NEXT_WAIT_TIME=0
echo "Trying init tables:"
until ${DOCKER_COMPOSE_COMMAND} exec postgre bash /init_tables.sh > /dev/null 2>&1 || [ $NEXT_WAIT_TIME -eq 60 ];
do
    sleep $(( NEXT_WAIT_TIME++ ))
done

popd