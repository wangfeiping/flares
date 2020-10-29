package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	flareskeeper "github.com/wangfeiping/flares/x/flares/keeper"
	flarestypes "github.com/wangfeiping/flares/x/flares/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

type (
	Keeper struct {
		cdc          codec.Marshaler
		storeKey     sdk.StoreKey
		memKey       sdk.StoreKey
		FlaresKeeper flareskeeper.Keeper
		bankKeeper   bankkeeper.Keeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey,
	k flareskeeper.Keeper, bankKeeper bankkeeper.Keeper) *Keeper {
	sealedKeeper := &Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		bankKeeper:   bankKeeper,
		FlaresKeeper: k,
	}
	k.RegisterContractClearing(types.ModuleName, sealedKeeper.ContractClearing)
	return sealedKeeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func NewContract(creator sdk.AccAddress) *flarestypes.MsgContract {
	return flarestypes.NewMsgContract(
		creator,
		types.ModuleName, types.ModuleName, "", -1, "")
}

func (k Keeper) ContractClearing(ctx sdk.Context,
	contract flarestypes.MsgContract) bool {

	return true
}
