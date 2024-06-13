package keeper_test

import (
	"math/rand"
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"

	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim"
	"github.com/ignite/modules/x/claim/keeper"
	"github.com/ignite/modules/x/claim/types"
)

var r *rand.Rand

// initialize random generator
func init() {
	s := rand.NewSource(1)
	r = rand.New(s)
}

func createClaimKeeper(t *testing.T) (sdk.Context, *keeper.Keeper) {
	t.Helper()

	encCfg := moduletestutil.MakeTestEncodingConfig(claim.AppModuleBasic{})
	key := storetypes.NewKVStoreKey(types.StoreKey)
	memKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	testCtx := testutil.DefaultContextWithKeys(map[string]*storetypes.KVStoreKey{types.ModuleName: key}, map[string]*storetypes.TransientStoreKey{types.ModuleName: storetypes.NewTransientStoreKey("transient_test")}, map[string]*storetypes.MemoryStoreKey{types.ModuleName: memKey})

	return testCtx, keeper.NewKeeper(encCfg.Codec, key, memKey, nil, nil, nil, sample.AccAddress())
}
