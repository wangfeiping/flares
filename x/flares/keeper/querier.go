package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/wangfeiping/flares/x/flares/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListContractTransferRecord:
			ctx.Logger().With("module", types.ModuleName).
				Debug("query contract transfers: ", string(req.Data))
			return listContractTransfers(ctx, k, legacyQuerierCdc, "")
		case types.QueryListBoard:
			return listBoard(ctx, k, legacyQuerierCdc)
		case types.QueryListContract:
			return listContract(ctx, k, legacyQuerierCdc)
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}
