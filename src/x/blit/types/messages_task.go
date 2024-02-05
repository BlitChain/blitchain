package types

import (
	time "time"

	errorsmod "cosmossdk.io/errors"
	sdktypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
)

var _ sdk.Msg = &MsgCreateTask{}
var _ sdktypes.UnpackInterfacesMessage = MsgCreateTask{}

func (msg MsgCreateTask) UnpackInterfaces(unpacker sdktypes.AnyUnpacker) error {

	err := sdktx.UnpackInterfaces(unpacker, msg.Messages)
	if err != nil {
		return err
	}

	return nil
}

func NewMsgCreateTask(
	creator string,
	activateAfter time.Time,
	expireAfter time.Time,
	interval time.Duration,
	frequency time.Duration,
	maxRuns uint64,
	disableOnError bool,
	enabled bool,
	taskGasLimit uint64,
	taskGasFee sdk.Coin,
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
		TaskGasLimit:   taskGasLimit,
		TaskGasFee:     taskGasFee,
		Messages:       anys,
	}, nil
}

func (msg *MsgCreateTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	msgs, err := sdktx.GetMsgs(msg.Messages, "blit.MsgCreateTask")
	if err != nil {
		return err
	}

	err = validateMsgs(msgs)

	if err != nil {
		return err
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

func validateMsgs(msgs []sdk.Msg) error {
	for i, msg := range msgs {
		m, ok := msg.(sdk.HasValidateBasic)
		if !ok {
			continue
		}

		if err := m.ValidateBasic(); err != nil {
			return errorsmod.Wrapf(err, "msg %d", i)
		}
	}

	return nil
}
