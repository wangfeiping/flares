package handler

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

// BeginBlockHandle check for infraction evidence or downtime of validators
// on every begin block
func BeginBlockHandle(ctx sdk.Context, req abci.RequestBeginBlock,
	k keeper.Keeper) {
	contracts := k.GetAllContract(ctx)
	for _, c := range contracts {
		if k.IsAuctions(&c) {
			if uint64(ctx.BlockHeight()) >= c.Height+uint64(c.DurationHeight) {
				k.Logger(ctx).Info("contract expires", "height", ctx.BlockHeight())
				// the highest price is traded.
				// contract clearing
				k.ClearingContract(ctx, types.ModuleName, &c)
			}
		}
	}
	k.Logger(ctx).Debug("Begin block handle",
		"height", ctx.BlockHeight(),
		"contracts", len(contracts))
}
