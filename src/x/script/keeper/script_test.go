package keeper_test

import (
	"strconv"
	"testing"

	keepertest "blit/testutil/keeper"
	"blit/testutil/nullify"
	"blit/x/script/keeper"
	"blit/x/script/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNScript(keeper keeper.Keeper, ctx sdk.Context, n int) []types.Script {
	items := make([]types.Script, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetScript(ctx, items[i])
	}
	return items
}

func TestScriptGet(t *testing.T) {
	keeper, ctx := keepertest.ScriptKeeper(t)
	items := createNScript(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetScript(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestScriptRemove(t *testing.T) {
	keeper, ctx := keepertest.ScriptKeeper(t)
	items := createNScript(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveScript(ctx,
			item.Index,
		)
		_, found := keeper.GetScript(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestScriptGetAll(t *testing.T) {
	keeper, ctx := keepertest.ScriptKeeper(t)
	items := createNScript(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllScript(ctx)),
	)
}
