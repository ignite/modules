<!--
order: 7
-->

# Client

## CLI

A user can query and interact with the `claim` module using the chain CLI.

### Query

The `query` commands allow users to query `claim` state.

```sh
modulesd q claim
```

#### `params`

Shows the params of the module.

```sh
modulesd q claim params
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
modulesd q claim show-airdrop-supply
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
modulesd q claim show-initial-claim
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
modulesd q claim list-claim-record
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
modulesd q claim show-claim-record [address]
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
modulesd q claim list-mission
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
modulesd q claim show-mission [mission-id]
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
modulesd tx claim
```

#### `claim-initial`

Claim the initial airdrop allocation for the user.

```sh
modulesd tx claim claim-initial
```

Example:

```sh
modulesd tx claim claim-initial --from alice
```

#### `claim`

Claim the airdrop allocation for the user and mission.

```sh
modulesd tx claim claim 2
```

Example:

```sh
modulesd tx claim claim 3 --from alice
```
