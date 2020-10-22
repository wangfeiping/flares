package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/exported"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"

	"github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

var _ bank.Keeper = (*BankKeeperWrapper)(nil)

type BankKeeperWrapper struct {
	k            bank.Keeper
	flaresKeeper keeper.Keeper
}

func NewBankKeeperWrapper(bankKeeper bank.Keeper, flaresK keeper.Keeper) bank.Keeper {
	return BankKeeperWrapper{k: bankKeeper,
		flaresKeeper: flaresK}
}

// github.com/cosmos/cosmos-sdk/x/bank/keeper.ViewKeeper interface

func (k BankKeeperWrapper) ValidateBalance(ctx sdk.Context, addr sdk.AccAddress) error {
	return k.k.ValidateBalance(ctx, addr)
}
func (k BankKeeperWrapper) HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool {
	return k.k.HasBalance(ctx, addr, amt)
}
func (k BankKeeperWrapper) GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k.k.GetAllBalances(ctx, addr)
}
func (k BankKeeperWrapper) GetAccountsBalances(ctx sdk.Context) []banktypes.Balance {
	return k.k.GetAccountsBalances(ctx)
}
func (k BankKeeperWrapper) GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return k.k.GetBalance(ctx, addr, denom)
}
func (k BankKeeperWrapper) LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k.k.LockedCoins(ctx, addr)
}
func (k BankKeeperWrapper) SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins {
	return k.k.SpendableCoins(ctx, addr)
}
func (k BankKeeperWrapper) IterateAccountBalances(ctx sdk.Context,
	addr sdk.AccAddress, cb func(coin sdk.Coin) (stop bool)) {
	k.k.IterateAccountBalances(ctx, addr, cb)
}
func (k BankKeeperWrapper) IterateAllBalances(ctx sdk.Context,
	cb func(address sdk.AccAddress, coin sdk.Coin) (stop bool)) {
	k.k.IterateAllBalances(ctx, cb)
}

// github.com/cosmos/cosmos-sdk/x/bank/keeper.SendKeeper interface

func (k BankKeeperWrapper) InputOutputCoins(ctx sdk.Context,
	inputs []banktypes.Input, outputs []banktypes.Output) error {
	if err := k.k.InputOutputCoins(ctx, inputs, outputs); err != nil {
		return err
	}
	ctx.Logger().With("module", "flares/x/bank").
		Info("transfer: InputOutputCoins", "height", ctx.BlockHeight())
	return nil
}
func (k BankKeeperWrapper) SendCoins(ctx sdk.Context,
	fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error {
	if err := k.k.SendCoins(ctx, fromAddr, toAddr, amt); err != nil {
		return err
	}
	// Check if the address belongs to a contract.
	if k.flaresKeeper.CheckContractReceiver(ctx, toAddr) {
		rec := types.MsgContractTransferRecord{
			From:   fromAddr.String(),
			To:     toAddr.String(),
			Amount: amt.String()}
		k.flaresKeeper.CreateContractTransferRecord(ctx, rec)
		ctx.Logger().With("module", "flares/x/bank").
			Info("SendCoins to a flares contract",
				"height", ctx.BlockHeight(), "receiver", toAddr.String())
	}
	return nil
}
func (k BankKeeperWrapper) SubtractCoins(ctx sdk.Context,
	addr sdk.AccAddress, amt sdk.Coins) error {
	return k.k.SubtractCoins(ctx, addr, amt)
}
func (k BankKeeperWrapper) AddCoins(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coins) error {
	return k.k.AddCoins(ctx, addr, amt)
}
func (k BankKeeperWrapper) SetBalance(ctx sdk.Context, addr sdk.AccAddress, balance sdk.Coin) error {
	return k.k.SetBalance(ctx, addr, balance)
}
func (k BankKeeperWrapper) SetBalances(ctx sdk.Context, addr sdk.AccAddress, balances sdk.Coins) error {
	return k.k.SetBalances(ctx, addr, balances)
}
func (k BankKeeperWrapper) GetParams(ctx sdk.Context) banktypes.Params {
	return k.k.GetParams(ctx)
}
func (k BankKeeperWrapper) SetParams(ctx sdk.Context, params banktypes.Params) {
	k.k.SetParams(ctx, params)
}
func (k BankKeeperWrapper) SendEnabledCoin(ctx sdk.Context, coin sdk.Coin) bool {
	return k.k.SendEnabledCoin(ctx, coin)
}
func (k BankKeeperWrapper) SendEnabledCoins(ctx sdk.Context, coins ...sdk.Coin) error {
	return k.k.SendEnabledCoins(ctx, coins...)
}
func (k BankKeeperWrapper) BlockedAddr(addr sdk.AccAddress) bool {
	return k.k.BlockedAddr(addr)
}

// github.com/cosmos/cosmos-sdk/x/bank/keeper.Keeper interface

func (k BankKeeperWrapper) InitGenesis(ctx sdk.Context, state banktypes.GenesisState) {
	k.k.InitGenesis(ctx, state)
}
func (k BankKeeperWrapper) ExportGenesis(ctx sdk.Context) *banktypes.GenesisState {
	return k.k.ExportGenesis(ctx)
}
func (k BankKeeperWrapper) GetSupply(ctx sdk.Context) exported.SupplyI {
	return k.k.GetSupply(ctx)
}
func (k BankKeeperWrapper) SetSupply(ctx sdk.Context, supply exported.SupplyI) {
	k.k.SetSupply(ctx, supply)
}
func (k BankKeeperWrapper) GetDenomMetaData(ctx sdk.Context, denom string) banktypes.Metadata {
	return k.k.GetDenomMetaData(ctx, denom)
}
func (k BankKeeperWrapper) SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata) {
	k.k.SetDenomMetaData(ctx, denomMetaData)
}
func (k BankKeeperWrapper) IterateAllDenomMetaData(ctx sdk.Context, cb func(banktypes.Metadata) bool) {
	k.k.IterateAllDenomMetaData(ctx, cb)
}
func (k BankKeeperWrapper) SendCoinsFromModuleToAccount(ctx sdk.Context,
	senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	if err := k.k.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt); err != nil {
		return err
	}
	if strings.EqualFold(senderModule, ibctypes.ModuleName) {
		// Check if the address belongs to a contract.
		if k.flaresKeeper.CheckContractReceiver(ctx, recipientAddr) {
			// Because this function can not get the sending address,
			// so it's not allowed to IBC transfer to a contract receiver address.
			ctx.Logger().With("module", "flares/x/bank").
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
func (k BankKeeperWrapper) SendCoinsFromModuleToModule(ctx sdk.Context,
	senderModule, recipientModule string, amt sdk.Coins) error {
	return k.k.SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt)
}
func (k BankKeeperWrapper) SendCoinsFromAccountToModule(ctx sdk.Context,
	senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return k.k.SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}
func (k BankKeeperWrapper) DelegateCoinsFromAccountToModule(ctx sdk.Context,
	senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return k.k.DelegateCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt)
}
func (k BankKeeperWrapper) UndelegateCoinsFromModuleToAccount(ctx sdk.Context,
	senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.k.UndelegateCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}
func (k BankKeeperWrapper) MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	return k.k.MintCoins(ctx, moduleName, amt)
}
func (k BankKeeperWrapper) BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error {
	return k.k.BurnCoins(ctx, moduleName, amt)
}
func (k BankKeeperWrapper) DelegateCoins(ctx sdk.Context,
	delegatorAddr, moduleAccAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.k.DelegateCoins(ctx, delegatorAddr, moduleAccAddr, amt)
}
func (k BankKeeperWrapper) UndelegateCoins(ctx sdk.Context,
	moduleAccAddr, delegatorAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.k.UndelegateCoins(ctx, moduleAccAddr, delegatorAddr, amt)
}
func (k BankKeeperWrapper) MarshalSupply(supplyI exported.SupplyI) ([]byte, error) {
	return k.k.MarshalSupply(supplyI)
}
func (k BankKeeperWrapper) UnmarshalSupply(bz []byte) (exported.SupplyI, error) {
	return k.k.UnmarshalSupply(bz)
}

// github.com/cosmos/cosmos-sdk/x/bank/keeper.QueryServer interface

func (k BankKeeperWrapper) Balance(ctx context.Context,
	req *banktypes.QueryBalanceRequest) (*banktypes.QueryBalanceResponse, error) {
	return k.k.Balance(ctx, req)
}
func (k BankKeeperWrapper) AllBalances(ctx context.Context,
	req *banktypes.QueryAllBalancesRequest) (*banktypes.QueryAllBalancesResponse, error) {
	return k.k.AllBalances(ctx, req)
}
func (k BankKeeperWrapper) TotalSupply(ctx context.Context,
	req *banktypes.QueryTotalSupplyRequest) (*banktypes.QueryTotalSupplyResponse, error) {
	return k.k.TotalSupply(ctx, req)
}
func (k BankKeeperWrapper) SupplyOf(ctx context.Context,
	req *banktypes.QuerySupplyOfRequest) (*banktypes.QuerySupplyOfResponse, error) {
	return k.k.SupplyOf(ctx, req)
}
func (k BankKeeperWrapper) Params(ctx context.Context,
	req *banktypes.QueryParamsRequest) (*banktypes.QueryParamsResponse, error) {
	return k.k.Params(ctx, req)
}
