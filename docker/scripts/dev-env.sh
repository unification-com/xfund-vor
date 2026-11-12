#!/usr/bin/env bash
# A local chain (http://127.0.0.1:8545, chain id 696969) with the contracts deployed, xFUND for accounts 2-19,
# and an oracle serving it (its admin API on port 8445, its key in docker/assets/pass).
set -euo pipefail
source "$(dirname "$0")/lib.sh"

start_chain 8545 696969 /tmp/anvil.log
deploy_contracts http://127.0.0.1:8545
for i in $(seq 2 19); do
	send http://127.0.0.1:8545 "$XFUND" "transfer(address,uint256)" "$(cast wallet address --mnemonic "$MNEMONIC" --mnemonic-index "$i")" 100000000000
done
echo "xFUND sent to accounts 2-19"
cd "$ROOT/docker/assets"
exec "$ORACLE" start -c config.json -k pass
