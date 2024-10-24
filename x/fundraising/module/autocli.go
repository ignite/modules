package fundraising

import (
	"fmt"
	"strings"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	modulev1 "github.com/ignite/modules/api/modules/fundraising/v1"
	"github.com/ignite/modules/x/fundraising/keeper"
	"github.com/ignite/modules/x/fundraising/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	bech32PrefixValAddr := sdk.GetConfig().GetBech32ValidatorAddrPrefix()
	moduloOpts := &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "ListAllowedBidder",
					Use:            "list-allowed-bidder [auction-id]",
					Short:          "List all AllowedBidder",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}},
				},
				{
					RpcMethod:      "GetAllowedBidder",
					Use:            "get-allowed-bidder [auction-id] [bidder]",
					Short:          "Gets a AllowedBidder",
					Alias:          []string{"show-allowed-bidder"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}, {ProtoField: "bidder"}},
				},
				{
					RpcMethod:      "ListVestingQueue",
					Use:            "list-vesting-queue [auction-id]",
					Short:          "List all VestingQueue",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}},
				},
				{
					RpcMethod:      "ListBid",
					Use:            "list-bid [auction-id]",
					Short:          "List all Bid",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}},
				},
				{
					RpcMethod:      "GetBid",
					Use:            "get-bid [auction-id] [bid-id]",
					Short:          "Gets a Bid by id",
					Alias:          []string{"show-bid"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}, {ProtoField: "bid_id"}},
				},
				{
					RpcMethod: "ListAuction",
					Use:       "list-auction",
					Short:     "List all auction",
				},
				{
					RpcMethod:      "GetAuction",
					Use:            "get-auction [auction-id]",
					Short:          "Gets a auction by id",
					Alias:          []string{"show-auction"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod: "CancelAuction",
					Use:       "cancel [auction-id]",
					Short:     "Cancel the auction",
					Long: strings.TrimSpace(
						fmt.Sprintf(`Cancel the auction with the id. 
		
Example:
$ %s tx %s cancel 1 --from mykey 
`,
							version.AppName, types.ModuleName,
						),
					),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}},
				},
				{
					RpcMethod: "PlaceBid",
					Use:       "place-bid [auction-id] [bid-type] [price] [coin]",
					Short:     "Bid for the auction",
					Long: strings.TrimSpace(
						fmt.Sprintf(`Bid for the auction with what price and amount of coin you want to bid for. 

Bid Type Options:
1. fixed-price (fp or f)
2. batch-worth (bw or w) 
3. batch-many  (bm or m)

Example:
$ %s tx %s place-bid 1 fixed-price 0.55 100000000denom2 --from mykey 
$ %s tx %s place-bid 1 batch-worth 0.55 100000000denom2 --from mykey 
$ %s tx %s place-bid 1 batch-many 0.55 100000000denom1 --from mykey 
$ %s tx %s place-bid 1 fp 0.55 100000000denom2 --from mykey 
$ %s tx %s place-bid 1 bw 0.55 100000000denom2 --from mykey 
$ %s tx %s place-bid 1 bm 0.55 100000000denom1 --from mykey 

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
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}, {ProtoField: "bid_type"}, {ProtoField: "price"}, {ProtoField: "coin"}},
				},
				{
					RpcMethod: "ModifyBid",
					Use:       "modify-bid [auction-id] [bid-id] [price] [coin]",
					Short:     "Modify the bid",
					Long: strings.TrimSpace(
						fmt.Sprintf(`Modify the bid with new price and coin.
Either price or coin must be higher than the existing bid.

Example:
$ %s tx %s bid 1 1 1.0 100000000denom2 --from mykey
`,
							version.AppName, types.ModuleName,
						),
					),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}, {ProtoField: "bid_id"}, {ProtoField: "price"}, {ProtoField: "coin"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
	if keeper.EnableAddAllowedBidder {
		moduloOpts.Tx.RpcCommandOptions = append(moduloOpts.Tx.RpcCommandOptions, &autocliv1.RpcCommandOptions{
			RpcMethod: "AddAllowedBidder",
			Use:       "add-allowed-bidder [auction-id] [bidder] [max-bid-amount]",
			Short:     "(Testing) Add an allowed bidder for the auction",
			Long: strings.TrimSpace(
				fmt.Sprintf(`Add an allowed bidder for the auction.
This message is available for testing purpose and it is only accessible when you build the binary with testing mode.
		
Example:
$ %s tx %s add-allowed-bidder 1 %s1mzgucqnfr2l8cj5apvdpllhzt4zeuh2cshz5xu 10000000000 --from mykey 
`,
					version.AppName, types.ModuleName, bech32PrefixValAddr,
				),
			),
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}, {ProtoField: "bidder"}, {ProtoField: "max_bid_amount"}},
		})
	}

	return moduloOpts
}
