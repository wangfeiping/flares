package sealedmonsters

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/sealedmonsters/keeper"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

func handleMsgCreateMonster(ctx sdk.Context, k keeper.Keeper,
	monster *types.MsgMonster) (*sdk.Result, error) {
	if err := k.CreateMonster(ctx, *monster); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateMonster),
			sdk.NewAttribute(types.AttributeKeyCreator, monster.Creator.String()),
			sdk.NewAttribute(types.AttributeSolutionHash, monster.SolutionHash),
			sdk.NewAttribute(types.AttributeDescription, monster.Description),
			sdk.NewAttribute(types.AttributeReward, monster.Reward),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
