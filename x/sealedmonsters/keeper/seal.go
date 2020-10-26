package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func (k Keeper) CreateSeal(ctx sdk.Context, seal types.MsgSeal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SealKey))
	b := k.cdc.MustMarshalBinaryBare(&seal)
	store.Set(types.KeyPrefix(types.SealKey), b)
}

func (k Keeper) GetAllSeal(ctx sdk.Context) (msgs []types.MsgSeal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SealKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.SealKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgSeal
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
