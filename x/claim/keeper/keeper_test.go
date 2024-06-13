package keeper_test

import (
	"math/rand"
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/golang/mock/gomock"

	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/claim"
	"github.com/ignite/modules/x/claim/keeper"
	claimtestutil "github.com/ignite/modules/x/claim/testutil"
	"github.com/ignite/modules/x/claim/types"
)

var r *rand.Rand

// initialize random generator
func init() {
	s := rand.NewSource(1)
	r = rand.New(s)
}

type testSuite struct {
	ctx           sdk.Context
	tk            *keeper.Keeper
	bankKeeper    types.BankKeeper
	accountKeeper types.AccountKeeper
	distrKeeper   types.DistrKeeper
}

func createClaimKeeper(t *testing.T) testSuite {
	t.Helper()

	encCfg := moduletestutil.MakeTestEncodingConfig(claim.AppModuleBasic{})
	key := storetypes.NewKVStoreKey(types.StoreKey)
	memKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)
	testCtx := testutil.DefaultContextWithKeys(map[string]*storetypes.KVStoreKey{types.ModuleName: key}, map[string]*storetypes.TransientStoreKey{types.ModuleName: storetypes.NewTransientStoreKey("transient_test")}, map[string]*storetypes.MemoryStoreKey{types.ModuleName: memKey})

	ctrl := gomock.NewController(t)
	bankKeeper := claimtestutil.NewMockBankKeeper(ctrl)
	accountKeeper := claimtestutil.NewMockAccountKeeper(ctrl)
	distrKeeper := claimtestutil.NewMockDistrKeeper(ctrl)

	return testSuite{
		ctx:           testCtx,
		tk:            keeper.NewKeeper(encCfg.Codec, key, memKey, accountKeeper, distrKeeper, bankKeeper, sample.AccAddress()),
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		distrKeeper:   distrKeeper,
	}
}
