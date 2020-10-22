package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/wangfeiping/flares/x/flares/types"
)

var (
	ErrContractExists   = sdkerrors.Register(types.ModuleName, 10001, "contract already exists")
	ErrContractNotFound = sdkerrors.Register(types.ModuleName, 10002, "contract not found")
)

func (k Keeper) CreateContract(ctx sdk.Context, contract types.MsgContract) error {
	store := k.getContractStore(ctx)
	contractKey := fmt.Sprintf("%s-%s", types.ContractKey, contract.Key)

	if store.Has(types.KeyPrefix(contractKey)) {
		ctx.Logger().With("module", types.ModuleName).
			Error(ErrContractExists.Error(), ": ", contractKey)
		return ErrContractExists
	}

	contract.Receiver = AccAddressString(types.ModuleName,
		fmt.Sprintf("%s%s", contractKey, contract.Id)).String()
	contract.Height = uint64(ctx.BlockHeight())

	b := k.cdc.MustMarshalBinaryBare(&contract)
	store.Set(types.KeyPrefix(contractKey), b)

	k.getContractReceiverStore(ctx).
		Set(types.KeyPrefix(contract.Receiver), []byte(contractKey))

	return nil
}

func (k Keeper) GetContract(ctx sdk.Context,
	contractKey string) (types.MsgContract, error) {
	store := k.getContractStore(ctx)

	var msg types.MsgContract
	if bz := store.Get(types.KeyPrefix(contractKey)); bz != nil {
		k.cdc.MustUnmarshalBinaryBare(bz, &msg)
		return msg, nil
	}
	return msg, ErrContractNotFound
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

func (k Keeper) CheckContractReceiver(ctx sdk.Context, addr sdk.AccAddress) []byte {
	return k.getContractReceiverStore(ctx).
		Get(types.KeyPrefix(addr.String()))
}

func (k Keeper) getContractStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.ContractKey))
}

func (k Keeper) getContractReceiverStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.ContractReceiverKey))
}
