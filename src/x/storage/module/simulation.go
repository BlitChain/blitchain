package storage

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"blit/testutil/sample"
	storagesimulation "blit/x/storage/simulation"
	"blit/x/storage/types"
)

// avoid unused import issue
var (
	_ = storagesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateStorage = "op_weight_msg_storage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStorage int = 100

	opWeightMsgUpdateStorage = "op_weight_msg_storage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStorage int = 100

	opWeightMsgDeleteStorage = "op_weight_msg_storage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteStorage int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	storageGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		StorageList: []types.Storage{
			{
				Address: sample.AccAddress(),
				Index:   "0",
			},
			{
				Address: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&storageGenesis)
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

	var weightMsgCreateStorage int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateStorage, &weightMsgCreateStorage, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStorage = defaultWeightMsgCreateStorage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStorage,
		storagesimulation.SimulateMsgCreateStorage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStorage int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateStorage, &weightMsgUpdateStorage, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStorage = defaultWeightMsgUpdateStorage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStorage,
		storagesimulation.SimulateMsgUpdateStorage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteStorage int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteStorage, &weightMsgDeleteStorage, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteStorage = defaultWeightMsgDeleteStorage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteStorage,
		storagesimulation.SimulateMsgDeleteStorage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateStorage,
			defaultWeightMsgCreateStorage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				storagesimulation.SimulateMsgCreateStorage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateStorage,
			defaultWeightMsgUpdateStorage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				storagesimulation.SimulateMsgUpdateStorage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteStorage,
			defaultWeightMsgDeleteStorage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				storagesimulation.SimulateMsgDeleteStorage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
