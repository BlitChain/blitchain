package keeper

import (
	"blit/x/storage/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStorage set a specific storage in the store from its index
func (k Keeper) SetStorage(ctx sdk.Context, storage types.Storage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageKeyPrefix))
	b := k.cdc.MustMarshal(&storage)
	store.Set(types.StorageKey(
		storage.Address,
		storage.Index,
	), b)
}

// GetStorage returns a storage from its index
func (k Keeper) GetStorage(
	ctx sdk.Context,
	address string,
	index string,

) (val types.Storage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageKeyPrefix))

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
	ctx sdk.Context,
	address string,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageKeyPrefix))
	store.Delete(types.StorageKey(
		address,
		index,
	))
}

// GetAllStorage returns all storage
func (k Keeper) GetAllStorage(ctx sdk.Context) (list []types.Storage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StorageKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Storage
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
