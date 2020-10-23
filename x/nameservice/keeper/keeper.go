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

func (k Keeper) ContractClearing(ctx sdk.Context, msg flarestypes.MsgContract) bool {
	return true
}
