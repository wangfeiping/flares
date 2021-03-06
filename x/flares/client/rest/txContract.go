package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/wangfeiping/flares/x/flares/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createContractRequest struct {
	BaseReq        rest.BaseReq `json:"base_req"`
	Module         string       `json:"module"`
	Creator        string       `json:"creator"`
	Key            string       `json:"key"`
	Receiver       string       `json:"receiver"`
	Accept         string       `json:"accept"`
	DurationHeight int32        `json:"durationHeight"`
	Bottom         string       `json:"bottom"`
}

func createContractHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createContractRequest
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

		module := req.Module
		parsedKey := req.Key
		parsedAccept := req.Accept
		parsedDurationHeight := req.DurationHeight
		parsedBottom := req.Bottom

		msg := types.NewMsgContract(
			creator,
			module,
			parsedKey,
			parsedAccept,
			parsedDurationHeight,
			parsedBottom,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
