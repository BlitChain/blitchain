package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
)

const TypeMsgRun = "run"

var (
	_ sdk.Msg = &MsgRun{}

	_ cdctypes.UnpackInterfacesMessage = &MsgRun{}
)

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

// GetMessages returns the cache values from the MsgExecAuthorized.Msgs if present.
func (msg *MsgRun) GetMessages() ([]sdk.Msg, error) {
	return sdktx.GetMsgs(msg.Msgs, "blit.script.MsgRun")

}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (msg *MsgRun) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	for _, x := range msg.Msgs {
		var msg sdk.Msg
		err := unpacker.UnpackAny(x, &msg)
		if err != nil {
			return err
		}
	}

	return nil
}
