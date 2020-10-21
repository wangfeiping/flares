package flares

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

func handleMsgCreateContract(ctx sdk.Context, k keeper.Keeper, contract *types.MsgContract) (*sdk.Result, error) {
	k.CreateContract(ctx, *contract)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
