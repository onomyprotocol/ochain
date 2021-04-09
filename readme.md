# ochain

**ochain** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Launch

To launch your blockchain live on mutliple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)

## Development testnet docker command
`docker run --name ochain-testnet -p 26656-26659:26656-26659 -p 6061:6061 -p 9091:9091 -p 1317:1317 onomy/ochain-testnet:latest ochaind start --rpc.laddr tcp://0.0.0.0:26657 --grpc.address 0.0.0.0:9091`