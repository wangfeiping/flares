package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllSeal(c context.Context, req *types.QueryAllSealRequest) (*types.QueryAllSealResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var seals []*types.MsgSeal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sealStore := prefix.NewStore(store, types.KeyPrefix(types.SealKey))

	pageRes, err := query.Paginate(sealStore, req.Pagination, func(key []byte, value []byte) error {
		var seal types.MsgSeal
		if err := k.cdc.UnmarshalBinaryBare(value, &seal); err != nil {
			return err
		}

		seals = append(seals, &seal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSealResponse{Seal: seals, Pagination: pageRes}, nil
}
