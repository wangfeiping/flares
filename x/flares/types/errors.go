package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/flares module sentinel errors
var (
	ErrContractExists           = sdkerrors.Register(ModuleName, 1100, "contract already exists")
	ErrContractNotFound         = sdkerrors.Register(ModuleName, 1101, "contract not found")
	ErrContractClearingNotFound = sdkerrors.Register(ModuleName, 1102, "contract-clearing not found")
	ErrContractClearingFailed   = sdkerrors.Register(ModuleName, 1103, "contract-clearing was failed")
	ErrInvalidAmount            = sdkerrors.Register(ModuleName, 1104, "invalid amount")
)
