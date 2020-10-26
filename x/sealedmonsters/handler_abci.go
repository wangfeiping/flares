package sealedmonsters

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/wangfeiping/flares/x/sealedmonsters/keeper"
	// "github.com/wangfeiping/flares/x/sealedmonsters/types"
)

// BeginBlockHandle check for infraction evidence or downtime of validators
// on every begin block
func BeginBlockHandle(ctx sdk.Context, req abci.RequestBeginBlock,
	k keeper.Keeper) {
	monsters := k.GetAllMonster(ctx)
	k.Logger(ctx).Debug("Begin block handle",
		"height", ctx.BlockHeight(),
		"monsters", len(monsters))
}
