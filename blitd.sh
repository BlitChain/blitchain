#!/bin/bash

# Docker image name from your docker-compose file
IMAGE_NAME="blitchain/blitchain:${TAG-dev}"

# Mapping the daemon home directory from host to container
DAEMON_HOME_MAP="${DEAMON_HOME}:/home/user/.blit"

# Collect all arguments for the blitd command
BLITD_ARGS="$@"

# Run docker command
docker run -it --rm \
    -e TAG \
    -v $DAEMON_HOME_MAP \
    -p 26656:26656 -p 26657:26657 -p 1317:1317 \
    $IMAGE_NAME \
    ./blitd $BLITD_ARGS
