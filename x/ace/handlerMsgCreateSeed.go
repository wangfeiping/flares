package ace

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/ace/x/ace/types"
	"github.com/wangfeiping/ace/x/ace/keeper"
)

func handleMsgCreateSeed(ctx sdk.Context, k keeper.Keeper, seed *types.MsgSeed) (*sdk.Result, error) {
	k.CreateSeed(ctx, *seed)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
