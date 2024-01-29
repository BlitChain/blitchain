package keeper

import (
	"context"

	"blit/x/blit/types"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateFutureTask(goCtx context.Context, msg *types.MsgCreateFutureTask) (*types.MsgCreateFutureTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetFutureTask(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var futureTask = types.FutureTask{
		Creator:     msg.Creator,
		Index:       msg.Index,
		ScheduledOn: msg.ScheduledOn,
		TaskId:      msg.TaskId,
	}

	k.SetFutureTask(
		ctx,
		futureTask,
	)
	return &types.MsgCreateFutureTaskResponse{}, nil
}

func (k msgServer) UpdateFutureTask(goCtx context.Context, msg *types.MsgUpdateFutureTask) (*types.MsgUpdateFutureTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetFutureTask(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var futureTask = types.FutureTask{
		Creator:     msg.Creator,
		Index:       msg.Index,
		ScheduledOn: msg.ScheduledOn,
		TaskId:      msg.TaskId,
	}

	k.SetFutureTask(ctx, futureTask)

	return &types.MsgUpdateFutureTaskResponse{}, nil
}

func (k msgServer) DeleteFutureTask(goCtx context.Context, msg *types.MsgDeleteFutureTask) (*types.MsgDeleteFutureTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetFutureTask(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveFutureTask(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteFutureTaskResponse{}, nil
}
