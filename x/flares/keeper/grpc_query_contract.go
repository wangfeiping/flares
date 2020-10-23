package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/wangfeiping/flares/x/flares/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllContract(c context.Context, req *types.QueryAllContractRequest) (*types.QueryAllContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contracts []*types.MsgContract
	ctx := sdk.UnwrapSDKContext(c)

	contractStore := k.getContractStore(ctx)

	pageRes, err := query.Paginate(contractStore, req.Pagination, func(key []byte, value []byte) error {
		var contract types.MsgContract
		if err := k.cdc.UnmarshalBinaryBare(value, &contract); err != nil {
			return err
		}
		if (strings.EqualFold(req.State, "success") && strings.EqualFold(contract.Result, "success")) ||
			(strings.EqualFold(req.State, "fail") && contract.Code != 0) ||
			((req.State == "" || strings.EqualFold(req.State, "pending")) && contract.Result == "") {
			contracts = append(contracts, &contract)
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractResponse{Contract: contracts, Pagination: pageRes}, nil
}
