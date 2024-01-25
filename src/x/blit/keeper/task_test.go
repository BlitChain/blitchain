package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "blit/testutil/keeper"
	"blit/testutil/nullify"
	"blit/x/blit/keeper"
	"blit/x/blit/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTask(keeper keeper.Keeper, ctx context.Context, n int) []types.Task {
	items := make([]types.Task, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTask(ctx, items[i])
	}
	return items
}

func TestTaskGet(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	items := createNTask(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTask(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTaskRemove(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	items := createNTask(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTask(ctx,
			item.Index,
		)
		_, found := keeper.GetTask(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTaskGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	items := createNTask(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTask(ctx)),
	)
}
