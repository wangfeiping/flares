package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/nameservice/keeper"
	"github.com/wangfeiping/flares/x/nameservice/types"
)

func handleMsgBuyName(ctx sdk.Context, k keeper.Keeper, whois *types.MsgWhois) (*sdk.Result, error) {
	k.CreateWhois(ctx, *whois)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
