package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func (k Keeper) CreateReveal(ctx sdk.Context, reveal types.MsgReveal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevealKey))
	b := k.cdc.MustMarshalBinaryBare(&reveal)
	store.Set(types.KeyPrefix(types.RevealKey), b)
}

func (k Keeper) GetAllReveal(ctx sdk.Context) (msgs []types.MsgReveal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RevealKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.RevealKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgReveal
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
