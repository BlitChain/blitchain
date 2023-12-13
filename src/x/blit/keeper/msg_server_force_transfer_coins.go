package keeper

import (
	"context"

	"blit/x/blit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k msgServer) ForceTransferCoins(goCtx context.Context, msg *types.MsgForceTransferCoins) (*types.MsgForceTransferCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := ValidateDenomOwner(msg.Amount.Denom)
	if err != nil {
		return nil, err
	}

	ownerAddress, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return nil, err
	}

	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	toAddress, err := sdk.AccAddressFromBech32(msg.ToAddress)
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

		var execResp types.MsgForceTransferCoinsResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &execResp)
		return &execResp, nil
	}

	bankMsgKeeper := bankkeeper.NewMsgServerImpl(k.bankKeeper.(bankkeeper.Keeper))
	msgSend := banktypes.NewMsgSend(fromAddress, toAddress, sdk.NewCoins(msg.Amount))
	_, err = bankMsgKeeper.Send(goCtx, msgSend)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &types.MsgForceTransferCoinsResponse{}, nil

}
