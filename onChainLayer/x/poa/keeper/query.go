package keeper

import (
	"onChainLayer/x/poa/types"
)

var _ types.QueryServer = Keeper{}
