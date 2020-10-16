#!/bin/sh

GAIA_DATA=./demo/runtime
chainid1=flares

# echo "Generating gaia configurations..."
mkdir -p $GAIA_DATA/${chainid1}/config/gentx/
mkdir -p $GAIA_DATA/${chainid1}/data/

./demo/scripts/one-chain ${chainid1} $GAIA_DATA 26657 26656 6060 9090
