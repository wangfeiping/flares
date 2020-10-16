#!/bin/sh

GAIA_DATA=./demo/runtime
chainid0=ibc0

# echo "Generating gaia configurations..."
mkdir -p $GAIA_DATA/${chainid0}/config/gentx/
mkdir -p $GAIA_DATA/${chainid0}/data/

./demo/scripts/one-chain ${chainid0} $GAIA_DATA 26557 26556 6050 9050
