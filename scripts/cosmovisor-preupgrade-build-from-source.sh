#!/usr/bin/env bash

set -xe

UPGRADE_NAME=$1
UPGRADE_HEIGHT=$2 # unused

# assert that the upgrade name is not empty
if [ -z "$UPGRADE_NAME" ]; then
  echo "UPGRADE_NAME is empty"
  exit 1
fi

# assert DAEMON_HOME is set
if [ -z "$DAEMON_HOME" ]; then
  echo "DAEMON_HOME is empty"
  exit 1
fi


rm -rf  $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME

git clone --depth 1 --branch $UPGRADE_NAME https://github.com/BlitChain/blitchain.git $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME
cd $DAEMON_HOME/cosmovisor/upgrades/$UPGRADE_NAME

goenv install -s
exec make installdeps build
