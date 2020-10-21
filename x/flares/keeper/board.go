package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/flares/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) CreateBoard(ctx sdk.Context, board types.MsgBoard) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	b := k.cdc.MustMarshalBinaryBare(&board)
	store.Set(types.KeyPrefix(types.BoardKey), b)
}

func (k Keeper) GetAllBoard(ctx sdk.Context) (msgs []types.MsgBoard) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.BoardKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgBoard
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
