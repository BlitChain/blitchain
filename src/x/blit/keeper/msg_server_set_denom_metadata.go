package keeper

import (
	"context"

	"blit/x/blit/types"

	errorsmod "cosmossdk.io/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetDenomMetadata(goCtx context.Context, msg *types.MsgSetDenomMetadata) (*types.MsgSetDenomMetadataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	if k.GetAuthority() != msg.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), msg.Authority)
	}

	baseDenomUnit := banktypes.DenomUnit{
		Denom:    msg.Base,
		Exponent: 0,
	}
	displayDenomUnit := banktypes.DenomUnit{
		Denom:    msg.Display,
		Exponent: msg.Exponent,
	}

	denomMetadata := banktypes.Metadata{
		Description: msg.Description,
		Base:        msg.Base,
		Display:     msg.Display,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		DenomUnits:  []*banktypes.DenomUnit{&baseDenomUnit, &displayDenomUnit},
	}
	k.bankKeeper.SetDenomMetaData(goCtx, denomMetadata)

	return &types.MsgSetDenomMetadataResponse{}, nil
}
