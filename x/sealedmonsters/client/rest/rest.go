package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers sealedmonsters-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("custom/sealedmonsters/"+types.QueryListMonster, listMonsterHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/sealedmonsters/monster", createMonsterHandler(clientCtx)).Methods("POST")

}
