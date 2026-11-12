#!/usr/bin/env bash
# Deploys the dev and test contracts to a fresh local chain: deploy-contracts.sh [rpc-url]
set -euo pipefail
source "$(dirname "$0")/lib.sh"
deploy_contracts "${1:-http://127.0.0.1:8545}"
