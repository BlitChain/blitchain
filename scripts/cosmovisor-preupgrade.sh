#!/usr/bin/env bash


UPGRADE_NAME=$1
UPGRADE_HEIGHT=$2 // unused

git clone --depth 1 --branch $UPGRADE_NAME --recursive https://github.com/BlitChain/blitchain.git $DAEMON_HOME/upgrades/$UPGRADE_NAME
cd $DAEMON_HOME/upgrades/$UPGRADE_NAME
make build

