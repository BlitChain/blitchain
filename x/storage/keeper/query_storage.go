package keeper

import (
	"context"

	"blit/x/storage/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StorageAll(goCtx context.Context, req *types.QueryAllStorageRequest) (*types.QueryAllStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storages []types.Storage
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)

	storageStore := prefix.NewStore(prefix.NewStore(store, types.KeyPrefix(types.StorageKeyPrefix)), types.StorageKeyFilterPrefix(req.FilterAddress, req.FilterIndexPrefix))

	pageRes, err := query.Paginate(storageStore, req.Pagination, func(key []byte, value []byte) error {
		var storage types.Storage
		if err := k.cdc.Unmarshal(value, &storage); err != nil {
			return err
		}

		storages = append(storages, storage)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStorageResponse{Storage: storages, Pagination: pageRes}, nil
}

func (k Keeper) Storage(goCtx context.Context, req *types.QueryGetStorageRequest) (*types.QueryGetStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStorage(
		ctx,
		req.Address,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStorageResponse{Storage: val}, nil
}
