package sealedmonsters_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	flaresbank "github.com/wangfeiping/flares/x/bank/keeper"

	sealedkeeper "github.com/wangfeiping/flares/x/sealedmonsters/keeper"
	sealedtypes "github.com/wangfeiping/flares/x/sealedmonsters/types"

	"github.com/wangfeiping/flares/app"
	"github.com/wangfeiping/flares/app/params"
	"github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

var encCfg params.EncodingConfig = app.MakeEncodingConfig() // Ideally, we would reuse the one created by NewRootCmd.

var flaresStoreKey *sdk.KVStoreKey = sdk.NewKVStoreKey(types.StoreKey)
var flaresMemStoreKey *sdk.KVStoreKey = sdk.NewKVStoreKey(types.MemStoreKey)
var authStoreKey *sdk.KVStoreKey = sdk.NewKVStoreKey(authtypes.StoreKey)
var bankStoreKey *sdk.KVStoreKey = sdk.NewKVStoreKey(banktypes.StoreKey)
var sealedStoreKey *sdk.KVStoreKey = sdk.NewKVStoreKey(sealedtypes.StoreKey)
var sealedMemStoreKey *sdk.KVStoreKey = sdk.NewKVStoreKey(sealedtypes.MemStoreKey)

// module account permissions
var maccPerms map[string][]string = map[string][]string{
	authtypes.FeeCollectorName: nil,
}

func TestKeeper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sealed-Monsters Game Demo Suite")
}

func MockSdkContext() sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)

	cms.MountStoreWithDB(flaresStoreKey, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(authStoreKey, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(bankStoreKey, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(sealedStoreKey, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(sealedMemStoreKey, sdk.StoreTypeIAVL, db)

	cms.LoadLatestVersion()
	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
}

func MockFlaresKeeper() *keeper.Keeper {
	return keeper.NewKeeper(encCfg.Marshaler, flaresStoreKey, flaresMemStoreKey)
}

func MockParamsKeeper() paramskeeper.Keeper {
	return paramskeeper.NewKeeper(encCfg.Marshaler, encCfg.Amino,
		sdk.NewKVStoreKey(paramstypes.StoreKey),
		sdk.NewTransientStoreKey(paramstypes.TStoreKey))
}

func MockAccountKeeper(pk paramskeeper.Keeper) authkeeper.AccountKeeper {
	subspace := pk.Subspace(authtypes.ModuleName)

	return authkeeper.NewAccountKeeper(
		encCfg.Marshaler, authStoreKey, subspace,
		authtypes.ProtoBaseAccount, maccPerms,
	)
}

func MockBankKeeper(ak authkeeper.AccountKeeper, pk paramskeeper.Keeper,
	flaresKeeper keeper.Keeper) bankkeeper.Keeper {
	// return bank.NewBaseKeeper(
	// 	ak,
	// 	pk.Subspace(bank.ModuleName),
	// 	// app.ModuleAccountAddrs(),
	// 	make(map[string]bool),
	// )
	return flaresbank.NewBankKeeperWrapper(bankkeeper.NewBaseKeeper(
		encCfg.Marshaler, bankStoreKey,
		ak, pk.Subspace(banktypes.ModuleName),
		// app.BlockedAddrs(),
		make(map[string]bool),
	), flaresKeeper)
}

func MockSealedMonstersKeeper(flaresKeeper *keeper.Keeper, bk bankkeeper.Keeper) *sealedkeeper.Keeper {
	return sealedkeeper.NewKeeper(
		encCfg.Marshaler, sealedStoreKey, sealedMemStoreKey,
		*flaresKeeper, bk)
}
