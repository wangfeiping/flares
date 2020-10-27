package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctransfer "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/keeper"
	ibctransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"

	"github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

// var _ ibctransfer.Keeper = (*IBCTransferKeeperWrapper)(nil)

var ModuleName = "flares/x/ibc"

type IBCTransferKeeperWrapper struct {
	ibctransfer.Keeper
	flaresKeeper keeper.Keeper
}

// func NewIBCTransferKeeperWrapper(ibctransferKeeper ibctransfer.Keeper, flaresK keeper.Keeper) ibctransfer.Keeper {
// 	return IBCTransferKeeperWrapper{
// 		Keeper:       ibctransferKeeper,
// 		flaresKeeper: flaresK}
// }

func (k IBCTransferKeeperWrapper) OnRecvPacket(ctx sdk.Context,
	packet channeltypes.Packet, data ibctransfertypes.FungibleTokenPacketData) error {
	if err := k.Keeper.OnRecvPacket(ctx, packet, data); err != nil {
		return err
	}
	// Check if the address belongs to a contract.
	if ck := k.flaresKeeper.CheckContractReceiver(ctx, data.Receiver); ck != nil {
		coin := sdk.NewInt64Coin(data.Denom, int64(data.Amount))
		// stores the record of transfer
		rec := types.MsgContractTransferRecord{
			From:   data.Sender,
			To:     data.Receiver,
			Amount: coin.String()}
		k.flaresKeeper.CreateContractTransferRecord(ctx, rec)
		k.Logger(ctx).
			Info("IBC transfer to a flares contract",
				"height", ctx.BlockHeight(), "receiver", data.Receiver)
		c, err := k.flaresKeeper.GetContract(ctx, string(ck))
		if err != nil {
			return err
		}
		if k.flaresKeeper.IsPayment(&c) {
			// contract clearing
			// check contract bottom
			// check if the base price is met
			k.flaresKeeper.ClearingContract(ctx, ModuleName, &c)
		}
	}
	return nil
}
