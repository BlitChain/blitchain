package keeper

import (
	"context"
	"time"

	"blit/x/blit/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTask(goCtx context.Context, msg *types.MsgCreateTask) (*types.MsgCreateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}
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

	if msg.Messages == nil || len(msg.Messages) == 0 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "messages must be non-empty")
	}

	if msg.MinimumInterval.Seconds() < 1 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "interval must be at least 1 second")
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

	if msg.TaskGasFee.Denom != "ublit" {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "task gas fee must be in ublit")
	}

	if msg.MaxRuns < 1 {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "max runs must be positive")
	}

	var task = &types.Task{
		Creator:         msg.Creator,
		Id:              taskId,
		ActivateAfter:   msg.ActivateAfter,
		ExpireAfter:     msg.ExpireAfter,
		MinimumInterval: msg.MinimumInterval,
		MaxRuns:         msg.MaxRuns,
		DisableOnError:  msg.DisableOnError,
		Enabled:         msg.Enabled,
		TaskGasLimit:    msg.TaskGasLimit,
		TaskGasFee:      msg.TaskGasFee,
		Messages:        msg.Messages,
		TotalRuns:       0,
	}

	// Calculate gas price from gas fee / gas limit
	gasPrice := sdk.NewDecCoinsFromCoins((task.TaskGasFee)).QuoDec(math.LegacyNewDec(int64(task.TaskGasLimit)))[0]

	futureTask := &types.FutureTask{
		TaskId:      task.Id,
		ScheduledOn: task.ActivateAfter.Truncate(60 * time.Second),
		Status:      types.FutureTaskStatus_PENDING,
		GasPrice:    &gasPrice,
	}
	futureTaskIndex, err := k.SetFutureTask(ctx, futureTask)

	task.FutureTaskIndex = futureTaskIndex

	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to set future task")
	}

	err = k.SetTask(
		ctx,
		task,
	)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to set task")
	}
	return &types.MsgCreateTaskResponse{
		Id: taskId,
	}, nil
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
