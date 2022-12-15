<!--
order: 2
-->

# Messages

### `MsgClaimInitial`

Claim the initial claim amount for airdrop defined in `InitialClaim`

```protobuf
message MsgClaimInitial {
  string claimer = 1;
}
```

**State transition**

- Complete the initial claim mission for the address
- Transfer the initial claim amount to the claimer balance

**Fails if**

- Initial claim is not enabled
- The claimer is not eligible
- The initial claim mission doesn't exist
- The initial claim mission has already been completed

### `MsgClaim`

Claim completed mission amount for airdrop

```protobuf
message MsgClaim {
  string claimer = 1;
  uint64 missionID = 2;
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
