package keeper

import (
	"context"

	"blit/x/blit/types"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TaskAll(ctx context.Context, req *types.QueryAllTaskRequest) (*types.QueryAllTaskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr, err := k.accountKeeper.AddressCodec().StringToBytes(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %s", err.Error())
	}

	tasks, pageRes, err := query.CollectionPaginate(
		ctx,
		k.Tasks,
		req.Pagination,
		func(key collections.Pair[sdk.AccAddress, uint64], value types.Task) (types.Task, error) {
			return value, nil
		},
		query.WithCollectionPaginationPairPrefix[sdk.AccAddress, uint64](addr),
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "paginate: %v", err)
	}

	return &types.QueryAllTaskResponse{Task: tasks, Pagination: pageRes}, nil

}

func (k Keeper) Task(ctx context.Context, req *types.QueryGetTaskRequest) (*types.QueryGetTaskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	task, err := k.GetTaskById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.QueryGetTaskResponse{Task: *task}, nil
}
