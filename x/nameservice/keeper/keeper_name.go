package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/nameservice/types"
)

func (k Keeper) CreateName(ctx sdk.Context, name types.MsgName) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	b := k.cdc.MustMarshalBinaryBare(&name)
	store.Set(types.KeyPrefix(types.NameKey), b)
}

func (k Keeper) GetAllName(ctx sdk.Context) (msgs []types.MsgName) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NameKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.NameKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgName
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
