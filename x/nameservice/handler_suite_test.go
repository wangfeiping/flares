package nameservice_test

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

	nskeeper "github.com/wangfeiping/flares/x/nameservice/keeper"
	nstypes "github.com/wangfeiping/flares/x/nameservice/types"

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
var nsStoreKey *sdk.KVStoreKey = sdk.NewKVStoreKey(nstypes.StoreKey)
var nsMemStoreKey *sdk.KVStoreKey = sdk.NewKVStoreKey(nstypes.MemStoreKey)

// module account permissions
var maccPerms map[string][]string = map[string][]string{
	authtypes.FeeCollectorName: nil,
}

func TestKeeper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Name-Service Payment & Auctions Demo Suite")
}

func MockSdkContext() sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)

	cms.MountStoreWithDB(flaresStoreKey, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(authStoreKey, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(bankStoreKey, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(nsStoreKey, sdk.StoreTypeIAVL, db)
	cms.MountStoreWithDB(nsMemStoreKey, sdk.StoreTypeIAVL, db)

	cms.LoadLatestVersion()
	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
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

func MockBankKeeper(ak authkeeper.AccountKeeper,
	pk paramskeeper.Keeper) bankkeeper.Keeper {
	return bankkeeper.NewBaseKeeper(
		encCfg.Marshaler, bankStoreKey,
		ak, pk.Subspace(banktypes.ModuleName),
		// app.BlockedAddrs(),
		make(map[string]bool),
	)
}

func MockFlaresKeeper(bank bankkeeper.Keeper) *keeper.Keeper {
	return keeper.NewKeeper(encCfg.Marshaler,
		flaresStoreKey, flaresMemStoreKey, bank)
}

func MockFlaresBankKeeper(bank bankkeeper.Keeper,
	flaresKeeper keeper.Keeper) bankkeeper.Keeper {
	// return bank.NewBaseKeeper(
	// 	ak,
	// 	pk.Subspace(bank.ModuleName),
	// 	// app.ModuleAccountAddrs(),
	// 	make(map[string]bool),
	// )
	return flaresbank.NewBankKeeperWrapper(bank, flaresKeeper)
}

func MockNameServiceKeeper(flaresKeeper *keeper.Keeper) *nskeeper.Keeper {
	return nskeeper.NewKeeper(
		encCfg.Marshaler, nsStoreKey, nsMemStoreKey,
		*flaresKeeper)
}
