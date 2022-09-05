<!--
order: 5
-->

# Parameters

The parameters of the module contain information about decaying for airdrop.

```
message Params {
  DecayInformation decayInformation = 1 [(gogoproto.nullable) = false];
}
```

### `DecayInformation`

This parameter determines if the airdrop starts to decay at a specific time.

```
message DecayInformation {
  bool enabled = 1;
  google.protobuf.Timestamp decayStart = 2 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
  google.protobuf.Timestamp decayEnd = 3 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true];
}
```

When enabled, the claimable amount for each eligible address will start to decrease starting from `decayStart` till `decayEnd` where the airdrop is ended and the remaining fund is transferred to the community pool.

The decrease is linear.

If `decayStart == decayEnd`, there is no decay for the airdrop but the airdrop ends at `decayEnd` and the remaining fund is transferred to the community pool.
