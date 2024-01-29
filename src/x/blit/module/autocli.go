package blit

import (
	modulev1 "blit/api/blit/blit"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
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
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "TaskResultAll",
					Use:       "list-task-result",
					Short:     "List all TaskResult",
				},
				{
					RpcMethod:      "TaskResult",
					Use:            "show-task-result [id]",
					Short:          "Shows a TaskResult",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "FutureTaskAll",
					Use:       "list-future-task",
					Short:     "List all FutureTask",
				},
				{
					RpcMethod:      "FutureTask",
					Use:            "show-future-task [index]",
					Short:          "Shows a FutureTask",
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
					RpcMethod: "CreateTask",
					Use:       "create-task [creator] [activate_after] [expire_after] [interval] [max_runs] [disable_on_error] [enabled] [gas_limit] [gas_price]",
					Short:     "Create a new Task",
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"messages": {Name: "message"},
					},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "creator"},
						{ProtoField: "activate_after"},
						{ProtoField: "expire_after"},
						{ProtoField: "interval"},
						{ProtoField: "max_runs"},
						{ProtoField: "disable_on_error"},
						{ProtoField: "enabled"},
						{ProtoField: "gas_limit"},
						{ProtoField: "gas_price"},
					},
				},
				{
					RpcMethod: "UpdateTask",
					Use:       "update-task [id] [creator] [activate_after] [expire_after] [interval] [max_runs] [disable_on_error] [enabled] [gas_limit] [gas_price] [messages]",
					Short:     "Update Task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "id"},
						{ProtoField: "creator"},
						{ProtoField: "activate_after"},
						{ProtoField: "expire_after"},
						{ProtoField: "interval"},
						{ProtoField: "max_runs"},
						{ProtoField: "disable_on_error"},
						{ProtoField: "enabled"},
						{ProtoField: "gas_limit"},
						{ProtoField: "gas_price"},
						{ProtoField: "messages"},
					},
				},
				{
					RpcMethod:      "DeleteTask",
					Use:            "delete-task [id]",
					Short:          "Delete Task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateTaskResult",
					Use:            "create-task-result [id] [executed_on]",
					Short:          "Create a new TaskResult",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "executed_on"}},
				},
				{
					RpcMethod:      "UpdateTaskResult",
					Use:            "update-task-result [id] [executed_on]",
					Short:          "Update TaskResult",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "executed_on"}},
				},
				{
					RpcMethod:      "DeleteTaskResult",
					Use:            "delete-task-result [id]",
					Short:          "Delete TaskResult",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateFutureTask",
					Use:            "create-future-task [index] [scheduled_on] [task_id]",
					Short:          "Create a new FutureTask",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "scheduled_on"}, {ProtoField: "task_id"}},
				},
				{
					RpcMethod:      "UpdateFutureTask",
					Use:            "update-future-task [index] [scheduled_on] [task_id]",
					Short:          "Update FutureTask",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "scheduled_on"}, {ProtoField: "task_id"}},
				},
				{
					RpcMethod:      "DeleteFutureTask",
					Use:            "delete-future-task [index]",
					Short:          "Delete FutureTask",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
