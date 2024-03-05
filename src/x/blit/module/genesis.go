package blit

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"blit/x/blit/keeper"
	"blit/x/blit/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {

	err := k.TaskID.Set(ctx, genState.StartingTaskId)
	if err != nil {
		panic(err)
	}

	// Set all the task
	for _, elem := range genState.TaskList {
		k.SetTask(ctx, &elem)
	}
	// Set all the futureTask
	for _, elem := range genState.FutureTaskList {
		k.SetFutureTask(ctx, &elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {

	startingTaskId, err := k.TaskID.Peek(ctx)
	if err != nil {
		panic(err)
	}

	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.StartingTaskId = startingTaskId
	genesis.TaskList, err = k.GetAllTask(ctx)
	if err != nil {
		panic(err)
	}
	genesis.FutureTaskList, err = k.GetAllFutureTask(ctx)
	if err != nil {
		panic(err)
	}

	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
