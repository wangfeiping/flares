package ace

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/ace/x/ace/keeper"
	"github.com/wangfeiping/ace/x/ace/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
        // this line is used by starport scaffolding # 1
case *types.MsgSeed:
return handleMsgCreateSeed(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
