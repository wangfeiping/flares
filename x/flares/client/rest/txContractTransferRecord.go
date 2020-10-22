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

type createContractTransferRecordRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Hash string `json:"hash"`
	From string `json:"from"`
	To string `json:"to"`
	Amount string `json:"amount"`
	
}

func createContractTransferRecordHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createContractTransferRecordRequest
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

		
		parsedHash := req.Hash
		
		parsedFrom := req.From
		
		parsedTo := req.To
		
		parsedAmount := req.Amount
		

		msg := types.NewMsgContractTransferRecord(
			creator,
			parsedHash,
			parsedFrom,
			parsedTo,
			parsedAmount,
			
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
