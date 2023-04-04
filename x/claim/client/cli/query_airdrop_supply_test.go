package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"github.com/ignite/modules/testutil/nullify"
	"github.com/ignite/modules/x/claim/client/cli"
	"github.com/ignite/modules/x/claim/types"
)

func (suite *QueryTestSuite) TestShowAirdropSupply() {
	ctx := suite.Network.Validators[0].ClientCtx
	airdropSupply := suite.ClaimState.AirdropSupply

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		name string
		args []string
		err  error
		obj  sdk.Coin
	}{
		{
			name: "should allow get",
			args: common,
			obj:  airdropSupply,
		},
	}
	for _, tc := range tests {
		suite.T().Run(tc.name, func(t *testing.T) {
			_, err := suite.Network.WaitForHeight(0)
			require.NoError(t, err)

			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowAirdropSupply(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
				return
			}

			require.NoError(t, err)
			var resp types.QueryGetAirdropSupplyResponse
			require.NoError(t, suite.Network.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NotNil(t, resp.AirdropSupply)
			require.Equal(t,
				nullify.Fill(&tc.obj),
				nullify.Fill(&resp.AirdropSupply),
			)
		})
	}
}
