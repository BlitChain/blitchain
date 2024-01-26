package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
)

var _ sdk.Msg = &MsgCreateTask{}

func NewMsgCreateTask(
	creator string,
	activateAfter string,
	expireAfter string,
	interval string,
	maxRuns int32,
	disableOnError bool,
	enabled bool,
	gasLimit int32,
	gasPrice string,
	messages []sdk.Msg,

) (*MsgCreateTask, error) {
	anys, err := sdktx.SetMsgs(messages)
	if err != nil {
		return nil, err
	}

	return &MsgCreateTask{
		Creator:        creator,
		ActivateAfter:  activateAfter,
		ExpireAfter:    expireAfter,
		Interval:       interval,
		MaxRuns:        maxRuns,
		DisableOnError: disableOnError,
		Enabled:        enabled,
		GasLimit:       gasLimit,
		GasPrice:       gasPrice,
		Messages:       anys,
	}, nil
}

func (msg *MsgCreateTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTask{}

func NewMsgUpdateTask(
	creator string,
	id uint64,
	activateAfter string,
	expireAfter string,
	interval string,
	maxRuns int32,
	disableOnError bool,
	enabled bool,
	gasLimit int32,
	gasPrice string,
	messages []sdk.Msg,

) (*MsgUpdateTask, error) {
	anys, err := sdktx.SetMsgs(messages)
	if err != nil {
		return nil, err
	}

	return &MsgUpdateTask{
		Creator:        creator,
		Id:             id,
		ActivateAfter:  activateAfter,
		ExpireAfter:    expireAfter,
		Interval:       interval,
		MaxRuns:        maxRuns,
		DisableOnError: disableOnError,
		Enabled:        enabled,
		GasLimit:       gasLimit,
		GasPrice:       gasPrice,
		Messages:       anys,
	}, nil
}

func (msg *MsgUpdateTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTask{}

func NewMsgDeleteTask(
	creator string,
	id uint64,

) *MsgDeleteTask {
	return &MsgDeleteTask{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeleteTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
