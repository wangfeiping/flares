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

### The command line executes the test case (Sealed Monsters)

** Prepare the environment **

```
# build

$ git clone https://github.com/wangfeiping/flares.git

$ cd flares/

$ make all

# version

$ flaresd version --long

name: flares
server_name: flaresd
version: 0.0.1-1-g6e89fd6
commit: 6e89fd63c6402baa46ca4bd4240156e52396e460
build_tags: ""
go: go version go1.15.2 linux/amd64

$ rm -rf ./demo/runtime/

# start flares

$ ./demo/scripts/start-flares.sh

# check init accounts

$ flaresd keys list --keyring-backend=test --home ./demo/runtime/flares

- name: user
- name: user1
- name: user2
- name: user3
- name: user4
- name: user5
- name: validator

# query all balances of account

$ flaresd q bank balances \
    $(flaresd keys show user --keyring-backend=test --home ./demo/runtime/flares -a)

# create three boards

$ flaresd tx flares create-board base aaa "" \
    --from $(flaresd keys show user --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

$ flaresd tx flares create-board base bbb "" \
    --from $(flaresd keys show user --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

$ flaresd tx flares create-board base ccc "" \
    --from $(flaresd keys show user --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

$ flaresd tx flares create-board base base "" \
    --from $(flaresd keys show user --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

# list all boards

$ flaresd q flares list-board

board:
- acceptDenom: aaa
  address: cosmos1ntkwggq4g46gzetly3k4mree469u0f62wsjd92
  baseDenom: base
  creator: cosmos1x2mve3jj3relg0f6ed8hew6w0rx73dxfgywmzh
  id: 4ea630f0-76c2-4512-a565-922036d403db
  source: ""
- acceptDenom: bbb
  address: cosmos1k2q46m50nakgcw62hnwdc9fl4dcwzep7g6hptn
  baseDenom: base
  creator: cosmos1x2mve3jj3relg0f6ed8hew6w0rx73dxfgywmzh
  id: 1725e4b2-e451-4b28-9ee1-2c5d8bfdb651
  source: ""
- acceptDenom: ccc
  address: cosmos1cvn2sc94jjp5xu7npl5nq0paugjrqw7ujczsu0
  baseDenom: base
  creator: cosmos1x2mve3jj3relg0f6ed8hew6w0rx73dxfgywmzh
  id: 7ca33f06-b67a-44cd-9022-d280ac0613d5
  source: ""
- acceptDenom: base
  address: cosmos1xmu47zz08fhyzpumprjd5yjdpg3y4pj4eus9xl
  baseDenom: base
  creator: cosmos1x2mve3jj3relg0f6ed8hew6w0rx73dxfgywmzh
  id: 29f3806a-1a07-4997-ab0d-21fc5a825a00
  source: ""
pagination:
  next_key: null
  total: "4"

# send base & accept tokens to the address of board

$ flaresd tx bank send user cosmos1ntkwggq4g46gzetly3k4mree469u0f62wsjd92 \
    1000base \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id=flares

$ flaresd tx bank send user cosmos1ntkwggq4g46gzetly3k4mree469u0f62wsjd92 \
    100000aaa \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id=flares

# query the balances of board

$ flaresd q bank balances cosmos1ntkwggq4g46gzetly3k4mree469u0f62wsjd92

balances:
- amount: "100000"
  denom: aaa
- amount: "1000"
  denom: base
pagination:
  next_key: null
  total: "0"

# other board do the same

$ flaresd tx bank send user cosmos1k2q46m50nakgcw62hnwdc9fl4dcwzep7g6hptn \
    1000base \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id=flares

$ flaresd tx bank send user cosmos1k2q46m50nakgcw62hnwdc9fl4dcwzep7g6hptn \
    50000bbb \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id=flares

$ flaresd tx bank send user cosmos1cvn2sc94jjp5xu7npl5nq0paugjrqw7ujczsu0 \
    1000base \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id=flares

$ flaresd tx bank send user cosmos1cvn2sc94jjp5xu7npl5nq0paugjrqw7ujczsu0 \
    10000ccc \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id=flares

# start the game
# create a monster (It will auto create a contract)

$ flaresd tx sealedmonsters summon-monster \
    "It has a mane of snakes, three heads, and a tail which was a living serpent." \
    "kerberos" 1base \
    --from $(flaresd keys show user --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

# list monsters

$ flaresd q sealedmonsters list-monster

# list contracts

$ flaresd q flares list-contract

contract:
- accept: ""
  bottom: 1base
  code: 0
  creator: cosmos1x2mve3jj3relg0f6ed8hew6w0rx73dxfgywmzh
  durationHeight: -1
  height: "179"
  id: 4e320bea-bd2b-4875-9276-9406623f76be
  key: sealedmonsters
  module: sealedmonsters
  receiver: cosmos1jes6vhf9terentzxv8aym8raa2rg3yethenjw8
  result: ""
pagination:
  next_key: null
  total: "1"

$ flaresd tx sealedmonsters create-seal cubolous 1000aaa \
    --from $(flaresd keys show user1 --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

$ flaresd tx sealedmonsters create-seal kerberos 9bbb \
    --from $(flaresd keys show user2 --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

$ flaresd tx sealedmonsters create-seal kerberos 999ccc \
    --from $(flaresd keys show user3 --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

$ flaresd tx sealedmonsters create-seal kerbero 100aaa \
    --from $(flaresd keys show user4 --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

$ flaresd tx sealedmonsters create-seal kerbos 100ccc \
    --from $(flaresd keys show user5 --keyring-backend=test --home ./demo/runtime/flares -a) \
    --keyring-backend=test \
    --home ./demo/runtime/flares \
    --chain-id flares

# check the total supply

$ flaresd q bank total

supply:
- amount: "121080000"
  denom: VOUCHER
...

# the game would be over after 100 blocks

$ flaresd q flares list-contract

contract: []
pagination:
  next_key: null
  total: "1"

$ flaresd q flares list-contract success

It's that the contract has been successfully processed.

contract:
- accept: ""
  bottom: 1base
  code: 0
  creator: cosmos1x2mve3jj3relg0f6ed8hew6w0rx73dxfgywmzh
  durationHeight: -1
  height: "179"
  id: 4e320bea-bd2b-4875-9276-9406623f76be
  key: sealedmonsters
  module: sealedmonsters
  receiver: cosmos1jes6vhf9terentzxv8aym8raa2rg3yethenjw8
  result: success
pagination:
  next_key: null
  total: "1"

# check rewards

$ flaresd q bank balances \
    $(flaresd keys show user1 --keyring-backend=test --home ./demo/runtime/flares -a)

user1 lost 1000aaa

- amount: "0"
  denom: VOUCHER
- amount: "99999999000"
  denom: aaa

$ flaresd q bank balances \
    $(flaresd keys show user2 --keyring-backend=test --home ./demo/runtime/flares -a)

user2 won 2aaa

- amount: "0"
  denom: VOUCHER
- amount: "100000000002"
  denom: aaa

$ flaresd q bank balances \
    $(flaresd keys show user3 --keyring-backend=test --home ./demo/runtime/flares -a)

user3 won 1098aaa 100ccc and 1base(monster rewards)

- amount: "0"
  denom: VOUCHER
- amount: "100000001098"
  denom: aaa
- amount: "100000000001"
  denom: base
- amount: "100000000100"
  denom: ccc

$ flaresd q bank balances \
    $(flaresd keys show user4 --keyring-backend=test --home ./demo/runtime/flares -a)

user4 lost 100aaa

- amount: "0"
  denom: VOUCHER
- amount: "99999999900"
  denom: aaa

$ flaresd q bank balances \
    $(flaresd keys show user5 --keyring-backend=test --home ./demo/runtime/flares -a)

user5 lost 100ccc

- amount: "0"
  denom: VOUCHER
- amount: "99999999900"
  denom: ccc

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
