package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nameservice module sentinel errors
var (
	ErrWhoisNotFound = sdkerrors.Register(ModuleName, 2100, "whois not found")
)
