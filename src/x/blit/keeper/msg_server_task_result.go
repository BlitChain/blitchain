package keeper

import (
	"context"

	"blit/x/blit/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTaskResult(goCtx context.Context, msg *types.MsgCreateTaskResult) (*types.MsgCreateTaskResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTaskResult(
		ctx,
		msg.Id,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var taskResult = types.TaskResult{
		Creator:    msg.Creator,
		Id:         msg.Id,
		ExecutedOn: msg.ExecutedOn,
	}

	k.SetTaskResult(
		ctx,
		taskResult,
	)
	return &types.MsgCreateTaskResultResponse{}, nil
}

func (k msgServer) UpdateTaskResult(goCtx context.Context, msg *types.MsgUpdateTaskResult) (*types.MsgUpdateTaskResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTaskResult(
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

	var taskResult = types.TaskResult{
		Creator:    msg.Creator,
		Id:         msg.Id,
		ExecutedOn: msg.ExecutedOn,
	}

	k.SetTaskResult(ctx, taskResult)

	return &types.MsgUpdateTaskResultResponse{}, nil
}

func (k msgServer) DeleteTaskResult(goCtx context.Context, msg *types.MsgDeleteTaskResult) (*types.MsgDeleteTaskResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTaskResult(
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

	k.RemoveTaskResult(
		ctx,
		msg.Id,
	)

	return &types.MsgDeleteTaskResultResponse{}, nil
}
