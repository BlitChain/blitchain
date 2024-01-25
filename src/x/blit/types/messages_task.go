package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTask{}

func NewMsgCreateTask(
	creator string,
	index string,
	taskId int32,
	totalRunCount int32,
	nextTaskResultIndex string,
	activateOn string,
	expireOn string,
	interval string,
	maxRuns int32,
	disableOnError bool,
	enabled bool,
	gasLimit int32,
	gasPrice string,
	messagesMo string,
	dule string,
	blit string,

) *MsgCreateTask {
	return &MsgCreateTask{
		Creator:             creator,
		Index:               index,
		TaskId:              taskId,
		TotalRunCount:       totalRunCount,
		NextTaskResultIndex: nextTaskResultIndex,
		ActivateOn:          activateOn,
		ExpireOn:            expireOn,
		Interval:            interval,
		MaxRuns:             maxRuns,
		DisableOnError:      disableOnError,
		Enabled:             enabled,
		GasLimit:            gasLimit,
		GasPrice:            gasPrice,
		MessagesMo:          messagesMo,
		Dule:                dule,
		Blit:                blit,
	}
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
	index string,
	taskId int32,
	totalRunCount int32,
	nextTaskResultIndex string,
	activateOn string,
	expireOn string,
	interval string,
	maxRuns int32,
	disableOnError bool,
	enabled bool,
	gasLimit int32,
	gasPrice string,
	messagesMo string,
	dule string,
	blit string,

) *MsgUpdateTask {
	return &MsgUpdateTask{
		Creator:             creator,
		Index:               index,
		TaskId:              taskId,
		TotalRunCount:       totalRunCount,
		NextTaskResultIndex: nextTaskResultIndex,
		ActivateOn:          activateOn,
		ExpireOn:            expireOn,
		Interval:            interval,
		MaxRuns:             maxRuns,
		DisableOnError:      disableOnError,
		Enabled:             enabled,
		GasLimit:            gasLimit,
		GasPrice:            gasPrice,
		MessagesMo:          messagesMo,
		Dule:                dule,
		Blit:                blit,
	}
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
	index string,

) *MsgDeleteTask {
	return &MsgDeleteTask{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
