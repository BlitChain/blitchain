package storage_test

import (
	"testing"

	keepertest "blit/testutil/keeper"
	"blit/testutil/nullify"
	"blit/x/storage"
	"blit/x/storage/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		StorageList: []types.Storage{
			{
				Address: "0",
				Index:   "0",
			},
			{
				Address: "1",
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StorageKeeper(t)
	storage.InitGenesis(ctx, k, genesisState)
	got := storage.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StorageList, got.StorageList)
	// this line is used by starport scaffolding # genesis/test/assert
}
