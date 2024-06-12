package keeper_test

import (
	gocontext "context"
	"testing"

	"github.com/stretchr/testify/suite"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"

	"github.com/ignite/modules/testutil/sample"
	"github.com/ignite/modules/x/mint/keeper"
	"github.com/ignite/modules/x/mint/types"
)

type MintTestSuite struct {
	suite.Suite

	mintKeeper  keeper.Keeper
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *MintTestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(suite.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	encCfg := moduletestutil.MakeTestEncodingConfig()

	k := keeper.NewKeeper(encCfg.Codec, key, sample.AccAddress(), nil, nil, nil, nil, "")
	queryHelper := baseapp.NewQueryServerTestHelper(testCtx.Ctx, encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, k)
	queryClient := types.NewQueryClient(queryHelper)

	suite.mintKeeper = k
	suite.ctx = testCtx.Ctx
	suite.queryClient = queryClient
}

func (suite *MintTestSuite) TestGRPCParams() {
	k, ctx, queryClient := suite.mintKeeper, suite.ctx, suite.queryClient

	params, err := queryClient.Params(gocontext.Background(), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(params.Params, k.GetParams(ctx))

	inflation, err := queryClient.Inflation(gocontext.Background(), &types.QueryInflationRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(inflation.Inflation, k.GetMinter(ctx).Inflation)

	annualProvisions, err := queryClient.AnnualProvisions(gocontext.Background(), &types.QueryAnnualProvisionsRequest{})
	suite.Require().NoError(err)
	suite.Require().Equal(annualProvisions.AnnualProvisions, k.GetMinter(ctx).AnnualProvisions)
}

func TestMintTestSuite(t *testing.T) {
	suite.Run(t, new(MintTestSuite))
}
