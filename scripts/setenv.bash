#!/bin/bash
#
# This is the Blit state sync script, which is based off the Regen Network State sync file, which is based on the Gaia State sync file, which is based on scripts written by Bitcanna and Microtick.
# http://192.46.233.149:26657/


RPC_HOST="192.46.233.149"

set -uxe

HOME_DIR=~/.blit

# GET TRUST HASH AND TRUST HEIGHT

NODE_ID=$(curl "http://$RPC_HOST:26657/status" | jq -r '.result.node_info.id')

BLOCK=`curl -s "http://$RPC_HOST:26657/commit" | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"` 


BLOCK_HEIGHT=$(echo $BLOCK | jq -r .height)
TRUST_HASH=$(echo $BLOCK | jq -r .hash)


# TELL USER WHAT WE ARE DOING
echo "TRUST HEIGHT: $BLOCK_HEIGHT"
echo "TRUST HASH: $TRUST_HASH"

# GET GENESIS FILE
curl http://$RPC_HOST:26657/genesis | jq '.result.genesis' > $HOME_DIR/config/genesis.json

# UPDATE CONFIG FILE
BLIT_P2P_MAX_NUM_OUTBOUND_PEERS=100
BLIT_P2P_MAX_NUM_INBOUND_PEERS=100
BLIT_STATESYNC_RPC_SERVERS="http://$RPC_HOST:26657,http://$RPC_HOST:26657"
BLIT_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
BLIT_STATESYNC_TRUST_HASH=$TRUST_HASH
BLIT_P2P_SEEDS="$NODE_ID@$RPC_HOST:26656"
BLIT_P2P_PERSISTENT_PEERS="$NODE_ID@$RPC_HOST:26656"

sed -i '/persistent_peers =/c\persistent_peers = "'"$BLIT_P2P_PERSISTENT_PEERS"'"' $HOME_DIR/config/config.toml
sed -i '/seeds =/c\seeds = "'"$BLIT_P2P_SEEDS"'"' $HOME_DIR/config/config.toml
sed -i '/max_num_outbound_peers =/c\max_num_outbound_peers = '$BLIT_P2P_MAX_NUM_OUTBOUND_PEERS'' $HOME_DIR/config/config.toml
sed -i '/max_num_inbound_peers =/c\max_num_inbound_peers = '$BLIT_P2P_MAX_NUM_INBOUND_PEERS'' $HOME_DIR/config/config.toml
sed -i '/enable =/c\enable = true' $HOME_DIR/config/config.toml
sed -i '/rpc_servers =/c\rpc_servers = "'"$BLIT_STATESYNC_RPC_SERVERS"'"' $HOME_DIR/config/config.toml
sed -i '/trust_height =/c\trust_height = '$BLIT_STATESYNC_TRUST_HEIGHT'' $HOME_DIR/config/config.toml
sed -i '/trust_hash =/c\trust_hash = "'"$BLIT_STATESYNC_TRUST_HASH"'"' $HOME_DIR/config/config.toml
sed -i '/double_sign_check_height =/c\double_sign_check_height = 1' $HOME_DIR/config/config.toml
sed -i '/discovery_time =/c\discovery_time = "5s"' $HOME_DIR/config/config.toml
