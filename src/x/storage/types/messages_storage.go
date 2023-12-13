package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateStorage{}

func NewMsgCreateStorage(
	address string,
	index string,
	data string,
	grantee string,
	force bool,
) *MsgCreateStorage {
	return &MsgCreateStorage{
		Address: address,
		Index:   index,
		Data:    data,
		Grantee: grantee,
		Force:   force,
	}
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
	force bool,
) *MsgUpdateStorage {
	return &MsgUpdateStorage{
		Address: address,
		Index:   index,
		Data:    data,
		Grantee: grantee,
		Force:   force,
	}
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
	force bool,
) *MsgDeleteStorage {
	return &MsgDeleteStorage{
		Address: address,
		Index:   index,
		Grantee: grantee,
		Force:   force,
	}
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
