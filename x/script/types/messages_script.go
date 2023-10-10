package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateScript = "create_script"
	TypeMsgUpdateScript = "update_script"
)

var _ sdk.Msg = &MsgCreateScript{}

func NewMsgCreateScript(
	creator string,
	code string,
	msgTypePermissions []string,
	grantee string,
) *MsgCreateScript {
	return &MsgCreateScript{
		Creator:            creator,
		Code:               code,
		MsgTypePermissions: msgTypePermissions,
		Grantee:            grantee,
	}
}

func (msg *MsgCreateScript) Route() string {
	return RouterKey
}

func (msg *MsgCreateScript) Type() string {
	return TypeMsgCreateScript
}

func (msg *MsgCreateScript) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateScript) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateScript) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateScript{}

func NewMsgUpdateScript(
	address string,
	code string,
	grantee string,

) *MsgUpdateScript {
	return &MsgUpdateScript{
		Address: address,
		Code:    code,
		Grantee: grantee,
	}
}

func (msg *MsgUpdateScript) Route() string {
	return RouterKey
}

func (msg *MsgUpdateScript) Type() string {
	return TypeMsgUpdateScript
}

func (msg *MsgUpdateScript) GetSigners() []sdk.AccAddress {

	if msg.Grantee != "" {
		address, err := sdk.AccAddressFromBech32(msg.Grantee)
		if err != nil {
			panic(err)
		}
		return []sdk.AccAddress{address}
	}

	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgUpdateScript) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateScript) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}
	if msg.Grantee != "" {
		_, err = sdk.AccAddressFromBech32(msg.Grantee)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee (%s)", err)
		}
	}
	return nil
}
