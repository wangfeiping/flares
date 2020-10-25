package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	flareskeeper "github.com/wangfeiping/flares/x/flares/keeper"
	flarestypes "github.com/wangfeiping/flares/x/flares/types"

	"github.com/wangfeiping/flares/x/nameservice/types"
)

type (
	Keeper struct {
		cdc          codec.Marshaler
		storeKey     sdk.StoreKey
		memKey       sdk.StoreKey
		flareskeeper flareskeeper.Keeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey,
	k flareskeeper.Keeper) *Keeper {
	nskeeper := &Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		flareskeeper: k,
	}
	k.RegisterContractClearing(types.ModuleName, nskeeper.ContractClearing)
	return nskeeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ContractClearing(ctx sdk.Context,
	contract flarestypes.MsgContract) bool {
	recs := k.flareskeeper.GetContractTransfers(ctx, contract.Receiver)
	if len(recs) <= 0 {
		return false
	}
	// check board & compare
	max := recs[0]
	k.flareskeeper.CheckBoard(contract, &max)
	for record := range recs[1:] {
		k.flareskeeper.CheckBoard(contract, &record)
		if record > max {
			k.flareskeeper.Return(contract, max)
			max = record
		}
	}

	// check base price
	// TODO maybe it needs to be exchanged

	// transfer/send
	k.flareskeeper.Deal(contract, max)
	// set the owner of name

	return true
}
