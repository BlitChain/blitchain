package keeper

import (
	"context"
	"fmt"

	"blit/x/script/types"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const MAX_QUERY_GAS = uint64(10000000)

func (k Keeper) Web(goCtx context.Context, req *types.QueryWebRequest) (*types.QueryWebResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ctx, _ = ctx.CacheContext()
	if ctx.GasMeter().Limit() < 1 || ctx.GasMeter().Limit() > MAX_QUERY_GAS {
		ctx = ctx.WithGasMeter(storetypes.NewGasMeter(MAX_QUERY_GAS))
	}
	val, found := k.RunWeb(ctx, req.Address, req.Httprequest)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Script at address '%v' not set", req.Address))
	}

	return &types.QueryWebResponse{Httpresponse: val}, nil
}
