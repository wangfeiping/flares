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

type createMonsterRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Creator      string       `json:"creator"`
	Description  string       `json:"description"`
	SolutionHash string       `json:"solutionHash"`
	Reward       string       `json:"reward"`
	Solution     string       `json:"solution"`
	Scavenger    string       `json:"scavenger"`
}

func createMonsterHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createMonsterRequest
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

		parsedDescription := req.Description

		parsedSolutionHash := req.SolutionHash

		parsedReward := req.Reward

		parsedSolution := req.Solution

		parsedScavenger := req.Scavenger

		msg := types.NewMsgMonster(
			creator,
			parsedDescription,
			parsedSolutionHash,
			parsedReward,
			parsedSolution,
			parsedScavenger,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
