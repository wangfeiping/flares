package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/flares/types"
)

func (k Keeper) CreateBoard(ctx sdk.Context, board types.MsgBoard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	key := fmt.Sprintf("%s-%s-%s/%s-%s", types.BoardKey,
		board.Creator.String(), board.Accept, board.Base, board.Id)

	board.BaseAddress = AccAddressString(types.ModuleName,
		fmt.Sprintf("%s-%s-%s/%s-%s", types.BoardKey,
			board.Creator.String(), board.Base, board.Accept, board.Id)).String()
	board.AcceptAddress = AccAddressString(types.ModuleName, key).String()

	b := k.cdc.MustMarshalBinaryBare(&board)
	store.Set(types.KeyPrefix(key), b)
}

func (k Keeper) CheckBoard(ctx sdk.Context,
	contract *types.MsgContract, record *types.MsgContractTransferRecord) {
	// TODO board logic: value discover
	record.Voucher = 100
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
