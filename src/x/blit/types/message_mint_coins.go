package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintCoins{}

func NewMsgMintCoins(owner sdk.AccAddress, amount sdk.Coin, grantee sdk.AccAddress) *MsgMintCoins {
	return &MsgMintCoins{
		Amount:  amount,
		Grantee: grantee.String(),
	}
}

func (msg *MsgMintCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
	}
	return nil
}
