package keeper

import (
	"github.com/wangfeiping/flares/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
