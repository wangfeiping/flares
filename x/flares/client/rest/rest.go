package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
    "github.com/wangfeiping/flares/x/flares/types"
)

const (
    MethodGet = "GET"
)

// RegisterRoutes registers flares-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 3
    r.HandleFunc("custom/flares/" + types.QueryListContractTransferRecord, listContractTransferRecordHandler(clientCtx)).Methods("GET")

    r.HandleFunc("custom/flares/" + types.QueryListBoard, listBoardHandler(clientCtx)).Methods("GET")

    r.HandleFunc("custom/flares/" + types.QueryListContract, listContractHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 4
    r.HandleFunc("/flares/contractTransferRecord", createContractTransferRecordHandler(clientCtx)).Methods("POST")

    r.HandleFunc("/flares/board", createBoardHandler(clientCtx)).Methods("POST")

    r.HandleFunc("/flares/contract", createContractHandler(clientCtx)).Methods("POST")

}

