<!--
order: 1
-->

# What is Ochain?

`ochain` is the name of the Cosmos SDK application for the Onomy. It comes with 2 main entrypoints:

- `ochaind`: The Ochain Daemon and command-line interface (CLI). runs a full-node of the `ochain` application.

`ochain` is built on the Cosmos SDK using the following modules:

- `x/auth`: Accounts and signatures.
- `x/bank`: Token transfers.
- `x/staking`: Staking logic.
- `x/mint`: Inflation logic.
- `x/distribution`: Fee distribution logic.
- `x/slashing`: Slashing logic.
- `x/gov`: Governance logic.
- `ibc-go/modules`: Inter-blockchain communication. Hosted in the `cosmos/ibc-go` repository. 
- `x/params`: Handles app-level parameters.

About the Onomy: The Onomy is the first Hub to be launched in the Cosmos Network. The role of a Hub is to facilitate transfers between blockchains. If a blockchain connects to a Hub via IBC, it automatically gains access to all the other blockchains that are connected to it. The Onomy is a public Proof-of-Stake chain. Its staking token is called the Atom.

Next, learn how to [install Ochain](./installation.md).
