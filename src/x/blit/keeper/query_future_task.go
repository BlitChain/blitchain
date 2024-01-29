package keeper

import (
	"context"

	"blit/x/blit/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FutureTaskAll(ctx context.Context, req *types.QueryAllFutureTaskRequest) (*types.QueryAllFutureTaskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var futureTasks []types.FutureTask

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	futureTaskStore := prefix.NewStore(store, types.KeyPrefix(types.FutureTaskKeyPrefix))

	pageRes, err := query.Paginate(futureTaskStore, req.Pagination, func(key []byte, value []byte) error {
		var futureTask types.FutureTask
		if err := k.cdc.Unmarshal(value, &futureTask); err != nil {
			return err
		}

		futureTasks = append(futureTasks, futureTask)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFutureTaskResponse{FutureTask: futureTasks, Pagination: pageRes}, nil
}

func (k Keeper) FutureTask(ctx context.Context, req *types.QueryGetFutureTaskRequest) (*types.QueryGetFutureTaskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetFutureTask(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFutureTaskResponse{FutureTask: val}, nil
}
