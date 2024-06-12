package claim

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	claimsimulation "github.com/ignite/modules/x/claim/simulation"
	"github.com/ignite/modules/x/claim/types"
)

const (
	airdropDenom = "drop"

	opWeightMsgClaim          = "op_weight_msg_claim"
	defaultWeightMsgClaim int = 50

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	claimRecords := make([]types.ClaimRecord, len(simState.Accounts))
	totalSupply := sdkmath.ZeroInt()
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()

		// fill claim records from simulation accounts
		accSupply := sdkmath.NewIntFromUint64(simState.Rand.Uint64() % 1000)
		claimRecords[i] = types.ClaimRecord{
			Claimable: accSupply,
			Address:   acc.Address.String(),
		}
		totalSupply = totalSupply.Add(accSupply)
	}

	// define some decimal numbers for mission weights
	dec1, err := sdkmath.LegacyNewDecFromStr("0.4")
	if err != nil {
		panic(err)
	}
	dec2, err := sdkmath.LegacyNewDecFromStr("0.3")
	if err != nil {
		panic(err)
	}

	claimGenesis := types.GenesisState{
		Params:        types.DefaultParams(),
		AirdropSupply: sdk.NewCoin(airdropDenom, totalSupply),
		Missions: []types.Mission{
			{
				MissionID:   0,
				Description: "initial claim",
				Weight:      dec1,
			},
			{
				MissionID:   1,
				Description: "mission 1",
				Weight:      dec2,
			},
			{
				MissionID:   2,
				Description: "mission 2",
				Weight:      dec2,
			},
		},
		InitialClaim: types.InitialClaim{
			Enabled:   true,
			MissionID: 0,
		},
		ClaimRecords: claimRecords,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&claimGenesis)
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.LegacyParamChange {
	return []simtypes.LegacyParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgClaim int
	simState.AppParams.GetOrGenerate(opWeightMsgClaim, &weightMsgClaim, nil,
		func(_ *rand.Rand) {
			weightMsgClaim = defaultWeightMsgClaim
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaim,
		claimsimulation.SimulateMsgClaim(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgClaim,
			defaultWeightMsgClaim,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				claimsimulation.SimulateMsgClaim(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
