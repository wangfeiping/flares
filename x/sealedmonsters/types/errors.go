package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sealedmonsters module sentinel errors
var (
	ErrMonsterExists = sdkerrors.Register(ModuleName, 3100, "monster already exists")
)
