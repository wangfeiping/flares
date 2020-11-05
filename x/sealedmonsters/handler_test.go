package sealedmonsters_test

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/wangfeiping/flares/x/sealedmonsters"
	sealedtypes "github.com/wangfeiping/flares/x/sealedmonsters/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/wangfeiping/flares/x/flares/handler"
	flareskeeper "github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

var _ = Describe("x/sealedmonsters", func() {

	var (
		voucher string   = types.VoucherDenom
		denoms  []string = []string{"base", "aaa", "bbb", "ccc"}
		balance int64    = 9999999
		num     int      = 6

		addrs []sdk.AccAddress
	)
	ctx := MockSdkContext()
	paramsKeeper := MockParamsKeeper()
	accountKeeper := MockAccountKeeper(paramsKeeper)
	bankKeeper := MockBankKeeper(accountKeeper, paramsKeeper)
	keeper := MockFlaresKeeper(bankKeeper)
	bankKeeper = MockFlaresBankKeeper(bankKeeper, *keeper)
	sealedKeeper := MockSealedMonstersKeeper(keeper, bankKeeper)
	handle := sealedmonsters.NewHandler(*sealedKeeper)
	flareshandle := handler.NewHandler(*keeper)

	supply := sdk.NewCoins()
	bankKeeper.SetSupply(ctx, banktypes.NewSupply(supply))

	BeforeEach(func() {

	})

	Describe("Mock cosmos-sdk.Context & Keepers", func() {
		Context("create sdk.Context", func() {
			It("should be success", func() {
				Expect(ctx).NotTo(BeNil())
			})
		})

		Context("create flares.Keeper", func() {
			It("should be success", func() {
				Expect(keeper).NotTo(BeNil())
			})
		})

		Context("create flares.Handler", func() {
			It("should be success", func() {
				Expect(handle).NotTo(BeNil())
			})
		})

		Context("create AccountKeeper", func() {
			It("should be success", func() {
				Expect(accountKeeper).NotTo(BeNil())
			})
		})

		Context("create flares.BankKeeper", func() {
			It("should be success", func() {
				Expect(bankKeeper).NotTo(BeNil())
			})
		})

		Context("create sealedmonsters Keeper", func() {
			It("should be success", func() {
				Expect(sealedKeeper).NotTo(BeNil())
			})
		})

		Context("create accounts", func() {
			It("should be success", func() {
				for i := 0; i < num; i++ {
					addrs = append(addrs, flareskeeper.AccAddressString(
						types.ModuleName, fmt.Sprintf("addr.test.%d", i)))
					acc := accountKeeper.NewAccountWithAddress(ctx, addrs[i])
					accountKeeper.SetAccount(ctx, acc)
					var coins sdk.Coins
					for _, d := range denoms {
						coins = append(coins, sdk.NewInt64Coin(d, balance))
					}
					bankKeeper.SetBalances(ctx, addrs[i], coins)
					// sdk.NewCoins(sdk.NewInt64Coin(denom, balance)))
					// sdk.NewInt64Coin(denom, balance))
				}
				Expect(len(addrs)).Should(Equal(num))

				for _, addr := range addrs {
					coin := bankKeeper.GetBalance(ctx, addr, denoms[1])
					// fmt.Println("account: ", addr.String(), "; ",
					// 	coin.Amount, "; ", coin.Denom)
					Expect(coin.Amount.Int64()).Should(Equal(balance))
				}
			})
		})

		Context("create boards", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 98, Time: time.Unix(10, 0)})
				msgBoard := types.NewMsgBoard(addrs[0], "base", "aaa", "local")
				_, err := flareshandle(ctx, msgBoard)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err := sdk.ParseCoins("1000base")
				Expect(err).ShouldNot(HaveOccurred())
				boardAcc, err := sdk.AccAddressFromBech32(msgBoard.Address)
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[0], boardAcc, coins)
				coins, err = sdk.ParseCoins("100000aaa")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[0], boardAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				msgBoard = types.NewMsgBoard(addrs[0], "base", "bbb", "local")
				_, err = flareshandle(ctx, msgBoard)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err = sdk.ParseCoins("1000base")
				Expect(err).ShouldNot(HaveOccurred())
				boardAcc, err = sdk.AccAddressFromBech32(msgBoard.Address)
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[0], boardAcc, coins)
				coins, err = sdk.ParseCoins("50000bbb")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[0], boardAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				msgBoard = types.NewMsgBoard(addrs[0], "base", "ccc", "local")
				_, err = flareshandle(ctx, msgBoard)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err = sdk.ParseCoins("1000base")
				Expect(err).ShouldNot(HaveOccurred())
				boardAcc, err = sdk.AccAddressFromBech32(msgBoard.Address)
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[0], boardAcc, coins)
				coins, err = sdk.ParseCoins("10000ccc")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[0], boardAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				msgBoard = types.NewMsgBoard(addrs[0], "base", "base", "local")
				_, err = flareshandle(ctx, msgBoard)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err = sdk.ParseCoins("1base")
				Expect(err).ShouldNot(HaveOccurred())
				boardAcc, err = sdk.AccAddressFromBech32(msgBoard.Address)
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[0], boardAcc, coins)
				coins, err = sdk.ParseCoins("1base")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[0], boardAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				boards := keeper.GetAllBoard(ctx)
				Expect(4).Should(Equal(len(boards)))
			})
		})
	})

	Describe("Test sealed monsters game", func() {
		Context("create a monster", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 99, Time: time.Unix(10, 0)})
				msg := sealedtypes.NewMsgMonster(addrs[0],
					"a monster", "kerberos", "1base")
				_, err := handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("create second monster", func() {
			It("should be failed", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 100, Time: time.Unix(10, 0)})
				msg := sealedtypes.NewMsgMonster(addrs[0],
					"a monster", "diablo", "1base")
				_, err := handle(ctx, msg)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("checks voucher", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 101, Time: time.Unix(10, 0)})
				coins := bankKeeper.GetSupply(ctx).GetTotal()
				Expect(0).Should(Equal(len(coins)))
			})
		})

		Context("seal the monster", func() {
			It("should be failed", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 111, Time: time.Unix(10, 0)})
				monsters := sealedKeeper.GetAllMonster(ctx)
				Expect(len(monsters)).Should(Equal(1))
				msg := sealedtypes.NewMsgSeal(addrs[1], "cubolous", "1000vvv")
				_, err := handle(ctx, msg)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("seal the monster", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 115, Time: time.Unix(10, 0)})
				msg := sealedtypes.NewMsgSeal(addrs[1], "cubolous", "1000aaa")
				_, err := handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
				msg = sealedtypes.NewMsgSeal(addrs[2], "kerberos", "9bbb")
				_, err = handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
				msg = sealedtypes.NewMsgSeal(addrs[3], "kerberos", "999ccc")
				_, err = handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
				msg = sealedtypes.NewMsgSeal(addrs[4], "kebero", "100aaa")
				_, err = handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
				msg = sealedtypes.NewMsgSeal(addrs[5], "kerbos", "100ccc")
				_, err = handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("checks total supply of the voucher", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 101, Time: time.Unix(10, 0)})
				coins := bankKeeper.GetSupply(ctx).GetTotal()
				Expect(1).Should(Equal(len(coins)))
				Expect(voucher).Should(Equal(coins[0].Denom))
				Expect(int64(121080000)).Should(Equal(coins[0].Amount.Int64()))
			})
		})

		Context("check all account balances before the game over", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 116, Time: time.Unix(10, 0)})
				contracts := keeper.GetAllContract(ctx)
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coin := bankKeeper.GetBalance(ctx, moduleAcc, denoms[1])
				Expect(int64(1100)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, moduleAcc, denoms[2])
				Expect(int64(9)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, moduleAcc, denoms[3])
				Expect(int64(1099)).Should(Equal(coin.Amount.Int64()))

				coin = bankKeeper.GetBalance(ctx, addrs[1], denoms[1])
				Expect(int64(9998999)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[1], voucher)
				Expect(int64(10000000)).Should(Equal(coin.Amount.Int64()))

				coin = bankKeeper.GetBalance(ctx, addrs[2], denoms[2])
				Expect(int64(9999990)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[2], voucher)
				Expect(int64(180000)).Should(Equal(coin.Amount.Int64()))

				coin = bankKeeper.GetBalance(ctx, addrs[3], denoms[3])
				Expect(int64(9999000)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[3], voucher)
				Expect(int64(99900000)).Should(Equal(coin.Amount.Int64()))

				coin = bankKeeper.GetBalance(ctx, addrs[4], denoms[1])
				Expect(int64(9999899)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[4], voucher)
				Expect(int64(1000000)).Should(Equal(coin.Amount.Int64()))

				coin = bankKeeper.GetBalance(ctx, addrs[5], denoms[3])
				Expect(int64(9999899)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[5], voucher)
				Expect(int64(10000000)).Should(Equal(coin.Amount.Int64()))
			})
		})

		Context("game over", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 199, Time: time.Unix(10, 0)})
				sealedmonsters.BeginBlockHandle(ctx, abci.RequestBeginBlock{}, *sealedKeeper)

				monsters := sealedKeeper.GetAllMonster(ctx)
				Expect(0).Should(Equal(len(monsters)))
				reveals := sealedKeeper.GetAllReveal(ctx)
				Expect(1).Should(Equal(len(reveals)))

				grpcReq := &types.QueryAllContractRequest{
					State: "success"}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(1).Should(Equal(len(resp.Contract)))
			})
		})

		Context("check all account balances after the game over", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 200, Time: time.Unix(10, 0)})

				coin := bankKeeper.GetBalance(ctx, addrs[2], denoms[1])
				Expect(int64(10000001)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[2], denoms[2])
				Expect(int64(9999999)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[2], voucher)
				Expect(int64(0)).Should(Equal(coin.Amount.Int64()))

				coin = bankKeeper.GetBalance(ctx, addrs[3], denoms[0])
				Expect(int64(10000001)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[3], denoms[1])
				Expect(int64(10001097)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[3], denoms[2])
				Expect(int64(9999999)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[3], denoms[3])
				Expect(int64(10000099)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[3], voucher)
				Expect(int64(0)).Should(Equal(coin.Amount.Int64()))

				total := bankKeeper.GetSupply(ctx).GetTotal()
				supply := total.AmountOf(types.VoucherDenom)
				Expect(int64(0)).Should(Equal(supply.Int64()))

			})
		})

		Context("create third monster", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 300, Time: time.Unix(10, 0)})
				msg := sealedtypes.NewMsgMonster(addrs[0],
					"a monster", "diablo", "1base")
				_, err := handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("game over", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 400, Time: time.Unix(10, 0)})
				sealedmonsters.BeginBlockHandle(ctx, abci.RequestBeginBlock{}, *sealedKeeper)

				monsters := sealedKeeper.GetAllMonster(ctx)
				Expect(0).Should(Equal(len(monsters)))
				reveals := sealedKeeper.GetAllReveal(ctx)
				Expect(2).Should(Equal(len(reveals)))

				grpcReq := &types.QueryAllContractRequest{
					State: "success"}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(1).Should(Equal(len(resp.Contract)))
			})
		})
	})
})
