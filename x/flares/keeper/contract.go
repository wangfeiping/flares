package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/wangfeiping/flares/x/flares/types"
)

type ContractClearing func(ctx sdk.Context, msg types.MsgContract) bool

var (
	cacheContractClearings map[string]ContractClearing = make(map[string]ContractClearing, 0)
)

func (k Keeper) RegisterContractClearing(module string, cc ContractClearing) {
	cacheContractClearings[module] = cc
}

func (k Keeper) CreateContract(ctx sdk.Context, contract types.MsgContract) error {
	store := k.getContractStore(ctx)
	contractKey := BuildContractKey(&contract)

	if store.Has(types.KeyPrefix(contractKey)) {
		k.Logger(ctx).
			Error(types.ErrContractExists.Error(), ": ", contractKey)
		return types.ErrContractExists
	}

	contract.Receiver = AccAddressString(types.ModuleName,
		fmt.Sprintf("%s%s", contractKey, contract.Id)).String()
	contract.Height = uint64(ctx.BlockHeight())
	bz := k.cdc.MustMarshalBinaryBare(&contract)

	// create the contract
	store.Set(types.KeyPrefix(contractKey), bz)

	// save a record for the contract receiver query
	k.getContractReceiverStore(ctx).
		Set(types.KeyPrefix(contract.Receiver), []byte(contractKey))

	return nil
}

func (k Keeper) closeContract(ctx sdk.Context,
	msg *types.MsgContract, err *sdkerrors.Error) {
	store := k.getContractStore(ctx)
	contractKey := BuildContractKey(msg)

	store.Delete(types.KeyPrefix(contractKey))

	if err != nil {
		msg.Code = err.ABCICode()
		msg.Result = err.Error()
		bz := k.cdc.MustMarshalBinaryBare(msg)

		contractKey = BuildFailContractKey(msg)
		store.Set(types.KeyPrefix(contractKey), bz)
		return
	}
	msg.Code = 0
	msg.Result = "success"
	bz := k.cdc.MustMarshalBinaryBare(msg)

	contractKey = BuildSuccessContractKey(msg)
	store.Set(types.KeyPrefix(contractKey), bz)
}

func (k Keeper) ClearingContract(ctx sdk.Context,
	moduleName string, msg *types.MsgContract) error {
	contractClearing := cacheContractClearings[msg.Module]
	if contractClearing == nil {
		ctx.Logger().With("module", moduleName).
			Error(types.ErrContractClearingNotFound.Error(),
				"height", ctx.BlockHeight(),
				"module", msg.Module, "contract", msg.Key)
		k.closeContract(ctx, msg, types.ErrContractClearingNotFound)
		return types.ErrContractClearingNotFound
	}
	if contractClearing(ctx, *msg) {
		ctx.Logger().With("module", moduleName).
			Info("the contract clearing was successful",
				"height", ctx.BlockHeight(),
				"module", msg.Module, "contract", msg.Key)
		k.closeContract(ctx, msg, nil)
		return nil
	}
	ctx.Logger().With("module", moduleName).
		Error(types.ErrContractClearingFailed.Error(),
			"height", ctx.BlockHeight(),
			"module", msg.Module, "contract", msg.Key)
	k.closeContract(ctx, msg, types.ErrContractClearingFailed)
	return types.ErrContractClearingFailed
}

// CheckContractBottom do some checks:
// check contract bottom
// check if the base price is met.
func (k Keeper) CheckContractBottom(msg *types.MsgContract, amount sdk.Coin) error {

	// return types.ErrInvalidAmount
	return nil
}

func (k Keeper) Return(ctx sdk.Context,
	contract *types.MsgContract, record *types.MsgContractTransferRecord) error {
	return nil
}

func (k Keeper) Deal(ctx sdk.Context,
	contract *types.MsgContract, record *types.MsgContractTransferRecord) error {
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
	return msg, types.ErrContractNotFound
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

func (k Keeper) CheckContractReceiver(ctx sdk.Context, accAddr string) []byte {
	return k.getContractReceiverStore(ctx).
		Get(types.KeyPrefix(accAddr))
}

func BuildContractKey(contract *types.MsgContract) string {
	return fmt.Sprintf("%s-%s", types.ContractKey, contract.Key)
}

func BuildSuccessContractKey(contract *types.MsgContract) string {
	return fmt.Sprintf("%s-%s", types.SuccessContractKey, contract.Key)
}

func BuildFailContractKey(contract *types.MsgContract) string {
	return fmt.Sprintf("%s-%s", types.FailContractKey, contract.Key)
}

func (k Keeper) getContractStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.ContractKey))
}

func (k Keeper) getContractReceiverStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey),
		types.KeyPrefix(types.ContractReceiverKey))
}
