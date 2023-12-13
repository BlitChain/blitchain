package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgForceTransferCoins{}

func NewMsgForceTransferCoins(grantee sdk.AccAddress, amount sdk.Coin, fromAddress sdk.AccAddress, toAddress sdk.AccAddress) *MsgForceTransferCoins {
	return &MsgForceTransferCoins{
		Grantee:     grantee.String(),
		Amount:      amount,
		FromAddress: fromAddress.String(),
		ToAddress:   toAddress.String(),
	}
}

func (msg *MsgForceTransferCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Grantee)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
	}
	return nil
}
