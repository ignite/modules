<!--
order: 3
-->

# Methods

### `CompleteMission`

Complete a claim mission for an eligible address.
This method can be used by an external module importing `claim` in order to define customized mission for the chain.

```
CompleteMission(
    ctx sdk.Context,
    missionID uint64,
    address string,
) error
```

**State transition**

- Complete the mission `missionID` in the claim record `address`
- Transfer the claim amount related to the mission

**Fails if**

- The mission doesn't exist
- The address has no claim record
- The mission has already been completed for the address
