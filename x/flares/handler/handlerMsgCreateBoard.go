package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

func handleMsgCreateBoard(ctx sdk.Context, k keeper.Keeper, board *types.MsgBoard) (*sdk.Result, error) {
	k.CreateBoard(ctx, *board)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
