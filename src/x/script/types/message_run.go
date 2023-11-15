package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRun{}

func NewMsgRun(callerAddress string, scriptAddress string, extraCode string, functionName string, kwargs string, grantee string) *MsgRun {
	return &MsgRun{
		CallerAddress: callerAddress,
		ScriptAddress: scriptAddress,
		ExtraCode:     extraCode,
		FunctionName:  functionName,
		Kwargs:        kwargs,
		Grantee:       grantee,
	}
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
