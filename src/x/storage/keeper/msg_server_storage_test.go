package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "blit/testutil/keeper"
	"blit/x/storage/keeper"
	"blit/x/storage/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestStorageMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.StorageKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateStorage{Creator: creator,
			Address: strconv.Itoa(i),
			Index:   strconv.Itoa(i),
		}
		_, err := srv.CreateStorage(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetStorage(ctx,
			expected.Address,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestStorageMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateStorage
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateStorage{Creator: creator,
				Address: strconv.Itoa(0),
				Index:   strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateStorage{Creator: "B",
				Address: strconv.Itoa(0),
				Index:   strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateStorage{Creator: creator,
				Address: strconv.Itoa(100000),
				Index:   strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateStorage{Creator: creator,
				Address: strconv.Itoa(0),
				Index:   strconv.Itoa(0),
			}
			_, err := srv.CreateStorage(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateStorage(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetStorage(ctx,
					expected.Address,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestStorageMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteStorage
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteStorage{Creator: creator,
				Address: strconv.Itoa(0),
				Index:   strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteStorage{Creator: "B",
				Address: strconv.Itoa(0),
				Index:   strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteStorage{Creator: creator,
				Address: strconv.Itoa(100000),
				Index:   strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.StorageKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateStorage(wctx, &types.MsgCreateStorage{Creator: creator,
				Address: strconv.Itoa(0),
				Index:   strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteStorage(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetStorage(ctx,
					tc.request.Address,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
