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

func (k Keeper) AllWhois(c context.Context, req *types.QueryAllWhoisRequest) (*types.QueryAllWhoisResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var whoiss []*types.MsgWhois
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	whoisStore := prefix.NewStore(store, types.KeyPrefix(types.WhoisKey))

	pageRes, err := query.Paginate(whoisStore, req.Pagination, func(key []byte, value []byte) error {
		var whois types.MsgWhois
		if err := k.cdc.UnmarshalBinaryBare(value, &whois); err != nil {
			return err
		}

		whoiss = append(whoiss, &whois)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWhoisResponse{Whois: whoiss, Pagination: pageRes}, nil
}
