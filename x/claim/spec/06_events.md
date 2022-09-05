<!--
order: 6
-->

# Events

### `EventMissionCompleted`

This event is emitted when a mission `missionID` is completed for a specific eligible address `claimer`.

```
message EventMissionCompleted {
  uint64 missionID = 1;
  string claimer   = 2;
}
```
