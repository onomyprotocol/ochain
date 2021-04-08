# Simple usage with a mounted data directory:
# > docker build -t ochain .
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.ochain:/ochain/.ochain ochain ochaind init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.ochain:/ochain/.ochain ochain ochaind start
FROM golang:1.16-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/tendermint

# Add source files
RUN git clone https://github.com/tendermint/starport.git
WORKDIR /go/src/github.com/tendermint/starport

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES
RUN make install

# Final image
FROM alpine:edge

ENV STARPORT /starport

# Install ca-certificates
RUN apk add --update ca-certificates

RUN addgroup starport && \
    adduser -S -G starport starport -h "$STARPORT"

USER starport

WORKDIR $STARPORT

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/starport /usr/bin/starport


RUN ochaind init --chain-id=test-net validator1 && \
    ochaind keys add validator1 && \
    ochaind add-genesis-account $(ochaind keys show validator1 -a) 100000000stake,100000000validatortoken && \
    ochaind gentx validator1 100000000stake --chain-id test-net && \
    ochaind collect-gentxs


# Run ochaind by default, omit entrypoint to ease using container with ochaincli
CMD ["ochaind"]

