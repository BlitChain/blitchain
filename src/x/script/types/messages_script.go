package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (msg *MsgCreateScript) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.MsgTypePermissions) > 0 {
		_, err = sdk.AccAddressFromBech32(msg.Grantee)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
		}
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

func (msg *MsgUpdateScript) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address address (%s)", err)
	}
	return nil
}
