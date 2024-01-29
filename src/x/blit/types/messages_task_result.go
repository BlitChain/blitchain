package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTaskResult{}

func NewMsgCreateTaskResult(
	creator string,
	id uint64,
	executedOn string,

) *MsgCreateTaskResult {
	return &MsgCreateTaskResult{
		Creator:    creator,
		Id:         id,
		ExecutedOn: executedOn,
	}
}

func (msg *MsgCreateTaskResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTaskResult{}

func NewMsgUpdateTaskResult(
	creator string,
	id uint64,
	executedOn string,

) *MsgUpdateTaskResult {
	return &MsgUpdateTaskResult{
		Creator:    creator,
		Id:         id,
		ExecutedOn: executedOn,
	}
}

func (msg *MsgUpdateTaskResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTaskResult{}

func NewMsgDeleteTaskResult(
	creator string,
	id uint64,

) *MsgDeleteTaskResult {
	return &MsgDeleteTaskResult{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeleteTaskResult) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
