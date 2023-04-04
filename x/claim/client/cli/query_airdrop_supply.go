package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/ignite/modules/x/claim/types"
)

func CmdShowAirdropSupply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-airdrop-supply",
		Short: "shows the airdrop supply",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetAirdropSupplyRequest{}

			res, err := queryClient.AirdropSupply(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
