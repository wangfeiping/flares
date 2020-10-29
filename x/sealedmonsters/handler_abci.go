package sealedmonsters

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/wangfeiping/flares/x/sealedmonsters/keeper"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

// BeginBlockHandle check for infraction evidence or downtime of validators
// on every begin block
func BeginBlockHandle(ctx sdk.Context, req abci.RequestBeginBlock,
	k keeper.Keeper) {
	monsters := k.GetAllMonster(ctx)
	for _, monster := range monsters {
		if uint64(ctx.BlockHeight()) >= monster.Height+uint64(monster.DurationHeight) {
			k.Logger(ctx).Info("game over", "height", ctx.BlockHeight())
			c, err := k.FlaresKeeper.GetContract(ctx, monster.ContractKey)
			if err != nil {
				k.Logger(ctx).Error(err.Error())
				panic(err)
			}
			k.FlaresKeeper.ClearingContract(ctx, types.ModuleName, &c)
		}
	}
	k.Logger(ctx).Debug("Begin block handle",
		"height", ctx.BlockHeight(),
		"monsters", len(monsters))
}
