package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

func handleMsgCreateContract(ctx sdk.Context, k keeper.Keeper, contract *types.MsgContract) (*sdk.Result, error) {
	if err := k.CreateContract(ctx, *contract); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
