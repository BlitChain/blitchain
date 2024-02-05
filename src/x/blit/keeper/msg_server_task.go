package keeper

import (
	"context"
	"fmt"

	"blit/x/blit/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTask(goCtx context.Context, msg *types.MsgCreateTask) (*types.MsgCreateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	taskId, err := k.TaskID.Next(ctx)

	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next task id")
	}

	// Validate attributes

	if msg.ExpireAfter.Before(ctx.BlockTime()) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "expire after must be in the future")
	}

	if msg.ExpireAfter.Before(msg.ActivateAfter) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "expire after must be after activate after")
	}

	if msg.Interval < 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "interval must be positive")
	}

	if msg.Messages == nil || len(msg.Messages) == 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "messages must be non-empty")
	}

	// XOR frequency and interval must be set
	if msg.Frequency == 0 && msg.Interval == 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "frequency or interval must be set (not neither)")
	}

	// XOR frequency and interval must be set
	if msg.Frequency != 0 && msg.Interval != 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "frequency or interval must be set (not both)")
	}

	if msg.TaskGasLimit < 1 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "task gas limit must be positive")
	}

	err = msg.TaskGasFee.Validate()
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "task gas fee must be valid")
	}

	if msg.TaskGasFee.Amount.IsZero() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "task gas fee must be non-zero")
	}

	var task = &types.Task{
		Creator:        msg.Creator,
		Id:             taskId,
		ActivateAfter:  msg.ActivateAfter,
		ExpireAfter:    msg.ExpireAfter,
		Interval:       msg.Interval,
		Frequency:      msg.Frequency,
		MaxRuns:        msg.MaxRuns,
		DisableOnError: msg.DisableOnError,
		Enabled:        msg.Enabled,
		TaskGasLimit:   msg.TaskGasLimit,
		TaskGasFee:     msg.TaskGasFee,
		Messages:       msg.Messages,
	}

	err = k.SetTask(
		ctx,
		task,
	)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to set task")
	}

	// Calculate gas price from gas fee / gas limit

	//gasPrice := sdk.NewDecCoinFromCoin(msg.TaskGasFee).Amount.Quo(math.LegacyNewDec(int64(msg.TaskGasLimit)))

	gasPrice := sdk.NewDecCoinsFromCoins((msg.TaskGasFee)).QuoDec(math.LegacyNewDec(int64(msg.TaskGasLimit)))[0]

	futureTask := types.FutureTask{
		TaskId:      task.Id,
		ScheduledOn: task.ActivateAfter,
		Status:      types.FutureTaskStatus_PENDING,
		GasPrice:    &gasPrice,
	}
	index, err := k.SetFutureTask(ctx, futureTask)
	fmt.Println("index", index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to set future task")
	}

	return &types.MsgCreateTaskResponse{}, nil
}

func (k msgServer) DeleteTask(goCtx context.Context, msg *types.MsgDeleteTask) (*types.MsgDeleteTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, err := k.GetTaskById(
		ctx,
		msg.Id,
	)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "Task not found")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	err = k.RemoveTaskById(
		ctx,
		msg.Id,
	)

	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to remove task")
	}

	return &types.MsgDeleteTaskResponse{}, nil

}
