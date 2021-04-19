FROM ubuntu:bionic

RUN \
  apt-get update && \
  apt-get upgrade -y && \
  apt-get install -y curl build-essential nano netcat

ENV NVM_DIR /root/.nvm
ENV NODE_VERSION 12.18.3
ENV PATH=$PATH:/usr/local/go/bin:/root/.nvm/versions/node/v$NODE_VERSION/bin

# Install nvm, node, npm and yarn
RUN curl -sL https://raw.githubusercontent.com/creationix/nvm/v0.35.3/install.sh | bash && \
  . $NVM_DIR/nvm.sh && \
  nvm install $NODE_VERSION && \
  npm install --global yarn

# install Go
RUN curl -sL https://golang.org/dl/go1.16.3.linux-amd64.tar.gz -o /root/go1.16.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf /root/go1.16.3.linux-amd64.tar.gz

RUN mkdir -p /root/xfund-vor
WORKDIR /root/xfund-vor

# first, copy only essential files required for compiling contracts
COPY ./contracts ./contracts/
COPY ./migrations ./migrations/
COPY ./package.json ./yarn.lock ./truffle-config.js Makefile ./

# install node dependencies & compile contracts
RUN yarn install --frozen-lockfile && \
    npx truffle compile

# copy only go.mod and go.sum then pre-cache module dependencies
# to save future downloads when modifying code/tests
COPY ./oracle/go.mod ./oracle/go.sum ./oracle/
RUN cd /root/xfund-vor/oracle && go mod download -x

# finally, copy the oracle source. Unless go.mod is changed, only this
# layer will need to be run in future when code changes are made to the oracle.
COPY ./oracle ./oracle/

# default cmd to set up Ganache, deploy contracts and run go test
CMD cd /root/xfund-vor && \
    npx ganache-cli --deterministic --networkId 696969 --accounts 20 --quiet & \
    until nc -z 127.0.0.1 8545; do sleep 0.5; echo "wait for ganache"; done && \
    echo "deploying contracts, please wait..." && \
    npx truffle deploy --network=develop && \
    cd /root/xfund-vor/oracle && \
    echo "running tests, please wait..." && \
    go test ./...
