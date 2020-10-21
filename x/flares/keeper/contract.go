package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/flares/types"
)

func (k Keeper) CreateContract(ctx sdk.Context, contract types.MsgContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKey))
	key := fmt.Sprintf("%s-%s-%s-%s", types.ContractKey,
		contract.Creator.String(), contract.Key, contract.Id)
	contract.Receiver = AccAddressString(types.ModuleName, key).String()

	b := k.cdc.MustMarshalBinaryBare(&contract)
	store.Set(types.KeyPrefix(key), b)
}

func (k Keeper) GetAllContract(ctx sdk.Context) (msgs []types.MsgContract) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ContractKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.ContractKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgContract
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
