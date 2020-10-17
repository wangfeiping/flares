package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nameservice module sentinel errors
var (
	ErrNameDoesNotExist = sdkerrors.Register(ModuleName, 1001, "name does not exist")
)
