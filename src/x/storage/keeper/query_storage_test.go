package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "blit/testutil/keeper"
	"blit/testutil/nullify"
	"blit/x/storage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestStorageQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStorage(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetStorageRequest
		response *types.QueryGetStorageResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetStorageRequest{
				Address: msgs[0].Address,
				Index:   msgs[0].Index,
			},
			response: &types.QueryGetStorageResponse{Storage: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetStorageRequest{
				Address: msgs[1].Address,
				Index:   msgs[1].Index,
			},
			response: &types.QueryGetStorageResponse{Storage: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetStorageRequest{
				Address: strconv.Itoa(100000),
				Index:   strconv.Itoa(100000),
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
			response, err := keeper.Storage(wctx, tc.request)
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

func TestStorageQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.StorageKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStorage(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllStorageRequest {
		return &types.QueryAllStorageRequest{
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
			resp, err := keeper.StorageAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Storage), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Storage),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StorageAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Storage), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Storage),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.StorageAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Storage),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.StorageAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
