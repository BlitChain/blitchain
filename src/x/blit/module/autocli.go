package blit

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "blit/api/blit/blit"
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
					Short:     "Shows the parameters of the module.",
				},
				{
					RpcMethod: "TaskAll",
					Use:       "list-task",
					Short:     "List all Task",
				},
				{
					RpcMethod:      "Task",
					Use:            "show-task [id]",
					Short:          "Shows a Task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
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
					RpcMethod:      "MintCoins",
					Use:            "mint-coins [amount] [grantee]",
					Short:          "Send a mint_coins tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "grantee"}},
				},
				{
					RpcMethod:      "BurnCoins",
					Use:            "burn-coins [amount] [grantee]",
					Short:          "Send a burn_coins tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "grantee"}},
				},
				{
					RpcMethod:      "ForceTransferCoins",
					Use:            "force-transfer-coins [amount] [from-address] [to-address] [grantee]",
					Short:          "Send a force_transfer_coins tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "from_address"}, {ProtoField: "to_address"}, {ProtoField: "grantee"}},
				},
				{
					RpcMethod:      "SetDenomMetadata",
					Use:            "set-denom-metadata [base] [display] [name] [symbol] [uri] [uri-hash] [exponent] [description]",
					Short:          "Send a set_denom_metadata tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "base"}, {ProtoField: "display"}, {ProtoField: "name"}, {ProtoField: "symbol"}, {ProtoField: "uri"}, {ProtoField: "uri_hash"}, {ProtoField: "exponent"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "CreateTask",
					Use:            "create-task [index] [taskId] [totalRunCount] [nextTaskResultIndex] [activateOn] [expireOn] [interval] [maxRuns] [disableOnError] [enabled] [gasLimit] [gasPrice] [messagesMo] [dule] [blit]",
					Short:          "Create a new Task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "taskId"}, {ProtoField: "totalRunCount"}, {ProtoField: "nextTaskResultIndex"}, {ProtoField: "activateOn"}, {ProtoField: "expireOn"}, {ProtoField: "interval"}, {ProtoField: "maxRuns"}, {ProtoField: "disableOnError"}, {ProtoField: "enabled"}, {ProtoField: "gasLimit"}, {ProtoField: "gasPrice"}, {ProtoField: "messagesMo"}, {ProtoField: "dule"}, {ProtoField: "blit"}},
				},
				{
					RpcMethod:      "UpdateTask",
					Use:            "update-task [index] [taskId] [totalRunCount] [nextTaskResultIndex] [activateOn] [expireOn] [interval] [maxRuns] [disableOnError] [enabled] [gasLimit] [gasPrice] [messagesMo] [dule] [blit]",
					Short:          "Update Task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "taskId"}, {ProtoField: "totalRunCount"}, {ProtoField: "nextTaskResultIndex"}, {ProtoField: "activateOn"}, {ProtoField: "expireOn"}, {ProtoField: "interval"}, {ProtoField: "maxRuns"}, {ProtoField: "disableOnError"}, {ProtoField: "enabled"}, {ProtoField: "gasLimit"}, {ProtoField: "gasPrice"}, {ProtoField: "messagesMo"}, {ProtoField: "dule"}, {ProtoField: "blit"}},
				},
				{
					RpcMethod:      "DeleteTask",
					Use:            "delete-task [index]",
					Short:          "Delete Task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
