package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	flareskeeper "github.com/wangfeiping/flares/x/flares/keeper"
	flarestypes "github.com/wangfeiping/flares/x/flares/types"

	"github.com/wangfeiping/flares/x/nameservice/types"
)

type (
	Keeper struct {
		cdc          codec.Marshaler
		storeKey     sdk.StoreKey
		memKey       sdk.StoreKey
		flareskeeper flareskeeper.Keeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey,
	k flareskeeper.Keeper) *Keeper {
	nskeeper := &Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		flareskeeper: k,
	}
	k.RegisterContractClearing(types.ModuleName, nskeeper.ContractClearing)
	return nskeeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) ContractClearing(ctx sdk.Context,
	contract flarestypes.MsgContract) bool {
	recs := k.flareskeeper.GetContractTransfers(ctx, contract.Receiver)
	if len(recs) <= 0 {
		return false
	}
	// check board & compare
	max := &recs[0]
	maxVoucher, err := k.flareskeeper.CheckBoard(ctx, &contract, max)
	if err != nil {
		k.Logger(ctx).Error("check board failed: ", err.Error())
		return false
	}
	for _, record := range recs[1:] {
		recVoucher, err := k.flareskeeper.CheckBoard(ctx, &contract, &record)
		if err != nil {
			k.Logger(ctx).Error("check board failed: ", err.Error())
			return false
		}
		if recVoucher.Amount.GT(maxVoucher.Amount) {
			if err := k.flareskeeper.ReturnBack(ctx, &contract, max); err != nil {
				// fmt.Println("ReturnBack error: ", err.Error())
				k.Logger(ctx).Error("contract return failed: ", err.Error())
				return false
			}
			max = &record
			maxVoucher = recVoucher
		} else {
			if err := k.flareskeeper.ReturnBack(ctx, &contract, &record); err != nil {
				// fmt.Println("ReturnBack error: ", err.Error())
				k.Logger(ctx).Error("contract return failed: ", err.Error())
				return false
			}
		}
	}

	// check base price
	// TODO maybe it needs to be exchanged

	// transfer/send
	if err := k.flareskeeper.Deal(ctx, &contract, max); err != nil {
		k.Logger(ctx).Error("contract deal failed: ", err.Error())
		return false
	}
	// set the owner of name
	whois := types.MsgWhois{
		Owner: max.From,
		Value: contract.Key}
	k.CreateWhois(ctx, whois)
	return true
}
