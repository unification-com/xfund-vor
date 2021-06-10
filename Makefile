#!/usr/bin/make -f

DEFAULT_VERSION=0.0.1

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

# Nothing released yet - set a default
ifeq ($(strip $(VERSION)),)
VERSION=$(DEFAULT_VERSION)
endif

ldflags = -X oracle/version.Version=$(VERSION) \
		  -X oracle/version.Commit=$(COMMIT) \
          -X oraclecli/version.Version=$(VERSION) \
          -X oraclecli/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

abigen:
	npx truffle run abigen
	abigen --abi abigenBindings/abi/MockERC20.abi --pkg mock_erc20 --out oracle/contracts/mock_erc20/mock_erc20.go
	abigen --abi abigenBindings/abi/VORCoordinator.abi --pkg vor_coordinator --out oracle/contracts/vor_coordinator/vor_coordinator.go
	abigen --abi abigenBindings/abi/VORRandomnessRequestMock.abi --pkg vor_randomness_request_mock --out oracle/contracts/vor_randomness_request_mock/vor_randomness_request_mock.go
	abigen --abi abigenBindings/abi/VORD20.abi --pkg vord_20 --out oracle/contracts/vord_20/vord20.go

build-oracle:
	cd oracle && rm -f build/oracle && go build -mod=readonly $(BUILD_FLAGS) -o ./build/oracle

build-oracle-cli:
	cd oracle-cli && rm -f build/oraclecli && go build -mod=readonly $(BUILD_FLAGS) -o ./build/oraclecli

build: build-oracle build-oracle-cli

install-oracle:
	cd oracle && go install -mod=readonly $(BUILD_FLAGS)

install-oracle-cli:
	cd oracle-cli && go install -mod=readonly $(BUILD_FLAGS)

install: install-oracle install-oracle-cli

# 1. Create a new release tag on Github, e.g. v0.1.5
# 2. 'git checkout main && git pull' to get tag in local repo
# 3. run this target to generate archive & checksum for upload
build-release: build
	rm -rf "dist/vor-oracle_v${VERSION}"
	mkdir -p "dist/vor-oracle_v${VERSION}"
	cp oracle/build/oracle "dist/vor-oracle_v${VERSION}/oracle"
	cp oracle-cli/build/oraclecli "dist/vor-oracle_v${VERSION}/oraclecli"
	cp docs/guide/oracle.md "dist/vor-oracle_v${VERSION}/README.md"
	cd dist && tar -cpzf "vor-oracle_linux_v${VERSION}.tar.gz" "vor-oracle_v${VERSION}"
	cd dist && sha256sum "vor-oracle_linux_v${VERSION}.tar.gz" > "checksum_v${VERSION}.txt"
	cd dist && sha256sum --check "checksum_v${VERSION}.txt"

lint:
	@cd oracle && find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -w -s
	@cd oracle-cli && find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -w -s

.PHONY: abigen build-oracle build-oracle-cli build install-oracle install-oracle-cli install build-release lint

# Tests

test-oracle:
	docker build -t oracle_test -f docker/test.Dockerfile .
	docker run -it oracle_test

# Dev environment
dev-env:
	docker build -t vor_dev_env -f docker/dev.Dockerfile .
	docker run -it -p 8545:8545 -p 8445:8445 vor_dev_env

.PHONY: test-oracle dev-env
