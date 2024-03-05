package keeper

import (
	"context"
	"strings"

	"blit/x/blit/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FutureTaskAll(ctx context.Context, req *types.QueryAllFutureTaskRequest) (*types.QueryAllFutureTaskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	futureTasks, pageRes, err := query.CollectionFilteredPaginate(
		ctx,
		k.FutureTasks,
		req.Pagination,
		func(key string, value types.FutureTask) (bool, error) {
			return strings.HasPrefix(value.Index, req.Prefix), nil
		},
		func(key string, value types.FutureTask) (types.FutureTask, error) {
			return value, nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFutureTaskResponse{FutureTask: futureTasks, Pagination: pageRes}, err
}

func (k Keeper) FutureTask(ctx context.Context, req *types.QueryGetFutureTaskRequest) (*types.QueryGetFutureTaskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	futureTask, err := k.FutureTasks.Get(ctx, req.Index)
	if err == nil {
		return &types.QueryGetFutureTaskResponse{FutureTask: futureTask}, nil
	}
	if errors.IsOf(err, collections.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "FutureTask [%s] doesn't exist", req.Index)
	}
	return nil, status.Error(codes.Internal, err.Error())
}
