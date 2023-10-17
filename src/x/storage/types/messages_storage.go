package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateStorage = "create_storage"
	TypeMsgUpdateStorage = "update_storage"
	TypeMsgDeleteStorage = "delete_storage"
)

var _ sdk.Msg = &MsgCreateStorage{}

func NewMsgCreateStorage(
	address string,
	index string,
	data string,
	grantee string,

) *MsgCreateStorage {
	return &MsgCreateStorage{
		Address: address,
		Index:   index,
		Data:    data,
		Grantee: grantee,
	}
}

func (msg *MsgCreateStorage) Route() string {
	return RouterKey
}

func (msg *MsgCreateStorage) Type() string {
	return TypeMsgCreateStorage
}

func (msg *MsgCreateStorage) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Address)

	if msg.Grantee != "" {
		address, err := sdk.AccAddressFromBech32(msg.Grantee)
		if err != nil {
			panic(err)
		}
		return []sdk.AccAddress{address}
	}

	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{address}
}

func (msg *MsgCreateStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err)
	}
	if msg.Grantee != "" {
		_, err = sdk.AccAddressFromBech32(msg.Grantee)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
		}
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateStorage{}

func NewMsgUpdateStorage(
	address string,
	index string,
	data string,
	grantee string,

) *MsgUpdateStorage {
	return &MsgUpdateStorage{
		Address: address,
		Index:   index,
		Data:    data,
		Grantee: grantee,
	}
}

func (msg *MsgUpdateStorage) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStorage) Type() string {
	return TypeMsgUpdateStorage
}

func (msg *MsgUpdateStorage) GetSigners() []sdk.AccAddress {
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

func (msg *MsgUpdateStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid storage address (%s)", err)
	}

	if msg.Grantee != "" {
		_, err = sdk.AccAddressFromBech32(msg.Grantee)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
		}
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteStorage{}

func NewMsgDeleteStorage(
	address string,
	index string,
	grantee string,

) *MsgDeleteStorage {
	return &MsgDeleteStorage{
		Address: address,
		Index:   index,
		Grantee: grantee,
	}
}
func (msg *MsgDeleteStorage) Route() string {
	return RouterKey
}

func (msg *MsgDeleteStorage) Type() string {
	return TypeMsgDeleteStorage
}

func (msg *MsgDeleteStorage) GetSigners() []sdk.AccAddress {

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

func (msg *MsgDeleteStorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteStorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid storage address (%s)", err)
	}
	if msg.Grantee != "" {
		_, err = sdk.AccAddressFromBech32(msg.Grantee)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
		}
	}
	return nil
}
