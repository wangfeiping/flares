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
	ctx.Logger().With("module", types.ModuleName).
		Debug("Begin block handle", "height", ctx.BlockHeight())
	contracts := k.GetAllContract(ctx)
	for _, c := range contracts {
		if c.IsAuctions() {
			if uint64(ctx.BlockHeight()) >= c.Height+uint64(c.DurationHeight) {
				ctx.Logger().With("module", types.ModuleName).
					Info("contract expires", "height", ctx.BlockHeight())
				// the highest price is traded.
				// TODO contract clearing
				ctx.Logger().With("module", types.ModuleName).
					Info("contract clearing",
						"height", ctx.BlockHeight(), "contract", c.Key)
			}
		}
	}
}
