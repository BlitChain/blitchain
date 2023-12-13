package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetDenomMetadata{}

func NewMsgSetDenomMetadata(authority string, base string, display string, name string, symbol string, uri string, uriHash string, exponent uint32, description string) *MsgSetDenomMetadata {
	return &MsgSetDenomMetadata{
		Authority:   authority,
		Base:        base,
		Display:     display,
		Name:        name,
		Symbol:      symbol,
		Uri:         uri,
		UriHash:     uriHash,
		Exponent:    exponent,
		Description: description,
	}
}

func (msg *MsgSetDenomMetadata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	return nil
}
