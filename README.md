# flares

**flares** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Create

```
# github.com/tendermint/starport   @ latest of develop
# github.com/cosmos/gaia           @ stargate-4
# github.com/ovrclk/relayer        @ stargate-4
# github.com/cosmos/cosmos-sdk     @ v0.40.0-rc0
# github.com/tendermint/tendermint @ v0.34.0-rc4.0.20201005135527-d7d0ffea13c6

$ starport app github.com/wangfeiping/flares --sdk-version stargate

$ cd flares/

$ starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Relay to gaia

### init & start chains

```
$ gaiad version --long
name: gaia
server_name: gaiad
version: stargate-4
commit: 3a8b1b414004ccddfa255fd0cd1499bbf6659d71
build_tags: netgo,ledger
go: go version go1.15.2 linux/amd64

# clear the old data

$ rm -rf ./demo/runtime/

# start two chains

$ ./demo/scripts/start-gaia.sh
$ ./demo/scripts/start-flares.sh

# send funds

$ flaresd tx bank send \
    validator $(flaresd keys show user --keyring-backend=test --home ./demo/runtime/flares -a) 100000samoleans \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id=flares

$ flaresd q bank balances \
    $(flaresd keys show user --keyring-backend=test --home ./demo/runtime/flares -a)

$ gaiad tx bank send \
    validator $(gaiad keys show user --keyring-backend=test --home ./demo/runtime/ibc0 -a) 100000samoleans \
    --keyring-backend=test \
    --home ./demo/runtime/ibc0 \
    --chain-id=ibc0 --node=tcp://127.0.0.1:26557

$ gaiad q bank balances \
    $(gaiad keys show user --keyring-backend=test --home ./demo/runtime/ibc0 -a) \
    --node=tcp://127.0.0.1:26557

```

### Relayer

```

$ rly config init

$ rly cfg add-dir ./demo/relayer/config/

# Delete if the keys are already exist
$ rly keys delete flares testkey
$ rly keys delete ibc0 testkey

# Now, add the key seeds from each chain to the relayer to give it funds to work with
$ rly keys restore flares testkey "$(jq -r '.mnemonic' demo/runtime/flares/key_seed.json)"
$ rly k r ibc0 testkey "$(jq -r '.mnemonic' demo/runtime/ibc0/key_seed.json)"

$ rly light init flares -f
$ rly l i ibc0 -f

$ rly chains list

$ rly tx link demo -d -o 3s

# Then send some tokens between the chains
$ rly tx transfer ibc0 flares 5samoleans $(rly chains address flares)
$ rly tx relay demo -d

# See that the transfer has completed
$ rly q bal ibc0 | jq
$ rly q bal flares | jq

```

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)