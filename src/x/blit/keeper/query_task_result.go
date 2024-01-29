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

func (k Keeper) TaskResultAll(ctx context.Context, req *types.QueryAllTaskResultRequest) (*types.QueryAllTaskResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var taskResults []types.TaskResult

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	taskResultStore := prefix.NewStore(store, types.KeyPrefix(types.TaskResultKeyPrefix))

	pageRes, err := query.Paginate(taskResultStore, req.Pagination, func(key []byte, value []byte) error {
		var taskResult types.TaskResult
		if err := k.cdc.Unmarshal(value, &taskResult); err != nil {
			return err
		}

		taskResults = append(taskResults, taskResult)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTaskResultResponse{TaskResult: taskResults, Pagination: pageRes}, nil
}

func (k Keeper) TaskResult(ctx context.Context, req *types.QueryGetTaskResultRequest) (*types.QueryGetTaskResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetTaskResult(
		ctx,
		req.Id,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTaskResultResponse{TaskResult: val}, nil
}
