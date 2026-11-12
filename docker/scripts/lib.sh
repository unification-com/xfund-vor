# Shared by the dev environment, the tests and the end-to-end flow. Source it; it needs anvil, cast and jq.

ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)
ORACLE=$ROOT/oracle/build/oracle

# ganache's deterministic accounts, which the dev and test configs, keystores and docs use: 0 deploys and owns
# the contracts, 1 is the oracle
MNEMONIC="myth like bonus scare over problem client lizard pioneer submit female collect"
OWNER_KEY=0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d
ORACLE_ADDRESS=0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0

# where deploy_contracts puts them (account 0's first five nonces, in truffle's migration order)
XFUND=0xe78A0F7E598Cc8b0Bb87894B0F60dD2a88d6a8Ab
BLOCK_HASH_STORE=0x5b1869D9A4C187F2EAa108f3062412ecf0526b24
VOR_COORDINATOR=0xCfEB869F69431e42cdB54A4F4f105C19C080A601
VORD20=0x254dffcd3277C0b1660F6d42EFbB754edaBAbC2B
REQUEST_MOCK=0xC89Ce4735882C9F0f0FE26686c53074E09B0D550

# start_chain <port> <chain id> <log file> [instant]: a fresh local chain, ready when this returns. It mines a
# block a second, so an oracle's confirmations accrue; "instant" mines each transaction as it arrives instead,
# as ganache did for the integration tests.
start_chain() {
	local mining=(--block-time 1)
	[ "${4:-}" = instant ] && mining=()
	anvil --host 0.0.0.0 --port "$1" --chain-id "$2" --mnemonic "$MNEMONIC" --accounts 20 "${mining[@]}" > "$3" 2>&1 &
	wait_for 30 "chain $2" cast chain-id --rpc-url "http://127.0.0.1:$1"
}

# wait_for <seconds> <what> <command...>: runs the command every second until it succeeds
wait_for() {
	local seconds=$1 what=$2
	shift 2
	for ((i = 0; i < seconds; i++)); do
		if "$@" > /dev/null 2>&1; then
			return 0
		fi
		sleep 1
	done
	echo "timed out after ${seconds}s waiting for $what" >&2
	return 1
}

# send <rpc> <to> <signature> [args...]: a transaction from the owner, mined
send() {
	local rpc=$1
	shift
	cast send --rpc-url "$rpc" --private-key "$OWNER_KEY" "$@" > /dev/null
}

# deploy_contracts <rpc>: the contracts, from the compiled bytecode in abigenBindings/bin, at the addresses above
deploy_contracts() {
	local rpc=$1
	_deploy "$rpc" MockERC20 "$XFUND" "constructor(string,string,uint256)" xFUND xFUND 1000000000000000000000000000
	_deploy "$rpc" BlockhashStore "$BLOCK_HASH_STORE"
	_deploy "$rpc" VORCoordinator "$VOR_COORDINATOR" "constructor(address,address)" "$XFUND" "$BLOCK_HASH_STORE"
	_deploy "$rpc" VORD20 "$VORD20" "constructor(address,address,bytes32,uint256)" "$VOR_COORDINATOR" "$XFUND" \
		"$(cast format-bytes32-string keyHash)" 1000000000000000000
	_deploy "$rpc" VORRandomnessRequestMock "$REQUEST_MOCK"
}

_deploy() {
	local rpc=$1 name=$2 expected=$3 code address
	shift 3
	code=0x$(cat "$ROOT/abigenBindings/bin/$name.bin")
	if [ $# -gt 0 ]; then
		local signature=$1
		shift
		code=$code$(cast abi-encode "$signature" "$@" | cut -c3-)
	fi
	address=$(cast send --rpc-url "$rpc" --private-key "$OWNER_KEY" --json --create "$code" | jq -r .contractAddress)
	if [ "$(cast to-check-sum-address "$address")" != "$expected" ]; then
		echo "$name was deployed at $address, not $expected: deploy to a fresh chain" >&2
		return 1
	fi
	echo "deployed $name at $expected"
}
