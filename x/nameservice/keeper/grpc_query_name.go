package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/wangfeiping/flares/x/nameservice/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllName(c context.Context, req *types.QueryAllNameRequest) (*types.QueryAllNameResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var names []*types.MsgName
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nameStore := prefix.NewStore(store, types.KeyPrefix(types.NameKey))

	pageRes, err := query.Paginate(nameStore, req.Pagination, func(key []byte, value []byte) error {
		var name types.MsgName
		if err := k.cdc.UnmarshalBinaryBare(value, &name); err != nil {
			return err
		}

		names = append(names, &name)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNameResponse{Name: names, Pagination: pageRes}, nil
}
