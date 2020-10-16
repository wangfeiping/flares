# flares

**flares** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

### Relay to gaia

```
# init & start gaiad

$ gaiad version --long
name: gaia
server_name: gaiad
version: stargate-3
commit: 5f28583b510a86fac671e4590d0863d56989735d
build_tags: netgo,ledger
go: go version go1.14.7 linux/amd64

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

# https://github.com/ovrclk/relayer @ stargate-3

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