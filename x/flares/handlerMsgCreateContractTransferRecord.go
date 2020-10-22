package flares

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/flares/types"
	"github.com/wangfeiping/flares/x/flares/keeper"
)

func handleMsgCreateContractTransferRecord(ctx sdk.Context, k keeper.Keeper, contractTransferRecord *types.MsgContractTransferRecord) (*sdk.Result, error) {
	k.CreateContractTransferRecord(ctx, *contractTransferRecord)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
