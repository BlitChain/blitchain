package blit_test

import (
	"testing"

	keepertest "blit/testutil/keeper"
	"blit/testutil/nullify"
	"blit/x/blit"
	"blit/x/blit/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlitKeeper(t)
	blit.InitGenesis(ctx, k, genesisState)
	got := blit.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
