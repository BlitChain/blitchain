#!/usr/bin/env bash

set -xe

UPGRADE_NAME=$1
UPGRADE_HEIGHT=$2 # unused

rm -rf  $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME

mkdir -p $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin

tee $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/blitd > /dev/null <<EOF
#!/usr/bin/env bash

set -xe

UPGRADE_NAME=\$1
UPGRADE_HEIGHT=\$2 # unused

# Docker image name from your docker-compose file
IMAGE_NAME="blitchain/blitchain:\$UPGRADE_NAME"

# Mapping the daemon home directory from host to container
DAEMON_HOME_VOLUME="-v \$DAEMON_HOME:/home/user/.blit"

# Collect all arguments for the blitd command
BLITD_ARGS="\$@"

# Initialize PORT_OPTIONS as an empty string
PORT_OPTIONS=""

# Determine if port mapping is needed based on \$BLITD_ARGS starting with "start"
if (echo "\$BLITD_ARGS" | grep -q "^start"); then
    # Assign port mappings directly to PORT_OPTIONS with default values, split into multiple lines for readability
    PORT_OPTIONS="-p \${BLIT_P2P_26656:-26656}:26656 \\
                  -p \${BLIT_JSONRPC_26657:-26657}:26657 \\
                  -p \${BLIT_REST_1317:-1317}:1317 \\
                  -p \${BLIT_GRPC_9090:-9090}:9090"
fi

# Check if script is run in TTY
TTY_OPTION=""
if [ -t 0 ]; then
    TTY_OPTION="-ti"
fi

# Run docker command
exec docker run --rm \\
    -u \$(id -u):\$(id -g) \\
    \$TTY_OPTION \\
    \$DAEMON_HOME_VOLUME \\
    \$PORT_OPTIONS \\
    \$IMAGE_NAME \\
    ./bin/blitd \$BLITD_ARGS
EOF

chmod +x $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME/bin/blitd
