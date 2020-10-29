package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	flareskeeper "github.com/wangfeiping/flares/x/flares/keeper"
	flarestypes "github.com/wangfeiping/flares/x/flares/types"

	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

type (
	Keeper struct {
		cdc          codec.Marshaler
		storeKey     sdk.StoreKey
		memKey       sdk.StoreKey
		FlaresKeeper flareskeeper.Keeper
		bankKeeper   bankkeeper.Keeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey,
	k flareskeeper.Keeper, bankKeeper bankkeeper.Keeper) *Keeper {
	sealedKeeper := &Keeper{
		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		bankKeeper:   bankKeeper,
		FlaresKeeper: k,
	}
	k.RegisterContractClearing(types.ModuleName, sealedKeeper.ContractClearing)
	return sealedKeeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func NewContract(creator sdk.AccAddress) *flarestypes.MsgContract {
	return flarestypes.NewMsgContract(
		creator,
		types.ModuleName, types.ModuleName, "", -1, "")
}

func (k Keeper) ContractClearing(ctx sdk.Context,
	contract flarestypes.MsgContract) bool {
	monsters := k.GetAllMonster(ctx)
	if len(monsters) != 1 {
		return false
	}
	total := k.bankKeeper.GetSupply(ctx).GetTotal()
	supply := total.AmountOf(flarestypes.VoucherDenom)
	contractAcc, err := sdk.AccAddressFromBech32(contract.Receiver)
	if err != nil {
		return false
	}
	monster := monsters[0]
	seals := k.GetAllSeal(ctx)
	var winners []types.MsgSeal
	for _, seal := range seals {
		if seal.SolutionHash == monster.SolutionHash {
			winners = append(winners, seal)
			// return back tokens of the winner's
			amt, err := sdk.ParseCoins(seal.Amount)
			if err != nil {
				k.Logger(ctx).Error("ParseCoins error: ", err.Error())
				return false
			}
			k.FlaresKeeper.SendCoins(ctx, contractAcc, seal.Creator, amt)
		} else {
			// Confiscate the vouchers
			amt := k.bankKeeper.GetBalance(ctx,
				seal.Creator, flarestypes.VoucherDenom)
			err := k.bankKeeper.SendCoinsFromAccountToModule(ctx,
				seal.Creator, flarestypes.ModuleName, sdk.NewCoins(amt))
			if err != nil {
				k.Logger(ctx).Error("SendCoinsFromAccountToModule error: ", err.Error())
				return false
			}
			k.RemoveSeal(ctx, &contract, &seal)
		}
	}
	lose := k.bankKeeper.GetBalance(ctx,
		authtypes.NewModuleAddress(flarestypes.ModuleName),
		flarestypes.VoucherDenom)
	win := supply.Sub(lose.Amount)
	all := k.bankKeeper.GetAllBalances(ctx, contractAcc)
	w := len(winners)
	for i, seal := range winners {
		hold := k.bankKeeper.GetBalance(ctx,
			seal.Creator, flarestypes.VoucherDenom)
		dec := sdk.NewDecFromBigInt(hold.Amount.BigInt())
		dec = dec.Quo(sdk.NewDecFromBigInt(win.BigInt()))
		for _, coin := range all {
			if i < w-1 {
				amt := dec.MulInt(coin.Amount).RoundInt()
				err := k.bankKeeper.SendCoins(ctx,
					contractAcc, seal.Creator,
					sdk.NewCoins(sdk.NewCoin(coin.Denom, amt)))
				if err != nil {
					k.Logger(ctx).Error("SendCoins error: ", err.Error())
					return false
				}
				// fmt.Println("win amt: ", seal.Creator, "; ", amt, "; ", coin)
			} else {
				amt := k.bankKeeper.GetBalance(ctx,
					contractAcc, coin.Denom)
				err := k.bankKeeper.SendCoins(ctx,
					contractAcc, seal.Creator, sdk.NewCoins(amt))
				if err != nil {
					k.Logger(ctx).Error("SendCoins error: ", err.Error())
					return false
				}
				// fmt.Println("win amt: ", seal.Creator, "; ", amt, "; ", coin)
			}
		}
		amt := k.bankKeeper.GetBalance(ctx,
			seal.Creator, flarestypes.VoucherDenom)
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx,
			seal.Creator, flarestypes.ModuleName, sdk.NewCoins(amt))
		if err != nil {
			k.Logger(ctx).Error("SendCoinsFromAccountToModule error: ", err.Error())
			return false
		}
		k.WinnerSeal(ctx, &contract, &seal)
	}
	k.SealedMonster(ctx, monster)
	reveal := types.NewMsgReveal(monster.Creator, monster.SolutionHash)
	k.CreateReveal(ctx, *reveal)
	return true
}
