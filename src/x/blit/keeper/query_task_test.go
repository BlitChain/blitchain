package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "blit/testutil/keeper"
	"blit/testutil/nullify"
	"blit/x/blit/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTaskQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	msgs := createNTask(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetTaskRequest
		response *types.QueryGetTaskResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetTaskRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetTaskResponse{Task: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetTaskRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetTaskResponse{Task: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetTaskRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Task(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestTaskQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.BlitKeeper(t)
	msgs := createNTask(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllTaskRequest {
		return &types.QueryAllTaskRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TaskAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Task), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Task),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.TaskAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Task), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Task),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.TaskAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Task),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.TaskAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
