package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/wangfeiping/flares/x/flares/types"
)

func (k Keeper) CreateBoard(ctx sdk.Context, board *types.MsgBoard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	key := fmt.Sprintf("%s-%s/%s-%s", types.BoardKey,
		board.AcceptDenom, board.BaseDenom, board.Creator.String())

	board.Address = AccAddressString(types.ModuleName, key).String()

	b := k.cdc.MustMarshalBinaryBare(board)
	store.Set(types.KeyPrefix(key), b)
}

func (k Keeper) GetBoard(ctx sdk.Context,
	acceptDenom, baseDenom string) (types.MsgBoard, sdk.Dec, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	key := fmt.Sprintf("%s-%s/%s", types.BoardKey,
		acceptDenom, baseDenom)

	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(key))

	defer iterator.Close()

	var board types.MsgBoard
	dec := sdk.NewDecWithPrec(0, 18)
	// dec := sdk.NewDec(0)
	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgBoard
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		addr, err := sdk.AccAddressFromBech32(msg.Address)
		if err != nil {
			return board, dec, err
		}
		coinBase := k.bank.GetBalance(ctx, addr, msg.BaseDenom)
		d := sdk.NewDecFromBigInt(coinBase.Amount.BigInt())
		coinAccept := k.bank.GetBalance(ctx, addr, msg.AcceptDenom)
		d = d.Quo(sdk.NewDecFromBigInt(coinAccept.Amount.BigInt()))

		if d.GT(dec) {
			dec = d
			board = msg
		}
	}

	return board, dec, nil
}

func (k Keeper) CheckBoard(ctx sdk.Context,
	contract *types.MsgContract, record *types.MsgContractTransferRecord) (sdk.Coin, error) {

	// TODO board logic: value discover
	coin, err := sdk.ParseCoin(record.Amount)
	if err != nil {
		return coin, err
	}
	if k.IsPayment(contract) {
		record.Voucher = 0
		return sdk.NewInt64Coin(types.VoycherKey, 0), nil
	}
	bottom, err := sdk.ParseCoin(contract.Bottom)
	if err != nil {
		return bottom, err
	}

	_, dec, err := k.GetBoard(ctx, coin.Denom, bottom.Denom)
	if err != nil {
		return coin, err
	}
	// fmt.Println("amount: ", coin.Amount)
	voucher := sdk.NewCoin(types.VoycherKey,
		sdk.NewIntFromBigInt(dec.Mul(sdk.NewDecFromIntWithPrec(coin.Amount, 12)).BigInt()))
	// sdk.NewIntFromBigInt(dec.Mul(sdk.NewDecFromInt(coin.Amount)).BigInt()))
	fmt.Println("\ndec: ", dec, "; record: ", record.Amount, "; vocher: ", voucher.Amount)
	return voucher, nil
}

func (k Keeper) GetAllBoard(ctx sdk.Context) (msgs []types.MsgBoard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.BoardKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgBoard
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
