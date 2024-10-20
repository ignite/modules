package claim

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/modules/testutil/sample"
	claimsimulation "github.com/ignite/modules/x/claim/simulation"
	"github.com/ignite/modules/x/claim/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
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
		AirdropSupply: types.AirdropSupply{Supply: sdk.NewCoin(airdropDenom, totalSupply)},
		MissionList: []types.Mission{
			{
				MissionId:   0,
				Description: "initial claim",
				Weight:      dec1,
			},
			{
				MissionId:   1,
				Description: "mission 1",
				Weight:      dec2,
			},
			{
				MissionId:   2,
				Description: "mission 2",
				Weight:      dec2,
			},
		},
		InitialClaim: types.InitialClaim{
			Enabled:   true,
			MissionId: 0,
		},
		ClaimRecordList: claimRecords,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&claimGenesis)
}

// RegisterStoreDecoder registers a decoder.
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
		claimsimulation.SimulateMsgClaim(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
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
				claimsimulation.SimulateMsgClaim(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
