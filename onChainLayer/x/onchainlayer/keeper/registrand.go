package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"onChainLayer/x/onchainlayer/types"
)

// GetRegistrandCount get the total number of registrand
func (k Keeper) GetRegistrandCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegistrandCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRegistrandCount set the total number of registrand
func (k Keeper) SetRegistrandCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RegistrandCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRegistrand appends a registrand in the store with a new id and update the count
func (k Keeper) AppendRegistrand(
	ctx sdk.Context,
	registrand types.Registrand,
) uint64 {
	// Create the registrand
	count := k.GetRegistrandCount(ctx)

	// Set the ID of the appended value
	registrand.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrandKey))
	appendedValue := k.cdc.MustMarshal(&registrand)
	store.Set(GetRegistrandIDBytes(registrand.Id), appendedValue)

	// Update registrand count
	k.SetRegistrandCount(ctx, count+1)

	return count
}

// SetRegistrand set a specific registrand in the store
func (k Keeper) SetRegistrand(ctx sdk.Context, registrand types.Registrand) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrandKey))
	b := k.cdc.MustMarshal(&registrand)
	store.Set(GetRegistrandIDBytes(registrand.Id), b)
}

// GetRegistrand returns a registrand from its id
func (k Keeper) GetRegistrand(ctx sdk.Context, id uint64) (val types.Registrand, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrandKey))
	b := store.Get(GetRegistrandIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRegistrand removes a registrand from the store
func (k Keeper) RemoveRegistrand(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrandKey))
	store.Delete(GetRegistrandIDBytes(id))
}

// GetAllRegistrand returns all registrand
func (k Keeper) GetAllRegistrand(ctx sdk.Context) (list []types.Registrand) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RegistrandKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Registrand
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRegistrandIDBytes returns the byte representation of the ID
func GetRegistrandIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRegistrandIDFromBytes returns ID in uint64 format from a byte array
func GetRegistrandIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
