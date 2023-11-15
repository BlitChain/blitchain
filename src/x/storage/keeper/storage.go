package keeper

import (
	"context"

	"blit/x/storage/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetStorage set a specific storage in the store from its index
func (k Keeper) SetStorage(ctx context.Context, storage types.Storage) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StorageKeyPrefix))
	b := k.cdc.MustMarshal(&storage)
	store.Set(types.StorageKey(
		storage.Address,
		storage.Index,
	), b)
}

// GetStorage returns a storage from its index
func (k Keeper) GetStorage(
	ctx context.Context,
	address string,
	index string,

) (val types.Storage, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StorageKeyPrefix))

	b := store.Get(types.StorageKey(
		address,
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStorage removes a storage from the store
func (k Keeper) RemoveStorage(
	ctx context.Context,
	address string,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StorageKeyPrefix))
	store.Delete(types.StorageKey(
		address,
		index,
	))
}

// GetAllStorage returns all storage
func (k Keeper) GetAllStorage(ctx context.Context) (list []types.Storage) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StorageKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Storage
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
