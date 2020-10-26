package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createSealRequest struct {
	BaseReq               rest.BaseReq `json:"base_req"`
	Creator               string       `json:"creator"`
	SolutionHash          string       `json:"solutionHash"`
	SolutionScavengerHash string       `json:"solutionScavengerHash"`
	Scavenger             string       `json:"scavenger"`
	Amount                string       `json:"amount"`
}

func createSealHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSealRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgSeal(
			creator,
			req.SolutionHash,
			req.Amount,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
