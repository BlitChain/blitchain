package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRun = "run"

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

func (msg *MsgRun) Route() string {
	return RouterKey
}

func (msg *MsgRun) Type() string {
	return TypeMsgRun
}

func (msg *MsgRun) GetSigners() []sdk.AccAddress {

	if msg.Grantee != "" {
		grantee, err := sdk.AccAddressFromBech32(msg.Grantee)
		if err != nil {
			panic(err)
		}
		return []sdk.AccAddress{grantee}
	}
	caller, err := sdk.AccAddressFromBech32(msg.CallerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{caller}
}

func (msg *MsgRun) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRun) ValidateBasic() error {
	if msg.Grantee != "" {
		_, err := sdk.AccAddressFromBech32(msg.Grantee)
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
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
