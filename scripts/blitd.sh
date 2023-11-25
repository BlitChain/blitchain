#!/bin/bash

# Docker image name from your docker-compose file
IMAGE_NAME="blitchain/blitchain:${TAG:-dev}"

# Mapping the daemon home directory from host to container
DAEMON_HOME_MAP="${DEAMON_HOME-$HOME/.blit}:/home/user/.blit"

# Collect all arguments for the blitd command
BLITD_ARGS="$@"

# External ports from environment variables with default values
# See: https://docs.cosmos.network/main/user/run-node/run-production#firewall
P2P_26656=${P2P_26656:-26656}
JSONRPC_26657=${JSONRPC_26657:-26657}
REST_1317=${REST_1317:-1317}
GRPC_9090=${GRPC_9090:-9090}

# Check if script is run in TTY
TTY_OPTION=""
if [ -t 0 ]; then
    TTY_OPTION="-ti"
fi

# Run docker command
exec docker run -it --rm \
    -e TAG \
    $TTY_OPTION \
    -v $DAEMON_HOME_MAP \
    -p $P2P_26656:26656 -p $JSONRPC_26657:26657 -p $REST_1317:1317 -p $GRPC_9090:9090 \
    $IMAGE_NAME \
    ./bin/blitd $BLITD_ARGS
