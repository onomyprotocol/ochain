#!/bin/sh
#
# Sifchain: Mainnet Genesis Entrypoint.
#

#
# Configure the node.
#
setup() {
  onomygen node create "$CHAINNET" "$MONIKER" "$MNEMONIC" --peer-address "$PEER_ADDRESSES" --genesis-url "$GENESIS_URL" --bind-ip-address "$BIND_IP_ADDRESS" --with-cosmovisor
}

#
# Run the node under cosmovisor.
#
run() {
  cosmovisor start --rpc.laddr tcp://0.0.0.0:26657 --minimum-gas-prices "$GAS_PRICE"
}

setup
run
