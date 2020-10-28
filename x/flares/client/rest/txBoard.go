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

type createBoardRequest struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Creator       string       `json:"creator"`
	Base          string       `json:"base"`
	BaseDenom     string       `json:"baseDenom"`
	BaseAddress   string       `json:"baseAddress"`
	Accept        string       `json:"accept"`
	AcceptDenom   string       `json:"acceptDenom"`
	AcceptAddress string       `json:"acceptAddress"`
	Source        string       `json:"source"`
}

func createBoardHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createBoardRequest
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

		parsedBaseDenom := req.BaseDenom
		parsedAcceptDenom := req.AcceptDenom
		parsedSource := req.Source

		msg := types.NewMsgBoard(
			creator,
			parsedBaseDenom,
			parsedAcceptDenom,
			parsedSource,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
