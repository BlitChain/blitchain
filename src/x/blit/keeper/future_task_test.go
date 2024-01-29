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

func createNFutureTask(keeper keeper.Keeper, ctx context.Context, n int) []types.FutureTask {
	items := make([]types.FutureTask, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetFutureTask(ctx, items[i])
	}
	return items
}

func TestFutureTaskGet(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	items := createNFutureTask(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFutureTask(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFutureTaskRemove(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	items := createNFutureTask(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFutureTask(ctx,
			item.Index,
		)
		_, found := keeper.GetFutureTask(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestFutureTaskGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	items := createNFutureTask(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFutureTask(ctx)),
	)
}
