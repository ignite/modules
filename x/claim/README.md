# `claim`

## Abstract

This document specifies the `claim` module developed by Ignite.

This module can be used by blockchains that wish to offer airdrops to eligible addresses upon the completion of specific actions.

Eligible addresses with airdrop allocations are listed in the genesis state of the module.

Initial claim, staking, and voting missions are natively supported. The developer can add custom missions related to their blockchain functionality. The `CompleteMission` method exposed by the module keeper can be used for blockchain specific missions.

## State

The state of the module stores data for the three following properties:

- Who: what is the list of eligible addresses, what is the allocation for each eligible address
- How: what are the missions to claim airdrops
- When: when does decaying for the airdrop start, and when does the airdrop end

The state of the module indexes the following values:

- `AirdropSupply`: the amount of tokens that remain for the airdrop
- `InitialClaim`: information about an initial claim, a portion of the airdrop that can be claimed without completing a specific task
- `ClaimRecords`: list of eligible addresses with allocated airdrop, and the current status of completed missions
- `Missions`: the list of missions to claim airdrop with their associated weight
- `Params`: parameter of the module

```
AirdropSupply:  [] -> sdk.Int
InitialClaim:   [] -> InitialClaim

ClaimRecords:   [address] -> ClaimRecord
Missions:       [id] -> Mission

Params:         [] -> Params
```

### `InitialClaim`

`InitialClaim` determines the rules for the initial claim, a portion of the airdrop that can be directly claimed without completing a specific task. The mission is completed by sending a `MsgClaim` message.

The structure determines if the initial claim is enabled for the chain, and what mission is completed when sending `MsgClaim`.

```protobuf
message InitialClaim {
  bool   enabled   = 1;
  uint64 missionID = 2;
}
```

### `ClaimRecord`

`ClaimRecord` contains information about an address eligible for airdrop, what amount the address is eligible for, and which missions have already been completed and claimed.

```protobuf
message ClaimRecord {
  string address   = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  string claimable = 2 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (cosmos_proto.scalar)  = "cosmos.Int"
  ];
  repeated uint64 completedMissions = 3;
  repeated uint64 claimedMissions = 4;
}
```

### `Mission`

`Mission` represents a mission to be completed to claim a percentage of the airdrop supply.

```protobuf
message Mission {
  uint64 missionID   = 1;
  string description = 2;
  string weight      = 3 [
    (gogoproto.nullable)   = false,
    (gogoproto.customtype) = "cosmossdk.io/math.LegacyDec",
    (cosmos_proto.scalar)  = "cosmos.Dec"
  ];
}
```

### `Params`

Described in **[Parameters](05_params.md)**

## Messages

### `MsgClaim`

Claim completed mission amount for airdrop

```protobuf
message MsgClaim {
  string claimer = 1;
  uint64 missionID = 2;
}

message MsgClaimResponse {
  string claimed = 1 [
    (gogoproto.nullable) = false,
    (gogoproto.customtype) = "cosmossdk.io/math.Int",
    (cosmos_proto.scalar) = "cosmos.Int"
  ];
}
```

**State transition**

- Complete the claim for the mission and address
- Transfer the claim amount to the claimer balance

**Fails if**

- Mission is not completed
- The mission doesn't exist
- The claimer is not eligible
- The airdrop start time not reached
- The mission has already been claimed

## Methods

### `CompleteMission`

Complete a mission for an eligible address.
This method can be used by an external chain importing `claim` in order to define customized mission for the chain.

```go
CompleteMission(
    ctx sdk.Context,
    missionID uint64,
    address string,
) error
```

**State transition**

- Complete the mission `missionID` in the claim record `address`

**Fails if**

- The mission doesn't exist
- The address has no claim record
- The mission has already been completed for the address

### `ClaimMission`

Claim mission for an eligible claim record and mission id.
This method can be used by an external module importing `claim` in order to define customized mission claims for the
chain.

```go
ClaimMission(
    ctx sdk.Context,
    claimRecord types.ClaimRecord,
    missionID uint64,
) error
```

**State transition**

- Transfer the claim amount related to the mission

**Fails if**

- The mission doesn't exist
- The address has no claim record
- The airdrop start time not reached
- The mission has not been completed for the address
- The mission has already been claimed for the address

## End-blocker

The end-blocker of the module verifies if the airdrop supply is non-null, decay is enabled and decay end has been reached.
Under these conditions, the remaining airdrop supply is transferred to the community pool.

### Pseudo-code

```go
airdropSupply = load(AirdropSupply)
decayInfo = load(Params).DecayInformation

if airdropSupply > 0 && decayInfo.Enabled && BlockTime > decayInfo.DecayEnd
    distrKeeper.FundCommunityPool(airdropSupply)
    airdropSupply = 0
    store(AirdropSupply, airdropSupply)
```

## Parameters

The parameters of the module contain information about time-based decay for the airdrop supply.

```protobuf
message Params {
  DecayInformation decayInformation = 1 [(gogoproto.nullable) = false];
  google.protobuf.Timestamp airdropStart = 2
  [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
}
```

### `DecayInformation`

This parameter determines if the airdrop starts to decay at a specific time.

```protobuf
message DecayInformation {
  bool enabled = 1;
  google.protobuf.Timestamp decayStart = 2 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
  google.protobuf.Timestamp decayEnd = 3 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
}
```

When enabled, the claimable amount for each eligible address will start to decrease starting from `decayStart` till `decayEnd` where the airdrop is ended and the remaining fund is transferred to the community pool.

The decrease is linear.

If `decayStart == decayEnd`, there is no decay for the airdrop but the airdrop ends at `decayEnd` and the remaining fund is transferred to the community pool.

### `AirdropStart`

This parameter determines the airdrop start time.
When set, the user cannot claim the airdrop after completing the mission. The airdrop will be available only after the block time reaches the airdrop start time.
If the mission was completed, the user could call the `MsgClaim` to claim an airdrop from a completed mission.
The claim will be called automatically if the mission is completed and the airdrop start time is already reached.

## Events

### `EventMissionCompleted`

This event is emitted when a mission `missionID` is completed for a specific eligible `address`.

```protobuf
message EventMissionCompleted {
  uint64 missionID = 1;
  string address = 2;
}
```

### `EventMissionClaimed`

This event is emitted when a mission `missionID` is claimed for a specific eligible address `claimer`.

```protobuf
message EventMissionClaimed {
  uint64 missionID = 1;
  string claimer = 2;
}
```

## Client

A user can query and interact with the `claim` module using the chain CLI.

### Query

The `query` commands allow users to query `claim` state.

```sh
testappd q claim
```

#### `params`

Shows the params of the module.

```sh
testappd q claim params
```

Example output:

```yml
params:
  decayInformation:
    decayEnd: "1970-01-01T00:00:00Z"
    decayStart: "1970-01-01T00:00:00Z"
    enabled: false
```

#### `show-airdrop-supply`

Shows the current airdrop supply.

```sh
testappd q claim show-airdrop-supply
```

Example output:

```yml
AirdropSupply:
  amount: "1000"
  denom: drop
```

#### `show-initial-claim`

Shows the information about the initial claim for airdrops.

```sh
testappd q claim show-initial-claim
```

Example output:

```yml
InitialClaim:
  enabled: true
  missionID: "0"
```

#### `list-claim-record`

Lists the claim records for eligible addresses for the aidrops.

```sh
testappd q claim list-claim-record
```

Example output:

```yml
claimRecord:
  - address: cosmos1aqn8ynvr3jmq67879qulzrwhchq5dtrvh6h4er
    claimable: "500"
    completedMissions: []
  - address: cosmos1ezptsm3npn54qx9vvpah4nymre59ykr9967vj9
    claimable: "400"
    completedMissions: []
  - address: cosmos1pkdk6m2nh77nlaep84cylmkhjder3areczme3w
    claimable: "100"
    completedMissions: []
pagination:
  next_key: null
  total: "0"
```

#### `show-claim-record`

Shows the claim record associated to an eligible address.

```sh
testappd q claim show-claim-record [address]
```

Example output:

```yml
claimRecord:
  address: cosmos1pkdk6m2nh77nlaep84cylmkhjder3areczme3w
  claimable: "100"
  completedMissions: []
```

#### `list-mission`

Lists the missions to complete to claim aidrop.

```sh
testappd q claim list-mission
```

Example output:

```yml
Mission:
  - description: initial claim
    missionID: "0"
    weight: "0.200000000000000000"
  - description: staking
    missionID: "1"
    weight: "0.500000000000000000"
  - description: voting
    missionID: "2"
    weight: "0.300000000000000000"
pagination:
  next_key: null
  total: "0"
```

#### `show-mission`

Shows information about a specific mission to claim a claimable amount of the airdrop.

```sh
testappd q claim show-mission [mission-id]
```

Example output:

```yml
Mission:
  description: staking
  missionID: "1"
  weight: "0.500000000000000000"
```

### Transactions

The `tx` commands allow users to interact with the `claim` module.

```sh
testappd tx claim
```

#### `claim-initial`

Claim the initial airdrop allocation for the user.

```sh
testappd tx claim claim-initial
```

Example:

```sh
testappd tx claim claim-initial --from alice
```

#### `claim`

Claim the airdrop allocation for the user and mission.

```sh
testappd tx claim claim 2
```

Example:

```sh
testappd tx claim claim 3 --from alice
```
