package simulation_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/x/mint/simulation"
	"github.com/ignite/modules/x/mint/types"
)

// TestRandomizedGenState tests the normal scenario of applying RandomizedGenState.
// Abnormal scenarios are not tested here.
func TestRandomizedGenState(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          cdc,
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)

	var mintGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &mintGenesis)

	var (
		dec1 = sdkmath.LegacyMustNewDecFromStr("0.670000000000000000")
		dec2 = sdkmath.LegacyMustNewDecFromStr("0.200000000000000000")
		dec3 = sdkmath.LegacyMustNewDecFromStr("0.070000000000000000")
		dec4 = sdkmath.LegacyMustNewDecFromStr("0.170000000000000000")
		dec5 = sdkmath.LegacyMustNewDecFromStr("0.70000000000000000")
		dec6 = sdkmath.LegacyMustNewDecFromStr("0.130000000000000000")
	)

	weightedAddresses := []types.WeightedAddress{
		{
			Address: "cosmos15mf334fgwp4fze4udr6l2wgwuxk24yps7we7dq",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.017458224291856355"),
		},
		{
			Address: "cosmos1repxmyy9mx4xq4fajgjxaahaw0yjlmh5uk64m6",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos1n6wnkglm8m3sxr2f7g9rmv0u6ekjc7e4t5ta2f",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.006141422857328581"),
		},
		{
			Address: "cosmos17q32k8jlkac2yapq7mxf5lak0huh06p0g6sqmy",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.007683902997839140"),
		},
		{
			Address: "cosmos1ls6shhp8swr5ards3fp2lst5ftysse25dqmk2x",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.006856440371419632"),
		},
		{
			Address: "cosmos1xq7lhaqxvkkljqp73xt5h3uqwe0f3gj2pwc0xv",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos14g2upxq6hu49lt544a0jrs6j376r334h9w6tus",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.016662840271533599"),
		},
		{
			Address: "cosmos1kcu2rlxffgpycjn0e2fzs9de3xvdw8gf5n8em0",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.016406308110665711"),
		},
		{
			Address: "cosmos14lh203x67r7a7vd9as2vqk7uc08uumlpkvuu5g",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.008878619785299237"),
		},
		{
			Address: "cosmos1dec5rwvvcygyx6eg0hjv35vxfr58hawx7xvekw",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.016267509677487911"),
		},
		{
			Address: "cosmos1tlxztuxxs8rxmxv97247vsn5cafkv9sgplhmag",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.017180571755071666"),
		},
		{
			Address: "cosmos17ktwhac3c75k7vr62lsvmtx8u9cs6t8g0t8ts6",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.011536944938043694"),
		},
		{
			Address: "cosmos1mtjrpensxply0kstkgfdjwnjakvvwdnhrg2at0",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.010610596013385205"),
		},
		{
			Address: "cosmos1uqv7xkqey6sv8c3ta8aqn08pd97qg5dyqrf6a4",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos1hc08vrfk2fyhe2wkdep4uswcqtepy73rxpe5me",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.010676596986934073"),
		},
		{
			Address: "cosmos1rl5exulw99m0mx43c09m0k4ydmzdurxu8c4mpm",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos12yeysqnpx43uddpuve72xs4w9ghnhkf2mggrhz",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.005312909508254276"),
		},
		{
			Address: "cosmos1r557dv93c3qtzl4ws5ne9rgk6pp3wqchycem0q",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.010491849933835756"),
		},
		{
			Address: "cosmos1uhp7u3jflawdn64ehe0g0gvlnrtz9uvzzc9uvk",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.006390850270813198"),
		},
		{
			Address: "cosmos13uh4804tdnrffhvxajlce7ru9wn84lgl94cs6v",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.019172715206289138"),
		},
		{
			Address: "cosmos1nulylqzrde5utkjrz0twh6e2kp2csat2f8n9je",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos1906mhe8x2xge0gm4neczdv7wvs0fzxnfh05j7r",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.015322247994609446"),
		},
		{
			Address: "cosmos1vxnfc3v6fx6vwxn0gees08n02cuaf8jrjq7hxv",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.014120472015057072"),
		},
		{
			Address: "cosmos1j2uxc248vlfw32ws0auyrjatgkh58c6zepkjku",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.005000000000000000"),
		},
		{
			Address: "cosmos1fqd5p03lv502gy24r5fxyexpq6zyccgc56yagm",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.006481158671130950"),
		},
		{
			Address: "cosmos1wl62xxnkeftjksvmx73wj2fa3vpreg5fyczurn",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.011313084918839887"),
		},
		{
			Address: "cosmos1pmuazlfeqjpv0fdktkuh20epw276zvg58advj8",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.014478223295274740"),
		},
		{
			Address: "cosmos1nnzlwke77snwzdq54ryaau9gzw47a7sg8mrng4",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.005000000000000000"),
		},
		{
			Address: "cosmos1xzyprgtm3hjrqnr75qclqk7ypqunmfl7k26dza",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.011906575459237459"),
		},
		{
			Address: "cosmos1hyt7g5533ypamrkhm29jqdzu9w3mws0dgpwzzd",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.013018175136719084"),
		},
		{
			Address: "cosmos1ct3n6rn6qvps6j2f5pfhsyszjw9r3z6ps89x4m",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos1rvp5k0g7lk5sgul6prt3hs6z55ddd6k2jxqxtp",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.012310359216781740"),
		},
		{
			Address: "cosmos1aye9n83a4eulzntgq25zkyq4faljg225x5hm7y",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.013831594424707402"),
		},
		{
			Address: "cosmos1q9puau53l07ys8en7k5ejjlu49plr607jqku2m",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.019075909099606813"),
		},
		{
			Address: "cosmos186js4cmqzqeesltj9sjspsj6xthehg5zuk8ryq",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.010506634053162084"),
		},
		{
			Address: "cosmos1h3ytxxhq5kerq080uv985jmrmpwmhjlq4cjgk5",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.011408884176143559"),
		},
		{
			Address: "cosmos1fdufk5dtcnf5r5ktfk9l997c7tlhcfln620t6f",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.014516652780079007"),
		},
		{
			Address: "cosmos1mh4yza5rszduv6jtpkhucggp2gqwm2k42h8nz8",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos1c7xwgn6x4pp0zj2c33m6j87jl9t02wkn3yt7s7",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.005000000000000000"),
		},
		{
			Address: "cosmos1h2qndghy9tvl7qz9y5z3433kvhuxcwt5sk43j7",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.011152746629363641"),
		},
		{
			Address: "cosmos19amx0kcpnrdx4g9pttcm2v04m76atlxztcq8sd",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos1g9rzaedphtthffy9nrh0ukjdehau4p0aq72k5k",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.015477169807850692"),
		},
		{
			Address: "cosmos1gf73lfe9tzn27xpzz6xksku0c08rvk3gekka5h",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.007218936958952564"),
		},
		{
			Address: "cosmos10l5my6s0z3j74jz5fn9a6hfldv0jt2aawdug5n",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.013692731693070881"),
		},
		{
			Address: "cosmos1zj8sm2fsazhcn2jg2h24084f2l3eeu20d396dg",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.014945145118083484"),
		},
		{
			Address: "cosmos1w0ln5gz9tldh5dun9qx8ed3ww82zfsp4vemh79",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.006580834237555860"),
		},
		{
			Address: "cosmos1s2kdm985deuj52yw9mvmrsafhe95qxaghr5ng2",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.020000000000000000"),
		},
		{
			Address: "cosmos1ryaxl3wjyqf793yx0upz0c5se8p3uzj98cev93",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.005000000000000000"),
		},
		{
			Address: "cosmos1tyjkd7g3r5txnxugexzusn69f6tc6m8cj6qlj2",
			Weight:  sdkmath.LegacyMustNewDecFromStr("0.374914161337716463"),
		},
	}

	require.Equal(t, uint64(6311520), mintGenesis.Params.BlocksPerYear)
	require.Equal(t, dec1, mintGenesis.Params.GoalBonded)
	require.Equal(t, dec2, mintGenesis.Params.InflationMax)
	require.Equal(t, dec3, mintGenesis.Params.InflationMin)
	require.Equal(t, "stake", mintGenesis.Params.MintDenom)
	require.Equal(t, dec4, mintGenesis.Params.DistributionProportions.Staking)
	require.Equal(t, dec5, mintGenesis.Params.DistributionProportions.FundedAddresses)
	require.Equal(t, dec6, mintGenesis.Params.DistributionProportions.CommunityPool)
	require.Equal(t, "0stake", mintGenesis.Minter.BlockProvision(mintGenesis.Params).String())
	require.Equal(t, "0.170000000000000000", mintGenesis.Minter.NextAnnualProvisions(mintGenesis.Params, sdkmath.OneInt()).String())
	require.Equal(t, "0.169999926644441493", mintGenesis.Minter.NextInflationRate(mintGenesis.Params, sdkmath.LegacyOneDec()).String())
	require.Equal(t, "0.170000000000000000", mintGenesis.Minter.Inflation.String())
	require.Equal(t, "0.000000000000000000", mintGenesis.Minter.AnnualProvisions.String())
	for _, addr := range mintGenesis.Params.FundedAddresses {
		fmt.Println(addr)
	}
	require.Equal(t, weightedAddresses, mintGenesis.Params.FundedAddresses)
}

// TestRandomizedGenState tests abnormal scenarios of applying RandomizedGenState.
func TestRandomizedGenState1(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)
	// all these tests will panic
	tests := []struct {
		simState module.SimulationState
		panicMsg string
	}{
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{}, "invalid memory address or nil pointer dereference"},
		{ // panic => reason: incomplete initialization of the simState
			module.SimulationState{
				AppParams: make(simtypes.AppParams),
				Cdc:       cdc,
				Rand:      r,
			}, "assignment to entry in nil map"},
	}

	for _, tt := range tests {
		require.Panicsf(t, func() { simulation.RandomizedGenState(&tt.simState) }, tt.panicMsg)
	}
}
