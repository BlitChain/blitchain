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

func TestFutureTaskMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.BlitKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateFutureTask{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateFutureTask(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetFutureTask(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestFutureTaskMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateFutureTask
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateFutureTask{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateFutureTask{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateFutureTask{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BlitKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateFutureTask{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateFutureTask(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateFutureTask(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetFutureTask(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestFutureTaskMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteFutureTask
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteFutureTask{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteFutureTask{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteFutureTask{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.BlitKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateFutureTask(ctx, &types.MsgCreateFutureTask{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteFutureTask(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetFutureTask(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
