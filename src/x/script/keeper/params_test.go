package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "blit/testutil/keeper"
	"blit/x/script/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ScriptKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
