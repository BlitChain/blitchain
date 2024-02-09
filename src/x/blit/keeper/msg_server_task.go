package keeper

import (
	"context"
	"fmt"
	"time"

	"blit/x/blit/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func (k msgServer) CreateTask(goCtx context.Context, msg *types.MsgCreateTask) (*types.MsgCreateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Grantee != msg.Address {

		// Prevent looping
		originalGrantee := sdk.MustAccAddressFromBech32(msg.Grantee)
		msg.Grantee = msg.Address
		execMsg := authz.NewMsgExec(
			originalGrantee,
			[]sdk.Msg{msg},
		)

		msgExecResp, err := k.authzKeeper.Exec(ctx, &execMsg)
		if err != nil {
			return nil, err
		}

		var execResp types.MsgCreateTaskResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &execResp)
		return &execResp, nil

	}

	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid msg.Address")
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

	var task = &types.Task{
		Address:         msg.Address,
		Id:              taskId,
		ActivateAfter:   msg.ActivateAfter,
		ExpireAfter:     msg.ExpireAfter,
		MinimumInterval: msg.MinimumInterval,
		MaxRuns:         msg.MaxRuns,
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

	if err := ctx.EventManager().EmitTypedEvent(&types.EventCreateTask{TaskId: taskId}); err != nil {
		return nil, err
	}

	return &types.MsgCreateTaskResponse{
		Id: taskId,
	}, nil
}
func (k msgServer) UpdateTask(goCtx context.Context, msg *types.MsgUpdateTask) (*types.MsgUpdateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	task, err := k.GetTaskById(
		ctx,
		msg.Id,
	)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "Task not found")
	}
	if task == nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "Task not found")
	}

	if msg.Grantee != msg.Address {

		// Prevent looping
		originalGrantee := sdk.MustAccAddressFromBech32(msg.Grantee)
		msg.Grantee = msg.Address
		execMsg := authz.NewMsgExec(
			originalGrantee,
			[]sdk.Msg{msg},
		)

		msgExecResp, err := k.authzKeeper.Exec(ctx, &execMsg)
		if err != nil {
			return nil, err
		}

		var execResp types.MsgUpdateTaskResponse

		data := msgExecResp.Results[0]
		k.cdc.MustUnmarshal(data, &execResp)
		return &execResp, nil

	}

	if task.Address != msg.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("Incorrect owner, got %s, expected %s", msg.Address, task.Address))
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

	task = &types.Task{
		Address:         msg.Address,
		Id:              msg.Id,
		ActivateAfter:   msg.ActivateAfter,
		ExpireAfter:     msg.ExpireAfter,
		MinimumInterval: msg.MinimumInterval,
		MaxRuns:         msg.MaxRuns,
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
	return &types.MsgUpdateTaskResponse{}, nil
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
	if msg.Address != valFound.Address {
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
