#!/usr/bin/env bash

mockgen_cmd="mockgen"
$mockgen_cmd -source=x/claim/types/expected_keepers.go -package testutil -destination x/claim/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=x/mint/types/expected_keepers.go -package testutil -destination x/mint/testutil/expected_keepers_mocks.go
