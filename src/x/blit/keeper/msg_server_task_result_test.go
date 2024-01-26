package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "blit/testutil/keeper"
	"blit/x/blit/keeper"
	"blit/x/blit/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTaskResultMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BlitKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateTaskResult{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateTaskResult(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetTaskResult(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestTaskResultMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateTaskResult
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateTaskResult{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateTaskResult{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateTaskResult{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BlitKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateTaskResult{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateTaskResult(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateTaskResult(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetTaskResult(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestTaskResultMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteTaskResult
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteTaskResult{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteTaskResult{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteTaskResult{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BlitKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateTaskResult(ctx, &types.MsgCreateTaskResult{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteTaskResult(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetTaskResult(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
