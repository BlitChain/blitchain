package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateFutureTask{}

func NewMsgCreateFutureTask(
	creator string,
	index string,
	scheduledOn string,
	taskId uint64,

) *MsgCreateFutureTask {
	return &MsgCreateFutureTask{
		Creator:     creator,
		Index:       index,
		ScheduledOn: scheduledOn,
		TaskId:      taskId,
	}
}

func (msg *MsgCreateFutureTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateFutureTask{}

func NewMsgUpdateFutureTask(
	creator string,
	index string,
	scheduledOn string,
	taskId uint64,

) *MsgUpdateFutureTask {
	return &MsgUpdateFutureTask{
		Creator:     creator,
		Index:       index,
		ScheduledOn: scheduledOn,
		TaskId:      taskId,
	}
}

func (msg *MsgUpdateFutureTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteFutureTask{}

func NewMsgDeleteFutureTask(
	creator string,
	index string,

) *MsgDeleteFutureTask {
	return &MsgDeleteFutureTask{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteFutureTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
