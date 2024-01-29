package keeper

import (
	"context"

	"blit/x/blit/types"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetFutureTask set a specific futureTask in the store from its index
func (k Keeper) SetFutureTask(ctx context.Context, futureTask types.FutureTask) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FutureTaskKeyPrefix))
	b := k.cdc.MustMarshal(&futureTask)
	store.Set(types.FutureTaskKey(
		futureTask.Index,
	), b)
}

// GetFutureTask returns a futureTask from its index
func (k Keeper) GetFutureTask(
	ctx context.Context,
	index string,

) (val types.FutureTask, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FutureTaskKeyPrefix))

	b := store.Get(types.FutureTaskKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFutureTask removes a futureTask from the store
func (k Keeper) RemoveFutureTask(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FutureTaskKeyPrefix))
	store.Delete(types.FutureTaskKey(
		index,
	))
}

// GetAllFutureTask returns all futureTask
func (k Keeper) GetAllFutureTask(ctx context.Context) (list []types.FutureTask) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FutureTaskKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FutureTask
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
