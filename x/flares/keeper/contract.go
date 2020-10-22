package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/flares/types"
)

func (k Keeper) CreateContract(ctx sdk.Context, contract types.MsgContract) {
	store := k.getContractStore(ctx)
	key := fmt.Sprintf("%s-%s-%s-%s", types.ContractKey,
		contract.Creator.String(), contract.Key, contract.Id)
	contract.Receiver = AccAddressString(types.ModuleName, key).String()

	b := k.cdc.MustMarshalBinaryBare(&contract)
	store.Set(types.KeyPrefix(key), b)

	k.getContractReceiverStore(ctx).
		Set(types.KeyPrefix(contract.Receiver), []byte(key))
}

func (k Keeper) GetAllContract(ctx sdk.Context) (msgs []types.MsgContract) {
	store := k.getContractStore(ctx)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.ContractKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgContract
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

func (k Keeper) CheckContractReceiver(ctx sdk.Context, addr sdk.AccAddress) bool {
	return k.getContractReceiverStore(ctx).
		Has(types.KeyPrefix(addr.String()))
}

func (k Keeper) CreateContractTransferRecord(ctx sdk.Context,
	fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error {
	store := k.getContractTransferStore(ctx, toAddr.String())

	bz := sha256.Sum256(ctx.TxBytes())
	hash := hex.EncodeToString(bz[:])

	store.Set(types.KeyPrefix(hash), []byte(""))
	return nil
}

func (k Keeper) getContractStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.ContractKey))
}

func (k Keeper) getContractReceiverStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.ContractReceiverKey))
}

func (k Keeper) getContractTransferStore(ctx sdk.Context, receiver string) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(fmt.Sprintf("%s%s", types.ContractTransferKey, receiver)))
}
