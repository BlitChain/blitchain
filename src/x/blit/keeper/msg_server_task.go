package keeper

import (
	"context"
	"fmt"

	"blit/x/blit/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTask(goCtx context.Context, msg *types.MsgCreateTask) (*types.MsgCreateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// log foos
	fmt.Printf("CreateTask: %v\n", msg)

	taskId, err := k.TaskID.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next task id")
	}

	var task = types.Task{
		Creator:        msg.Creator,
		Id:             taskId,
		ActivateAfter:  msg.ActivateAfter,
		ExpireAfter:    msg.ExpireAfter,
		Interval:       msg.Interval,
		MaxRuns:        msg.MaxRuns,
		DisableOnError: msg.DisableOnError,
		Enabled:        msg.Enabled,
		GasLimit:       msg.GasLimit,
		GasPrice:       msg.GasPrice,
		Messages:       msg.Messages,
	}

	k.SetTask(
		ctx,
		task,
	)
	return &types.MsgCreateTaskResponse{}, nil
}

func (k msgServer) UpdateTask(goCtx context.Context, msg *types.MsgUpdateTask) (*types.MsgUpdateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTask(
		ctx,
		msg.Id,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var task = types.Task{
		Creator:        msg.Creator,
		Id:             msg.Id,
		ActivateAfter:  msg.ActivateAfter,
		ExpireAfter:    msg.ExpireAfter,
		Interval:       msg.Interval,
		MaxRuns:        msg.MaxRuns,
		DisableOnError: msg.DisableOnError,
		Enabled:        msg.Enabled,
		GasLimit:       msg.GasLimit,
		GasPrice:       msg.GasPrice,
		Messages:       msg.Messages,
	}

	k.SetTask(
		ctx,
		task,
	)
	return &types.MsgUpdateTaskResponse{}, nil
}

func (k msgServer) DeleteTask(goCtx context.Context, msg *types.MsgDeleteTask) (*types.MsgDeleteTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTask(
		ctx,
		msg.Id,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTask(
		ctx,
		msg.Id,
	)

	return &types.MsgDeleteTaskResponse{}, nil
}
