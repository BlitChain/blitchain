package types

import (
	errorsmod "cosmossdk.io/errors"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRun{}

func NewMsgRun(callerAddress string, scriptAddress string, extraCode string, functionName string, kwargs string, grantee string, attachedMsgs []sdk.Msg) *MsgRun {
	msgsAny := make([]*cdctypes.Any, len(attachedMsgs))
	for i, msg := range attachedMsgs {
		any, err := cdctypes.NewAnyWithValue(msg)
		if err != nil {
			panic(err)
		}
		msgsAny[i] = any
	}

	return &MsgRun{
		CallerAddress:    callerAddress,
		ScriptAddress:    scriptAddress,
		ExtraCode:        extraCode,
		FunctionName:     functionName,
		Kwargs:           kwargs,
		Grantee:          grantee,
		AttachedMessages: msgsAny,
	}
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg MsgRun) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	for _, x := range msg.AttachedMessages {
		var attachedMsg sdk.Msg
		err := unpacker.UnpackAny(x, &attachedMsg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (msg *MsgRun) ValidateBasic() error {
	if msg.Grantee != "" {
		_, err := sdk.AccAddressFromBech32(msg.Grantee)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
		}
	}

	_, err := sdk.AccAddressFromBech32(msg.CallerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid caller address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.ScriptAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid script address (%s)", err)
	}
	return nil
}

// GetMessages returns the cache values from the MsgExecAuthorized.Msgs if present.
func (msg MsgRun) GetMessages() ([]sdk.Msg, error) {
	msgs := make([]sdk.Msg, len(msg.AttachedMessages))
	for i, msgAny := range msg.AttachedMessages {
		msg, ok := msgAny.GetCachedValue().(sdk.Msg)
		if !ok {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("AttachedMessages contains %T which is not a sdk.MsgRequest", msgAny)
		}
		msgs[i] = msg
	}

	return msgs, nil
}
