package keeper

import (
	"context"
	"fmt"

	"blit/x/blit/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetTask(ctx context.Context, task *types.Task) error {
	address, err := sdk.AccAddressFromBech32(task.Address)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get address from bech32")
	}
	return k.Tasks.Set(
		ctx,
		collections.Join(address, task.Id),
		*task,
	)
}

func (k Keeper) GetTask(ctx context.Context, creator string, id uint64) (types.Task, error) {
	address, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		return types.Task{}, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get address from bech32")
	}
	task, err := k.Tasks.Get(
		ctx,
		collections.Join(address, id),
	)
	if err != nil {
		return types.Task{}, err
	}
	return task, nil

}

func (k Keeper) GetTaskById(ctx context.Context, id uint64) (task *types.Task, err error) {

	key, err := k.Tasks.Indexes.Id.MatchExact(
		ctx,
		id,
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if !key.Valid() {
		return nil, status.Error(codes.NotFound, "not found")
	}
	fullKey, err := key.PrimaryKey()
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	t, err := k.Tasks.Get(ctx, fullKey)
	if err != nil {
		fmt.Println(fmt.Sprintf("err gettings task by fullkey %s", err))
		return nil, status.Error(codes.NotFound, "not found")
	}
	return &t, nil
}

// DeleteTask deletes a task
func (k Keeper) RemoveTaskById(ctx context.Context, id uint64) error {
	key, err := k.Tasks.Indexes.Id.MatchExact(
		ctx,
		id,
	)
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid request")
	}

	fullKey, err := key.PrimaryKey()
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid request")
	}
	err = k.Tasks.Remove(
		ctx,
		fullKey,
	)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to remove task")
	}
	return nil
}

// GetAllTask returns all task
func (k Keeper) GetAllTask(ctx context.Context) ([]types.Task, error) {
	iter, err := k.Tasks.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}
	tasks, err := iter.Values()
	if err != nil {
		return nil, err
	}

	return tasks, err
}
