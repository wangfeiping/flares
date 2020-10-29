package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	flarestypes "github.com/wangfeiping/flares/x/flares/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func (k Keeper) CreateMonster(ctx sdk.Context, monster types.MsgMonster) error {
	contract := NewContract(monster.Creator)
	contract.Bottom = monster.Reward
	contractKey, err := k.FlaresKeeper.CreateContract(ctx, contract)
	if err != nil &&
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
	if err = k.FlaresKeeper.SendCoins(ctx, monster.Creator, macc, amt); err != nil {
		return err
	}
	monster.Height = uint64(ctx.BlockHeight())
	monster.DurationHeight = 100
	monster.ContractKey = contractKey

	store := k.getMonsterStore(ctx)

	b := k.cdc.MustMarshalBinaryBare(&monster)
	key := types.KeyPrefix(types.MonsterKey)
	if store.Has(key) {
		return types.ErrMonsterExists
	}
	store.Set(key, b)
	k.Logger(ctx).Info("created a monster", "key", string(key))
	return nil
}

func (k Keeper) SealedMonster(ctx sdk.Context, monster types.MsgMonster) {
	store := k.getMonsterStore(ctx)

	bz := store.Get(types.KeyPrefix(types.MonsterKey))
	store.Delete(types.KeyPrefix(types.MonsterKey))

	store.Set(types.KeyPrefix(fmt.Sprintf("%s%s", types.SealedMonsterKey, monster.Id)), bz)
}

func (k Keeper) GetAllMonster(ctx sdk.Context) (msgs []types.MsgMonster) {
	store := k.getMonsterStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.MonsterKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgMonster
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

func (k Keeper) getMonsterStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MonsterKey))
}
