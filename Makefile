#!/usr/bin/make -f

DEFAULT_VERSION=0.0.1

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

# Nothing released yet - set a default
ifeq ($(strip $(VERSION)),)
VERSION=$(DEFAULT_VERSION)
endif

ldflags = -X oracle/version.Version=$(VERSION) \
		  -X oracle/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

# abigen is a tool of the oracle module (go.mod `tool`), so the bindings match the go-ethereum it builds with
ABIGEN := cd oracle && go tool abigen
abigen:
	npx truffle run abigen
	$(ABIGEN) --abi ../abigenBindings/abi/MockERC20.abi --pkg mock_erc20 --out contracts/mock_erc20/mock_erc20.go
	$(ABIGEN) --abi ../abigenBindings/abi/VORCoordinator.abi --pkg vor_coordinator --out contracts/vor_coordinator/vor_coordinator.go
	$(ABIGEN) --abi ../abigenBindings/abi/VORRandomnessRequestMock.abi --pkg vor_randomness_request_mock --out contracts/vor_randomness_request_mock/vor_randomness_request_mock.go
	$(ABIGEN) --abi ../abigenBindings/abi/VORD20.abi --pkg vord_20 --out contracts/vord_20/vord20.go
	$(ABIGEN) --abi ../abigenBindings/abi/BlockhashStore.abi --pkg block_hash_store --out contracts/block_hash_store/block_hash_store.go

build-oracle:
	cd oracle && rm -f build/oracle && go build -mod=readonly $(BUILD_FLAGS) -o ./build/oracle

build: build-oracle

install-oracle:
	cd oracle && go install -mod=readonly $(BUILD_FLAGS)

install: install-oracle

# 1. Create a new release tag on Github, e.g. v0.1.5
# 2. 'git checkout main && git pull' to get tag in local repo
# 3. run this target to generate archive & checksum for upload
build-release: build
	rm -rf "dist/vor-oracle_v${VERSION}"
	mkdir -p "dist/vor-oracle_v${VERSION}"
	cp oracle/build/oracle "dist/vor-oracle_v${VERSION}/oracle"
	cp docs/guide/oracle.md "dist/vor-oracle_v${VERSION}/README.md"
	cd dist && tar -cpzf "vor-oracle_linux_v${VERSION}.tar.gz" "vor-oracle_v${VERSION}"
	cd dist && sha256sum "vor-oracle_linux_v${VERSION}.tar.gz" > "checksum_v${VERSION}.txt"
	cd dist && sha256sum --check "checksum_v${VERSION}.txt"

lint:
	@cd oracle && find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -w -s

.PHONY: abigen build-oracle build install-oracle install build-release lint

# Tests

# the Go release the Docker image builds with: the go directive in oracle/go.mod
GO_VERSION := $(shell sed -n 's/^go //p' oracle/go.mod)

# Go, anvil and cast, with the oracle built from this tree
docker-image:
	docker build --build-arg GO_VERSION=$(GO_VERSION) -t vor_oracle -f docker/Dockerfile .

# unit and integration tests, then the end-to-end flow, on local chains
test-oracle: docker-image
	docker run --rm vor_oracle docker/scripts/test.sh

# a local chain with the contracts deployed, and an oracle serving it
dev-env: docker-image
	docker run -it --rm -p 8545:8545 -p 8445:8445 vor_oracle

.PHONY: docker-image test-oracle dev-env
