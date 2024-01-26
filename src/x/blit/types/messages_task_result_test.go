package types

import (
	"testing"

	"blit/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTaskResult_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTaskResult
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTaskResult{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTaskResult{
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

func TestMsgUpdateTaskResult_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTaskResult
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTaskResult{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTaskResult{
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

func TestMsgDeleteTaskResult_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTaskResult
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTaskResult{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTaskResult{
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
