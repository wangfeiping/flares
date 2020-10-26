package sealedmonsters

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/keeper"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func handleMsgCreateMonster(ctx sdk.Context, k keeper.Keeper, monster *types.MsgMonster) (*sdk.Result, error) {
	k.CreateMonster(ctx, *monster)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
