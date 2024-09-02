# Changelog

## Unreleased

## `v0.1.0`

### Features

- [#98](https://github.com/ignite/modules/pull/98): Cosmos-sdk to `v0.50.8`
- [#99](https://github.com/ignite/modules/pull/99): Move fundraising module to monorepo

### Changes

- [#100](https://github.com/ignite/modules/pull/100): Add changelog and fundraising docs

## `v0.0.2`

### Changes

- [#93](https://github.com/ignite/modules/pull/93): Bump cometbft

### Fixes

- [#95](https://github.com/ignite/modules/pull/95): Correct spelling error in `airdrop_supply.go`

## `v0.0.1`

### Features

- [#7](https://github.com/ignite/modules/pull/7): `testapp` initialization
- [#9](https://github.com/ignite/modules/pull/9): add `Makefile` and add CI files
- [#10](https://github.com/ignite/modules/pull/10): add `claim`
- [#17](https://github.com/ignite/modules/pull/17): add `mint` module
- [#32](https://github.com/ignite/modules/pull/32): use `cosmossdk.io/math`
- [#34](https://github.com/ignite/modules/pull/34): remove legacy querier for `mint`
- [#35](https://github.com/ignite/modules/pull/35): remove `handler.go` for `claim`
- [#33](https://github.com/ignite/modules/pull/33): use `cosmossdk.io/errors`
- [#38](https://github.com/ignite/modules/pull/38): define sdkerrors
- [#37](https://github.com/ignite/modules/pull/37): `claim` invariants
- [#52](https://github.com/ignite/modules/pull/52): use typed event for mint event
- [#43](https://github.com/ignite/modules/pull/43): alias sdk error types
- [#39](https://github.com/ignite/modules/pull/39): send remaining `airdropSupply` after `decayEnd` for `claim` module
- [#55](https://github.com/ignite/modules/pull/55): move errors to `pkg`
- [#58](https://github.com/ignite/modules/pull/58): cleanup `mint` module
- [#60](https://github.com/ignite/modules/pull/60): cleanup `claim` module.go
- [#63](https://github.com/ignite/modules/pull/63): remove `ibc-go` dep
- [#57](https://github.com/ignite/modules/pull/57): `mint` `types` package
- [#65](https://github.com/ignite/modules/pull/65): normalize `claim`
- [#81](https://github.com/ignite/modules/pull/81): register staking and gov hooks
- [#77](https://github.com/ignite/modules/pull/77): airdrop start
- [#82](https://github.com/ignite/modules/pull/82): add claim record mission invariant
- [#86](https://github.com/ignite/modules/pull/86): bump `cosmos-sdk` to `v0.47.1`
- [#87](https://github.com/ignite/modules/pull/87): add sim tests
- [#90](https://github.com/ignite/modules/pull/90): bump cosmos-sdk and ibc

### Changes

- [#8](https://github.com/ignite/modules/pull/8): create `CODEOWNERS` file
- [#15](https://github.com/ignite/modules/pull/15): add utility files
- [#19](https://github.com/ignite/modules/pull/19): bump dependencies
- [#18](https://github.com/ignite/modules/pull/18): create LICENSE
- [#25](https://github.com/ignite/modules/pull/25): initialize readme
- [#41](https://github.com/ignite/modules/pull/41): remove CLI dependency
- [#42](https://github.com/ignite/modules/pull/42): add tools to repo
- [#48](https://github.com/ignite/modules/pull/48): remove mention of ATOMs in comments
- [#51](https://github.com/ignite/modules/pull/51): add correct home in config
- [#45](https://github.com/ignite/modules/pull/45): specification for `claim`
- [#50](https://github.com/ignite/modules/pull/50): initialize specs
- [#54](https://github.com/ignite/modules/pull/54): fix gh badges
- [#56](https://github.com/ignite/modules/pull/56): repo updates
- [#67](https://github.com/ignite/modules/pull/67): dragonberry patch
- [#72](https://github.com/ignite/modules/pull/72): bump deps
- [#73](https://github.com/ignite/modules/pull/73): add Pantani as a code owner
- [#74](https://github.com/ignite/modules/pull/74): upgrade config and fix `gogo.proto` import issue
- [#75](https://github.com/ignite/modules/pull/75): do some cleanup in proto files
- [#78](https://github.com/ignite/modules/pull/78): update deps
- [#83](https://github.com/ignite/modules/pull/83): test all go packages and remove the non-testable files from coverage
- [#85](https://github.com/ignite/modules/pull/85): remove ibc
- [#89](https://github.com/ignite/modules/pull/89): bump `cosmos-sdk` and `cometbft`
- [#91](https://github.com/ignite/modules/pull/91): update buf version and files

### Fixes

- [#66](https://github.com/ignite/modules/pull/66): Update `root.go` with security fix
- [#76](https://github.com/ignite/modules/pull/76): Import "gogoproto/gogo.proto" was not found
