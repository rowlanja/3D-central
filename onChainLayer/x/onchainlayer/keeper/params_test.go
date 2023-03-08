package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "onChainLayer/testutil/keeper"
	"onChainLayer/x/onchainlayer/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.OnchainlayerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
