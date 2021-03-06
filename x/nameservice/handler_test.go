package nameservice_test

import (
	"context"
	"fmt"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/wangfeiping/flares/x/nameservice"
	nskeeper "github.com/wangfeiping/flares/x/nameservice/keeper"
	nstypes "github.com/wangfeiping/flares/x/nameservice/types"

	"github.com/wangfeiping/flares/x/flares/handler"
	flareskeeper "github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

var _ = Describe("x/nameservice", func() {

	var (
		// voucher string   = "voucher"
		denoms  []string = []string{"base", "aaa", "bbb", "ccc"}
		balance int64    = 9999999
		num     int      = 6

		addrs         []sdk.AccAddress
		ctx           sdk.Context
		paramsKeeper  paramskeeper.Keeper
		accountKeeper authkeeper.AccountKeeper
		bankKeeper    bankkeeper.Keeper
		keeper        *flareskeeper.Keeper
		nsKeeper      *nskeeper.Keeper
		handle        sdk.Handler
		nsHandle      sdk.Handler
	)

	ctx = MockSdkContext()
	paramsKeeper = MockParamsKeeper()
	accountKeeper = MockAccountKeeper(paramsKeeper)
	bankKeeper = MockBankKeeper(accountKeeper, paramsKeeper)
	keeper = MockFlaresKeeper(bankKeeper)
	bankKeeper = MockFlaresBankKeeper(bankKeeper, *keeper)
	nsKeeper = MockNameServiceKeeper(keeper)
	handle = handler.NewHandler(*keeper)
	nsHandle = nameservice.NewHandler(*nsKeeper)

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

		Context("create nameservice Keeper", func() {
			It("should be success", func() {
				Expect(nsKeeper).NotTo(BeNil())
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
	})

	Describe("Test payment for nameservice", func() {
		value := "pay.cosmos"

		Context("create a name", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 110, Time: time.Unix(10, 0)})
				msg := nstypes.NewMsgName(addrs[0],
					value, "99aaa")
				_, err := nsHandle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("create a contract", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 111, Time: time.Unix(10, 0)})
				msg := types.NewMsgContract(addrs[0],
					"nameservice", value, "aaa", 0, "99aaa")
				_, err := handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("send tokens (the contract is automatically clearing)", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 112, Time: time.Unix(10, 0)})

				contracts := keeper.GetAllContract(ctx)
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())

				coins, err := sdk.ParseCoins("99aaa")
				Expect(err).ShouldNot(HaveOccurred())
				cctx := ctx.WithTxBytes([]byte(addrs[1]))
				err = bankKeeper.SendCoins(cctx,
					addrs[1], moduleAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("check all account balances after the contract clearing", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 113, Time: time.Unix(10, 0)})

				grpcReq := &types.QueryAllContractRequest{}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(0).Should(Equal(len(resp.Contract)))

				grpcReq = &types.QueryAllContractRequest{
					State: "success"}
				cctx = context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err = keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(1).Should(Equal(len(resp.Contract)))

				c := resp.Contract[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coin := bankKeeper.GetBalance(ctx, moduleAcc, denoms[1])
				Expect(int64(0)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[0], denoms[1])
				Expect(int64(10000098)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[1], denoms[1])
				Expect(int64(9999900)).Should(Equal(coin.Amount.Int64()))
			})
		})

		Context("check whois", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 114, Time: time.Unix(10, 0)})

				whois, err := nsKeeper.GetWhois(ctx, value)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(whois).ShouldNot(BeNil())
				Expect(whois.Owner).Should(Equal(addrs[1].String()))
			})
		})
	})

	Describe("Test auctions for nameservice", func() {
		value := "xxx.cosmos"
		Context("create a name", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 199, Time: time.Unix(10, 0)})
				msg := nstypes.NewMsgName(addrs[0],
					value, "99bbb")
				_, err := nsHandle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("create a contract", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 200, Time: time.Unix(10, 0)})
				msg := types.NewMsgContract(addrs[0],
					"nameservice", value, "bbb", 80, "99bbb")
				_, err := handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("send tokens", func() {
			It("should be failed", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 201, Time: time.Unix(10, 0)})
				grpcReq := &types.QueryAllContractRequest{}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				contracts := resp.Contract
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				addr, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err := sdk.ParseCoins("100xxx")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[1], addr, coins)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("send tokens", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 202, Time: time.Unix(10, 0)})

				contracts := keeper.GetAllContract(ctx)
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())

				coins, err := sdk.ParseCoins("100bbb")
				Expect(err).ShouldNot(HaveOccurred())
				cctx := ctx.WithTxBytes([]byte(addrs[1]))
				err = bankKeeper.SendCoins(cctx,
					addrs[1], moduleAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				coins, err = sdk.ParseCoins("120bbb")
				Expect(err).ShouldNot(HaveOccurred())
				cctx = ctx.WithTxBytes([]byte(addrs[2]))
				err = bankKeeper.SendCoins(cctx,
					addrs[2], moduleAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("check all account balances before the contract clearing", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 203, Time: time.Unix(10, 0)})
				contracts := keeper.GetAllContract(ctx)
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coin := bankKeeper.GetBalance(ctx, moduleAcc, denoms[2])
				Expect(int64(220)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[0], denoms[2])
				Expect(int64(9999999)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[1], denoms[2])
				Expect(int64(9999899)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[2], denoms[2])
				Expect(int64(9999879)).Should(Equal(coin.Amount.Int64()))
			})
		})

		Context("clearing the contract", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 280, Time: time.Unix(10, 0)})
				handler.BeginBlockHandle(ctx, abci.RequestBeginBlock{}, *keeper)
				grpcReq := &types.QueryAllContractRequest{
					State: "success"}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(2).Should(Equal(len(resp.Contract)))

				grpcReq = &types.QueryAllContractRequest{
					State: "fail"}
				cctx = context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err = keeper.AllContract(cctx, grpcReq)
				// fmt.Println("err: ", resp.Contract[0].Result)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(0).Should(Equal(len(resp.Contract)))

				grpcReq = &types.QueryAllContractRequest{
					State: ""}
				cctx = context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err = keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(0).Should(Equal(len(resp.Contract)))
			})
		})

		Context("check all account balances after the contract clearing", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 291, Time: time.Unix(10, 0)})

				grpcReq := &types.QueryAllContractRequest{}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				// for _, c := range resp.Contract {
				// 	fmt.Println("contract: ", c.Key)
				// }
				Expect(0).Should(Equal(len(resp.Contract)))

				grpcReq = &types.QueryAllContractRequest{
					State: "success"}
				cctx = context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err = keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(2).Should(Equal(len(resp.Contract)))

				c := resp.Contract[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coin := bankKeeper.GetBalance(ctx, moduleAcc, denoms[2])
				Expect(int64(0)).Should(Equal(coin.Amount.Int64()))

				coin = bankKeeper.GetBalance(ctx, addrs[0], denoms[2])
				Expect(int64(10000119)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[1], denoms[2])
				Expect(int64(9999999)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[2], denoms[2])
				Expect(int64(9999879)).Should(Equal(coin.Amount.Int64()))
			})
		})

		Context("check whois", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 292, Time: time.Unix(10, 0)})

				whois, err := nsKeeper.GetWhois(ctx, value)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(whois).ShouldNot(BeNil())
				Expect(whois.Owner).Should(Equal(addrs[2].String()))
			})
		})
	})

	Describe("Test multi-tokens auctions for nameservice", func() {
		value := "multi.cosmos"

		Context("create boards & a name", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 299, Time: time.Unix(10, 0)})
				msg := nstypes.NewMsgName(addrs[0],
					value, "9base")
				_, err := nsHandle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
				msgBoard := types.NewMsgBoard(addrs[1], "base", "aaa", "local")
				_, err = handle(ctx, msgBoard)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err := sdk.ParseCoins("1000base")
				Expect(err).ShouldNot(HaveOccurred())
				boardAcc, err := sdk.AccAddressFromBech32(msgBoard.Address)
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[1], boardAcc, coins)
				coins, err = sdk.ParseCoins("100000aaa")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[1], boardAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				msgBoard = types.NewMsgBoard(addrs[1], "base", "bbb", "local")
				_, err = handle(ctx, msgBoard)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err = sdk.ParseCoins("1000base")
				Expect(err).ShouldNot(HaveOccurred())
				boardAcc, err = sdk.AccAddressFromBech32(msgBoard.Address)
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[1], boardAcc, coins)
				coins, err = sdk.ParseCoins("50000bbb")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[1], boardAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				msgBoard = types.NewMsgBoard(addrs[1], "base", "ccc", "local")
				_, err = handle(ctx, msgBoard)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err = sdk.ParseCoins("1000base")
				Expect(err).ShouldNot(HaveOccurred())
				boardAcc, err = sdk.AccAddressFromBech32(msgBoard.Address)
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[1], boardAcc, coins)
				coins, err = sdk.ParseCoins("10000ccc")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[1], boardAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				boards := keeper.GetAllBoard(ctx)
				Expect(3).Should(Equal(len(boards)))

			})
		})

		Context("create a contract", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 300, Time: time.Unix(10, 0)})
				msg := types.NewMsgContract(addrs[0],
					"nameservice", value, "@all", 80, "9base")
				_, err := handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("send tokens", func() {
			It("should be failed", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 301, Time: time.Unix(10, 0)})
				grpcReq := &types.QueryAllContractRequest{}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				contracts := resp.Contract
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				addr, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err := sdk.ParseCoins("100xxx")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[1], addr, coins)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("send tokens", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 302, Time: time.Unix(10, 0)})

				contracts := keeper.GetAllContract(ctx)
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())

				coins, err := sdk.ParseCoins("100aaa")
				Expect(err).ShouldNot(HaveOccurred())
				cctx := ctx.WithTxBytes([]byte(addrs[3]))
				err = bankKeeper.SendCoins(cctx,
					addrs[3], moduleAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				coins, err = sdk.ParseCoins("120bbb")
				Expect(err).ShouldNot(HaveOccurred())
				cctx = ctx.WithTxBytes([]byte(addrs[4]))
				err = bankKeeper.SendCoins(cctx,
					addrs[4], moduleAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				coins, err = sdk.ParseCoins("101ccc")
				Expect(err).ShouldNot(HaveOccurred())
				cctx = ctx.WithTxBytes([]byte(addrs[5]))
				err = bankKeeper.SendCoins(cctx,
					addrs[5], moduleAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("check all account balances before the contract clearing", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 303, Time: time.Unix(10, 0)})
				contracts := keeper.GetAllContract(ctx)
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coin := bankKeeper.GetBalance(ctx, moduleAcc, denoms[1])
				Expect(int64(100)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, moduleAcc, denoms[2])
				Expect(int64(120)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, moduleAcc, denoms[3])
				Expect(int64(101)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[3], denoms[1])
				Expect(int64(9999899)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[4], denoms[2])
				Expect(int64(9999879)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[5], denoms[3])
				Expect(int64(9999898)).Should(Equal(coin.Amount.Int64()))
			})
		})

		Context("clearing the contract", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 380, Time: time.Unix(10, 0)})
				handler.BeginBlockHandle(ctx, abci.RequestBeginBlock{}, *keeper)
				grpcReq := &types.QueryAllContractRequest{
					State: "success"}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(3).Should(Equal(len(resp.Contract)))

				grpcReq = &types.QueryAllContractRequest{
					State: "fail"}
				cctx = context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err = keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(0).Should(Equal(len(resp.Contract)))

				grpcReq = &types.QueryAllContractRequest{
					State: ""}
				cctx = context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err = keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(0).Should(Equal(len(resp.Contract)))
			})
		})

		Context("check all account balances after the contract clearing", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 391, Time: time.Unix(10, 0)})

				grpcReq := &types.QueryAllContractRequest{}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(0).Should(Equal(len(resp.Contract)))

				grpcReq = &types.QueryAllContractRequest{
					State: "success"}
				cctx = context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err = keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(3).Should(Equal(len(resp.Contract)))

				c := getContract(resp.Contract, value)
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coins := bankKeeper.GetAllBalances(ctx, moduleAcc)
				for _, coin := range coins {
					Expect(int64(0)).Should(Equal(coin.Amount.Int64()))
				}
				coin := bankKeeper.GetBalance(ctx, addrs[3], denoms[1])
				Expect(int64(9999999)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[4], denoms[2])
				Expect(int64(9999999)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[5], denoms[3])
				Expect(int64(9999898)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[0], denoms[3])
				Expect(int64(10000100)).Should(Equal(coin.Amount.Int64()))
			})
		})
	})
})

func getContract(contracts []*types.MsgContract, key string) *types.MsgContract {
	for _, c := range contracts {
		if strings.EqualFold(c.Key, key) {
			return c
		}
	}
	return nil
}
