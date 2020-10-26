package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func (k Keeper) CreateMonster(ctx sdk.Context, monster types.MsgMonster) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonsterKey))
	b := k.cdc.MustMarshalBinaryBare(&monster)
	key := types.KeyPrefix(fmt.Sprintf("%s-%s", types.MonsterKey, monster.SolutionHash))
	k.Logger(ctx).Info("created a monster", "key", string(key))
	store.Set(key, b)
}

func (k Keeper) GetAllMonster(ctx sdk.Context) (msgs []types.MsgMonster) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonsterKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.MonsterKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgMonster
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
