package keeper

import (
	"context"

	"blit/x/blit/types"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetTaskResult set a specific taskResult in the store from its index
func (k Keeper) SetTaskResult(ctx context.Context, taskResult types.TaskResult) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskResultKeyPrefix))
	b := k.cdc.MustMarshal(&taskResult)
	store.Set(types.TaskResultKey(
		taskResult.Index,
	), b)
}

// GetTaskResult returns a taskResult from its index
func (k Keeper) GetTaskResult(
	ctx context.Context,
	index string,

) (val types.TaskResult, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskResultKeyPrefix))

	b := store.Get(types.TaskResultKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTaskResult removes a taskResult from the store
func (k Keeper) RemoveTaskResult(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskResultKeyPrefix))
	store.Delete(types.TaskResultKey(
		index,
	))
}

// GetAllTaskResult returns all taskResult
func (k Keeper) GetAllTaskResult(ctx context.Context) (list []types.TaskResult) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskResultKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TaskResult
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
