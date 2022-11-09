package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/modules/x/mint/types"
)

func (suite *MintTestSuite) TestUpdateParams() {
	app, ctx, msgServer := suite.app, suite.ctx, suite.msgServer

	testCases := []struct {
		name      string
		request   *types.MsgUpdateParams
		expectErr bool
	}{
		{
			name: "set invalid authority",
			request: &types.MsgUpdateParams{
				Authority: "foo",
			},
			expectErr: true,
		},
		{
			name: "set invalid params",
			request: &types.MsgUpdateParams{
				Authority: app.MintKeeper.GetAuthority(),
				Params: types.Params{
					MintDenom:               sdk.DefaultBondDenom,
					InflationRateChange:     sdk.NewDecWithPrec(-13, 2),
					InflationMax:            sdk.NewDecWithPrec(20, 2),
					InflationMin:            sdk.NewDecWithPrec(7, 2),
					GoalBonded:              sdk.NewDecWithPrec(67, 2),
					BlocksPerYear:           uint64(60 * 60 * 8766 / 5),
					DistributionProportions: types.DefaultDistributionProportions,
					FundedAddresses:         types.DefaultFundedAddresses,
				},
			},
			expectErr: true,
		},
		{
			name: "set full valid params",
			request: &types.MsgUpdateParams{
				Authority: app.MintKeeper.GetAuthority(),
				Params: types.Params{
					MintDenom:               sdk.DefaultBondDenom,
					InflationRateChange:     sdk.NewDecWithPrec(8, 2),
					InflationMax:            sdk.NewDecWithPrec(20, 2),
					InflationMin:            sdk.NewDecWithPrec(2, 2),
					GoalBonded:              sdk.NewDecWithPrec(37, 2),
					BlocksPerYear:           uint64(60 * 60 * 8766 / 5),
					DistributionProportions: types.DefaultDistributionProportions,
					FundedAddresses:         types.DefaultFundedAddresses,
				},
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			_, err := msgServer.UpdateParams(ctx, tc.request)
			if tc.expectErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
