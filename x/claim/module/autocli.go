package claim

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ignite/modules/api/modules/claim/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListClaimRecord",
					Use:       "list-claim-record",
					Short:     "List all ClaimRecord",
				},
				{
					RpcMethod:      "GetClaimRecord",
					Use:            "get-claim-record [address]",
					Short:          "Gets a ClaimRecord",
					Alias:          []string{"show-claim-record"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod: "ListMission",
					Use:       "list-mission",
					Short:     "List all Mission",
				},
				{
					RpcMethod:      "GetMission",
					Use:            "get-mission [mission-id]",
					Short:          "Gets a Mission by id",
					Alias:          []string{"show-mission"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "missionID"}},
				},
				{
					RpcMethod: "GetInitialClaim",
					Use:       "get-initial-claim",
					Short:     "Gets a InitialClaim",
					Alias:     []string{"show-initial-claim"},
				},

				{
					RpcMethod: "GetAirdropSupply",
					Use:       "get-airdrop-supply",
					Short:     "Gets a AirdropSupply",
					Alias:     []string{"show-airdrop-supply"},
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
					RpcMethod:      "Claim",
					Use:            "claim [mission-id]",
					Short:          "Claim the airdrop allocation by mission id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "missionID"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
