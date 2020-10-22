package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/flares/types"
)

func (k Keeper) CreateContractTransferRecord(ctx sdk.Context,
	record types.MsgContractTransferRecord) {
	store := k.getContractTransferStore(ctx, record.To)

	bz := sha256.Sum256(ctx.TxBytes())
	record.Hash = hex.EncodeToString(bz[:])
	data := k.cdc.MustMarshalBinaryBare(&record)

	store.Set(types.KeyPrefix(record.Hash), data)

	return
}

func (k Keeper) GetContractTransfers(ctx sdk.Context, receiver string) (
	msgs []types.MsgContractTransferRecord) {
	store := k.getContractTransferStore(ctx, receiver)
	iterator := store.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgContractTransferRecord
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

func (k Keeper) getContractTransferStore(ctx sdk.Context, receiver string) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(fmt.Sprintf("%s%s", types.ContractTransferKey, receiver)))
}
