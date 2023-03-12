package keeper

import (
	"encoding/binary"

	"github.com/0xknstntn/fluxo/x/oracle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetValidatorCount get the total number of validator
func (k Keeper) GetValidatorCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ValidatorCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetValidatorCount set the total number of validator
func (k Keeper) SetValidatorCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ValidatorCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendValidator appends a validator in the store with a new id and update the count
func (k Keeper) AppendValidator(
	ctx sdk.Context,
	validator types.Validator,
) uint64 {
	// Create the validator
	count := k.GetValidatorCount(ctx)

	// Set the ID of the appended value
	validator.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	appendedValue := k.cdc.MustMarshal(&validator)
	store.Set(GetValidatorIDBytes(validator.Id), appendedValue)

	// Update validator count
	k.SetValidatorCount(ctx, count+1)

	return count
}

// SetValidator set a specific validator in the store
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	b := k.cdc.MustMarshal(&validator)
	store.Set(GetValidatorIDBytes(validator.Id), b)
}

// GetValidator returns a validator from its id
func (k Keeper) GetValidator(ctx sdk.Context, id uint64) (val types.Validator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	b := store.Get(GetValidatorIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveValidator removes a validator from the store
func (k Keeper) RemoveValidator(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	store.Delete(GetValidatorIDBytes(id))
}

// GetAllValidator returns all validator
func (k Keeper) GetAllValidator(ctx sdk.Context) (list []types.Validator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ValidatorKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Validator
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetValidatorIDBytes returns the byte representation of the ID
func GetValidatorIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetValidatorIDFromBytes returns ID in uint64 format from a byte array
func GetValidatorIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
