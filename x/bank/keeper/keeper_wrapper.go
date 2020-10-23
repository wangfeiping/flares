package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
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
	if ck := k.flaresKeeper.CheckContractReceiver(ctx, toAddr); ck != nil {
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

// github.com/cosmos/cosmos-sdk/x/bank/keeper.Keeper interface

func (k BankKeeperWrapper) SendCoinsFromModuleToAccount(ctx sdk.Context,
	senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	if err := k.Keeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt); err != nil {
		return err
	}
	if strings.EqualFold(senderModule, ibctypes.ModuleName) {
		// Check if the address belongs to a contract.
		if k.flaresKeeper.CheckContractReceiver(ctx, recipientAddr) != nil {
			// Because this function can not get the sending address,
			// so it's not allowed to IBC transfer to a contract receiver address.
			k.Logger(ctx).
				Error("IBC transfer: it's not allowed to IBC transfer to a contract receiver address",
					"height", ctx.BlockHeight(), "receiver", recipientAddr.String())
			err := fmt.Errorf("IBC transfer: %s: height=%s, receiver=%s",
				"it's not allowed to IBC transfer to a contract receiver address",
				ctx.BlockHeight(), recipientAddr.String())
			return err
		}
	}
	return nil
}
