package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/ace/x/ace/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) CreateSeed(ctx sdk.Context, seed types.MsgSeed) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SeedKey))
	b := k.cdc.MustMarshalBinaryBare(&seed)
	store.Set(types.KeyPrefix(types.SeedKey), b)
}

func (k Keeper) GetAllSeed(ctx sdk.Context) (msgs []types.MsgSeed) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SeedKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.SeedKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgSeed
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
