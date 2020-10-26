package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	flarestypes "github.com/wangfeiping/flares/x/flares/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func (k Keeper) CreateMonster(ctx sdk.Context, monster types.MsgMonster) error {
	contract := NewContract(monster.Creator)
	if err := k.flaresKeeper.CreateContract(ctx, contract); err != nil &&
		err != flarestypes.ErrContractExists {
		return err
	}
	macc, err := sdk.AccAddressFromBech32(contract.Receiver)
	if err != nil {
		return err
	}
	amt, err := sdk.ParseCoins(monster.Reward)
	if err != nil {
		return err
	}
	if err = k.bankKeeper.SendCoins(ctx, monster.Creator, macc, amt); err != nil {
		return err
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonsterKey))
	b := k.cdc.MustMarshalBinaryBare(&monster)
	key := types.KeyPrefix(types.MonsterKey)
	k.Logger(ctx).Info("created a monster", "key", string(key))
	store.Set(key, b)

	return nil
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
