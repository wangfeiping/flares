# Flares

**Flares** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

- Flares was designed to apply to scenes that are willing to
 accept more value consensus-based tokens, payment, auctions, games, DeFi, etc.  
- It is not intended to increase the use of just one token,
 but to discover the value of all crypto assets.  
- It's an IBC network.  
- It's a no-tokens-binding blockchain.  
- It's a value consensus.  
- It's a multi to multi contract or module.  

## Introduction

### Concept

- contract

As an intermediate module for payment, secure and realize the value consensus of multiple crypto assets through services such as DEC and DeFi

- voucher

Digital credentials for encrypted assets after value consensus

- board

Provides queries and exchange interfaces for encrypted assets

- clearing

This module realizes the functions of managing and serencing encrypted assets and reallocating them according to the contract application

## Demos

### Sealed Monsters - flares with games

Refer to a demo project [scavenge](https://tutorials.cosmos.network/scavenge/tutorial/01-background.html) and make some changes.

![Sealed Monsters - flares with games](https://raw.githubusercontent.com/wangfeiping/flares/hackatom2020/docs/flares-game-sealed-monsters.png)

### Name Service - flares with payment or auctions

Refer to a demo project [nameservice](https://tutorials.cosmos.network/nameservice/tutorial/00-intro.html) and make some changes.

![Name Service - flares with payment or auctions](https://raw.githubusercontent.com/wangfeiping/flares/hackatom2020/docs/flares-nameservice-payment-or-auctions.png)

### Future view

![Future view](https://raw.githubusercontent.com/wangfeiping/flares/hackatom2020/docs/flares-future-view.png)

### Test Cases

```
$ cd $GOPATH/src/github.com/wangfeiping/flares
$ go test ./...

or

$ go get -u github.com/onsi/ginkgo/ginkgo
$ cd $GOPATH/src/github.com/wangfeiping/flares
$ ginkgo ./...

```

### build

```
$ make all

```

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

$ flaresd keys list \
    --keyring-backend=test \
    --home ./demo/runtime/flares

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

## Protobuf

```
# changes the proto/*/*.proto file(s)

# just execute

$ ./scripts/protocgen

```

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
- [BDD Test](https://github.com/onsi/ginkgo)
