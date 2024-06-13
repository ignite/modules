# `mint`

## Abstract

This module is an enhanced version of [Cosmos SDK `mint` module](https://docs.cosmos.network/master/modules/mint/) where developers can use the minted coins from inflations for specific purposes other than staking rewards.

The developer can define proportions for minted coins purpose:

- Staking rewards
- Community pool
- Funded addresses

In the future, the module will suport defining custom purpose for minted coins.

## State

The state of the module indexes the following values:

- `Minter`: the minter is a space for holding current inflation information
- `Params`: parameter of the module

```
Minter: [] -> Minter
Params: [] -> Params
```

### `Minter`

`Minter` holds current inflation information, it contains the annual inflation rate, and the annual expected provisions

```proto
message Minter {
  string inflation = 1 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string annual_provisions = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
}
```

### `Params`

Described in **[Parameters](03_params.md)**

## Begin-block

Begin-block contains the logic to:

- recalculate minter parameters
- mint new coins
- distribute new coins depending on distribution proportions

### Pseudo-code

```go
minter = load(Minter)
params = load(Params)
minter = calculateInflationAndAnnualProvision(params)
store(Minter, minter)

mintedCoins = minter.BlockProvision(params)
Mint(mintedCoins)

DistributeMintedCoins(mintedCoin)
```

The inflation rate calculation follows the same logic as the [Cosmos SDK `mint` module](https://github.com/cosmos/cosmos-sdk/tree/main/x/mint#inflation-rate-calculation)

## Parameters

The parameters of the module contain information about inflation, and distribution of minted coins.

- `mint_denom`: the denom of the minted coins
- `inflation_rate_change`: maximum annual change in inflation rate
- `inflation_max`: maximum inflation rate
- `inflation_min`: minimum inflation rate
- `goal_bonded`: goal of percent bonded coins
- `blocks_per_year`: expected blocks per year
- `distribution_proportions`: distribution_proportions defines the proportion for minted coins distribution
- `funded_addresses`: list of funded addresses

```proto
message Params {
  string mint_denom = 1;
  string inflation_rate_change = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string inflation_max = 3 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string inflation_min = 4 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string goal_bonded = 5 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  uint64 blocks_per_year = 6;
  DistributionProportions distribution_proportions = 7 [(gogoproto.nullable) = false];
  repeated WeightedAddress funded_addresses = 8 [(gogoproto.nullable) = false];
}
```

### `DistributionProportions`

`DistributionProportions` contains propotions for the distributions.

```proto
message DistributionProportions {
  string staking = 1 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string funded_addresses = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
  string community_pool = 3 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
}
```

### `WeightedAddress`

`WeightedAddress` is an address with an associated weight to receive part the minted coins depending on the `funded_addresses` distribution proportion.

```proto
message WeightedAddress {
  string address = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string weight  = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
}
```

## Events

### `EventMint`

This event is emitted when new coins are minted. The event contains the amount of coins minted with the parameters of the minter at the current block.

```protobuf
message EventMint {
  string bondedRatio = 1 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar) = "cosmos.Dec"
  ];
  string inflation = 2 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar) = "cosmos.Dec"
  ];
  string annualProvisions = 3 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar) = "cosmos.Dec"
  ];
  string amount = 4 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (cosmos_proto.scalar) = "cosmos.Int"
  ];
}
```

## Client

### Query

The `query` commands allow users to query `mint` state.

```sh
testappd q mint
```

#### `params`

Shows the params of the module.

```sh
testappd q mint params
```

Example output:

```yml
blocks_per_year: "6311520"
distribution_proportions:
  community_pool: "0.300000000000000000"
  funded_addresses: "0.400000000000000000"
  staking: "0.300000000000000000"
funded_addresses:
  - address: cosmos1ezptsm3npn54qx9vvpah4nymre59ykr9967vj9
    weight: "0.400000000000000000"
  - address: cosmos1aqn8ynvr3jmq67879qulzrwhchq5dtrvh6h4er
    weight: "0.300000000000000000"
  - address: cosmos1pkdk6m2nh77nlaep84cylmkhjder3areczme3w
    weight: "0.300000000000000000"
goal_bonded: "0.670000000000000000"
inflation_max: "0.200000000000000000"
inflation_min: "0.070000000000000000"
inflation_rate_change: "0.130000000000000000"
mint_denom: stake
```

#### `annual-provisions`

Shows the current minting annual provisions valu

```sh
testappd q mint annual-provisions
```

Example output:

```yml
52000470.516851147993560400
```

#### `inflation`

Shows the current minting inflation value

```sh
testappd q mint inflation
```

Example output:

```yml
0.130001213701730800
```
