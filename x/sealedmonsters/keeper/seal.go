package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	flareskeeper "github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func (k Keeper) CreateSeal(ctx sdk.Context, seal types.MsgSeal) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SealKey))

	seal.Scavenger = seal.Creator.String()

	contract := *NewContract(seal.Creator)
	contract, err := k.flaresKeeper.GetContract(ctx, flareskeeper.BuildContractKey(&contract))
	macc, err := sdk.AccAddressFromBech32(contract.Receiver)
	if err != nil {
		return err
	}
	amt, err := sdk.ParseCoins(seal.Amount)
	if err != nil {
		return err
	}
	if err = k.bankKeeper.SendCoins(ctx, seal.Creator, macc, amt); err != nil {
		return err
	}

	b := k.cdc.MustMarshalBinaryBare(&seal)

	key := types.KeyPrefix(fmt.Sprintf("%s-%s", types.SealKey, seal.Id))
	k.Logger(ctx).Info("created a seal", "key", string(key))
	store.Set(key, b)

	return nil
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
