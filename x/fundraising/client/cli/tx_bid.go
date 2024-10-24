package cli

import (
	"fmt"
	"strconv"
	"strings"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/ignite/modules/x/fundraising/types"
)

func CmdPlaceBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bid [auction-id] [bid-type] [price] [coin]",
		Args:  cobra.ExactArgs(4),
		Short: "Bid for the auction",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Bid for the auction with what price and amount of coin you want to bid for. 

Bid Type Options:
1. fixed-price (fp or f)
2. batch-worth (bw or w) 
3. batch-many  (bm or m)

Example:
$ %s tx %s bid 1 fixed-price 0.55 100000000denom2 --from mykey 
$ %s tx %s bid 1 batch-worth 0.55 100000000denom2 --from mykey 
$ %s tx %s bid 1 batch-many 0.55 100000000denom1 --from mykey 
$ %s tx %s bid 1 fp 0.55 100000000denom2 --from mykey 
$ %s tx %s bid 1 bw 0.55 100000000denom2 --from mykey 
$ %s tx %s bid 1 bm 0.55 100000000denom1 --from mykey 

Note:
In case of placing a bid for a fixed price auction, you must provide [price] argument with the same price of the auction. 
In case of placing a bid for a batch auction, there are two bid type options; batch-worth and batch-many, which you can find more information
in our technical spec docs. https://github.com/tendermint/fundraising/blob/main/x/fundraising/spec/01_concepts.md
`,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			bidType, err := parseBidType(args[1])
			if err != nil {
				return fmt.Errorf("parse order direction: %w", err)
			}

			price, err := sdkmath.LegacyNewDecFromStr(args[2])
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgPlaceBid(
				auctionId,
				clientCtx.GetFromAddress().String(),
				bidType,
				price,
				coin,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdModifyBid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "modify-bid [auction-id] [bid-id] [price] [coin]",
		Args:  cobra.ExactArgs(4),
		Short: "Modify the bid",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Modify the bid with new price and coin.
Either price or coin must be higher than the existing bid.

Example:
$ %s tx %s bid 1 1 1.0 100000000denom2 --from mykey
`,
				version.AppName, types.ModuleName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			auctionId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			bidId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			price, err := sdkmath.LegacyNewDecFromStr(args[2])
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgModifyBid(
				auctionId,
				clientCtx.GetFromAddress().String(),
				bidId,
				price,
				coin,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
