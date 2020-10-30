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

func (k Keeper) AllMonster(c context.Context, req *types.QueryAllMonsterRequest) (*types.QueryAllMonsterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var monsters []*types.MsgMonster
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	monsterStore := prefix.NewStore(store, types.KeyPrefix(types.MonsterKey))
	monsterStore = prefix.NewStore(monsterStore, types.KeyPrefix(types.MonsterKey))

	pageRes, err := query.Paginate(monsterStore, req.Pagination, func(key []byte, value []byte) error {
		var monster types.MsgMonster
		if err := k.cdc.UnmarshalBinaryBare(value, &monster); err != nil {
			return err
		}

		monsters = append(monsters, &monster)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMonsterResponse{Monster: monsters, Pagination: pageRes}, nil
}
