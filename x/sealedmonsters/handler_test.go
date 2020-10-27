package sealedmonsters_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/wangfeiping/flares/x/sealedmonsters"
	sealedtypes "github.com/wangfeiping/flares/x/sealedmonsters/types"
	// "github.com/wangfeiping/flares/x/flares/handler"
	// "github.com/wangfeiping/flares/x/flares/types"
)

var _ = Describe("x/sealedmonsters", func() {

	var (
		denom   string = "voucher"
		balance int64  = 9999999
		num     int    = 5

		addrs []sdk.AccAddress
	)
	ctx := MockSdkContext()
	keeper := MockFlaresKeeper()
	paramsKeeper := MockParamsKeeper()
	accountKeeper := MockAccountKeeper(paramsKeeper)
	bankKeeper := MockBankKeeper(accountKeeper, paramsKeeper, *keeper)
	sealedKeeper := MockSealedMonstersKeeper(keeper, bankKeeper)
	handle := sealedmonsters.NewHandler(*sealedKeeper)

	BeforeEach(func() {

	})

	Describe("Mock cosmos-sdk.Context & Keepers", func() {
		Context("create sdk.Context", func() {
			It("should be success", func() {
				Expect(ctx).NotTo(BeZero())
			})
		})

		Context("create flares.Keeper", func() {
			It("should be success", func() {
				Expect(keeper).NotTo(BeZero())
			})
		})

		Context("create flares.Handler", func() {
			It("should be success", func() {
				Expect(handle).NotTo(BeZero())
			})
		})

		Context("create AccountKeeper", func() {
			It("should be success", func() {
				Expect(accountKeeper).NotTo(BeZero())
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
					addrs = append(addrs, sdk.AccAddress([]byte(
						fmt.Sprintf("addr.test.%d", i))))
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

	Describe("Test sealed monsters game", func() {
		Context("create a monster", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 99, Time: time.Unix(10, 0)})
				msg := sealedtypes.NewMsgMonster(addrs[0],
					"a monster", "kerberos", "99voucher")
				_, err := handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("seal the monster", func() {
			It("should be failed", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 111, Time: time.Unix(10, 0)})
				monsters := sealedKeeper.GetAllMonster(ctx)
				Expect(len(monsters)).Should(Equal(1))
				msg := sealedtypes.NewMsgSeal(addrs[1], "cubolous", "1000vou")
				_, err := handle(ctx, msg)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("seal the monster", func() {
			It("should be success", func() {
				ctx = ctx.WithBlockHeader(
					tmproto.Header{Height: 115, Time: time.Unix(10, 0)})
				msg := sealedtypes.NewMsgSeal(addrs[1], "cubolous", "1000voucher")
				_, err := handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
				msg = sealedtypes.NewMsgSeal(addrs[2], "kerberos", "9voucher")
				_, err = handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
				msg = sealedtypes.NewMsgSeal(addrs[3], "kerberos", "999voucher")
				_, err = handle(ctx, msg)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})
})
