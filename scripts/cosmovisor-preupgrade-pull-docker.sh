#!/usr/bin/env bash

set -xe

UPGRADE_NAME=$1
UPGRADE_HEIGHT=$2 # unused

mkdir -p $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin

tee $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/blitd > /dev/null <<EOF
#!/usr/bin/env bash

set -xe

UPGRADE_NAME=\$1
UPGRADE_HEIGHT=\$2 # unused

# Docker image name from your docker-compose file
IMAGE_NAME="blitchain/blitchain:$UPGRADE_NAME"

# Mapping the daemon home directory from host to container
DAEMON_HOME_VOLUME="$DAEMON_HOME:/home/user/.blit"

# Collect all arguments for the blitd command
BLITD_ARGS="\$@"

# External ports from environment variables with default values
# See: https://docs.cosmos.network/main/user/run-node/run-production#firewall
# NOTE: bind to 0.0.0.0 instead of localhost if you are having trouble connecting

BLIT_P2P_26656=\${BLIT_P2P_26656:-26656}
BLIT_JSONRPC_26657=\${BLIT_JSONRPC_26657:-26657}
BLIT_REST_1317=\${BLIT_REST_1317:-1317}
BLIT_GRPC_9090=\${BLIT_GRPC_9090:-9090}

# Check if script is run in TTY
TTY_OPTION=""
if [ -t 0 ]; then
    TTY_OPTION="-ti"
fi

# Run docker command
exec docker run --rm \\
    -u \$(id -u):\$(id -g) \\
    \$TTY_OPTION \\
    -v \$DAEMON_HOME_VOLUME \\
    -p \$BLIT_P2P_26656:26656 -p \$BLIT_JSONRPC_26657:26657 -p \$BLIT_REST_1317:1317 -p \$BLIT_GRPC_9090:9090 \\
    \$IMAGE_NAME \\
    ./bin/blitd \$BLITD_ARGS
EOF

chmod +x $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/blitd

