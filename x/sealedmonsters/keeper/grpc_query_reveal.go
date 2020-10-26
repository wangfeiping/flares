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

func (k Keeper) AllReveal(c context.Context, req *types.QueryAllRevealRequest) (*types.QueryAllRevealResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var reveals []*types.MsgReveal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	revealStore := prefix.NewStore(store, types.KeyPrefix(types.RevealKey))

	pageRes, err := query.Paginate(revealStore, req.Pagination, func(key []byte, value []byte) error {
		var reveal types.MsgReveal
		if err := k.cdc.UnmarshalBinaryBare(value, &reveal); err != nil {
			return err
		}

		reveals = append(reveals, &reveal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRevealResponse{Reveal: reveals, Pagination: pageRes}, nil
}
