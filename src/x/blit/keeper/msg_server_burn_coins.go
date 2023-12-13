package keeper

import (
	"context"

	"blit/x/blit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func (k msgServer) BurnCoins(goCtx context.Context, msg *types.MsgBurnCoins) (*types.MsgBurnCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := ValidateDenomOwner(msg.Amount.Denom)
	if err != nil {
		return nil, err
	}

	ownerAddress, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return nil, err
	}

	if msg.Grantee != owner {

		// Prevent looping
		originalGrantee := sdk.MustAccAddressFromBech32(msg.Grantee)
		msg.Grantee = ownerAddress.String()
		execMsg := authz.NewMsgExec(
			originalGrantee,
			[]sdk.Msg{msg},
		)

		msgExecResp, err := k.authzKeeper.Exec(ctx, &execMsg)
		if err != nil {
			return nil, err
		}

		var execResp types.MsgBurnCoinsResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &execResp)
		return &execResp, nil
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, ownerAddress, "transfer", sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, "transfer", sdk.NewCoins(msg.Amount))

	if err != nil {
		return nil, err
	}

	return &types.MsgBurnCoinsResponse{}, nil

}
