package keeper

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"

	"blit/x/blit/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdkprefix "cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// SetFutureTask set a specific futureTask in the store from its index
func (k Keeper) SetFutureTask(ctx context.Context, futureTask *types.FutureTask) (string, error) {

	futureTask.Index = string(types.FutureTaskKey(
		futureTask.Status, futureTask.ScheduledOn, futureTask.TaskId, futureTask.GasPrice,
	))

	err := k.FutureTasks.Set(ctx, futureTask.Index, *futureTask)
	if err != nil {
		return "", err
	}
	return futureTask.Index, nil

}

// GetFutureTask returns a futureTask from its index
func (k Keeper) GetFutureTask(
	ctx context.Context,
	index string,

) (val types.FutureTask, found bool) {

	val, err := k.FutureTasks.Get(ctx, index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return val, false
		}
	}

	return val, true
}

// RemoveFutureTask removes a futureTask from the store
func (k Keeper) RemoveFutureTask(
	ctx context.Context,
	index string,

) error {
	err := k.FutureTasks.Remove(ctx, index)
	if err != nil {
		return err
	}
	return nil

}

// GetAllFutureTask returns all futureTask
func (k Keeper) GetAllFutureTask(ctx context.Context,
) ([]types.FutureTask, error) {
	iter, err := k.FutureTasks.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}
	futureTasks, err := iter.Values()
	if err != nil {
		return nil, err
	}

	return futureTasks, err
}

// GetCurrentFutureTasks returns all FutureTasks that should be added to pool
func (k Keeper) GetCurrentFutureTasks(ctx context.Context) (list []*types.FutureTask) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := sdkprefix.NewStore(storeAdapter, types.FutureTasksKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, types.FutureTaskKeyPending())

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FutureTask
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		// if scheduledOn is in the past, then it is a current task
		currentBlockTime := sdkCtx.BlockTime()
		if val.ScheduledOn.Before(currentBlockTime) {
			list = append(list, &val)
		} else {
			break
		}
	}
	return
}

// GetPoolFutureTasks returns all FutureTasks that are in the pool
func (k Keeper) GetPoolFutureTasks(ctx context.Context) (list []*types.FutureTask) {

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := sdkprefix.NewStore(storeAdapter, types.FutureTasksKeyPrefix)
	iterator := storetypes.KVStoreReversePrefixIterator(store, types.FutureTaskKeyPool())

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FutureTask
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		list = append(list, &val)
	}

	return
}

// RunTask run the task
func (k Keeper) RunTask(ctx sdk.Context, task *types.Task) {

	taskStartingGas := ctx.BlockGasMeter().GasConsumed()
	cachedCtx, Write := ctx.CacheContext()
	cachedCtx = cachedCtx.WithGasMeter(storetypes.NewGasMeter(task.TaskGasLimit))

	func() {
		defer func() {
			lastExecutedOn := sdk.UnwrapSDKContext(cachedCtx).BlockTime()
			task.LastExecutedOn = &lastExecutedOn

			if r := recover(); r != nil {
				err := handleTaskRecovery(r, cachedCtx)

				k.Logger().Error("recovered from panic", "error", err)
				task.ErrorLog = fmt.Sprintf("Error: %s", err)
				task.Results = nil
				if task.DisableOnError {
					task.Enabled = false
				}
				// TODO disable task if needed
				// Set the error on the task result
			} else {
				// Set the next future task
				Write()
			}
			k.SetTask(ctx, task)

		}()

		// If BlockGasMeter() panics it will be caught by the above recover and will
		// return an error - in any case BlockGasMeter will consume gas past the limit.
		//
		// NOTE: This must exist in a separate defer function for the above recovery
		// to recover from this one.
		defer func() {
			cashedCtxConsumedGas := cachedCtx.GasMeter().GasConsumedToLimit()
			ctx.BlockGasMeter().ConsumeGas(
				cashedCtxConsumedGas, "block gas meter",
			)

			if ctx.BlockGasMeter().GasConsumed() < taskStartingGas {
				panic(storetypes.ErrorGasOverflow{Descriptor: "tx gas summation"})
			}
		}()

		// Charge the gas fee
		creatorAddress, err := sdk.AccAddressFromBech32(task.Address)
		if err != nil {
			task.ErrorLog = fmt.Sprintf("Failed to get task address: %s", err)
			if task.DisableOnError {
				task.Enabled = false
			}
			return
		}
		err = k.DeductTaskFee(ctx, creatorAddress, task.TaskGasFee)
		if err != nil {
			task.ErrorLog = fmt.Sprintf("Failed to deduct fee: %s", err)
			if task.DisableOnError {
				task.Enabled = false
			}
			return
		}

		// GetMessages returns the cache values from the MsgExecAuthorized.Msgs if present.
		msgs, err := task.GetTaskMessages()
		if err != nil {
			task.ErrorLog = fmt.Sprintf("Failed to get task messages: %s", err)
			if task.DisableOnError {
				task.Enabled = false
			}
			return
		}

		// Store the results as strings

		results := make([]*sdk.Result, len(msgs))
		// Run the messages
		for i, msg := range msgs {

			m, ok := msg.(sdk.HasValidateBasic)
			if ok {
				if err := m.ValidateBasic(); err != nil {
					task.ErrorLog = fmt.Sprintf("Errori: validatating message %d: %s", i, err)
					if task.DisableOnError {
						task.Enabled = false
					}
					return
				}
			}

			handler := k.Router.Handler(msg)
			if handler == nil {
				task.ErrorLog = fmt.Sprintf("Error: unrecognized task message type: %T", msg)
				if task.DisableOnError {
					task.Enabled = false
				}
				return
			}

			r, err := handler(cachedCtx, msg)
			if err != nil {
				task.ErrorLog = fmt.Sprintf("Error on message %d: %s", i, err)
				if task.DisableOnError {
					task.Enabled = false
				}
				return
			} // Handler should always return non-nil sdk.Result.

			if r == nil {
				task.ErrorLog = fmt.Sprintf("Error: handler %T returned nil Result", handler)
				if task.DisableOnError {
					task.Enabled = false
				}
				return
			}

			// Store the result
			results[i] = r
		}

		task.Results = results

		if task.Enabled && (task.MaxRuns == 0 || task.TotalRuns < task.MaxRuns) {
			gasPrice := sdk.NewDecCoinsFromCoins((task.TaskGasFee)).QuoDec(math.LegacyNewDec(int64(task.TaskGasLimit)))[0]
			nextTime := cachedCtx.BlockTime().Add(*task.MinimumInterval)
			futureTask := &types.FutureTask{
				TaskId:      task.Id,
				ScheduledOn: nextTime,
				Status:      types.FutureTaskStatus_PENDING,
				GasPrice:    &gasPrice,
			}
			futureTaskIndex, err := k.SetFutureTask(ctx, futureTask)
			if err != nil {
				task.ErrorLog = fmt.Sprintf("Error: failed to set future task: %s", err)
				if task.DisableOnError {
					task.Enabled = false
				}
				return
			}

			task.FutureTaskIndex = futureTaskIndex
		}

	}()

}

// RunTasks runs the tasks that are in the pool
func (k Keeper) RunTasks(goCtx context.Context) error {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Loop over current future tasks
	currentFutureTasks := k.GetCurrentFutureTasks(goCtx)
	for _, futureTask := range currentFutureTasks {

		k.RemoveFutureTask(ctx, futureTask.Index)
		// Set the future task to pool
		futureTask.Status = types.FutureTaskStatus_POOL
		index, err := k.SetFutureTask(ctx, futureTask)
		if err != nil {
			ctx.Logger().Error("Error: failed to set future task", "error", err)
			return err
		}

		task, err := k.GetTaskById(ctx, futureTask.TaskId)

		if err != nil {
			fmt.Println("task not found", futureTask.TaskId)
			k.RemoveFutureTask(ctx, futureTask.Index)
			continue
		}
		task.FutureTaskIndex = index
	}

	poolFutureTasks := k.GetPoolFutureTasks(goCtx)
	for i, futureTask := range poolFutureTasks {
		// max 15 tasks per block
		if i > 15 {
			ctx.Logger().Info("Max 15 tasks per block")
			break
		}

		fmt.Println(fmt.Sprintf("PoolFutureTask: %+v", futureTask))
		// get the task from the store
		task, err := k.GetTaskById(ctx, futureTask.TaskId)

		if err != nil {
			k.RemoveFutureTask(ctx, futureTask.Index)
			fmt.Println("task not found", futureTask.TaskId)
			continue
		}

		currentBlockTime := sdk.UnwrapSDKContext(ctx).BlockTime()
		if task.ExpireAfter.Before(currentBlockTime) {
			fmt.Println(fmt.Sprintf("Task Expired: ExpireAfter: %s < currentBlockTime: %s", task.ExpireAfter, currentBlockTime))
			k.RemoveFutureTask(ctx, futureTask.Index)
			task.FutureTaskIndex = ""
			if task.DisableOnError {
				task.Enabled = false
			}
			k.SetTask(ctx, task)
			continue
		}

		if !(task.MaxRuns == 0 || task.TotalRuns < task.MaxRuns) {
			fmt.Println(fmt.Sprintf("Task MaxRuns: TotalRuns: %d >= MaxRuns: %d", task.TotalRuns, task.MaxRuns))
			k.RemoveFutureTask(ctx, futureTask.Index)
			task.FutureTaskIndex = ""
			if task.DisableOnError {
				task.Enabled = false
			}
			k.SetTask(ctx, task)
			continue
		}

		if task.Enabled == false {
			fmt.Println(fmt.Sprintf("Task Disabled: Enabled: %t", task.Enabled))
			k.RemoveFutureTask(ctx, futureTask.Index)
			task.FutureTaskIndex = ""
			k.SetTask(ctx, task)
			continue
		}

		availableGas := ctx.BlockGasMeter().GasRemaining() / 2

		// Skip Task if gasRemaining < types.TaskGasLimit
		if availableGas < task.TaskGasLimit {
			fmt.Println(fmt.Sprintf("Skipping Task: availableGas: %d < TaskGasLimit: %d", availableGas, task.TaskGasLimit))
			continue
		}

		err = k.RemoveFutureTask(ctx, futureTask.Index)
		if err != nil {
			ctx.Logger().Error("Error: failed to remove future task", "error", err)
			task.Enabled = false
			task.ErrorLog = fmt.Sprintf("Error: failed to remove future task: %s", err)
			k.SetTask(ctx, task)
			continue
		}
		task.FutureTaskIndex = ""

		task.TotalRuns++
		// Run the task
		k.RunTask(ctx, task)

	}

	return nil
}

func handleTaskRecovery(r interface{}, ctx sdk.Context) error {
	switch r := r.(type) {
	case storetypes.ErrorOutOfGas:
		return errorsmod.Wrapf(sdkerrors.ErrOutOfGas,
			"ErrorOutofGas script out of gas in location: %v; gasWanted: %d, gasUsed: %d",
			r.Descriptor, ctx.GasMeter().Limit(), ctx.GasMeter().GasConsumed(),
		)

	default:
		ctx.Logger().Error("recovered from panic", "error", r)
		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		return sdkerrors.ErrPanic.Wrapf("Unhandled Excpetion")
	}
}

// DeductFees deducts fee from the given account.
func (k Keeper) DeductTaskFee(ctx sdk.Context, address sdk.AccAddress, fee sdk.Coin) error {
	if !fee.IsValid() {
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %s", fee)
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, authtypes.FeeCollectorName, sdk.NewCoins(fee))
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	return nil
}
