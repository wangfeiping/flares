package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/nameservice/types"
)

func (k Keeper) CreateWhois(ctx sdk.Context, whois types.MsgWhois) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	b := k.cdc.MustMarshalBinaryBare(&whois)
	whoisKey := fmt.Sprintf("%s-%s", types.WhoisKey, whois.Value)
	store.Set(types.KeyPrefix(whoisKey), b)
}

func (k Keeper) GetWhois(ctx sdk.Context, value string) (types.MsgWhois, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	whoisKey := fmt.Sprintf("%s-%s", types.WhoisKey, value)

	var msg types.MsgWhois
	if bz := store.Get(types.KeyPrefix(whoisKey)); bz != nil {
		k.cdc.MustUnmarshalBinaryBare(bz, &msg)
		return msg, nil
	}
	return msg, types.ErrWhoisNotFound
}

func (k Keeper) GetAllWhois(ctx sdk.Context) (msgs []types.MsgWhois) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.WhoisKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgWhois
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
