package nameservice_test

import (
	"context"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/wangfeiping/flares/x/nameservice"
	nstypes "github.com/wangfeiping/flares/x/nameservice/types"

	"github.com/wangfeiping/flares/x/flares/handler"
	flareskeeper "github.com/wangfeiping/flares/x/flares/keeper"
	"github.com/wangfeiping/flares/x/flares/types"
)

var _ = Describe("x/nameservice", func() {

	var (
		denom   string = "voucher"
		balance int64  = 9999999
		num     int    = 5

		addrs []sdk.AccAddress
	)
	ctx := MockSdkContext()
	paramsKeeper := MockParamsKeeper()
	accountKeeper := MockAccountKeeper(paramsKeeper)
	bankKeeper := MockBankKeeper(accountKeeper, paramsKeeper)
	keeper := MockFlaresKeeper(bankKeeper)
	bankKeeper = MockFlaresBankKeeper(bankKeeper, *keeper)
	nsKeeper := MockNameServiceKeeper(keeper)
	handle := handler.NewHandler(*keeper)
	nsHandle := nameservice.NewHandler(*nsKeeper)

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
					bankKeeper.SetBalances(ctx, addrs[i],
						sdk.NewCoins(sdk.NewInt64Coin(denom, balance)))
					// sdk.NewInt64Coin(denom, balance))
				}
				Expect(len(addrs)).Should(Equal(num))

				for _, addr := range addrs {
					coin := bankKeeper.GetBalance(ctx, addr, denom)
					// fmt.Println("account: ", addr.String(), "; ",
					// 	coin.Amount, "; ", coin.Denom)
					Expect(coin.Amount.Int64()).Should(Equal(balance))
				}
			})
		})
	})

	Describe("Test nameservice", func() {
		Context("create a name", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 99, Time: time.Unix(10, 0)})
				msg := nstypes.NewMsgName(addrs[0],
					"xxx.cosmos", "99aaa")
				_, err := nsHandle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("create a contract", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 110, Time: time.Unix(10, 0)})
				msg := types.NewMsgContract(addrs[0],
					"nameservice", "xxx.cosmos", "ccc", 100, "99voucher")
				_, err := handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("send tokens", func() {
			It("should be failed", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 111, Time: time.Unix(10, 0)})
				contracts := keeper.GetAllContract(ctx)
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				addr, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coins, err := sdk.ParseCoins("100aaa")
				Expect(err).ShouldNot(HaveOccurred())
				err = bankKeeper.SendCoins(ctx,
					addrs[1], addr, coins)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("send tokens", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 112, Time: time.Unix(10, 0)})

				contracts := keeper.GetAllContract(ctx)
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())

				coins, err := sdk.ParseCoins("100voucher")
				Expect(err).ShouldNot(HaveOccurred())
				cctx := ctx.WithTxBytes([]byte(addrs[1]))
				err = bankKeeper.SendCoins(cctx,
					addrs[1], moduleAcc, coins)
				Expect(err).ShouldNot(HaveOccurred())

				coins, err = sdk.ParseCoins("120voucher")
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
					tmproto.Header{Height: 113, Time: time.Unix(10, 0)})
				contracts := keeper.GetAllContract(ctx)
				Expect(1).Should(Equal(len(contracts)))
				c := contracts[0]
				moduleAcc, err := sdk.AccAddressFromBech32(c.Receiver)
				Expect(err).ShouldNot(HaveOccurred())
				coin := bankKeeper.GetBalance(ctx, moduleAcc, denom)
				Expect(int64(220)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[0], denom)
				Expect(int64(9999999)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[1], denom)
				Expect(int64(9999899)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[2], denom)
				Expect(int64(9999879)).Should(Equal(coin.Amount.Int64()))
			})
		})

		Context("clearing the contract", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 210, Time: time.Unix(10, 0)})
				handler.BeginBlockHandle(ctx, abci.RequestBeginBlock{}, *keeper)
				grpcReq := &types.QueryAllContractRequest{
					State: "success"}
				cctx := context.WithValue(context.Background(), sdk.SdkContextKey, ctx)
				resp, err := keeper.AllContract(cctx, grpcReq)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(1).Should(Equal(len(resp.Contract)))

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
					tmproto.Header{Height: 211, Time: time.Unix(10, 0)})

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
				coin := bankKeeper.GetBalance(ctx, moduleAcc, denom)
				Expect(int64(0)).Should(Equal(coin.Amount.Int64()))

				coin = bankKeeper.GetBalance(ctx, addrs[0], denom)
				Expect(int64(10000119)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[1], denom)
				Expect(int64(9999999)).Should(Equal(coin.Amount.Int64()))
				coin = bankKeeper.GetBalance(ctx, addrs[2], denom)
				Expect(int64(9999879)).Should(Equal(coin.Amount.Int64()))
			})
		})
	})
})
