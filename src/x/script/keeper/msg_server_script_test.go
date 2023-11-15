package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "blit/testutil/keeper"
	"blit/x/script/keeper"
	"blit/x/script/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestScriptMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ScriptKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	address := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateScript{Address: address,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateScript(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetScript(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Address, rst.Address)
	}
}

func TestScriptMsgServerUpdate(t *testing.T) {
	address := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateScript
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateScript{Address: address,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateScript{Address: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateScript{Address: address,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ScriptKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateScript{Address: address,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateScript(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateScript(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetScript(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Address, rst.Address)
			}
		})
	}
}
