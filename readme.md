# ochain

**ochain** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Development Testnet Launch Docker Command

```
docker run --name ochain-testnet -p 26656-26659:26656-26659 -p 6061:6061 -p 9091:9091 -p 1317:1317 onomy/ochain-testnet:latest ochaind start --rpc.laddr tcp://0.0.0.0:26657 --grpc.address 0.0.0.0:9091
```

## Launch

To launch your blockchain live on mutliple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

To use SPN you will need [Starport](https://github.com/tendermint/starport) installed. SPN is being actively developed, please, build Starport from source and use `develop` branch.

### Initiating a blockchain launch

To initiate a blockchain launch run the following command:

```
starport network chain create [chainID] [sourceURL]
```

`chainID` is a string that uniquely identifies your blockchain on SPN. `sourceURL` is a URL that can be used to clone the repository containing a Cosmos SDK blockchain node (for example, `https://github.com/tendermint/spn`). By running the `create` command you act as a "coordinator" and initiate the launch of a blockchain.

To start an Onomy chain the following command may be used
```
```
starport network chain create <your_chain_name> http://github.com/onomyprotocol/ochain
```

### Joining as a validator

Run the following command from a server to propose yourself as a validator:

```
starport network chain join [chainID]
```

Follow the prompts to provide information about the validator. Starport will download the source code of the blockchain node, build, initialize and create and send two proposals to SPN: to add an account and to add a validator with self-delegation. By running a `join` command you act as a "validator".

By default, a coordinator does not propose themselves as a validator. To do so, run `join` command and your proposals will be automatically approved.

Please visit the following site for more information such as current test-net or main-net chainID as well as genesis file:
https://github.com/onomyprotocol/onomy-launch

### Listing pending proposals

```
starport network proposal list [chainID]
```

This command lists all proposals. To filter the list of proposals use the `--status` flag (possible values are: `approved`, `pending` and `rejected`). Each proposal has a `proposalID` (integer, unique to the chain), this ID is used to approve and reject a proposal.

### Accepting and rejecting proposals

As a coordinator run the following command to approve proposals:

```
starport network proposal approve [chainID] 1,4,5,6
```

Replace comma-separated values with a list of `proposalID` being accepted. Replace `approve` with `reject` to reject proposals instead.

### Starting a blockchain node

Once validator proposals have been accepted, run the following command to start a blockchain node:

```
starport network chain start [chainID]
```

This command will use SPN to create a correct genesis file, configure and launch your blockchain node. Once the node is started and the required number of validators are online, you will see output with incrementing block height number, which means that the blockchain has been successfully started.

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/W8trcGV)

