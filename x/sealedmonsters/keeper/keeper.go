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
		flaresKeeper flareskeeper.Keeper
		bankKeeper   bankkeeper.Keeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey,
	flaresKeeper flareskeeper.Keeper, bankKeeper bankkeeper.Keeper) *Keeper {
	return &Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		flaresKeeper: flaresKeeper,
		bankKeeper:   bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func NewContract(creator sdk.AccAddress) *flarestypes.MsgContract {
	return flarestypes.NewMsgContract(
		creator,
		types.ModuleName, types.ModuleName, "", -1, "")
}
