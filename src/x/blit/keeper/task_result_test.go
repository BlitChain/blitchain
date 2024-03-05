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

func createNTaskResult(keeper keeper.Keeper, ctx context.Context, n int) []types.TaskResult {
	items := make([]types.TaskResult, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetTaskResult(ctx, items[i])
	}
	return items
}

func TestTaskResultGet(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	items := createNTaskResult(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTaskResult(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTaskResultRemove(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	items := createNTaskResult(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTaskResult(ctx,
			item.Index,
		)
		_, found := keeper.GetTaskResult(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestTaskResultGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	items := createNTaskResult(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTaskResult(ctx)),
	)
}
