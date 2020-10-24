package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

var _ bank.Keeper = (*BankKeeperWrapper)(nil)

type BankKeeperWrapper struct {
	bank.Keeper
	flaresKeeper keeper.Keeper
}

func NewBankKeeperWrapper(bankKeeper bank.Keeper, flaresK keeper.Keeper) bank.Keeper {
	return BankKeeperWrapper{
		Keeper:       bankKeeper,
		flaresKeeper: flaresK}
}

func (k BankKeeperWrapper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "flares/x/bank")
}

// github.com/cosmos/cosmos-sdk/x/bank/keeper.SendKeeper interface

func (k BankKeeperWrapper) InputOutputCoins(ctx sdk.Context,
	inputs []banktypes.Input, outputs []banktypes.Output) error {
	if err := k.Keeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
		return err
	}
	k.Logger(ctx).
		Info("transfer: InputOutputCoins", "height", ctx.BlockHeight())
	return nil
}
func (k BankKeeperWrapper) SendCoins(ctx sdk.Context,
	fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error {
	if err := k.Keeper.SendCoins(ctx, fromAddr, toAddr, amt); err != nil {
		return err
	}
	// Check if the address belongs to a contract.
	if ck := k.flaresKeeper.CheckContractReceiver(ctx, toAddr.String()); ck != nil {
		// TODO check contract bottom

		// stores the record of transfer
		rec := types.MsgContractTransferRecord{
			From:   fromAddr.String(),
			To:     toAddr.String(),
			Amount: amt.String()}
		k.flaresKeeper.CreateContractTransferRecord(ctx, rec)
		k.Logger(ctx).
			Info("SendCoins to a flares contract",
				"height", ctx.BlockHeight(), "receiver", toAddr.String())
		// check to see if the lowest price is met.
		c, err := k.flaresKeeper.GetContract(ctx, string(ck))
		if err != nil {
			return err
		}
		if !c.IsAuctions() {
			// it is traded
			// contract clearing
			k.flaresKeeper.ClearingContract(ctx, "flares/x/bank", &c)
		}
	}
	return nil
}
