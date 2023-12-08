package keeper

import (
	"context"

	"blit/x/blit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintCoins(goCtx context.Context, msg *types.MsgMintCoins) (*types.MsgMintCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgMintCoinsResponse{}, nil
}
