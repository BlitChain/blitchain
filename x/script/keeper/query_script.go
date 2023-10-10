package keeper

import (
	"context"

	"blit/x/script/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ScriptAll(goCtx context.Context, req *types.QueryAllScriptRequest) (*types.QueryAllScriptResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var scripts []types.Script
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	scriptStore := prefix.NewStore(store, types.KeyPrefix(types.ScriptKeyPrefix))

	pageRes, err := query.Paginate(scriptStore, req.Pagination, func(key []byte, value []byte) error {
		var script types.Script
		if err := k.cdc.Unmarshal(value, &script); err != nil {
			return err
		}

		scripts = append(scripts, script)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllScriptResponse{Script: scripts, Pagination: pageRes}, nil
}

func (k Keeper) Script(goCtx context.Context, req *types.QueryGetScriptRequest) (*types.QueryGetScriptResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetScript(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetScriptResponse{Script: val}, nil
}
