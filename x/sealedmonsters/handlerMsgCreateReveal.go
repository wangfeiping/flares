package sealedmonsters

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/keeper"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func handleMsgCreateReveal(ctx sdk.Context, k keeper.Keeper, reveal *types.MsgReveal) (*sdk.Result, error) {
	k.CreateReveal(ctx, *reveal)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
