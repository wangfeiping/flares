package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/wangfeiping/flares/x/flares/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AllContractTransferRecord(c context.Context,
	req *types.QueryAllContractTransferRecordRequest) (*types.QueryAllContractTransferRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var contractTransferRecords []*types.MsgContractTransferRecord
	ctx := sdk.UnwrapSDKContext(c)

	contractTransferRecordStore := k.getContractTransferStore(ctx, "")

	pageRes, err := query.Paginate(contractTransferRecordStore,
		req.Pagination, func(key []byte, value []byte) error {
			var contractTransferRecord types.MsgContractTransferRecord
			if err := k.cdc.UnmarshalBinaryBare(value, &contractTransferRecord); err != nil {
				return err
			}

			contractTransferRecords = append(contractTransferRecords, &contractTransferRecord)
			return nil
		})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContractTransferRecordResponse{
		ContractTransferRecord: contractTransferRecords,
		Pagination:             pageRes}, nil
}
