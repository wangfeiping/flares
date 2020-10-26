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

type createRevealRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Creator      string       `json:"creator"`
	SolutionHash string       `json:"solutionHash"`
	Solution     string       `json:"solution"`
	Scavenger    string       `json:"scavenger"`
}

func createRevealHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRevealRequest
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

		parsedSolution := req.Solution

		msg := types.NewMsgReveal(
			creator,
			parsedSolution,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
