package claim_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/stretchr/testify/require"

	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ClaimRecords: []types.ClaimRecord{
			{
				Address: sample.AccAddress(),
			},
			{
				Address: sample.AccAddress(),
			},
		},
		Missions: []types.Mission{
			{
				MissionID: 0,
			},
			{
				MissionID: 1,
			},
		},
		AirdropSupply: sdk.NewCoin("foo", sdkmath.NewInt(1000)),
		InitialClaim: types.InitialClaim{
			Enabled:   true,
			MissionID: 35,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	encCfg := moduletestutil.MakeTestEncodingConfig(claim.AppModuleBasic{})
	key := storetypes.NewKVStoreKey(types.StoreKey)
	memKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	ctx := testutil.DefaultContextWithKeys(map[string]*storetypes.KVStoreKey{types.ModuleName: key}, map[string]*storetypes.TransientStoreKey{types.ModuleName: storetypes.NewTransientStoreKey("transient_test")}, map[string]*storetypes.MemoryStoreKey{types.ModuleName: memKey})

	tk := keeper.NewKeeper(encCfg.Codec, key, memKey, nil, nil, nil, sample.AccAddress())

	t.Run("should allow import and export of genesis", func(t *testing.T) {
		claim.InitGenesis(ctx, *tk, genesisState)
		got := claim.ExportGenesis(ctx, *tk)
		require.NotNil(t, got)

		nullify.Fill(&genesisState)
		nullify.Fill(got)

		require.ElementsMatch(t, genesisState.ClaimRecords, got.ClaimRecords)
		require.ElementsMatch(t, genesisState.Missions, got.Missions)
		require.Equal(t, genesisState.AirdropSupply, got.AirdropSupply)
		require.Equal(t, genesisState.InitialClaim, got.InitialClaim)
		// this line is used by starport scaffolding # genesis/test/assert
	})
}
