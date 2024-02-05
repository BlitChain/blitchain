package types

import (
	"testing"

	"blit/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTask_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTask
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTask{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTask{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteTask_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTask
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTask{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTask{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
