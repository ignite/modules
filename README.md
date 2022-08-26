<div align="center">
  <h1> Ignite Modules </h1>
</div>

<div align="center">
  <a href="https://github.com/cosmos/cosmos-sdk/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/cosmos/cosmos-sdk.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/ignite/modules?tab=doc">
    <img alt="GoDoc" src="https://pkg.go.dev/badge/github.com/ignite/modules.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/ignite/modules">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/ignite/modules" />
  </a>
<!--
  <a href="https://codecov.io/gh/ignite/modules">
    <img alt="Code Coverage" src="https://codecov.io/gh/ignite/modules/branch/develop/graph/badge.svg" />
  </a>
-->
</div>
<div align="center">
  <a href="https://github.com/ignite/modules">
    <img alt="Lines Of Code" src="https://tokei.rs/b1/github/ignite/modules" />
  </a>
    <img alt="Test Status" src="https://github.com/ignite/modules/workflows/Test/badge.svg" />
    <img alt="Lint Status" src="https://github.com/ignite/modules/workflows/Lint/badge.svg" />
</div>

**Ignite Modules** contains Cosmos SDK modules developed by Ignite for common uses of cosmos blockchains.
Modules are aimed to be generic and can be imported in any project depending on their blockchain functionalities.

### ⚠️ Disclaimer

Modules are under active development and should not be currently considered for production.

## Modules

- [`claim`](x/claim): this module can be used by blockchains that wish to offer airdrops to eligible addresses upon the completion of specific actions. Eligible addresses with airdrop allocations are listed in the genesis state of the module. Initial claim, staking and voting missions are natively supported. The developer can add custom missions related to their blockchain functionality.


- [`mint`](x/mint): this module is an enhanced version of [Cosmos SDK `mint` module](https://docs.cosmos.network/master/modules/mint/) where developers can use the minted coins from inflations for specific purposes other than staking rewards.

## Testing

The repository comes with a sample Cosmos SDK app to test the different module features.

In order to launch a test app instance and interact with the modules, install [Ignite CLI](https://ignite.com) and run the following command under the repository:

```
ignite chain serve
```

You can interact with the modules with the native chain CLI: `testappd`

```
testappd q mint params
```

## Contributing

We welcome contributions from everyone. You can create a branch from `main` and create a pull request, or maintain your own fork and submit a cross-repository pull request.

**Important** Before you start implementing a new feature or making a fix, the first step is to create an issue on GitHub that describes the proposed changes.

## Community

Ignite Modules is a free and open source project maintained by [Ignite](https://ignite.com). Here's where you can find us. 

- [@ignite_dev on Twitter](https://twitter.com/ignite_dev)
- [Ignite Blog](https://ignite.com/blog/)
- [Ignite Discord](https://discord.com/invite/ignite)
- [Ignite Docs](https://docs.ignite.com/)
- [Ignite Jobs](https://ignite.com/careers)
- [Cosmos SDK Docs](https://docs.cosmos.network)
- [Cosmos Academy](https://tutorials.cosmos.network/academy/0-welcome/)
