package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/wangfeiping/flares/x/flares/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllBoard(c context.Context, req *types.QueryAllBoardRequest) (*types.QueryAllBoardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var boards []*types.MsgBoard
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	boardStore := prefix.NewStore(store, types.KeyPrefix(types.BoardKey))

	pageRes, err := query.Paginate(boardStore, req.Pagination, func(key []byte, value []byte) error {
		var board types.MsgBoard
		if err := k.cdc.UnmarshalBinaryBare(value, &board); err != nil {
			return err
		}

		boards = append(boards, &board)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBoardResponse{Board: boards, Pagination: pageRes}, nil
}
