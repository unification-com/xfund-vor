#!/usr/bin/make -f

abigen:
	npx truffle run abigen
	abigen --abi abigenBindings/abi/MockERC20.abi --pkg mock_erc20 --out oracle/contracts/mock_erc20/mock_erc20.go
	abigen --abi abigenBindings/abi/VORCoordinator.abi --pkg vor_coordinator --out oracle/contracts/vor_coordinator/vor_coordinator.go
	abigen --abi abigenBindings/abi/VORRandomnessRequestMock.abi --pkg vor_randomness_request_mock --out oracle/contracts/vor_randomness_request_mock/vor_randomness_request_mock.go
	abigen --abi abigenBindings/abi/VORD20.abi --pkg vord_20 --out oracle/contracts/vord_20/vord20.go

build-oracle:
	cd oracle && rm -f build/oracle && go build -mod=readonly -o ./build/oracle

build-oracle-cli:
	cd oracle-cli && rm -f build/oraclecli && go build -mod=readonly -o ./build/oraclecli

build: build-oracle build-oracle-cli

install-oracle:
	cd oracle && go install -mod=readonly

install-oracle-cli:
	cd oracle-cli && go install -mod=readonly

install: install-oracle install-oracle-cli

.PHONY: abigen build-oracle build-oracle-cli build install-oracle install-oracle-cli install
