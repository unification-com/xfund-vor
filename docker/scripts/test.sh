#!/usr/bin/env bash
# The oracle's unit and integration tests, against a local chain with the test contracts, then the end-to-end
# flow.
set -euo pipefail
source "$(dirname "$0")/lib.sh"

start_chain 8545 696969 /tmp/anvil-tests.log instant
deploy_contracts http://127.0.0.1:8545
# one package at a time: the integration tests in several packages share the chain's contracts and state
(cd "$ROOT/oracle" && go test -p 1 ./...)
"$ROOT/docker/scripts/e2e.sh"
