package keeper

import (
	"onChainLayer/x/onchainlayer/types"
)

var _ types.QueryServer = Keeper{}
