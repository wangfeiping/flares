package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/wangfeiping/ace/x/ace/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllSeed(c context.Context, req *types.QueryAllSeedRequest) (*types.QueryAllSeedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var seeds []*types.MsgSeed
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	seedStore := prefix.NewStore(store, types.KeyPrefix(types.SeedKey))

	pageRes, err := query.Paginate(seedStore, req.Pagination, func(key []byte, value []byte) error {
		var seed types.MsgSeed
		if err := k.cdc.UnmarshalBinaryBare(value, &seed); err != nil {
			return err
		}

		seeds = append(seeds, &seed)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSeedResponse{Seed: seeds, Pagination: pageRes}, nil
}
