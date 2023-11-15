package keeper

import (
	"context"

	"blit/x/script/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Scripts(ctx context.Context, req *types.QueryScriptsRequest) (*types.QueryScriptsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var scripts []types.Script

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
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

	return &types.QueryScriptsResponse{Script: scripts, Pagination: pageRes}, nil
}

func (k Keeper) Script(goCtx context.Context, req *types.QueryScriptRequest) (*types.QueryScriptResponse, error) {
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

	return &types.QueryScriptResponse{Script: val}, nil
}
