#!/usr/bin/env bash

set -xe

UPGRADE_NAME=$1
UPGRADE_HEIGHT=$2 # unused

rm -r  $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME || \
  git clone --depth 1 --branch $UPGRADE_NAME --recursive https://github.com/BlitChain/blitchain.git $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME && \
  cd $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME && 
  exec make build

