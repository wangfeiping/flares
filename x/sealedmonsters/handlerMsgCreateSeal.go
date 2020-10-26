package sealedmonsters

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/keeper"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func handleMsgCreateSeal(ctx sdk.Context, k keeper.Keeper, seal *types.MsgSeal) (*sdk.Result, error) {
	k.CreateSeal(ctx, *seal)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
