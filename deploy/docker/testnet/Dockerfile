# Simple usage with a mounted data directory:
# > docker build -t ochain .
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.ochain:/ochain/.ochain ochain ochaind init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.ochain:/ochain/.ochain ochain ochaind start
FROM golang:1.16-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/onomyprotocol/ochain

# Add source files
COPY . .

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES
RUN make install

# Final image
FROM alpine:edge

ENV OCHAIN /ochain

# Install ca-certificates
RUN apk add --update ca-certificates

RUN addgroup ochain && \
    adduser -S -G ochain ochain -h "$OCHAIN"

USER ochain

WORKDIR $OCHAIN

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/ochaind /usr/bin/ochaind


RUN ochaind init --chain-id=ochain-testnet validator1 && \
    ochaind keys add validator1 && \
    ochaind add-genesis-account $(ochaind keys show validator1 -a) 10000000nom && \
    ochaind gentx validator1 10000000nom --chain-id ochain-testnet && \
    ochaind collect-gentxs


# Run ochaind by default, omit entrypoint to ease using container with ochaincli
CMD ["ochaind"]

