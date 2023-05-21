package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "onChainLayer/testutil/keeper"
	"onChainLayer/x/poa/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PoaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
