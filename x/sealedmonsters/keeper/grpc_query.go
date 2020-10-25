package keeper

import (
	"github.com/wangfeiping/flares/x/sealedmonsters/types"
)

var _ types.QueryServer = Keeper{}
