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

func CmdCreateFixedPriceAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-fixed-price-auction [file]",
		Args:  cobra.ExactArgs(1),
		Short: "Create a fixed price auction",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a fixed price auction.
The auction details must be provided through a JSON file. 
		
Example:
$ %s tx %s create-fixed-price-auction <path/to/auction.json> --from mykey 

Where auction.json contains:

{
  "start_price": "1.000000000000000000",
  "selling_coin": {
    "denom": "denom1",
    "amount": "1000000000000"
  },
  "paying_coin_denom": "denom2",
  "vesting_schedules": [
    {
      "release_time": "2022-01-01T00:00:00Z",
      "weight": "0.500000000000000000"
    },
    {
      "release_time": "2022-06-01T00:00:00Z",
      "weight": "0.250000000000000000"
    },
    {
      "release_time": "2022-12-01T00:00:00Z",
      "weight": "0.250000000000000000"
    }
  ],
  "start_time": "2021-11-01T00:00:00Z",
  "end_time": "2021-12-01T00:00:00Z"
}

Description of the parameters:

[start_price]: the start price of the selling coin that is proportional to the paying coin denom 
[selling_coin]: the selling amount of coin for the auction
[paying_coin_denom]: the paying coin denom that the auctioneer wants to exchange with
[vesting_schedules]: the vesting schedules that release the paying coins to the auctioneer
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

			auction, err := parseFixedPriceAuctionRequest(args[0])
			if err != nil {
				return sdkerrors.Wrapf(errors.ErrInvalidRequest, "failed to parse %s file due to %v", args[0], err)
			}

			msg := types.NewMsgCreateFixedPriceAuction(
				clientCtx.GetFromAddress().String(),
				auction.StartPrice,
				auction.SellingCoin,
				auction.PayingCoinDenom,
				auction.VestingSchedules,
				auction.StartTime,
				auction.EndTime,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
