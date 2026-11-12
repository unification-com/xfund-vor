# The oracle with a local chain: Go, anvil and cast, and the oracle built from this tree. Its default command
# runs the dev environment (docker/scripts/dev-env.sh); docker/scripts/test.sh runs the tests.
# GO_VERSION comes from oracle/go.mod (see the Makefile).
ARG GO_VERSION
FROM ghcr.io/foundry-rs/foundry:v1.5.1 AS foundry
FROM golang:${GO_VERSION}

COPY --from=foundry /usr/local/bin/anvil /usr/local/bin/cast /usr/local/bin/
RUN apt-get update && apt-get install -y --no-install-recommends jq && rm -rf /var/lib/apt/lists/*

# the integration tests' configs expect the repository here
WORKDIR /root/xfund-vor

# modules first, so a code change doesn't download them again
COPY ./oracle/go.mod ./oracle/go.sum ./oracle/
RUN cd oracle && go mod download

COPY ./abigenBindings ./abigenBindings/
COPY ./docker ./docker/
COPY ./oracle ./oracle/
RUN cd oracle && go build -o build/oracle .

EXPOSE 8545 8445
CMD ["docker/scripts/dev-env.sh"]
