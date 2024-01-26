package keeper

import (
	"context"

	"blit/x/blit/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetTask set a specific task in the store from its index
func (k Keeper) SetTask(ctx context.Context, task types.Task) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskKeyPrefix))
	b := k.cdc.MustMarshal(&task)
	store.Set(types.TaskKey(
		task.Id,
	), b)
}

// GetTask returns a task from its index
func (k Keeper) GetTask(
	ctx context.Context,
	id uint64,

) (val types.Task, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskKeyPrefix))

	b := store.Get(types.TaskKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTask removes a task from the store
func (k Keeper) RemoveTask(
	ctx context.Context,
	id uint64,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskKeyPrefix))
	store.Delete(types.TaskKey(
		id,
	))
}

// GetAllTask returns all task
func (k Keeper) GetAllTask(ctx context.Context) (list []types.Task) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TaskKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Task
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
