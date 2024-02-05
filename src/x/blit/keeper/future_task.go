package keeper

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"

	"blit/x/blit/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	sdkprefix "cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetFutureTask set a specific futureTask in the store from its index
func (k Keeper) SetFutureTask(ctx context.Context, futureTask types.FutureTask) (string, error) {
	futureTask.Index = string(types.FutureTaskKey(
		futureTask.Status, futureTask.ScheduledOn, futureTask.TaskId, futureTask.GasPrice,
	))

	err := k.FutureTasks.Set(ctx, futureTask.Index, futureTask)
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

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.FutureTasksKeyPrefix)
	store.Delete([]byte(index))
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
func (k Keeper) GetCurrentFutureTasks(ctx context.Context) (list []types.FutureTask) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := sdkprefix.NewStore(storeAdapter, types.FutureTasksKeyPrefix)
	iterator := storetypes.KVStoreReversePrefixIterator(store, types.FutureTaskKeyPending())

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FutureTask
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		// if scheduledOn is in the past, then it is a current task
		currentBlockTime := sdkCtx.BlockTime()
		if val.ScheduledOn.Before(currentBlockTime) {
			list = append(list, val)
		} else {
			break
		}
	}
	return
}

// GetPoolFutureTasks returns all FutureTasks that are in the pool
func (k Keeper) GetPoolFutureTasks(ctx context.Context) (list []types.FutureTask) {

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := sdkprefix.NewStore(storeAdapter, types.FutureTasksKeyPrefix)
	iterator := storetypes.KVStoreReversePrefixIterator(store, types.FutureTaskKeyPool())

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FutureTask
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		list = append(list, val)
	}

	return
}

// RunTask run the task
func (k Keeper) RunTask(ctx sdk.Context, task *types.Task) {

	taskStartingGas := ctx.GasMeter().GasConsumed()
	cachedCtx, Write := ctx.CacheContext()
	cachedCtx = cachedCtx.WithGasMeter(storetypes.NewGasMeter(task.TaskGasLimit))

	func() {
		defer func() {
			if r := recover(); r != nil {
				err := handleTaskRecovery(r, cachedCtx)

				k.Logger().Error("recovered from panic", "error", err)
				task.ErrorLog = fmt.Sprintf("Err %s", err)
				task.Results = nil
				// TODO disable task if needed
				// Set the error on the task result
			} else {
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
			ctx.GasMeter().ConsumeGas(
				cashedCtxConsumedGas, "block gas meter",
			)

			if ctx.GasMeter().GasConsumed() < taskStartingGas {
				panic(storetypes.ErrorGasOverflow{Descriptor: "tx gas summation"})
			}
		}()

		// TODO: Run the task

		// GetMessages returns the cache values from the MsgExecAuthorized.Msgs if present.
		msgs, err := task.GetTaskMessages()
		if err != nil {
			panic(fmt.Errorf("failed to get task messages: %w", err))
		}

		// Store the results as strings

		results := make([]*sdk.Result, len(msgs))
		// Run the messages
		for i, msg := range msgs {
			fmt.Println("msg", i, msg)

			m, ok := msg.(sdk.HasValidateBasic)
			if ok {
				if err := m.ValidateBasic(); err != nil {
					task.ErrorLog = fmt.Sprintf("Error validatating message %d: %s", i, err)
					return
				}
			}

			handler := k.Router.Handler(msg)
			if handler == nil {
				panic(fmt.Errorf("unrecognized task message type: %T", msg))
			}

			r, err := handler(cachedCtx, msg)
			if err != nil {
				task.ErrorLog = fmt.Sprintf("Error on message %d: %s", i, err)
				return
			} // Handler should always return non-nil sdk.Result.
			if r == nil {
				panic(fmt.Errorf("handler %T returned nil Result", handler))
			}

			// Store the result
			results[i] = r
		}

		task.Results = results
	}()

}

// RunTasks runs the tasks that are in the pool
func (k Keeper) RunTasks(goCtx context.Context) error {

	ctx := sdk.UnwrapSDKContext(goCtx)
	// Loop over current future tasks
	currentFutureTasks := k.GetCurrentFutureTasks(ctx)
	for _, futureTask := range currentFutureTasks {
		fmt.Println("currentFutureTask", futureTask)
		k.RemoveFutureTask(ctx, futureTask.Index)
		futureTask.Status = types.FutureTaskStatus_POOL
		index, err := k.SetFutureTask(ctx, futureTask)
		fmt.Println("index", index)
		if err != nil {
			return err
		}
	}

	// Loop over pool future tasks
	poolFutureTasks := k.GetPoolFutureTasks(ctx)
	for _, futureTask := range poolFutureTasks {
		fmt.Println("poolFutureTask", futureTask)
		// get the task from the store
		k.RemoveFutureTask(ctx, futureTask.Index)
		task, err := k.GetTaskById(ctx, futureTask.TaskId)
		if err != nil {
			fmt.Println("error getting task", err)
			continue
		}

		if err != nil {
			return err
		}

		// if the task.ExpireAfter is in the past, then remove it from the pool
		currentBlockTime := sdk.UnwrapSDKContext(ctx).BlockTime()
		if task.ExpireAfter.Before(currentBlockTime) {
			fmt.Println("removed from pool", futureTask)
			// TODO write task result with error
			continue
		}

		// Run the task
		k.RunTask(ctx, task)

		fmt.Println("task", fmt.Sprintf("%#v", task))
		// Create the TaskResult
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
		// print stack debug
		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))

		return sdkerrors.ErrPanic.Wrapf("Unhandled Excpetion")
	}
}
