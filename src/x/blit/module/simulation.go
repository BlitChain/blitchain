package blit

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"blit/testutil/sample"
	blitsimulation "blit/x/blit/simulation"
	"blit/x/blit/types"
)

// avoid unused import issue
var (
	_ = blitsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateTaskResult = "op_weight_msg_task_result"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTaskResult int = 100

	opWeightMsgUpdateTaskResult = "op_weight_msg_task_result"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTaskResult int = 100

	opWeightMsgDeleteTaskResult = "op_weight_msg_task_result"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTaskResult int = 100

	opWeightMsgCreateFutureTask = "op_weight_msg_future_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateFutureTask int = 100

	opWeightMsgUpdateFutureTask = "op_weight_msg_future_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateFutureTask int = 100

	opWeightMsgDeleteFutureTask = "op_weight_msg_future_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteFutureTask int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	blitGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TaskResultList: []types.TaskResult{
			{
				Creator: sample.AccAddress(),
				Id:      0,
			},
			{
				Creator: sample.AccAddress(),
				Id:      1,
			},
		},
		FutureTaskList: []types.FutureTask{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&blitGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTaskResult int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTaskResult, &weightMsgCreateTaskResult, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTaskResult = defaultWeightMsgCreateTaskResult
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTaskResult,
		blitsimulation.SimulateMsgCreateTaskResult(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTaskResult int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateTaskResult, &weightMsgUpdateTaskResult, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTaskResult = defaultWeightMsgUpdateTaskResult
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTaskResult,
		blitsimulation.SimulateMsgUpdateTaskResult(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTaskResult int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteTaskResult, &weightMsgDeleteTaskResult, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTaskResult = defaultWeightMsgDeleteTaskResult
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTaskResult,
		blitsimulation.SimulateMsgDeleteTaskResult(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateFutureTask int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateFutureTask, &weightMsgCreateFutureTask, nil,
		func(_ *rand.Rand) {
			weightMsgCreateFutureTask = defaultWeightMsgCreateFutureTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateFutureTask,
		blitsimulation.SimulateMsgCreateFutureTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateFutureTask int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateFutureTask, &weightMsgUpdateFutureTask, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateFutureTask = defaultWeightMsgUpdateFutureTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateFutureTask,
		blitsimulation.SimulateMsgUpdateFutureTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteFutureTask int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteFutureTask, &weightMsgDeleteFutureTask, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteFutureTask = defaultWeightMsgDeleteFutureTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteFutureTask,
		blitsimulation.SimulateMsgDeleteFutureTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTaskResult,
			defaultWeightMsgCreateTaskResult,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blitsimulation.SimulateMsgCreateTaskResult(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTaskResult,
			defaultWeightMsgUpdateTaskResult,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blitsimulation.SimulateMsgUpdateTaskResult(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteTaskResult,
			defaultWeightMsgDeleteTaskResult,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blitsimulation.SimulateMsgDeleteTaskResult(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateFutureTask,
			defaultWeightMsgCreateFutureTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blitsimulation.SimulateMsgCreateFutureTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateFutureTask,
			defaultWeightMsgUpdateFutureTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blitsimulation.SimulateMsgUpdateFutureTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteFutureTask,
			defaultWeightMsgDeleteFutureTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				blitsimulation.SimulateMsgDeleteFutureTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
