#!/usr/bin/env bash
# The oracle's whole flow, on two local chains:
#  1. chain a (EIP-1559): a fresh oracle registers its proving key, and fulfils a dice roll on VORD20
#  2. chain b (legacy gas) added to the running set-up: the oracle resumes chain a where it left off, registers
#     on b, and fulfils a roll on each chain, one of them with the same request id on both
#  3. scan-from on the running oracle: chain b rescans (nothing recorded twice) and still fulfils, chain a skips
#     ahead, and an invalid target is refused
# Needs anvil, cast, jq and the oracle built at oracle/build/oracle. Exits non-zero on the first failure.
set -euo pipefail
source "$(dirname "$0")/lib.sh"

WORK=$(mktemp -d)
CHAIN_A=696969
RPC_A=http://127.0.0.1:8645
CHAIN_B=696970
RPC_B=http://127.0.0.1:8646
FEE=100000000
ROLLER1=0x0000000000000000000000000000000000000a01
ROLLER2=0x0000000000000000000000000000000000000a02
ORACLE_PID=

step() { printf '\n== %s\n' "$*"; }
fail() {
	echo "E2E FAILED: $*" >&2
	echo "--- oracle log (last 40 lines):" >&2
	tail -40 "$WORK/oracle.log" >&2 || true
	exit 1
}
cleanup() {
	[ -n "$ORACLE_PID" ] && kill "$ORACLE_PID" 2> /dev/null || true
	kill $(jobs -p) 2> /dev/null || true
}
trap cleanup EXIT

admin() { "$ORACLE" "$@" -c "$WORK/config.json"; }

# chain_json <name> <chain id> <rpc> [more settings]
chain_json() {
	echo "{\"name\":\"$1\",\"network_id\":$2,\"eth_http_host\":\"$3\",\"contract_address\":\"$VOR_COORDINATOR\"," \
		"\"blockhash_store_address\":\"$BLOCK_HASH_STORE\",\"check_duration\":1,\"wait_confirmations\":2,\"fee\":$FEE${4:+,$4}}"
}

# write_config <chain json>...
write_config() {
	local chains
	chains=$(IFS=,; echo "$*")
	cat > "$WORK/config.json" << JSON
{"chains": [$chains],
 "serve": {"host": "127.0.0.1", "port": 18445},
 "keystorage": {"file": "$WORK/keystore.json", "account": "oracle"},
 "database": {"dialect": "sqlite", "storage": "$WORK/oracle.db"}}
JSON
}

start_oracle() {
	"$ORACLE" start -c "$WORK/config.json" -k "$WORK/pass" >> "$WORK/oracle.log" 2>&1 &
	ORACLE_PID=$!
	wait_for 30 "the oracle's API" admin status || fail "the oracle didn't start"
}

stop_oracle() {
	admin stop > /dev/null
	wait "$ORACLE_PID" || fail "the oracle stopped with status $?"
	ORACLE_PID=
}

# wait_running <chain>: its worker runs once its proving key is registered
wait_running() { wait_for 90 "chain $1 to run" admin about --chain "$1" || fail "chain $1 didn't start"; }

provider() { cast call --rpc-url "$1" "$VOR_COORDINATOR" "getProviderAddress(bytes32)(address)" "$2"; }

# set_up_consumer <rpc> <key hash>: VORD20 asks this oracle for randomness, with xFUND to pay for it
set_up_consumer() {
	send "$1" "$XFUND" "transfer(address,uint256)" "$VORD20" 1000000000000
	send "$1" "$VORD20" "increaseVorAllowance(uint256)" 1000000000000
	send "$1" "$VORD20" "setKeyHash(bytes32)" "$2"
	send "$1" "$VORD20" "setFee(uint256)" "$FEE"
}

roll() { send "$1" "$VORD20" "rollDice(uint256,address)" "$2" "$3"; }
landed() { cast call --rpc-url "$1" "$VORD20" "house(address)(string)" "$2"; }
wait_landed() { wait_for 90 "the roll for $3 on chain $1" landed "$2" "$3" || fail "the roll for $3 on chain $1 wasn't fulfilled"; }

# requests <chain> [status]: the oracle's records of the chain's requests, oldest first
requests() { admin queryrequests --chain "$1" --order asc --limit 100 ${2:+--status "$2"}; }
count() { requests "$@" | jq '.pagination.num_records'; }

# wait_fulfilled <chain> <n>: the oracle records a fulfilment when it reads the fulfilled event, a poll after
# the transaction is mined
wait_fulfilled() {
	wait_for 30 "chain $1's oracle records to show $2 fulfilled" bash -c "[ \"\$($ORACLE queryrequests -c $WORK/config.json --chain $1 --status 4 | jq .pagination.num_records)\" = $2 ]" ||
		fail "chain $1: the oracle records $(count "$1" 4) fulfilled requests, not $2"
	echo "ok: chain $1's fulfilled requests = $2"
}

# fulfil_tx_type <chain> <rpc>: the type of the chain's first recorded fulfilment (2: EIP-1559, 0: legacy)
fulfil_tx_type() {
	printf '%d' "$(cast tx --rpc-url "$2" --json "$(requests "$1" 4 | jq -r '.requests[0].fulfill_tx_hash')" | jq -r .type)"
}

# expect <what> <got> <want>
expect() { [ "$2" = "$3" ] || fail "$1: got $2, want $3"; echo "ok: $1 = $3"; }

step "chain a ($CHAIN_A): start and deploy the contracts"
start_chain 8645 "$CHAIN_A" "$WORK/anvil-a.log"
deploy_contracts "$RPC_A"

step "oracle: API key, start on chain a, registration"
cp "$ROOT/docker/assets/keystore.json" "$ROOT/docker/assets/pass" "$WORK/"
write_config "$(chain_json a "$CHAIN_A" "$RPC_A")"
admin api-key
start_oracle
wait_running a
KEY_HASH=$(admin about --chain a | awk '/KeyHash:/ {print $2}')
expect "chain a's provider for key $KEY_HASH" "$(provider "$RPC_A" "$KEY_HASH")" "$ORACLE_ADDRESS"

step "chain a: a dice roll is fulfilled"
set_up_consumer "$RPC_A" "$KEY_HASH"
roll "$RPC_A" 1001 "$ROLLER1"
wait_landed a "$RPC_A" "$ROLLER1"
echo "landed: $(landed "$RPC_A" "$ROLLER1")"
wait_fulfilled a 1
FIRST_REQUEST=$(requests a 4 | jq -r '.requests[0].request_id')
expect "chain a's fulfilment transaction type (EIP-1559)" "$(fulfil_tx_type a "$RPC_A")" 2

step "chain b ($CHAIN_B, legacy gas): start, deploy, add to the oracle"
stop_oracle
start_chain 8646 "$CHAIN_B" "$WORK/anvil-b.log"
deploy_contracts "$RPC_B"
write_config "$(chain_json a "$CHAIN_A" "$RPC_A")" "$(chain_json b "$CHAIN_B" "$RPC_B" '"legacy_gas":true')"
start_oracle
wait_running a
wait_running b
expect "chain b's provider" "$(provider "$RPC_B" "$KEY_HASH")" "$ORACLE_ADDRESS"
RESUMED_AT=$(grep 'begin polling' "$WORK/oracle.log" | grep 'chain=a' | tail -1 | grep -oE 'from_block=[0-9]+' | cut -d= -f2)
[ "${RESUMED_AT:-0}" -gt 1 ] || fail "chain a restarted from block ${RESUMED_AT:-?}, not its saved cursor"
echo "ok: chain a resumed from its saved cursor, block $RESUMED_AT"
expect "chain a's requests after the restart" "$(count a)" 1

step "both chains: rolls fulfilled, one request id on both"
set_up_consumer "$RPC_B" "$KEY_HASH"
# the same consumer, nonce, seed and key as chain a's first roll: the same request id
roll "$RPC_B" 1001 "$ROLLER1"
roll "$RPC_A" 1002 "$ROLLER2"
wait_landed b "$RPC_B" "$ROLLER1"
wait_landed a "$RPC_A" "$ROLLER2"
wait_fulfilled a 2
wait_fulfilled b 1
expect "chain b's request id (the same as chain a's first)" "$(requests b 4 | jq -r '.requests[0].request_id')" "$FIRST_REQUEST"
expect "chain b's fulfilment transaction type (legacy)" "$(fulfil_tx_type b "$RPC_B")" 0

step "scan-from: chain b rescans from block 1, chain a skips to its head, a bad target is refused"
admin scan-from 1 --chain b
wait_for 60 "chain b to rescan from block 1" grep -q 'msg=scanned chain=b from_block=1 ' "$WORK/oracle.log" ||
	fail "chain b didn't rescan from block 1"
expect "chain b's requests after the rescan (none recorded twice)" "$(count b)" 1
roll "$RPC_B" 1003 "$ROLLER2"
wait_landed b "$RPC_B" "$ROLLER2"
wait_fulfilled b 2
admin scan-from head --chain a
admin scan-from head-999999999 --chain a > /dev/null 2>&1 && fail "an invalid scan target was accepted"
echo "ok: an invalid scan target is refused"

step "stop"
stop_oracle
echo
echo "E2E PASSED"
