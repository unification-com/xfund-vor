# Running an Oracle

The VOR Oracle software monitors the `VORCoordinator` smart contract for
emitted events, and acts on those events to generate random numbers and the
required proof, before sending the data back to `VORCoordinator`.

Running a VOR Provider Oracle is quick and easy, and this guide will run
you through getting set up as a VOR Provider.

The software is a single application, `oracle`: the server, which can serve
several chains from one process, and the commands that administer it.

## Prerequisites

This guide uses Linux as the platform for running a VOR Oracle. The software
has been tested on Ubuntu and CentOS distros. It assumes your system has some
basic applications installed such as `git`, `curl`, `nano` and `make` etc.

Depending on your database backend of choice, you will also need either
**SQLite** or **PostgreSQL** installing on your host.

**Note:** SQLite recommended mainly for dev/testing

## Download pre-build binaries

The latest pre-built (linux) release can be downloaded from
[https://github.com/unification-com/xfund-vor/releases](https://github.com/unification-com/xfund-vor/releases)

A SHA256 `checksum.txt` is supplied with the archive. Download it to
the same location as the archive, and check:

```bash
$ sha256sum --check checksum.txt 
vor-oracle_linux_v0.1.1.tar.gz: OK
```

## Compile from source

::: warning
You will require **Go 1.27** or later to compile from source - the latest version can be
installed by following the [official instructions](https://go.dev/doc/install). Build with the
latest Go patch release (for example `GOTOOLCHAIN=go1.27.4 make build`): it carries the standard
library's security fixes.
:::

Clone the repo at the latest release tag, from
[https://github.com/unification-com/xfund-vor/releases](https://github.com/unification-com/xfund-vor/releases):

```bash
git clone -b [latest-release-tag] https://github.com/unification-com/xfund-vor
```

Then install the `oracle` binary in your `$GOPATH/bin`:

```bash
cd xfund-vor
make install
```

or compile it to `xfund-vor/oracle/build/oracle` with `make build`.

## Configuration

The `oracle` reads a `config.json` file, `./config.json` by default; set its path with `-c`.

```bash
mkdir $HOME/vor
nano $HOME/vor/config.json
```

List each chain the oracle serves under `chains`:

```json
{
  "chains": [
    {
      "name": "polygon",
      "network_id": 137,
      "eth_http_host": "https://polygon-rpc.example",
      "contract_address": "0x9e84d9DB...",
      "blockhash_store_address": "0x...",
      "first_block": 63000000,
      "max_gas_price": 300
    },
    {
      "name": "mainnet",
      "network_id": 1,
      "eth_http_host": "https://eth-rpc.example",
      "contract_address": "0xCeB679DE...",
      "blockhash_store_address": "0x...",
      "first_block": 21000000,
      "max_gas_price": 120
    }
  ],
  "serve": {
    "host": "127.0.0.1",
    "port": 8445
  },
  "keystorage": {
    "file": "/home/username/vor/keystore.json",
    "account": "oracle"
  },
  "database": {
    "dialect": "postgres",
    "host": "127.0.0.1",
    "port": 5432,
    "user": "vor",
    "password": "...",
    "database": "vor"
  }
}
```

An oracle serving one chain may instead keep that chain's settings at the top level, as configs
for v0.1 did.

Each chain's settings:

- `name` - selects the chain in the admin commands (`--chain`). Default: the network id
- `network_id` - the chain's id, e.g. 1 = Ethereum, 137 = Polygon
- `eth_http_host` - HTTP(S) JSON-RPC endpoint for the chain
- `contract_address` - the `VORCoordinator` contract. See
  [contracts](https://vor.unification.io/contracts.html)
- `blockhash_store_address` - the `BlockHashStore` contract: retries of requests older than ~250
  blocks need it
- `first_block` - where the oracle first scans the chain from; afterwards it resumes where it
  stopped. Set it near the current block for a new chain. Default `1`
- `check_duration` - seconds between polls. Default `15`
- `max_gas_price` - the most, in gwei, the oracle pays per unit of gas (the gas price, or an EIP-1559
  transaction's fee cap); `0` for no cap. It must stay above the chain's base fee, or transactions
  won't be mined. Default `150`
- `legacy_gas` - `true` to send legacy-priced transactions even where the chain supports EIP-1559.
  Default `false`: EIP-1559 where the chain's blocks have a base fee
- `wait_confirmations` - blocks to wait before fulfilling a request. Default `10`
- `event_scan_batch_blocks` - the most blocks one `eth_getLogs` call asks for: many RPCs refuse wide
  ranges. Default `2000`
- `fee` - the fee, in xFUND's smallest unit, to register your proving key at on this chain, if it isn't
  registered there yet (e.g. a chain added to a running oracle). Ignored once registered: change the
  fee with `oracle changefee`

A setting a chain leaves out takes its default; one it sets to `0` stays `0`. Gas limits are always
estimated.

Process-wide settings:

- `serve.host`, `serve.port` - where the oracle's admin API listens. Keep it off public interfaces.
  Default `127.0.0.1:8445`
- `keystorage.file` - the keystore, holding your encrypted private key. Created on the first run.
  Default `./keystore.json`
- `keystorage.account` - the account in the keystore the oracle signs with (default: the first)
- `database.dialect` - `postgres` or `sqlite` (dev/testing). Default `sqlite`
- `database.storage` - (`sqlite` only) the database file. Default `./oracle.db`
- `database.host`, `.port`, `.user`, `.password`, `.database` - (`postgres` only)

One database and one keystore serve every chain: each record carries its chain's id, and the same
provider key is used on every chain.

## First run

The first run asks for some input, so run it in a terminal:

```bash
/path/to/oracle start -c $HOME/vor/config.json
```

You will be asked for:

1. Username - the account name in the keystore; put it in `keystorage.account`
2. Fee - your initial xFUND fee, for example 100000000 is 0.1 xFUND
3. Add an existing key or create a new one - import an existing private key (typed without being
   shown), or have the `oracle` generate one (it prints the new key's address)

The `oracle` then prints your **keystore key**, the only copy:

```bash
Your keystore key:   3f9c...
Start the oracle with it (-k). KEEP THIS KEY SAFE! YOU WILL LOSE YOUR KEYSTORAGE DATA WITHOUT IT!

The admin commands' API key is in /home/username/vor/api_key
```

::: tip
Save the keystore key in a file only your user can read, for `-k`, and back up `keystore.json`.
Each private key in it is a standard go-ethereum (v3, scrypt) keystore, usable with any Ethereum
tool.
:::

At start-up the `oracle` registers your proving key with `VORCoordinator` on each chain where it
isn't registered yet, at the chain's `fee` (or the fee entered on the first run), waiting for the
transaction to succeed, so the account needs gas on every chain.

## Running the Oracle

```bash
/path/to/oracle start -c $HOME/vor/config.json -k /path/to/keystore-key-file
```

`-k` takes a file holding the keystore key (or the key itself, for testing only); without it, the
`oracle` asks for the key in a terminal, and refuses to start without one. A chain whose RPC is
unreachable at start-up is retried every minute without holding up the others.

`oracle version` (or `--version`) prints the version.

### Running the oracle as a service

```bash
sudo nano /etc/systemd/system/vor_oracle.service
```

```
[Unit]
Description=VOR Provider Oracle

[Service]
User=USERNAME
Group=USERNAME
WorkingDirectory=/home/USERNAME
ExecStart=/home/USERNAME/go/bin/oracle start -c /path/to/config.json -k /path/to/keystore-key-file
Restart=on-failure
RestartSec=30

[Install]
WantedBy=default.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl start vor_oracle
```

## Keys

The oracle has two keys:

- the **keystore key** (`-k`) decrypts the keystore
- the **API key** authenticates the admin commands. The keystore holds only its hash; the key
  itself is in `api_key` beside the config. Issue a new one, with the oracle stopped, with:

```bash
oracle api-key -c $HOME/vor/config.json [--out FILE]
```

### Upgrading from v0.1

v0.1 keystores used weak encryption under a guessable key, and the `oracle` refuses to start on
one. With the oracle stopped, back up its directory, then convert the keystore under a new key and
issue an API key:

```bash
cd $HOME/vor
openssl rand -hex 32 > password.new && chmod 600 password.new
oracle migrate-keystore -c config.json -k password --new-key password.new && mv password.new password
oracle api-key -c config.json
```

Every key is converted and checked from disk before the old keystore is replaced. Once the oracle
runs, delete the backup: the old keystore and its key together expose your private key.

The first start also brings a v0.1 database up to date, assigning its records to the chain the
config names, so start each v0.1 database with a config for its one chain the first time.

## Administering the oracle

Run on the oracle's host, the admin commands find the oracle from its config (`-c`) and use the API
key in `api_key` beside it (or `--api-key FILE`). When the oracle serves several chains, name one
with `--chain <name or network id>`.

```bash
oracle status -c config.json                              # is it running?
oracle about -c config.json --chain polygon               # address, key hash, balances, fee
oracle queryfees -c config.json --chain polygon [consumer]
oracle changefee -c config.json --chain polygon 100000000 # 0.1 xFUND
oracle changegranularfee -c config.json --chain polygon 50000000 0xConsumer...
oracle querywithdrawable -c config.json --chain polygon
oracle withdraw -c config.json --chain polygon 100000000 0xRecipient...
oracle queryrequests -c config.json --chain polygon [--page 1] [--limit 10] [--status -1] [--order desc]
oracle gettx -c config.json --chain polygon 0xTxHash...
oracle register -c config.json --chain polygon <account> <fee>  # asks for the private key
oracle stop -c config.json                                # stops the whole oracle
```

Amounts and fees are in xFUND's smallest unit (9 decimals).

### Analytics

```bash
oracle analytics -c config.json --chain polygon 100
oracle analytics sim -c config.json --chain polygon 100 --if-gas 150 --if-fees 0.05 [--consumer 0x...]
oracle analytics consumers -c config.json --chain polygon [consumer]
```

`analytics` summarises gas, costs and fees over the last fulfilled requests; `sim` recalculates them
at a simulated gas price and fee; `consumers` breaks them down per consumer contract. Costs use
xFUND's price from CoinGecko.
