# Running an Oracle

The VOR Oracle software monitors the `VORCoordinator` smart contract for
emitted events, and acts on those events to generate random numbers and the
required proof, before sending the data back to `VORCoordinator`.

Running a VOR Provider Oracle is quick and easy, and this guide will run
you through getting set up as a VOR Provider.

The software is comprised of two applications:

1. `oracle` - the server application
2. `oraclecli` - the CLI tool for `oracle` server administration

## Prerequisites

This guide uses Linux as the platform for running a VOR Oracle. The software
has been tested on Ubuntu and CentOS distros. It assumes your system has some
basic applications installed such as `git`, `curl`, `nano` and `make` etc.

You will also require **Golang** - the latest version can be installed by following
the [official instructions](https://golang.org/doc/install).

Depending on your database backend of choice, you will also need either
**SQLite** or **PostgreSQL** installing on your host.

**Note:** SQLite recommended mainly for dev/testing

## Compile from source

Clone the repo:

```bash
git clone https://github.com/unification-com/xfund-vor
```

Then compile the binaries using the `make install` target:

```bash
cd xfund-vor
make install
```

This will install the `oracle` and `oraclecli` binaries in your `$GOPATH/bin`

Alternatively, run the `build` target:

```bash
cd xfund-vor/oracle
make build
```

which will compile the binaries to `xfund-vor/oracle/build` and
`xfund-vor/oracle-cli/build`

## Configuration

The `oracle` requires a `config.json` file, which will contain all the
information it needs to run. By default it will look for `config.json` in 
the same directory as the `oracle` binary, but the location can be set
at runtime with the `-c` flag.

Create your `config.json`:

```bash
mkdir $HOME/vor
nano $HOME/vor/config.json
```

```json
{
  "contract_address": "0x00.....",
  "eth_http_host": "http://127.0.0.1:8545",
  "eth_ws_host": "http://127.0.0.1:8545",
  "network_id": 696969,
  "serve": {
    "host": "0.0.0.0",
    "port": 8446
  },
  "first_block": 1,
  "keystorage": {
    "file": "/home/username/vor/keystore.json",
    "account": "oracle"
  },
  "gas_price_limit": 1000000,
  "database": {
    "dialect": "sqlite",
    "storage": "/home/username/vor/oracle.db"
  }
}
```

The config options are as follows:

- `contract_address` - the address of the `VORCoordinator` smart contract. See
[contracts](../contracts.md).
- `eth_http_host` - HTTP(S) host for your Eth provider. E.g. Infura
- `eth_ws_host` - WS(S) host for your Eth provider. E.g. Infura
- `network_id` - Eth network ID, e.g. 1 = mainnet, 4 = Rinkeby etc.
- `serve.host` - host to serve the `oracle` on. This is used by the `oracle-cli`
  tool, and should not be publicly exposed.
- `serve.port` - port to serve the `oracle` on. This is used by the `oracle-cli`
  tool, and should not be publicly exposed.
- `first_block` - default first block to scan the contract from. Should be 
  a value near to the block your oracle will do it's first run from. Only used
  when the oracle first connects.
- `keystorage.file` - path to the `keystore` which will contain your encrypted
  private key and other runtime info. It will be created on the oracle's first
  run if one does not exist.
- `keystore.account` - account name used to identify the private key. Set on 
  first run.
- `gas_price_limit` - max gas you are willing to pay (in Wei) for fulfilling
  a request.
- `database.dialect` - `postgres` or `sqlite`
- `database.storage` - (`sqlite` only) - path to the DB file. It will be created on the oracle's first
  run if one does not exist.
- `database.host` - (`postgres` only) - DB host IP/name
- `database.port` - (`postgres` only) - DB Port
- `database.user` - (`postgres` only) - DB username
- `database.password` - (`postgres` only) - DB password
- `database.database` - (`postgres` only) - DB name

## First run

The first time `oracle` is run, it will prompt for some input to further 
configure the environment.

Start the `oracle`, assuming `config.json` has been saved to 
`$HOME/vor/config.json`:

```bash
/path/to/oracle start -c $HOME/vor/config.json
```
 You will be asked for:

1. Username - this will be used as the account name in `keystore.json`,
and should be entered as the value for `keystore.account` in `config.json`
2. Fee - your initial xFUND fee, for example 100000000 will be 0.1 xFUND
3. Add existing key/Create a new key - you can either import an existing
Eth private key, or have the `oracle` create a new one for you.
   
::: tip
If you elect to have the `oracle` generate a new private key for you, you
will need to manually register your key with `VORCOOrdinator`. 
:::

You need to ensure the wallet address has sufficient Eth to
send the `registerProvingKey` transaction to `VORCOOrdinator`. If you
opted for importing an existing key, this process is automatically run
for you, and your key will be registered on the `oracle`'s first run.

Once complete, you should see something along the lines of:

```bash
# some info about your config and key...

Your daemon api key:   rkf.....

Use this key to login via cli/HTTP (command: oracle-cli settings)
KEEP THIS KEYS SAFE!
```

::: tip
It is important to key the **daemon api key** safe, as this is used to start
the oracle, and access its api via the `oraclecli`. The `oracle` also uses
it to encrypt/decrypt your `keystore`.

It is also advisable to backup the `keystore.json` file.
:::

Finally, it will ask you to enter the **api key** previously output in order to start running
the `oracle`.

The `oracle` can be stopped with <key>CTRL+C</key>, or the safer api method:

```bash
/path/to/oraclecli stop
```

which will signal the `oracle` to stop via its API.

::: tip
`oraclecli` has its own configuration file, which will be covered later
in this guide.
:::

## Running the Oracle

Once configured, the `oracle` can be run using:

```bash
/path/to/oracle start -c $HOME/vor/config.json
```

It will prompt for the **api key** which was output for oyu during the first run.

Optionally, the key can be passed to the `oracle` using the `-k` flag. It can be
passed as either the key itself in plaintext (not recommended outside of testing), or
as a filepath to a file containing the password.

```bash
/path/to/oracle start -c $HOME/vor/config.json -k /path/to/password/file
```

which is useful if running the `oracle` as a service using `systemd`.

### Flags

`-v`: Output `oracle` version info and exit:
```bash
oracle -v
```

`-k`: Pass the keystore decryption key as a filepath or plaintext key:

```bash
oracle start -k /path/to/password/file
oracle start -k fg1acljv8int8g5hutc3cr2kk24lpg2s
```

`-c`: Pass the filepath to `config.json`:

```bash
oracle start -c /home/user/vor/config.json
```

### Running the oracle as a service

It is recommended to run the `oracle` as a background service, for example using
`systemd`:

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
ExecStart=/home/USERNAME/go/bin/oracle start -c /path/to/config.json -k /path/to/decrypt/pass

[Install]
WantedBy=default.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl start vor_oracle
```

## oraclecli

The `oraclecli` acts as a client to run administrative tasks on the `oracle` daemon.
It allows an `oracle` operator to change their fees, withdraw earned fees from 
`VORCoordinator`, stop the `oracle`, and query your fees and withdrawable xFUND.

### help

Outputs a list of commands

```bash
oraclecli help
```

### settings

The `settings` command allows you to set the configuration options for the CLI.
Configuration is stored in `$HOME/.oracle-cli_settings.json`.

```bash
oraclecli settings
```

You can configure the host & port the CLI will use to communicate with a running `oracle`
and should equate to the values used in your `oracle`'s `config.json`

You can also set the **api key** which the CLI will need to authenticate when
communicating with the `oracle`. This is the same as you use to start the `oracle`

### about

Outputs data about your `oracle`, such as your `keyHash`, wallet address, IP/PORT etc.

```bash
oraclecli about
```

### changefee

Allows you to change the base fee for fulfilling requests. The base fee is used
for all consumers for whom you have not set a granular fee.

Fee should be `amount * 10^9`, for example, if you want a fee of 0.2 xFUND,
enter 200000000.

```bash
oraclecli changefee
```

### changegranularfee

Allows you to change the fee for fulfilling requests at a granular level for
a particular consumer contract address. This allows you to set fees dependent on the
Tx cost for fulfilling requests, and thus higher fees for more expensive consumer contracts.

Fee should be `amount * 10^9`, for example, if you want a fee of 0.2 xFUND,
enter 200000000.

You will additionally be prompted for the contract address you are
applying the granular fee to.

```bash
oraclecli changegranularfee
```

### queryfees

Query the fees you have currently set. Optionally pass a consumer contract address
to query fees at a granular level. If you have not set a granular fee for the input
contract address, or do not input a contract address, the base fee will be returned.

```bash
oraclecli queryfees
oraclecli queryfees 0xD833215cBcc3f914bD1C9ece3EE7BF8B14f841bb
```

### querywithdrawable

Query the amount of fees you have accumulated, and are currently held by the 
`VORCoordinator`. You can withdraw these at any time.

```bash
oraclecli querywithdrawable
```

### register

Register a new proving key with `VORCoordinator`. You will need to run this, for example
if you change key, or have generated a new key on the `oracle`'s first run.

```bash
oraclecli register
```

### stop

Stops the `oracle` daemon.

```bash
oraclecli stop
```

### withdraw

Withdraw any accumulated tokens held by the `VORCoordinator` to the selected 
recipient address. The recipient can be your Oracle's wallet address, or any other 
beneficiary you choose.

Withdrawal amount should be `amount * 10^9`, for example, if you wish to withdraw
2 xFUND, enter 2000000000. The amount must not exceed the value output by the
`querywithdrawable` command.

```bash
oraclecli withdraw
```
