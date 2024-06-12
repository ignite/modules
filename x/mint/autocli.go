package mint

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "", // mintv1beta1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "shows the parameters of the module",
				},
				{
					RpcMethod: "AirdropSupply",
					Use:       "show-airdrop-supply",
					Short:     "shows the airdrop supply",
				},
				{
					RpcMethod: "ClaimRecordAll",
					Use:       "list-claim-record",
					Short:     "list all claim records",
				},
				{
					RpcMethod:      "ClaimRecord",
					Use:            "show-claim-record [address]",
					Short:          "shows a claim record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod: "InitialClaim",
					Use:       "show-initial-claim",
					Short:     "shows information about initial claim",
					Long:      "shows if initial claim is enabled and what is the mission ID completed by initial claim",
				},
				{
					RpcMethod: "MissionAll",
					Use:       "list-mission",
					Short:     "list all missions to claim airdrop",
				},
				{
					RpcMethod:      "Mission",
					Use:            "show-mission [mission-id]",
					Short:          "shows a mission to claim airdrop",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "mission-id"}},
				},
			},
		},
	}
}
