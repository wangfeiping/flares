package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/nameservice/keeper"
	"github.com/wangfeiping/flares/x/nameservice/types"
)

func handleMsgCreateName(ctx sdk.Context, k keeper.Keeper, name *types.MsgName) (*sdk.Result, error) {
	k.CreateName(ctx, *name)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
