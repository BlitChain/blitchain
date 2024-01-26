package blit_test

import (
	"testing"

	keepertest "blit/testutil/keeper"
	"blit/testutil/nullify"
	"blit/x/blit/module"
	"blit/x/blit/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TaskList: []types.Task{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		TaskResultList: []types.TaskResult{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlitKeeper(t)
	blit.InitGenesis(ctx, k, genesisState)
	got := blit.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TaskList, got.TaskList)
	require.ElementsMatch(t, genesisState.TaskResultList, got.TaskResultList)
	// this line is used by starport scaffolding # genesis/test/assert
}
