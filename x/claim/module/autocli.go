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
					Short:     "List all claim records",
				},
				{
					RpcMethod:      "GetClaimRecord",
					Use:            "get-claim-record [address]",
					Short:          "Gets a claim record by address",
					Alias:          []string{"show-claim-record"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod: "ListMission",
					Use:       "list-mission",
					Short:     "list all missions to claim airdrop",
				},
				{
					RpcMethod:      "GetMission",
					Use:            "get-mission [mission-id]",
					Short:          "Gets a Mission by id",
					Alias:          []string{"show-mission"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "mission_id"}},
				},
				{
					RpcMethod: "GetInitialClaim",
					Use:       "get-initial-claim",
					Short:     "Gets information about initial claim",
					Long:      "Gets if initial claim is enabled and what is the mission ID completed by initial claim",
					Alias:     []string{"show-initial-claim"},
				},

				{
					RpcMethod: "GetAirdropSupply",
					Use:       "get-airdrop-supply",
					Short:     "Gets the airdrop supply",
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
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "mission_id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
