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

func TestTaskMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BlitKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateTask{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateTask(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetTask(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestTaskMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateTask
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateTask{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateTask{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateTask{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BlitKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateTask{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateTask(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateTask(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetTask(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestTaskMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteTask
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteTask{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteTask{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteTask{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BlitKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateTask(ctx, &types.MsgCreateTask{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteTask(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetTask(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
