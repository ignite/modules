package cli

import (
	"fmt"
	"strings"

	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/ignite/modules/x/fundraising/types"
)

func CmdCreateBatchAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-batch-auction [file]",
		Args:  cobra.ExactArgs(1),
		Short: "Create a batch auction",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a batch auction.
The auction details must be provided through a JSON file. 
		
Example:
$ %s tx %s create-batch-auction <path/to/auction.json> --from mykey 

Where auction.json contains:
{
  "start_price": "0.500000000000000000",
  "min_bid_price": "0.100000000000000000",
  "selling_coin": {
    "denom": "denom1",
    "amount": "1000000000000"
  },
  "paying_coin_denom": "denom2",
  "vesting_schedules": [
    {
      "release_time": "2023-06-01T00:00:00Z",
      "weight": "0.500000000000000000"
    },
    {
      "release_time": "2023-12-01T00:00:00Z",
      "weight": "0.500000000000000000"
    }
  ],
  "max_extended_round": 2,
  "extended_round_rate": "0.150000000000000000",
  "start_time": "2022-02-01T00:00:00Z",
  "end_time": "2022-06-20T00:00:00Z"
}

Description of the parameters:

[start_price]: the start price of the selling coin that is proportional to the paying coin denom 
[selling_coin]: the selling amount of coin for the auction
[paying_coin_denom]: the paying coin denom that the auctioneer wants to exchange with
[vesting_schedules]: the vesting schedules that release the paying coins to the autioneer
[max_extended_round]: the number of extended rounds
[extended_round_rate]: the rate that determines if the auction needs to run another round
[start_time]: the start time of the auction
[end_time]: the end time of the auction
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			auction, err := ParseBatchAuctionRequest(args[0])
			if err != nil {
				return sdkerrors.Wrapf(errors.ErrInvalidRequest, "failed to parse %s file due to %v", args[0], err)
			}

			msg := types.NewMsgCreateBatchAuction(
				clientCtx.GetFromAddress().String(),
				auction.StartPrice,
				auction.MinBidPrice,
				auction.SellingCoin,
				auction.PayingCoinDenom,
				auction.VestingSchedules,
				auction.MaxExtendedRound,
				auction.ExtendedRoundRate,
				auction.StartTime,
				auction.EndTime,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
